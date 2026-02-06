// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paopao",
	Short: `an artistic "twitter like" community`,
	Long:  `an artistic "twitter like" community`,
}

// Setup set root command name,short-describe, long-describe
// return &cobra.Command to custom other options
func Setup(use, short, long string) *cobra.Command {
	rootCmd.Use = use
	rootCmd.Short = short
	rootCmd.Long = long
	return rootCmd
}

// Register add sub-command
func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
func Execute() {
	rootCmd.Execute()
}
