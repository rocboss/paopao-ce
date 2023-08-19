// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	Register(versionCmd)
}

func versionRun(_cmd *cobra.Command, _args []string) {
	utils.PrintHelloBanner(version.VersionInfo())
}
