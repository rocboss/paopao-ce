package api

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/model/rest"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

// Login 用户登录
func Login(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, err := service.DoLogin(c, &param)
	if err != nil {
		logrus.Errorf("service.DoLogin err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}

// Register 用户注册
func Register(c *gin.Context) {

	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 用户名检查
	err := service.ValidUsername(param.Username)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 密码检查
	err = service.CheckPassword(param.Password)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	user, err := service.Register(
		param.Username,
		param.Password,
	)

	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterFailed)
		return
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

func RequestingFriend(c *gin.Context) {
	param := rest.RequestingFriendReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, ok := userFrom(c)
	if !ok {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}
	if user.ID == param.UserId {
		response.ToErrorResponse(errcode.NoRequestingFriendToSelf)
		return
	}

	if err := service.RequestingFriend(user, &param); err != nil {
		logrus.Errorf("service.RequestingFriend err: %v", err)
		response.ToErrorResponse(errcode.SendRequestingFriendFailed)
		return
	}

	response.ToResponse(nil)
}

func AddFriend(c *gin.Context) {
	param := rest.AddFriendReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	user, ok := userFrom(c)
	if !ok {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}
	if user.ID == param.UserId {
		response.ToErrorResponse(errcode.NoActionToSelf)
		return
	}

	if err := service.AddFriend(user, &param); err != nil {
		logrus.Errorf("service.AddFriend err: %v", err)
		response.ToErrorResponse(errcode.AddFriendFailed)
		return
	}

	response.ToResponse(nil)
}

func RejectFriend(c *gin.Context) {
	param := rest.RejectFriendReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	user, ok := userFrom(c)
	if !ok {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}
	if user.ID == param.UserId {
		response.ToErrorResponse(errcode.NoActionToSelf)
		return
	}

	if err := service.RejectFriend(user, &param); err != nil {
		logrus.Errorf("service.RejectFriend err: %v", err)
		response.ToErrorResponse(errcode.RejectFriendFailed)
		return
	}

	response.ToResponse(nil)
}

func DeleteFriend(c *gin.Context) {
	param := rest.DeleteFriendReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	user, ok := userFrom(c)
	if !ok {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}
	if user.ID == param.UserId {
		response.ToErrorResponse(errcode.NoActionToSelf)
		return
	}

	if err := service.DeleteFriend(user, &param); err != nil {
		logrus.Errorf("service.DeleteFriend err: %v", err)
		response.ToErrorResponse(errcode.DeleteFriendFailed)
		return
	}

	response.ToResponse(nil)
}

func GetContacts(c *gin.Context) {
	response := app.NewResponse(c)
	user, ok := userFrom(c)
	if !ok {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	}

	offset, limit := app.GetPageOffset(c)
	resp, err := service.GetContacts(user, offset, limit)
	if err != nil {
		logrus.Errorf("service.DeleteFriend err: %v", err)
		response.ToErrorResponse(errcode.GetContactsFailed)
		return
	}
	response.ToResponseList(resp.Contacts, resp.Total)
}

// 获取用户基本信息
func GetUserInfo(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	if username, exists := c.Get("USERNAME"); exists {
		param.Username = username.(string)
	}

	user, err := service.GetUserInfo(&param)

	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	phone := ""
	if user.Phone != "" && len(user.Phone) == 11 {
		phone = user.Phone[0:3] + "****" + user.Phone[7:]
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"nickname": user.Nickname,
		"username": user.Username,
		"status":   user.Status,
		"avatar":   user.Avatar,
		"balance":  user.Balance,
		"phone":    phone,
		"is_admin": user.IsAdmin,
	})
}

// 修改密码
func ChangeUserPassword(c *gin.Context) {
	param := service.ChangePasswordReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	// 密码检查
	err := service.CheckPassword(param.Password)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 旧密码校验
	if !service.ValidPassword(user.Password, param.OldPassword, user.Salt) {
		response.ToErrorResponse(errcode.ErrorOldPassword)
		return
	}

	// 更新入库
	user.Password, user.Salt = service.EncryptPasswordAndSalt(param.Password)
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改昵称
func ChangeNickname(c *gin.Context) {
	param := service.ChangeNicknameReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	if utf8.RuneCountInString(param.Nickname) < 2 || utf8.RuneCountInString(param.Nickname) > 12 {
		response.ToErrorResponse(errcode.NicknameLengthLimit)
		return
	}

	// 执行绑定
	user.Nickname = param.Nickname
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改头像
func ChangeAvatar(c *gin.Context) {
	param := service.ChangeAvatarReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, exist := userFrom(c)
	if !exist {
		response.ToErrorResponse(errcode.UnauthorizedTokenError)
		return
	}

	if err := service.ChangeUserAvatar(user, param.Avatar); err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(nil)
}

// 用户绑定手机号
func BindUserPhone(c *gin.Context) {
	param := service.UserPhoneBindReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	// 手机重复性检查
	if service.CheckPhoneExist(user.ID, param.Phone) {
		response.ToErrorResponse(errcode.ExistedUserPhone)
		return
	}

	if err := service.CheckPhoneCaptcha(param.Phone, param.Captcha); err != nil {
		logrus.Errorf("service.CheckPhoneCaptcha err: %v\n", err)
		response.ToErrorResponse(err)
		return
	}

	// 执行绑定
	user.Phone = param.Phone
	if err := service.UpdateUserInfo(user); err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(nil)
}

// 修改用户状态
func ChangeUserStatus(c *gin.Context) {
	param := service.ChangeUserStatusReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if param.Status != model.UserStatusNormal && param.Status != model.UserStatusClosed {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	user, err := service.GetUserByID(param.ID)
	if err != nil {
		logrus.Errorf("service.GetUserByID err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	// 执行更新
	user.Status = param.Status
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

func GetUserProfile(c *gin.Context) {
	response := app.NewResponse(c)
	username := c.Query("username")
	user, _ := userFrom(c)

	resp, err := service.GetUserByUsername(user, username)
	if err != nil {
		logrus.Errorf("service.GetUserByUsername err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	response.ToResponse(resp)
}

func GetUserPosts(c *gin.Context) {
	response := app.NewResponse(c)
	username := c.Query("username")
	user, exists := userFrom(c)

	other, err := service.GetUserByUsername(user, username)
	if err != nil {
		logrus.Errorf("service.GetUserByUsername err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	visibilities := []model.PostVisibleT{model.PostVisitPublic}
	if exists {
		if user.ID == other.ID || user.IsAdmin {
			visibilities = append(visibilities, model.PostVisitPrivate, model.PostVisitFriend)
		} else if other.IsFriend {
			visibilities = append(visibilities, model.PostVisitFriend)
		}
	}
	conditions := model.ConditionsT{
		"user_id":         other.ID,
		"visibility IN ?": visibilities,
		"ORDER":           "latest_replied_on DESC",
	}

	_, posts, err := service.GetPostList(&service.PostListReq{
		Conditions: &conditions,
		Offset:     (app.GetPage(c) - 1) * app.GetPageSize(c),
		Limit:      app.GetPageSize(c),
	})
	if err != nil {
		logrus.Errorf("service.GetPostList err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostsFailed)
		return
	}
	totalRows, _ := service.GetPostCount(&conditions)

	response.ToResponseList(posts, totalRows)
}

func GetUserCollections(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	posts, totalRows, err := service.GetUserCollections(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		logrus.Errorf("service.GetUserCollections err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(posts, totalRows)
}

func GetUserStars(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	posts, totalRows, err := service.GetUserStars(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))
	if err != nil {
		logrus.Errorf("service.GetUserStars err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(posts, totalRows)
}

func GetSuggestUsers(c *gin.Context) {
	keyword := c.Query("k")
	response := app.NewResponse(c)

	usernames, err := service.GetSuggestUsers(keyword)
	if err != nil {
		logrus.Errorf("service.GetSuggestUsers err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponse(usernames)
}

func GetSuggestTags(c *gin.Context) {
	keyword := c.Query("k")
	response := app.NewResponse(c)

	tags, err := service.GetSuggestTags(keyword)
	if err != nil {
		logrus.Errorf("service.GetSuggestTags err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponse(tags)
}

func GetUserRechargeLink(c *gin.Context) {
	param := service.RechargeReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 下单
	userID, _ := c.Get("UID")
	recharge, err := service.CreateRecharge(userID.(int64), param.Amount)
	if err != nil {
		logrus.Errorf("service.CreateRecharge err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	p := alipay.TradePreCreate{}
	p.OutTradeNo = fmt.Sprintf("%d", recharge.ID)
	p.Subject = "PaoPao用户钱包充值"
	p.TotalAmount = fmt.Sprintf("%.2f", float64(recharge.Amount)/100.0)
	p.NotifyURL = "https://" + c.Request.Host + "/v1/alipay/notify"

	rsp, err := alipayClient.TradePreCreate(p)
	if err != nil {
		logrus.Errorf("client.TradePreCreate err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	response.ToResponse(gin.H{
		"id":  recharge.ID,
		"pay": rsp.Content.QRCode,
	})
}

func GetUserRechargeResult(c *gin.Context) {
	response := app.NewResponse(c)

	id := c.Query("id")
	userID, _ := c.Get("UID")

	recharge, err := service.GetRechargeByID(convert.StrTo(id).MustInt64())
	if err != nil {
		response.ToErrorResponse(errcode.GetRechargeFailed)
		return
	}

	if recharge.UserID != userID.(int64) {
		response.ToErrorResponse(errcode.GetRechargeFailed)
		return
	}

	response.ToResponse(gin.H{
		"id":     recharge.ID,
		"status": recharge.TradeStatus,
	})
}

func AlipayNotify(c *gin.Context) {
	response := app.NewResponse(c)
	c.Request.ParseForm()

	_, err := alipayClient.GetTradeNotification(c.Request)
	if err != nil {
		logrus.Errorf("alipayClient.GetTradeNotification err: %v\n", err)
		logrus.Infoln(c.Request.Form)
		response.ToErrorResponse(errcode.RechargeNotifyError)
		return
	}

	id := c.Request.Form.Get("out_trade_no")
	tradeNo := c.Request.Form.Get("trade_no")
	tradeStatus := c.Request.Form.Get("trade_status")
	if tradeStatus == "TRADE_SUCCESS" {
		// 交易支付成功
		err = service.FinishRecharge(c, convert.StrTo(id).MustInt64(), tradeNo)
		if err != nil {
			logrus.Errorf("service.FinishRecharge err: %v\n", err)
			response.ToErrorResponse(errcode.RechargeNotifyError)
			return
		}
	}
	response.Ctx.String(http.StatusOK, "success")
}

func GetUserWalletBills(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	bills, totalRows, err := service.GetUserWalletBills(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		logrus.Errorf("service.GetUserWalletBills err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(bills, totalRows)
}

func userFrom(c *gin.Context) (*model.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*model.User)
		return user, ok
	}
	return nil, false
}

func userIdFrom(c *gin.Context) (int64, bool) {
	if u, exists := c.Get("UID"); exists {
		uid, ok := u.(int64)
		return uid, ok
	}
	return -1, false
}
