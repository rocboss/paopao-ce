package security

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/json"
	"gopkg.in/resty.v1"
)

var (
	_ core.PhoneVerifyService = (*juheSmsServant)(nil)
)

type juhePhoneCaptchaRsp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

type juheSmsServant struct {
	gateway string
	key     string
	tplID   string
	tplVal  string
}

// SendPhoneCaptcha 发送短信验证码
func (s *juheSmsServant) SendPhoneCaptcha(phone string, captcha string, expire time.Duration) error {
	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		SetFormData(map[string]string{
			"mobile":    phone,
			"tpl_id":    s.tplID,
			"tpl_value": fmt.Sprintf(s.tplVal, captcha, expire),
			"key":       s.key,
		}).Post(s.gateway)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	result := &juhePhoneCaptchaRsp{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return err
	}
	if result.ErrorCode != 0 {
		return errors.New(result.Reason)
	}
	return nil
}

func newJuheSmsServant() *juheSmsServant {
	return &juheSmsServant{
		gateway: conf.SmsJuheSetting.Gateway,
		key:     conf.SmsJuheSetting.Key,
		tplID:   conf.SmsJuheSetting.TplID,
		tplVal:  conf.SmsJuheSetting.TplVal,
	}
}
