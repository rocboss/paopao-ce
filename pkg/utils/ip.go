// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import "github.com/rocboss/paopao-ce/pkg/utils/iploc"

func GetIPLoc(ip string) string {
	country, _ := iploc.Find(ip)
	return country
}
