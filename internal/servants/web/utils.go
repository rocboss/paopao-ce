// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"crypto/subtle"
	"image"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofrs/uuid/v5"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

var defaultAvatars = []string{
	"https://paopao-demo.vercel.app/avatar/default/zoe.png",
	"https://paopao-demo.vercel.app/avatar/default/william.png",
	"https://paopao-demo.vercel.app/avatar/default/walter.png",
	"https://paopao-demo.vercel.app/avatar/default/thomas.png",
	"https://paopao-demo.vercel.app/avatar/default/taylor.png",
	"https://paopao-demo.vercel.app/avatar/default/sophia.png",
	"https://paopao-demo.vercel.app/avatar/default/sam.png",
	"https://paopao-demo.vercel.app/avatar/default/ryan.png",
	"https://paopao-demo.vercel.app/avatar/default/ruby.png",
	"https://paopao-demo.vercel.app/avatar/default/quinn.png",
	"https://paopao-demo.vercel.app/avatar/default/paul.png",
	"https://paopao-demo.vercel.app/avatar/default/owen.png",
	"https://paopao-demo.vercel.app/avatar/default/olivia.png",
	"https://paopao-demo.vercel.app/avatar/default/norman.png",
	"https://paopao-demo.vercel.app/avatar/default/nora.png",
	"https://paopao-demo.vercel.app/avatar/default/natalie.png",
	"https://paopao-demo.vercel.app/avatar/default/naomi.png",
	"https://paopao-demo.vercel.app/avatar/default/miley.png",
	"https://paopao-demo.vercel.app/avatar/default/mike.png",
	"https://paopao-demo.vercel.app/avatar/default/lucas.png",
	"https://paopao-demo.vercel.app/avatar/default/kylie.png",
	"https://paopao-demo.vercel.app/avatar/default/julia.png",
	"https://paopao-demo.vercel.app/avatar/default/joshua.png",
	"https://paopao-demo.vercel.app/avatar/default/john.png",
	"https://paopao-demo.vercel.app/avatar/default/jane.png",
	"https://paopao-demo.vercel.app/avatar/default/jackson.png",
	"https://paopao-demo.vercel.app/avatar/default/ivy.png",
	"https://paopao-demo.vercel.app/avatar/default/isaac.png",
	"https://paopao-demo.vercel.app/avatar/default/henry.png",
	"https://paopao-demo.vercel.app/avatar/default/harry.png",
	"https://paopao-demo.vercel.app/avatar/default/harold.png",
	"https://paopao-demo.vercel.app/avatar/default/hanna.png",
	"https://paopao-demo.vercel.app/avatar/default/grace.png",
	"https://paopao-demo.vercel.app/avatar/default/george.png",
	"https://paopao-demo.vercel.app/avatar/default/freddy.png",
	"https://paopao-demo.vercel.app/avatar/default/frank.png",
	"https://paopao-demo.vercel.app/avatar/default/finn.png",
	"https://paopao-demo.vercel.app/avatar/default/emma.png",
	"https://paopao-demo.vercel.app/avatar/default/emily.png",
	"https://paopao-demo.vercel.app/avatar/default/edward.png",
	"https://paopao-demo.vercel.app/avatar/default/clara.png",
	"https://paopao-demo.vercel.app/avatar/default/claire.png",
	"https://paopao-demo.vercel.app/avatar/default/chloe.png",
	"https://paopao-demo.vercel.app/avatar/default/audrey.png",
	"https://paopao-demo.vercel.app/avatar/default/arthur.png",
	"https://paopao-demo.vercel.app/avatar/default/anna.png",
	"https://paopao-demo.vercel.app/avatar/default/andy.png",
	"https://paopao-demo.vercel.app/avatar/default/alfred.png",
	"https://paopao-demo.vercel.app/avatar/default/alexa.png",
	"https://paopao-demo.vercel.app/avatar/default/abigail.png",
}

func getRandomAvatar() string {
	rand.Seed(time.Now().UnixMicro())
	return defaultAvatars[rand.Intn(len(defaultAvatars))]
}

// checkPassword 密码检查
func checkPassword(password string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(password) < 6 || utf8.RuneCountInString(password) > 16 {
		return web.ErrPasswordLengthLimit
	}
	return nil
}

// ValidPassword 检查密码是否一致
func validPassword(secret, password, salt string) bool {
	expected := utils.EncodeMD5(utils.EncodeMD5(password) + salt)

	// Check that the secret matches the expected value.
	// Use constant time comparison to avoid timing attacks.
	return subtle.ConstantTimeCompare([]byte(secret), []byte(expected)) == 1
}

// encryptPasswordAndSalt 密码加密&生成salt
func encryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	secret := utils.EncodeMD5(utils.EncodeMD5(password) + salt)
	return secret, salt
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
		case ms.ContentTypeImage,
			ms.ContentTypeVideo,
			ms.ContentTypeAudio,
			ms.ContentTypeAttachment,
			ms.ContentTypeChargeAttachment:
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

func fileCheck(uploadType string, size int64) error {
	if uploadType != "public/video" &&
		uploadType != "public/image" &&
		uploadType != "public/avatar" &&
		uploadType != "attachment" {
		return xerror.InvalidParams
	}
	if size > 1024*1024*100 {
		return web.ErrFileInvalidSize.WithDetails("最大允许100MB")
	}
	return nil
}

func getFileExt(s string) (string, error) {
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
		return "", web.ErrFileInvalidExt.WithDetails("仅允许 png/jpg/gif/mp4/mov/zip 类型")
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
func checkPermision(user *ms.User, targetUserId int64) error {
	if user == nil || (user.ID != targetUserId && !user.IsAdmin) {
		return web.ErrNoPermission
	}
	return nil
}

// checkPostViewPermission 检查当前用户是否可读指定post
func checkPostViewPermission(user *ms.User, post *ms.Post, ds core.DataService) error {
	if post.Visibility == core.PostVisitPublic {
		return nil
	}

	if user == nil {
		return web.ErrNoPermission
	}

	if user.IsAdmin || user.ID == post.UserID {
		return nil
	}

	if post.Visibility == core.PostVisitPrivate {
		return web.ErrNoPermission
	}

	if post.Visibility == core.PostVisitFriend {
		if !ds.IsFriend(post.UserID, user.ID) && !ds.IsFriend(user.ID, post.UserID) {
			return web.ErrNoPermission
		}
	}
	// TODO: add following check logic
	return nil
}
