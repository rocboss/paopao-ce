// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
)

var version, commitID, buildDate string

type BuildInfo struct {
	Version   string `json:"version"`
	Sum       string `json:"sum"`
	BuildDate string `json:"build_date"`
}

func VersionInfo() string {
	return fmt.Sprintf("paopao %s (build:%s %s)", version, commitID, buildDate)
}

func ReadBuildInfo() *BuildInfo {
	return &BuildInfo{
		Version:   version,
		Sum:       commitID,
		BuildDate: buildDate,
	}
}
