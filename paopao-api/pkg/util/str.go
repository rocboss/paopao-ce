package util

import (
	"math/rand"
	"time"
)

type StrType int

const (
	NUM   StrType = iota // 数字
	LOWER                // 小写字母
	UPPER                // 大写字母
	ALL                  // 全部
	CLEAR                // 去除部分易混淆的字符
)

var fontKinds = [][]int{{10, 48}, {26, 97}, {26, 65}}
var letters = []byte("34578acdefghjkmnpqstwxyABCDEFGHJKMNPQRSVWXY")

// 生成随机字符串
// size 个数 kind 模式
func RandStr(size int, kind StrType) []byte {
	ikind, result := kind, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = StrType(rand.Intn(int(ALL)))
		}
		scope, base := fontKinds[ikind][0], fontKinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
		// 不易混淆字符模式：重新生成字符
		if kind == 4 {
			result[i] = letters[rand.Intn(len(letters))]
		}
	}
	return result
}
