// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/Masterminds/semver/v3"
)

// VersionInfo 版本信息
type VersionInfo interface {
	Name() string
	Version() *semver.Version
}
