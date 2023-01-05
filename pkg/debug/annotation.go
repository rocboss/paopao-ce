// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package debug

import (
	"github.com/sirupsen/logrus"
)

func TODO() {
	logrus.Fatalln("in todo progress")
}

func NotImplemented() {
	logrus.Fatalln("not implemented")
}
