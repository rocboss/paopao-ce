package service

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/rocboss/paopao-ce/pkg/util"
)

const MAX_CAPTCHA_TIMES = 2

type PhoneCaptchaReq struct {
	Phone        string `json:"phone" form:"phone" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type UserPhoneBindReq struct {
	Phone   string `json:"phone" form:"phone" binding:"required"`
	Captcha string `json:"captcha" form:"captcha" binding:"required"`
}

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ChangePasswordReq struct {
	Password    string `json:"password" form:"password" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
}

type ChangeNicknameReq struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
}

type ChangeAvatarReq struct {
	Avatar string `json:"avatar" form:"avatar" binding:"required"`
}

const LOGIN_ERR_KEY = "PaoPaoUserLoginErr"
const MAX_LOGIN_ERR_TIMES = 10

// DoLogin 用户认证
func DoLogin(ctx *gin.Context, param *AuthRequest) (*model.User, error) {
	user, err := ds.GetUserByUsername(param.Username)
	if err != nil {
		return nil, errcode.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if errTimes, err := conf.Redis.Get(ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID)).Result(); err == nil {
			if convert.StrTo(errTimes).MustInt() >= MAX_LOGIN_ERR_TIMES {
				return nil, errcode.TooManyLoginError
			}
		}

		// 对比密码是否正确
		if ValidPassword(user.Password, param.Password, user.Salt) {

			if user.Status == model.UserStatusClosed {
				return nil, errcode.UserHasBeenBanned
			}

			// 清空登录计数
			conf.Redis.Del(ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID))
			return user, nil
		}

		// 登录错误计数
		_, err = conf.Redis.Incr(ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID)).Result()
		if err == nil {
			conf.Redis.Expire(ctx, fmt.Sprintf("%s:%d", LOGIN_ERR_KEY, user.ID), time.Hour).Result()
		}

		return nil, errcode.UnauthorizedAuthFailed
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

// ValidPassword 检查密码是否一致
func ValidPassword(dbPassword, password, salt string) bool {
	return strings.Compare(dbPassword, util.EncodeMD5(util.EncodeMD5(password)+salt)) == 0
}

// CheckStatus 检测用户权限
func CheckStatus(user *model.User) bool {
	return user.Status == model.UserStatusNormal
}

// ValidUsername 验证用户
func ValidUsername(username string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return errcode.UsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return errcode.UsernameCharLimit
	}

	// 重复检查
	user, _ := ds.GetUserByUsername(username)

	if user.Model != nil && user.ID > 0 {
		return errcode.UsernameHasExisted
	}

	return nil
}

// CheckPassword 密码检查
func CheckPassword(password string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 16 {
		return errcode.PasswordLengthLimit
	}

	return nil
}

// CheckPhoneCaptcha 验证手机验证码
func CheckPhoneCaptcha(phone, captcha string) *errcode.Error {
	// 如果禁止phone verify 则允许通过任意验证码
	if DisablePhoneVerify {
		return nil
	}

	c, err := ds.GetLatestPhoneCaptcha(phone)
	if err != nil {
		return errcode.ErrorPhoneCaptcha
	}

	if c.Captcha != captcha {
		return errcode.ErrorPhoneCaptcha
	}

	if c.ExpiredOn < time.Now().Unix() {
		return errcode.ErrorPhoneCaptcha
	}

	if c.UseTimes >= MAX_CAPTCHA_TIMES {
		return errcode.MaxPhoneCaptchaUseTimes
	}

	// 更新检测次数
	ds.UsePhoneCaptcha(c)

	return nil
}

// CheckPhoneExist 检测手机号是否存在
func CheckPhoneExist(uid int64, phone string) bool {
	u, err := ds.GetUserByPhone(phone)
	if err != nil {
		return false
	}

	if u.Model == nil || u.ID == 0 {
		return false
	}

	if u.ID == uid {
		return false
	}

	return true
}

// EncryptPasswordAndSalt 密码加密&生成salt
func EncryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = util.EncodeMD5(util.EncodeMD5(password) + salt)

	return password, salt
}

// Register 用户注册
func Register(username, password string) (*model.User, error) {
	password, salt := EncryptPasswordAndSalt(password)

	user := &model.User{
		Nickname: username,
		Username: username,
		Password: password,
		Avatar:   GetRandomAvatar(),
		Salt:     salt,
		Status:   model.UserStatusNormal,
	}

	user, err := ds.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(param *AuthRequest) (*model.User, error) {
	user, err := ds.GetUserByUsername(param.Username)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

func GetUserByUsername(username string) (*model.User, error) {
	user, err := ds.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.NoExistUsername
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(user *model.User) error {
	return ds.UpdateUser(user)
}

// GetUserCollections 获取用户收藏列表
func GetUserCollections(userID int64, offset, limit int) ([]*model.PostFormated, int64, error) {
	collections, err := ds.GetUserPostCollections(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := ds.GetUserPostCollectionCount(userID)
	if err != nil {
		return nil, 0, err
	}
	postIDs := []int64{}
	for _, collection := range collections {
		postIDs = append(postIDs, collection.PostID)
	}

	// 获取Posts
	posts, err := ds.GetPosts(&model.ConditionsT{
		"id IN ?": postIDs,
		"ORDER":   "id DESC",
	}, 0, 0)
	if err != nil {
		return nil, 0, err
	}

	postsFormated, err := FormatPosts(posts)
	if err != nil {
		return nil, 0, err
	}

	return postsFormated, totalRows, nil
}

// GetUserStars 获取用户点赞列表
func GetUserStars(userID int64, offset, limit int) ([]*model.PostFormated, int64, error) {
	stars, err := ds.GetUserPostStars(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := ds.GetUserPostStarCount(userID)
	if err != nil {
		return nil, 0, err
	}
	postIDs := []int64{}
	for _, star := range stars {
		postIDs = append(postIDs, star.PostID)
	}

	// 获取Posts
	posts, err := ds.GetPosts(&model.ConditionsT{
		"id IN ?": postIDs,
		"ORDER":   "id DESC",
	}, 0, 0)
	if err != nil {
		return nil, 0, err
	}

	postsFormated, err := FormatPosts(posts)
	if err != nil {
		return nil, 0, err
	}

	return postsFormated, totalRows, nil
}

// GetUserWalletBills 获取用户账单列表
func GetUserWalletBills(userID int64, offset, limit int) ([]*model.WalletStatement, int64, error) {
	bills, err := ds.GetUserWalletBills(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := ds.GetUserWalletBillCount(userID)
	if err != nil {
		return nil, 0, err
	}

	return bills, totalRows, nil
}

// SendPhoneCaptcha 发送短信验证码
func SendPhoneCaptcha(ctx *gin.Context, phone string) error {

	err := ds.SendPhoneCaptcha(phone)
	if err != nil {
		return err
	}

	// 写入计数缓存
	conf.Redis.Incr(ctx, "PaoPaoSmsCaptcha:"+phone).Result()

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	conf.Redis.Expire(ctx, "PaoPaoSmsCaptcha:"+phone, endTime.Sub(currentTime))

	return nil
}

// GetSuggestUsers 根据关键词获取用户推荐
func GetSuggestUsers(keyword string) ([]string, error) {
	users, err := ds.GetUsersByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	return usernames, nil
}

// GetSuggestTags 根据关键词获取标签推荐
func GetSuggestTags(keyword string) ([]string, error) {
	tags, err := ds.GetTagsByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	ts := []string{}
	for _, t := range tags {
		ts = append(ts, t.Tag)
	}

	return ts, nil
}
