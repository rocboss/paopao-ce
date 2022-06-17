package api

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"image/color"
	"image/png"
	"time"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/debug"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/rocboss/paopao-ce/pkg/util"
	"github.com/sirupsen/logrus"
)

//go:embed assets/comic.ttf
var comic []byte

const MAX_PHONE_CAPTCHA = 10

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"BuildInfo": debug.ReadBuildInfo(),
	})
}

func SyncSearchIndex(c *gin.Context) {
	response := app.NewResponse(c)

	user, _ := c.Get("USER")

	if user.(*model.User).IsAdmin {
		go service.PushPostsToSearch(c)
	}

	response.ToResponse(nil)
}

func GetCaptcha(c *gin.Context) {
	cap := captcha.New()

	if err := cap.AddFontFromBytes(comic); err != nil {
		panic(err.Error())
	}

	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	_ = png.Encode(emptyBuff, img)

	key := util.EncodeMD5(uuid.Must(uuid.NewV4()).String())

	// 五分钟有效期
	conf.Redis.SetEX(c, "PaoPaoCaptcha:"+key, password, time.Minute*5)

	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"id":   key,
		"b64s": "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	})
}

func PostCaptcha(c *gin.Context) {
	param := service.PhoneCaptchaReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 验证图片验证码
	if res, err := conf.Redis.Get(c.Request.Context(), "PaoPaoCaptcha:"+param.ImgCaptchaID).Result(); err != nil || res != param.ImgCaptcha {
		response.ToErrorResponse(errcode.ErrorCaptchaPassword)
		return
	}
	conf.Redis.Del(c.Request.Context(), "PaoPaoCaptcha:"+param.ImgCaptchaID).Result()

	// 今日频次限制
	if res, _ := conf.Redis.Get(c.Request.Context(), "PaoPaoSmsCaptcha:"+param.Phone).Result(); convert.StrTo(res).MustInt() >= MAX_PHONE_CAPTCHA {
		response.ToErrorResponse(errcode.TooManyPhoneCaptchaSend)
		return
	}

	err := service.SendPhoneCaptcha(c, param.Phone)
	if err != nil {
		logrus.Errorf("app.SendPhoneCaptcha errs: %v", errs)
		response.ToErrorResponse(errcode.GetPhoneCaptchaError)
		return
	}

	response.ToResponse(nil)
}
