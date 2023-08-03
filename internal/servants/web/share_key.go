package web

import (
	"fmt"
	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/chain"
	"github.com/sirupsen/logrus"
)

var (
	_ api.KeyQuery = (*shareKey)(nil)
)

type shareKey struct {
	api.UnimplementedShareKeyServant
	*base.DaoServant
}

func (s *shareKey) Chain() gin.HandlersChain {
	return gin.HandlersChain{chain.JWT()}
}

func (s *shareKey) GetKeyDetail(req *web.GetUserKeysReq) (*web.GetUserKeysResp, mir.Error) {
	//判断GetUserKeys是否为nil
	if s.Ds == nil {
		logrus.Errorf("GetKeyDetail err: %s", web.ErrDsNil)
		return nil, web.ErrDsNil
	} else {
		logrus.Info("GetKeyDetail success")
	}
	if req.UserName == "" {
		logrus.Errorf("GetKeyDetail err: %s", web.ErrUserNameEmpty)
		return nil, web.ErrUserNameEmpty
	}
	// 调用数据源的方法查询用户的所有key信息
	keys, err := s.Ds.GetUserKeys(req.UserName)
	if err != nil {
		logrus.Errorf("GetUserKeys err: %s", err)
		return nil, web.ErrGetUserKeysFailed
	}

	// 将查询到的key信息转换为KeyInfo结构体
	var keyInfos []web.KeyInfo
	for _, key := range keys {
		keyInfos = append(keyInfos, web.KeyInfo{
			ShareKey:    key.ShareKey,
			Name:        key.Name,
			Description: key.Description,
		})
	}

	// 构建返回结果并返回
	resp := &web.GetUserKeysResp{
		ShareKeys: keyInfos,
	}
	return resp, nil
}

func (s *shareKey) DeleteKeyDetail(req *web.DeleteKeyReq) (*web.DeleteKeyResp, mir.Error) {
	//判断GetUserKeys是否为nil
	if s.Ds == nil {
		logrus.Errorf("GetKeyDetail err: %s", web.ErrDsNil)
		return nil, web.ErrDsNil
	} else {
		logrus.Info("GetKeyDetail success")
	}
	if req.UserName == "" {
		logrus.Errorf("GetKeyDetail err: %s", web.ErrUserNameEmpty)
		return nil, web.ErrUserNameEmpty
	}

	fmt.Print(req.UserName, req.ShareKey)
	sof, err := s.Ds.DeleteUserKey(req.UserName, req.ShareKey)
	if err != nil {
		logrus.Errorf("DeleteUserKey err: %s", err)
		return nil, web.ErrDeleteUserKeyFailed
	}
	if sof.SuccessOrFail == false {
		logrus.Errorf("DeleteUserKey err: %s", err)
		return nil, web.ErrDeleteUserKeyFailed
	}

	// 构建返回结果并返回
	resp := &web.DeleteKeyResp{
		Status: "DELETE_SUCCESS",
	}
	return resp, nil
}

// NewShareKeyServant 创建share_key服务
func NewShareKeyServant(s *base.DaoServant) api.KeyQuery {
	return &shareKey{
		DaoServant: s,
	}
}
