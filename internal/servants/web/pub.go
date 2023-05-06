// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/color"
	"image/png"
	"regexp"
	"unicode/utf8"

	"github.com/afocus/captcha"
	"github.com/alimy/mir/v3"
	"github.com/gofrs/uuid"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/web/assets"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/version"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var (
	_ api.Pub        = (*pubSrv)(nil)
	_ api.PubBinding = (*pubBinding)(nil)
	_ api.PubRender  = (*pubRender)(nil)
)

const (
	_MaxLoginErrTimes = 10
	_MaxPhoneCaptcha  = 10
)

type pubSrv struct {
	api.UnimplementedPubServant
	*base.DaoServant
}

type pubBinding struct {
	*api.UnimplementedPubBinding
}

type pubRender struct {
	*api.UnimplementedPubRender
}

func (s *pubSrv) TweetDetail(req *web.TweetDetailReq) (*web.TweetDetailResp, mir.Error) {
	post, err := s.Ds.GetPostByID(req.TweetId)
	if err != nil {
		return nil, _errGetPostFailed
	}
	postContents, err := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, _errGetPostFailed
	}
	users, err := s.Ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, _errGetPostFailed
	}
	// 数据整合
	postFormated := post.Format()
	for _, user := range users {
		postFormated.User = user.Format()
	}
	for _, content := range postContents {
		if content.PostID == post.ID {
			postFormated.Contents = append(postFormated.Contents, content.Format())
		}
	}
	return (*web.TweetDetailResp)(postFormated), nil
}

func (s *pubSrv) SendCaptcha(req *web.SendCaptchaReq) mir.Error {
	ctx := context.Background()

	// 验证图片验证码
	if captcha, err := s.Redis.GetImgCaptcha(ctx, req.ImgCaptchaID); err != nil || string(captcha) != req.ImgCaptcha {
		logrus.Debugf("get captcha err:%s expect:%s got:%s", err, captcha, req.ImgCaptcha)
		return _errErrorCaptchaPassword
	}
	s.Redis.DelImgCaptcha(ctx, req.ImgCaptchaID)

	// 今日频次限制
	if count, _ := s.Redis.GetCountSmsCaptcha(ctx, req.Phone); count >= _MaxPhoneCaptcha {
		return _errTooManyPhoneCaptchaSend
	}

	if err := s.Ds.SendPhoneCaptcha(req.Phone); err != nil {
		return xerror.ServerError
	}
	// 写入计数缓存
	s.Redis.IncrCountSmsCaptcha(ctx, req.Phone)

	return nil
}

func (s *pubSrv) GetCaptcha() (*web.GetCaptchaResp, mir.Error) {
	cap := captcha.New()
	if err := cap.AddFontFromBytes(assets.ComicBytes); err != nil {
		logrus.Errorf("cap.AddFontFromBytes err:%s", err)
		return nil, xerror.ServerError
	}
	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	if err := png.Encode(emptyBuff, img); err != nil {
		logrus.Errorf("png.Encode err:%s", err)
		return nil, xerror.ServerError
	}
	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())
	// 五分钟有效期
	s.Redis.SetImgCaptcha(context.Background(), key, password)
	return &web.GetCaptchaResp{
		Id:      key,
		Content: "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	}, nil
}

func (s *pubSrv) Register(req *web.RegisterReq) (*web.RegisterResp, mir.Error) {
	if _disallowUserRegister {
		return nil, _errDisallowUserRegister
	}
	// 用户名检查
	if err := s.validUsername(req.Username); err != nil {
		return nil, err
	}
	// 密码检查
	if err := checkPassword(req.Password); err != nil {
		logrus.Errorf("scheckPassword err: %v", err)
		return nil, _errUserRegisterFailed
	}
	password, salt := encryptPasswordAndSalt(req.Password)
	user := &core.User{
		Nickname: req.Username,
		Username: req.Username,
		Password: password,
		Avatar:   getRandomAvatar(),
		Salt:     salt,
		Status:   core.UserStatusNormal,
	}
	user, err := s.Ds.CreateUser(user)
	if err != nil {
		logrus.Errorf("Ds.CreateUser err: %s", err)
		return nil, _errUserRegisterFailed
	}
	return &web.RegisterResp{
		UserId:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *pubSrv) Login(req *web.LoginReq) (*web.LoginResp, mir.Error) {
	ctx := context.Background()
	user, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err:%s", err)
		return nil, xerror.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if count, err := s.Redis.GetCountLoginErr(ctx, user.ID); err == nil && count >= _MaxLoginErrTimes {
			return nil, _errTooManyLoginError
		}
		// 对比密码是否正确
		if validPassword(user.Password, req.Password, user.Salt) {
			if user.Status == core.UserStatusClosed {
				return nil, _errUserHasBeenBanned
			}
			// 清空登录计数
			s.Redis.DelCountLoginErr(ctx, user.ID)
		} else {
			// 登录错误计数
			s.Redis.IncrCountLoginErr(ctx, user.ID)
			return nil, xerror.UnauthorizedAuthFailed
		}
	} else {
		return nil, xerror.UnauthorizedAuthNotExist
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		return nil, xerror.UnauthorizedTokenGenerate
	}
	return &web.LoginResp{
		Token: token,
	}, nil
}

func (s *pubSrv) Version() (*web.VersionResp, mir.Error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

// validUsername 验证用户
func (s *pubSrv) validUsername(username string) mir.Error {
	// 检测用户是否合规
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return _errUsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return _errUsernameCharLimit
	}

	// 重复检查
	user, _ := s.Ds.GetUserByUsername(username)
	if user.Model != nil && user.ID > 0 {
		return _errUsernameHasExisted
	}
	return nil
}

func newPubSrv(s *base.DaoServant) api.Pub {
	return &pubSrv{
		DaoServant: s,
	}
}

func newPubBinding() api.PubBinding {
	return &pubBinding{
		UnimplementedPubBinding: &api.UnimplementedPubBinding{
			BindAny: base.NewBindAnyFn(),
		},
	}
}

func newPubRender() api.PubRender {
	return &pubRender{
		UnimplementedPubRender: &api.UnimplementedPubRender{
			RenderAny: base.RenderAny,
		},
	}
}
