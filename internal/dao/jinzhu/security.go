package jinzhu

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/json"
	"gopkg.in/resty.v1"
	"gorm.io/gorm"
)

var (
	_ core.SecurityService = (*securityServant)(nil)
)

type securityServant struct {
	db *gorm.DB
}

func newSecurityService(db *gorm.DB) core.SecurityService {
	return &securityServant{
		db: db,
	}
}

type juhePhoneCaptchaRsp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

// 获取最新短信验证码
func (s *securityServant) GetLatestPhoneCaptcha(phone string) (*model.Captcha, error) {
	return (&model.Captcha{
		Phone: phone,
	}).Get(s.db)
}

// 更新短信验证码
func (s *securityServant) UsePhoneCaptcha(captcha *model.Captcha) error {
	captcha.UseTimes++
	return captcha.Update(s.db)
}

// 发送短信验证码
func (s *securityServant) SendPhoneCaptcha(phone string) error {
	rand.Seed(time.Now().UnixNano())
	captcha := rand.Intn(900000) + 100000
	m := 5

	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		SetFormData(map[string]string{
			"mobile":    phone,
			"tpl_id":    conf.SmsJuheSetting.TplID,
			"tpl_value": fmt.Sprintf(conf.SmsJuheSetting.TplVal, captcha, m),
			"key":       conf.SmsJuheSetting.Key,
		}).Post(conf.SmsJuheSetting.Gateway)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	result := &juhePhoneCaptchaRsp{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return err
	}

	if result.ErrorCode != 0 {
		return errors.New(result.Reason)
	}

	// 写入表
	captchaModel := &model.Captcha{
		Phone:     phone,
		Captcha:   strconv.Itoa(captcha),
		ExpiredOn: time.Now().Add(time.Minute * time.Duration(m)).Unix(),
	}
	captchaModel.Create(s.db)
	return nil
}
