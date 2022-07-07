package core

import (
	"github.com/Masterminds/semver/v3"
)

// VersionInfo 版本信息
type VersionInfo interface {
	Name() string
	Version() *semver.Version
}
