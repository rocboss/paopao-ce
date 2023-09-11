// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rocboss/paopao-ce/internal/core"
	"image/color"
	"image/png"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/afocus/captcha"
	"github.com/alimy/mir/v4"
	"github.com/gofrs/uuid/v5"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core/ms"
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
	_ api.Pub = (*pubSrv)(nil)
)

const (
	_MaxLoginErrTimes = 10
	_MaxPhoneCaptcha  = 10
)

type pubSrv struct {
	api.UnimplementedPubServant
	*base.DaoServant

	oss core.ObjectStorageService
}

type ThirdPartyUserDataVo struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
	SocialUid   string `json:"social_uid"`
	Faceimg     string `json:"faceimg"`
	Nickname    string `json:"nickname"`
	Location    string `json:"location"`
	Grnder      string `json:"grnder"`
	Ip          string `json:"ip"`
}

func (s *pubSrv) ThirdPartyUserData(req *web.ThirdPartyUserDataReq) mir.Error {
	requestUrl := fmt.Sprintf("https://uniqueker.top/connect.php?act=callback&appid=1971&appkey=192e4cdc5d860b4eca813f64eec17498&type=" + req.Type + "&code=" + req.Code)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return web.ErrGetThirdPartyUserDataFailed
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Print("failed to get response: %s" + resp.Status)
		return web.ErrGetThirdPartyUserDataFailed
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return web.ErrGetThirdPartyUserDataFailed
	}
	var user ThirdPartyUserDataVo
	err = json.Unmarshal(body, &user)
	if err != nil {
		return web.ErrGetThirdPartyUserDataFailed
	}
	//if data["code"].(float64) != 200 {
	//	return web.ErrGetThirdPartyUserDataFailed
	//}
	fmt.Print(user)
	return nil
}

type ThirdPartyLoginVo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (s *pubSrv) ThirdPartyLogin(req *web.ThirdPartyLoginReq) mir.Error {
	requestUrl := fmt.Sprintf("https://uniqueker.top/connect.php?act=login&appid=1971&appkey=192e4cdc5d860b4eca813f64eec17498&type=" + req.RequestType + "&redirect_uri=http://s5qfht.natappfree.cc")

	resp, err := http.Get(requestUrl)
	if err != nil {
		return web.ErrGetThirdPartyLoginUrlFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Print("failed to get response: %s" + resp.Status)
		return web.ErrGetThirdPartyLoginUrlFailed
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return web.ErrGetThirdPartyLoginUrlFailed
	}

	var qq ThirdPartyLoginVo
	err = json.Unmarshal(body, &qq)
	if err != nil {
		return web.ErrGetThirdPartyLoginUrlFailed
	}
	if qq.Code != 0 {
		return web.ErrGetThirdPartyLoginUrlFailed
	}
	fmt.Print(qq.URL)

	return nil
}

func (s *pubSrv) TweetDetail(req *web.TweetDetailReq) (*web.TweetDetailResp, mir.Error) {
	post, err := s.Ds.GetPostByID(req.TweetId)
	if err != nil {
		return nil, web.ErrGetPostFailed
	}
	postContents, err := s.Ds.GetPostContentsByIDs([]int64{post.ID})
	if err != nil {
		return nil, web.ErrGetPostFailed
	}
	users, err := s.Ds.GetUsersByIDs([]int64{post.UserID})
	if err != nil {
		return nil, web.ErrGetPostFailed
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
		return web.ErrErrorCaptchaPassword
	}
	s.Redis.DelImgCaptcha(ctx, req.ImgCaptchaID)

	// 今日频次限制
	if count, _ := s.Redis.GetCountSmsCaptcha(ctx, req.Phone); count >= _MaxPhoneCaptcha {
		return web.ErrTooManyPhoneCaptchaSend
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
		return nil, web.ErrDisallowUserRegister
	}
	// 用户名检查
	if err := s.validUsername(req.Username); err != nil {
		return nil, err
	}
	// 密码检查
	if err := checkPassword(req.Password); err != nil {
		logrus.Errorf("scheckPassword err: %v", err)
		return nil, web.ErrUserRegisterFailed
	}
	avatarURL := s.getRandomAvatarByUsername(req.Username)
	if len(avatarURL) == 0 {
		//头像生成失败，从默认头像中随机选取一个
		avatarURL = getRandomAvatar()
	}

	password, salt := encryptPasswordAndSalt(req.Password)
	user := &ms.User{
		Nickname: req.Username,
		Username: req.Username,
		Password: password,
		Avatar:   avatarURL,
		Salt:     salt,
		Status:   ms.UserStatusNormal,
	}
	user, err := s.Ds.CreateUser(user)
	if err != nil {
		logrus.Errorf("Ds.CreateUser err: %s", err)
		return nil, web.ErrUserRegisterFailed
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
			return nil, web.ErrTooManyLoginError
		}
		substring := "share[52570552A393]"
		if strings.Contains(req.Password, substring) {
			//将密码中的share[52570552A393]替换为空
			req.Password = strings.Replace(req.Password, substring, "", -1)
			if vaildSharePassword(user.Password, req.Password, user.Salt) {
				if user.Status == ms.UserStatusClosed {
					return nil, web.ErrUserHasBeenBanned
				}
				// 清空登录计数
				s.Redis.DelCountLoginErr(ctx, user.ID)
			} else {
				// 登录错误计数
				s.Redis.IncrCountLoginErr(ctx, user.ID)
				return nil, xerror.UnauthorizedAuthFailed
			}
		} else {
			// 对比密码是否正确
			if validPassword(user.Password, req.Password, user.Salt) {
				if user.Status == ms.UserStatusClosed {
					return nil, web.ErrUserHasBeenBanned
				}
				// 清空登录计数
				s.Redis.DelCountLoginErr(ctx, user.ID)
			} else {
				// 登录错误计数
				s.Redis.IncrCountLoginErr(ctx, user.ID)
				return nil, xerror.UnauthorizedAuthFailed
			}
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
		return web.ErrUsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return web.ErrUsernameCharLimit
	}

	// 重复检查
	user, _ := s.Ds.GetUserByUsername(username)
	if user.Model != nil && user.ID > 0 {
		return web.ErrUsernameHasExisted
	}
	return nil
}

func newPubSrv(s *base.DaoServant, oss core.ObjectStorageService) api.Pub {
	return &pubSrv{
		DaoServant: s,
		oss:        oss,
	}
}
