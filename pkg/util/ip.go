package util

import "github.com/rocboss/paopao-ce/pkg/util/iploc"

func GetIPLoc(ip string) string {
	country, _ := iploc.Find(ip)
	return country
}
