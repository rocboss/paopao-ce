package dao

import (
	"github.com/rocboss/paopao-ce/internal/conf"
)

func getOssDomain() string {
	uri := "https://"
	if conf.CfgIf("AliOSS") {
		return uri + conf.AliOSSSetting.Domain + "/"
	} else if conf.CfgIf("MinIO") {
		if !conf.MinIOSetting.Secure {
			uri = "http://"
		}
		return uri + conf.MinIOSetting.Domain + "/" + conf.MinIOSetting.Bucket + "/"
	} else if conf.CfgIf("S3") {
		if !conf.S3Setting.Secure {
			uri = "http://"
		}
		// TODO: will not work well need test in real world
		return uri + conf.S3Setting.Domain + "/" + conf.S3Setting.Bucket + "/"
	} else if conf.CfgIf("LocalOSS") {
		if !conf.LocalOSSSetting.Secure {
			uri = "http://"
		}
		return uri + conf.LocalOSSSetting.Domain + "/oss/" + conf.LocalOSSSetting.Bucket + "/"
	}
	return uri + conf.AliOSSSetting.Domain + "/"
}
