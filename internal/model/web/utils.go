// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/mir/v5"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

func fileCheck(uploadType string, size int64) mir.Error {
	if uploadType != "public/video" &&
		uploadType != "public/image" &&
		uploadType != "public/avatar" &&
		uploadType != "attachment" {
		return xerror.InvalidParams
	}
	if size > 1024*1024*100 {
		return ErrFileInvalidSize.WithDetails("最大允许100MB")
	}
	return nil
}

func getFileExt(s string) (string, mir.Error) {
	switch s {
	case "image/webp":
		return ".webp", nil
	case "image/png":
		return ".png", nil
	case "image/jpg":
		return ".jpg", nil
	case "image/jpeg":
		return ".jpeg", nil
	case "image/gif":
		return ".gif", nil
	case "video/mp4":
		return ".mp4", nil
	case "video/quicktime":
		return ".mov", nil
	case "application/zip",
		"application/x-zip",
		"application/octet-stream",
		"application/x-zip-compressed":
		return ".zip", nil
	default:
		return "", ErrFileInvalidExt.WithDetails("仅允许 webp/png/jpg/gif/mp4/mov/zip 类型")
	}
}
