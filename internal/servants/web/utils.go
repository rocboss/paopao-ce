package web

import (
	"strings"
	"unicode/utf8"

	"github.com/alimy/mir/v3"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/pkg/util"
)

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
	return strings.Compare(dbPassword, util.EncodeMD5(util.EncodeMD5(password)+salt)) == 0
}

// encryptPasswordAndSalt 密码加密&生成salt
func encryptPasswordAndSalt(password string) (string, string) {
	salt := uuid.Must(uuid.NewV4()).String()[:8]
	password = util.EncodeMD5(util.EncodeMD5(password) + salt)
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
