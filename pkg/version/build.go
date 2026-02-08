// Copyright 2025 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package version

import (
	"debug/buildinfo"
	"os"
	"time"
)

var BuildTime = time.Now()

func init() {
	exe, err := os.Executable()
	if err != nil {
		return
	}
	info, err := buildinfo.ReadFile(exe)
	if err != nil {
		return
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.time" && s.Value != "" {
			if t, err := time.Parse(time.RFC3339, s.Value); err == nil {
				BuildTime = t
			}
			break
		}
	}
}
