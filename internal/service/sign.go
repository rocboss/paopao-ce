package service

import (
	"fmt"
	"sort"

	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/pkg/util"
)

func (svc *Service) GetParamSign(param map[string]interface{}, secretKey string) string {
	signRaw := ""

	rawStrs := []string{}
	for k, v := range param {
		if k != "sign" {
			rawStrs = append(rawStrs, k+"="+fmt.Sprintf("%v", v))
		}
	}

	sort.Strings(rawStrs)
	for _, v := range rawStrs {
		signRaw += v
	}

	if global.ServerSetting.RunMode == "debug" {
		global.Logger.Info(map[string]string{
			"signRaw": signRaw,
			"sysSign": util.EncodeMD5(signRaw + secretKey),
		})
	}

	return util.EncodeMD5(signRaw + secretKey)
}
