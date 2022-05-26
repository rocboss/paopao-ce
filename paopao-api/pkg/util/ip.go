package util

import (
	"github.com/yinheli/qqwry"
)

func GetIPLoc(ip string) string {
	q := qqwry.NewQQwry("qqwry.dat")
	q.Find(ip)

	return q.Country
}
