package api

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-api/global"
	"github.com/rocboss/paopao-api/internal/model"
	"github.com/rocboss/paopao-api/internal/service"
	"github.com/rocboss/paopao-api/pkg/app"
	"github.com/rocboss/paopao-api/pkg/convert"
	"github.com/rocboss/paopao-api/pkg/errcode"
	"github.com/smartwalle/alipay/v3"
)

// 用户登录
func Login(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c)
	user, err := svc.DoLogin(&param)
	if err != nil {
		global.Logger.Errorf("svc.DoLogin err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}

// 用户注册
func Register(c *gin.Context) {

	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c)

	// 用户名检查
	err := svc.ValidUsername(param.Username)
	if err != nil {
		global.Logger.Errorf("svc.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 密码检查
	err = svc.CheckPassword(param.Password)
	if err != nil {
		global.Logger.Errorf("svc.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	user, err := svc.Register(
		param.Username,
		param.Password,
	)

	if err != nil {
		global.Logger.Errorf("svc.Register err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterFailed)
		return
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

// 获取用户基本信息
func GetUserInfo(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	svc := service.New(c)

	if username, exists := c.Get("USERNAME"); exists {
		param.Username = username.(string)
	}

	user, err := svc.GetUserInfo(&param)

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
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	svc := service.New(c)

	// 密码检查
	err := svc.CheckPassword(param.Password)
	if err != nil {
		global.Logger.Errorf("svc.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 旧密码校验
	if !svc.ValidPassword(user.Password, param.OldPassword, user.Salt) {
		response.ToErrorResponse(errcode.ErrorOldPassword)
		return
	}

	// 更新入库
	user.Password, user.Salt = svc.EncryptPasswordAndSalt(param.Password)
	svc.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改昵称
func ChangeNickname(c *gin.Context) {
	param := service.ChangeNicknameReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}
	svc := service.New(c)

	if utf8.RuneCountInString(param.Nickname) < 2 || utf8.RuneCountInString(param.Nickname) > 12 {
		response.ToErrorResponse(errcode.NicknameLengthLimit)
		return
	}

	// 执行绑定
	user.Nickname = param.Nickname
	svc.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改头像
func ChangeAvatar(c *gin.Context) {
	param := service.ChangeAvatarReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}
	svc := service.New(c)

	if strings.Index(param.Avatar, "https://"+global.AliossSetting.AliossDomain) != 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	// 执行绑定
	user.Avatar = param.Avatar
	svc.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 用户绑定手机号
func BindUserPhone(c *gin.Context) {
	param := service.UserPhoneBindReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}
	svc := service.New(c)

	// 手机重复性检查
	if svc.CheckPhoneExist(user.ID, param.Phone) {
		response.ToErrorResponse(errcode.ExistedUserPhone)
		return
	}

	// 验证短信验证码
	err := svc.CheckPhoneCaptcha(param.Phone, param.Captcha)
	if err != nil {
		global.Logger.Errorf("svc.CheckPhoneCaptcha err: %v\n", err)
		response.ToErrorResponse(err)
		return
	}

	// 执行绑定
	user.Phone = param.Phone
	svc.UpdateUserInfo(user)

	response.ToResponse(nil)
}

func GetUserProfile(c *gin.Context) {
	response := app.NewResponse(c)
	username := c.Query("username")

	svc := service.New(c)
	user, err := svc.GetUserByUsername(username)
	if err != nil {
		global.Logger.Errorf("svc.GetUserByUsername err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"nickname": user.Nickname,
		"username": user.Username,
		"status":   user.Status,
		"avatar":   user.Avatar,
		"is_admin": user.IsAdmin,
	})
}

func GetUserPosts(c *gin.Context) {
	response := app.NewResponse(c)
	username := c.Query("username")

	svc := service.New(c)
	user, err := svc.GetUserByUsername(username)
	if err != nil {
		global.Logger.Errorf("svc.GetUserByUsername err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	conditions := &model.ConditionsT{
		"user_id": user.ID,
		"ORDER":   "latest_replied_on DESC",
	}

	posts, err := svc.GetPostList(&service.PostListReq{
		Conditions: conditions,
		Offset:     (app.GetPage(c) - 1) * app.GetPageSize(c),
		Limit:      app.GetPageSize(c),
	})
	if err != nil {
		global.Logger.Errorf("svc.GetPostList err: %v\n", err)
		response.ToErrorResponse(errcode.GetPostsFailed)
		return
	}
	totalRows, _ := svc.GetPostCount(conditions)

	response.ToResponseList(posts, totalRows)
}

func GetUserCollections(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	svc := service.New(c)
	posts, totalRows, err := svc.GetUserCollections(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		global.Logger.Errorf("svc.GetUserCollections err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(posts, totalRows)
}

func GetUserStars(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	svc := service.New(c)
	posts, totalRows, err := svc.GetUserStars(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))
	if err != nil {
		global.Logger.Errorf("svc.GetUserStars err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(posts, totalRows)
}

func GetSuggestUsers(c *gin.Context) {
	keyword := c.Query("k")
	response := app.NewResponse(c)

	svc := service.New(c)
	usernames, err := svc.GetSuggestUsers(keyword)
	if err != nil {
		global.Logger.Errorf("svc.GetSuggestUsers err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponse(usernames)
}

func GetSuggestTags(c *gin.Context) {
	keyword := c.Query("k")
	response := app.NewResponse(c)

	svc := service.New(c)
	tags, err := svc.GetSuggestTags(keyword)
	if err != nil {
		global.Logger.Errorf("svc.GetSuggestTags err: %v\n", err)
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
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 下单
	userID, _ := c.Get("UID")
	svc := service.New(c)
	recharge, err := svc.CreateRecharge(userID.(int64), param.Amount)
	if err != nil {
		global.Logger.Errorf("svc.CreateRecharge err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return

	}

	client, err := alipay.New(global.AppSetting.AlipayAppID, global.AppSetting.AlipayPrivateKey, true)
	// 将 key 的验证调整到初始化阶段
	if err != nil {
		global.Logger.Errorf("alipay.New err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	err = client.LoadAppPublicCertFromFile("configs/alipayAppCertPublicKey.crt") // 加载应用公钥证书
	if err != nil {
		global.Logger.Errorf("client.LoadAppPublicCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	err = client.LoadAliPayRootCertFromFile("configs/alipayRootCert.crt") // 加载支付宝根证书
	if err != nil {
		global.Logger.Errorf("client.LoadAliPayRootCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	err = client.LoadAliPayPublicCertFromFile("configs/alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
	if err != nil {
		global.Logger.Errorf("client.LoadAliPayPublicCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	p := alipay.TradePreCreate{}
	p.OutTradeNo = fmt.Sprintf("%d", recharge.ID)
	p.Subject = "PaoPao用户钱包充值"
	p.TotalAmount = fmt.Sprintf("%.2f", float64(recharge.Amount)/100.0)
	p.NotifyURL = "https://" + c.Request.Host + "/alipay/notify"

	rsp, err := client.TradePreCreate(p)
	if err != nil {
		global.Logger.Errorf("client.TradePreCreate err: %v\n", err)
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

	svc := service.New(c)
	recharge, err := svc.GetRechargeByID(convert.StrTo(id).MustInt64())
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

	aliClient, err := alipay.New(global.AppSetting.AlipayAppID, global.AppSetting.AlipayPrivateKey, true)
	// 将 key 的验证调整到初始化阶段
	if err != nil {
		global.Logger.Errorf("alipay.New err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}
	err = aliClient.LoadAppPublicCertFromFile("configs/alipayAppCertPublicKey.crt") // 加载应用公钥证书
	if err != nil {
		global.Logger.Errorf("client.LoadAppPublicCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	err = aliClient.LoadAliPayRootCertFromFile("configs/alipayRootCert.crt") // 加载支付宝根证书
	if err != nil {
		global.Logger.Errorf("client.LoadAliPayRootCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	err = aliClient.LoadAliPayPublicCertFromFile("configs/alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
	if err != nil {
		global.Logger.Errorf("client.LoadAliPayPublicCertFromFile err: %v\n", err)
		response.ToErrorResponse(errcode.RechargeReqFail)
		return
	}

	_, err = aliClient.GetTradeNotification(c.Request)
	if err != nil {
		global.Logger.Errorf("aliClient.GetTradeNotification err: %v\n", err)
		global.Logger.Infoln(c.Request.Form)
		response.ToErrorResponse(errcode.RechargeNotifyError)
		return
	}

	svc := service.New(c)
	id := c.Request.Form.Get("out_trade_no")
	tradeNo := c.Request.Form.Get("trade_no")
	tradeStatus := c.Request.Form.Get("trade_status")
	if tradeStatus == "TRADE_SUCCESS" {
		// 交易支付成功
		err = svc.FinishRecharge(convert.StrTo(id).MustInt64(), tradeNo)
		if err != nil {
			global.Logger.Errorf("svc.FinishRecharge err: %v\n", err)
			response.ToErrorResponse(errcode.RechargeNotifyError)
			return
		}
	}
	response.Ctx.String(http.StatusOK, "success")
}

func GetUserWalletBills(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	svc := service.New(c)
	bills, totalRows, err := svc.GetUserWalletBills(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		global.Logger.Errorf("svc.GetUserWalletBills err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(bills, totalRows)
}
