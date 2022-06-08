package dao

import (
	"github.com/rocboss/paopao-ce/global"
)

func getOssDomain() string {
	uri := "https://"
	if global.CfgIf("AliOSS") {
		return uri + global.AliOSSSetting.Domain + "/"
	} else if global.CfgIf("MinIO") {
		if !global.MinIOSetting.Secure {
			uri = "http://"
		}
		return uri + global.MinIOSetting.Domain + "/" + global.MinIOSetting.Bucket + "/"
	} else if global.CfgIf("S3") {
		if !global.S3Setting.Secure {
			uri = "http://"
		}
		// TODO: will not work well need test in real world
		return uri + global.S3Setting.Domain + "/" + global.S3Setting.Bucket + "/"
	} else if global.CfgIf("LocalOSS") {
		if !global.LocalOSSSetting.Secure {
			uri = "http://"
		}
		return uri + global.LocalOSSSetting.Domain + "/oss/" + global.LocalOSSSetting.Bucket + "/"
	}
	return uri + global.AliOSSSetting.Domain + "/"
}
