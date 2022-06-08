package dao

import (
	"github.com/rocboss/paopao-ce/global"
)

func getOssDomain() string {
	if global.CfgIf("AliOSS") {
		return "https://" + global.AliOSSSetting.Domain + "/"
	} else if global.CfgIf("MinIO") {
		domain := "http://"
		if global.MinIOSetting.Secure {
			domain = "http://"
		}
		return domain + global.MinIOSetting.Endpoint + "/" + global.MinIOSetting.Bucket + "/"
	} else if global.CfgIf("S3") {
		domain := "http://"
		if global.S3Setting.Secure {
			domain = "https://"
		}
		return domain + global.S3Setting.Endpoint + "/" + global.S3Setting.Bucket + "/"
	}
	return ""
}
