// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package migrate

import (
	"fmt"

	"github.com/rocboss/paopao-ce/cmd"
	"github.com/spf13/cobra"
)

func init() {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database data",
		Long:  "miegrate database data when paopao-ce upgrade",
		Run:   migrateRun,
	}
	cmd.Register(migrateCmd)
}

func migrateRun(_cmd *cobra.Command, _args []string) {
	// TODO: add some logic for migrate cmd feature
	fmt.Println("sorry, this feature is not implemented yet.")
}
