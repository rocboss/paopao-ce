package core

import (
	"github.com/Masterminds/semver/v3"
)

type VersionInfo interface {
	Name() string
	Version() *semver.Version
}
