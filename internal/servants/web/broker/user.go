// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/sirupsen/logrus"
)

const _MaxCaptchaTimes = 2

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

type ChangeUserStatusReq struct {
	ID     int64 `json:"id" form:"id" binding:"required"`
	Status int   `json:"status" form:"status" binding:"required"`
}

type RequestingFriendReq struct {
	UserId    int64  `json:"user_id" binding:"required"`
	Greetings string `json:"greetings" binding:"required"`
}

type AddFriendReq struct {
	UserId int64 `json:"user_id" binding:"required"`
}

type RejectFriendReq struct {
	UserId int64 `json:"user_id" binding:"required"`
}

type DeleteFriendReq struct {
	UserId int64 `json:"user_id"`
}

type UserProfileResp struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"is_admin"`
	IsFriend bool   `json:"is_friend"`
}

const (
	_LoginErrKey      = "PaoPaoUserLoginErr"
	_MaxLoginErrTimes = 10
)

// DoLogin 用户认证
func DoLogin(ctx *gin.Context, param *AuthRequest) (*core.User, error) {
	user, err := ds.GetUserByUsername(param.Username)
	if err != nil {
		return nil, errcode.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if errTimes, err := redisClient.Get(ctx, fmt.Sprintf("%s:%d", _LoginErrKey, user.ID)).Result(); err == nil {
			if convert.StrTo(errTimes).MustInt() >= _MaxLoginErrTimes {
				return nil, errcode.TooManyLoginError
			}
		}

		// 对比密码是否正确
		if ValidPassword(user.Password, param.Password, user.Salt) {

			if user.Status == core.UserStatusClosed {
				return nil, errcode.UserHasBeenBanned
			}

			// 清空登录计数
			redisClient.Del(ctx, fmt.Sprintf("%s:%d", _LoginErrKey, user.ID))
			return user, nil
		}

		// 登录错误计数
		_, err = redisClient.Incr(ctx, fmt.Sprintf("%s:%d", _LoginErrKey, user.ID)).Result()
		if err == nil {
			redisClient.Expire(ctx, fmt.Sprintf("%s:%d", _LoginErrKey, user.ID), time.Hour).Result()
		}

		return nil, errcode.UnauthorizedAuthFailed
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

// ValidPassword 检查密码是否一致
func ValidPassword(dbPassword, password, salt string) bool {
	return strings.Compare(dbPassword, utils.EncodeMD5(utils.EncodeMD5(password)+salt)) == 0
}

// CheckStatus 检测用户权限
func CheckStatus(user *core.User) bool {
	return user.Status == core.UserStatusNormal
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

	if c.UseTimes >= _MaxCaptchaTimes {
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
	password = utils.EncodeMD5(utils.EncodeMD5(password) + salt)

	return password, salt
}

// Register 用户注册
func Register(username, password string) (*core.User, error) {
	password, salt := EncryptPasswordAndSalt(password)

	user := &core.User{
		Nickname: username,
		Username: username,
		Password: password,
		Avatar:   GetRandomAvatar(),
		Salt:     salt,
		Status:   core.UserStatusNormal,
	}

	user, err := ds.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func RequestingFriend(user *core.User, param *RequestingFriendReq) error {
	if _, err := ds.GetUserByID(param.UserId); err != nil {
		return errcode.NotExistFriendId
	}
	return ds.RequestingFriend(user.ID, param.UserId, param.Greetings)
}

func AddFriend(user *core.User, param *AddFriendReq) error {
	if _, err := ds.GetUserByID(param.UserId); err != nil {
		return errcode.NotExistFriendId
	}
	return ds.AddFriend(user.ID, param.UserId)
}

func RejectFriend(user *core.User, param *RejectFriendReq) error {
	if _, err := ds.GetUserByID(param.UserId); err != nil {
		return errcode.NotExistFriendId
	}
	return ds.RejectFriend(user.ID, param.UserId)
}

func DeleteFriend(user *core.User, param *DeleteFriendReq) error {
	if _, err := ds.GetUserByID(param.UserId); err != nil {
		return errcode.NotExistFriendId
	}
	return ds.DeleteFriend(user.ID, param.UserId)
}

func GetContacts(user *core.User, offset int, limit int) (*core.ContactList, error) {
	return ds.GetContacts(user.ID, offset, limit)
}

// GetUserInfo 获取用户信息
func GetUserInfo(param *AuthRequest) (*core.User, error) {
	user, err := ds.GetUserByUsername(param.Username)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.UnauthorizedAuthNotExist
}

func GetUserByID(id int64) (*core.User, error) {
	user, err := ds.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	if user.Model != nil && user.ID > 0 {
		return user, nil
	}

	return nil, errcode.NoExistUsername
}

func GetUserByUsername(user *core.User, username string) (*UserProfileResp, error) {
	other, err := ds.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	var resp *UserProfileResp
	if other.Model != nil && other.ID > 0 {
		resp = &UserProfileResp{
			ID:       other.ID,
			Nickname: other.Nickname,
			Username: other.Username,
			Status:   other.Status,
			Avatar:   other.Avatar,
			IsAdmin:  other.IsAdmin,
			IsFriend: !(user == nil || user.ID == other.ID),
		}
	} else {
		return nil, errcode.NoExistUsername
	}

	if user != nil && user.ID != other.ID {
		resp.IsFriend = ds.IsFriend(user.ID, other.ID)
	}
	return resp, nil
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(user *core.User) *errcode.Error {
	if err := ds.UpdateUser(user); err != nil {
		return errcode.ServerError
	}
	return nil
}

func ChangeUserAvatar(user *core.User, avatar string) (err *errcode.Error) {
	defer func() {
		if err != nil {
			deleteOssObjects([]string{avatar})
		}
	}()

	if err := ds.CheckAttachment(avatar); err != nil {
		return errcode.InvalidParams
	}

	if err := oss.PersistObject(oss.ObjectKey(avatar)); err != nil {
		logrus.Errorf("service.ChangeUserAvatar persist object failed: %s", err)
		return errcode.ServerError
	}

	user.Avatar = avatar
	err = UpdateUserInfo(user)
	return
}

// GetUserCollections 获取用户收藏列表
func GetUserCollections(userID int64, offset, limit int) ([]*core.PostFormated, int64, error) {
	collections, err := ds.GetUserPostCollections(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := ds.GetUserPostCollectionCount(userID)
	if err != nil {
		return nil, 0, err
	}
	var posts []*core.Post
	for _, collection := range collections {
		posts = append(posts, collection.Post)
	}
	postsFormated, err := ds.MergePosts(posts)
	if err != nil {
		return nil, 0, err
	}

	return postsFormated, totalRows, nil
}

// GetUserStars 获取用户点赞列表
func GetUserStars(userID int64, offset, limit int) ([]*core.PostFormated, int64, error) {
	stars, err := ds.GetUserPostStars(userID, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := ds.GetUserPostStarCount(userID)
	if err != nil {
		return nil, 0, err
	}

	var posts []*core.Post
	for _, star := range stars {
		posts = append(posts, star.Post)
	}
	postsFormated, err := ds.MergePosts(posts)
	if err != nil {
		return nil, 0, err
	}

	return postsFormated, totalRows, nil
}

// GetUserWalletBills 获取用户账单列表
func GetUserWalletBills(userID int64, offset, limit int) ([]*core.WalletStatement, int64, error) {
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
	redisClient.Incr(ctx, "PaoPaoSmsCaptcha:"+phone).Result()

	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	redisClient.Expire(ctx, "PaoPaoSmsCaptcha:"+phone, endTime.Sub(currentTime))

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
	tags, err := ds.TagsByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	ts := []string{}
	for _, t := range tags {
		ts = append(ts, t.Tag)
	}

	return ts, nil
}

func IsFriend(userId, friendId int64) bool {
	return ds.IsFriend(userId, friendId)
}

// checkPermision 检查是否拥有者或管理员
func checkPermision(user *core.User, targetUserId int64) *errcode.Error {
	if user == nil || (user.ID != targetUserId && !user.IsAdmin) {
		return errcode.NoPermission
	}
	return nil
}
