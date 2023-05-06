// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

func PrintHelloBanner(text string) {

	fmt.Println(" ____   __    __  ____   __    __  ")
	fmt.Println("(  _ \\ / _\\  /  \\(  _ \\ / _\\  /  \\ ")
	fmt.Println(" ) __//    \\(  O )) __//    \\(  O )")
	fmt.Println("(__)  \\_/\\_/ \\__/(__)  \\_/\\_/ \\__/ ")
	fmt.Println(text)
}

func SidStr(name string, version *semver.Version, size int) string {
	return fmt.Sprintf(fmt.Sprintf("%%s@%%-%ds", size-len(name+version.String())+4), name, version)
}
