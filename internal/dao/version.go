package dao

import (
	"github.com/Masterminds/semver/v3"
)

// versionInfo 版本信息
type versionInfo interface {
	name() string
	version() *semver.Version
}
