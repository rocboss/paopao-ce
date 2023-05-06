// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"image"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/alimy/mir/v3"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var defaultAvatars = []string{
	"https://assets.paopao.info/public/avatar/default/zoe.png",
	"https://assets.paopao.info/public/avatar/default/william.png",
	"https://assets.paopao.info/public/avatar/default/walter.png",
	"https://assets.paopao.info/public/avatar/default/thomas.png",
	"https://assets.paopao.info/public/avatar/default/taylor.png",
	"https://assets.paopao.info/public/avatar/default/sophia.png",
	"https://assets.paopao.info/public/avatar/default/sam.png",
	"https://assets.paopao.info/public/avatar/default/ryan.png",
	"https://assets.paopao.info/public/avatar/default/ruby.png",
	"https://assets.paopao.info/public/avatar/default/quinn.png",
	"https://assets.paopao.info/public/avatar/default/paul.png",
	"https://assets.paopao.info/public/avatar/default/owen.png",
	"https://assets.paopao.info/public/avatar/default/olivia.png",
	"https://assets.paopao.info/public/avatar/default/norman.png",
	"https://assets.paopao.info/public/avatar/default/nora.png",
	"https://assets.paopao.info/public/avatar/default/natalie.png",
	"https://assets.paopao.info/public/avatar/default/naomi.png",
	"https://assets.paopao.info/public/avatar/default/miley.png",
	"https://assets.paopao.info/public/avatar/default/mike.png",
	"https://assets.paopao.info/public/avatar/default/lucas.png",
	"https://assets.paopao.info/public/avatar/default/kylie.png",
	"https://assets.paopao.info/public/avatar/default/julia.png",
	"https://assets.paopao.info/public/avatar/default/joshua.png",
	"https://assets.paopao.info/public/avatar/default/john.png",
	"https://assets.paopao.info/public/avatar/default/jane.png",
	"https://assets.paopao.info/public/avatar/default/jackson.png",
	"https://assets.paopao.info/public/avatar/default/ivy.png",
	"https://assets.paopao.info/public/avatar/default/isaac.png",
	"https://assets.paopao.info/public/avatar/default/henry.png",
	"https://assets.paopao.info/public/avatar/default/harry.png",
	"https://assets.paopao.info/public/avatar/default/harold.png",
	"https://assets.paopao.info/public/avatar/default/hanna.png",
	"https://assets.paopao.info/public/avatar/default/grace.png",
	"https://assets.paopao.info/public/avatar/default/george.png",
	"https://assets.paopao.info/public/avatar/default/freddy.png",
	"https://assets.paopao.info/public/avatar/default/frank.png",
	"https://assets.paopao.info/public/avatar/default/finn.png",
	"https://assets.paopao.info/public/avatar/default/emma.png",
	"https://assets.paopao.info/public/avatar/default/emily.png",
	"https://assets.paopao.info/public/avatar/default/edward.png",
	"https://assets.paopao.info/public/avatar/default/clara.png",
	"https://assets.paopao.info/public/avatar/default/claire.png",
	"https://assets.paopao.info/public/avatar/default/chloe.png",
	"https://assets.paopao.info/public/avatar/default/audrey.png",
	"https://assets.paopao.info/public/avatar/default/arthur.png",
	"https://assets.paopao.info/public/avatar/default/anna.png",
	"https://assets.paopao.info/public/avatar/default/andy.png",
	"https://assets.paopao.info/public/avatar/default/alfred.png",
	"https://assets.paopao.info/public/avatar/default/alexa.png",
	"https://assets.paopao.info/public/avatar/default/abigail.png",
}

func getRandomAvatar() string {
	rand.Seed(time.Now().UnixMicro())
	return defaultAvatars[rand.Intn(len(defaultAvatars))]
}

// checkPassword 密码检查
func checkPassword(password string) mir.Error {
	// 检测用户是否合规
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 16 {
		return _errPasswordLengthLimit
	}
	return nil
}

// ValidPassword 检查密码是否一致
func validPassword(dbPassword, password, salt string) bool {
	return strings.Compare(dbPassword, utils.EncodeMD5(utils.EncodeMD5(password)+salt)) == 0
}

// encryptPasswordAndSalt 密码加密&生成salt
func encryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = utils.EncodeMD5(utils.EncodeMD5(password) + salt)
	return password, salt
}

// deleteOssObjects 删除推文的媒体内容, 宽松处理错误(就是不处理), 后续完善
func deleteOssObjects(oss core.ObjectStorageService, mediaContents []string) {
	mediaContentsSize := len(mediaContents)
	if mediaContentsSize > 1 {
		objectKeys := make([]string, 0, mediaContentsSize)
		for _, cUrl := range mediaContents {
			objectKeys = append(objectKeys, oss.ObjectKey(cUrl))
		}
		// TODO: 优化处理尽量使用channel传递objectKeys使用可控数量的Goroutine集中处理object删除动作，后续完善
		go oss.DeleteObjects(objectKeys)
	} else if mediaContentsSize == 1 {
		oss.DeleteObject(oss.ObjectKey(mediaContents[0]))
	}
}

// persistMediaContents 获取媒体内容并持久化
func persistMediaContents(oss core.ObjectStorageService, contents []*web.PostContentItem) (items []string, err error) {
	items = make([]string, 0, len(contents))
	for _, item := range contents {
		switch item.Type {
		case core.ContentTypeImage,
			core.ContentTypeVideo,
			core.ContentTypeAudio,
			core.ContentTypeAttachment,
			core.ContentTypeChargeAttachment:
			items = append(items, item.Content)
			if err != nil {
				continue
			}
			if err = oss.PersistObject(oss.ObjectKey(item.Content)); err != nil {
				logrus.Errorf("service.persistMediaContents failed: %s", err)
			}
		}
	}
	return
}

func fileCheck(uploadType string, size int64) mir.Error {
	if uploadType != "public/video" &&
		uploadType != "public/image" &&
		uploadType != "public/avatar" &&
		uploadType != "attachment" {
		return xerror.InvalidParams
	}
	if size > 1024*1024*100 {
		return _errFileInvalidSize.WithDetails("最大允许100MB")
	}
	return nil
}

func getFileExt(s string) (string, mir.Error) {
	switch s {
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
		return "", _errFileInvalidExt.WithDetails("仅允许 png/jpg/gif/mp4/mov/zip 类型")
	}
}

func generatePath(s string) string {
	n := len(s)
	if n <= 2 {
		return s
	}
	return generatePath(s[:n-2]) + "/" + s[n-2:]
}

func getImageSize(img image.Rectangle) (int, int) {
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	return width, height
}

func tagsFrom(originTags []string) []string {
	tags := make([]string, 0, len(originTags))
	for _, tag := range originTags {
		// TODO: 优化tag有效性检测
		if tag = strings.TrimSpace(tag); len(tag) > 0 {
			tags = append(tags, tag)
		}
	}
	return tags
}

// checkPermision 检查是否拥有者或管理员
func checkPermision(user *core.User, targetUserId int64) mir.Error {
	if user == nil || (user.ID != targetUserId && !user.IsAdmin) {
		return _errNoPermission
	}
	return nil
}
