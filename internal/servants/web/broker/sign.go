// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package broker

import (
	"fmt"
	"sort"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/util"
	"github.com/sirupsen/logrus"
)

func GetParamSign(param map[string]any, secretKey string) string {
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

	if conf.ServerSetting.RunMode == "debug" {
		logrus.Info(map[string]string{
			"signRaw": signRaw,
			"sysSign": util.EncodeMD5(signRaw + secretKey),
		})
	}

	return util.EncodeMD5(signRaw + secretKey)
}
