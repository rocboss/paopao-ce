package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"gopkg.in/resty.v1"
)

type JuhePhoneCaptchaRsp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
}

// 根据用户ID获取用户
func (d *dataServant) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{
		Model: &model.Model{
			ID: id,
		},
	}

	return user.Get(d.engine)
}

// 根据用户名获取用户
func (d *dataServant) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}

	return user.Get(d.engine)
}

// 根据手机号获取用户
func (d *dataServant) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{
		Phone: phone,
	}

	return user.Get(d.engine)
}

// 根据IDs获取用户列表
func (d *dataServant) GetUsersByIDs(ids []int64) ([]*model.User, error) {
	user := &model.User{}

	return user.List(d.engine, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

// 根据关键词模糊获取用户列表
func (d *dataServant) GetUsersByKeyword(keyword string) ([]*model.User, error) {
	user := &model.User{}

	if strings.Trim(keyword, "") == "" {
		return user.List(d.engine, &model.ConditionsT{
			"ORDER": "id ASC",
		}, 0, 6)
	} else {

		return user.List(d.engine, &model.ConditionsT{
			"username LIKE ?": strings.Trim(keyword, "") + "%",
		}, 0, 6)
	}
}

// 根据关键词模糊获取用户列表
func (d *dataServant) GetTagsByKeyword(keyword string) ([]*model.Tag, error) {
	tag := &model.Tag{}

	if strings.Trim(keyword, "") == "" {
		return tag.List(d.engine, &model.ConditionsT{
			"ORDER": "quote_num DESC",
		}, 0, 6)
	} else {

		return tag.List(d.engine, &model.ConditionsT{
			"tag LIKE ?": "%" + strings.Trim(keyword, "") + "%",
			"ORDER":      "quote_num DESC",
		}, 0, 6)
	}
}

// 创建用户
func (d *dataServant) CreateUser(user *model.User) (*model.User, error) {
	return user.Create(d.engine)
}

// 更新用户
func (d *dataServant) UpdateUser(user *model.User) error {
	return user.Update(d.engine)
}

// 获取最新短信验证码
func (d *dataServant) GetLatestPhoneCaptcha(phone string) (*model.Captcha, error) {
	return (&model.Captcha{
		Phone: phone,
	}).Get(d.engine)
}

// 更新短信验证码
func (d *dataServant) UsePhoneCaptcha(captcha *model.Captcha) error {
	captcha.UseTimes++
	return captcha.Update(d.engine)
}

// 发送短信验证码
func (d *dataServant) SendPhoneCaptcha(phone string) error {
	rand.Seed(time.Now().UnixNano())
	captcha := rand.Intn(900000) + 100000
	m := 5

	gateway := "https://v.juhe.cn/sms/send"

	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		SetFormData(map[string]string{
			"mobile":    phone,
			"tpl_id":    global.SmsJuheSetting.TplID,
			"tpl_value": fmt.Sprintf(global.SmsJuheSetting.TplVal, captcha, m),
			"key":       global.SmsJuheSetting.Key,
		}).Post(gateway)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	result := &JuhePhoneCaptchaRsp{}
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
	captchaModel.Create(d.engine)
	return nil
}
