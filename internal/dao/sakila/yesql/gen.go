// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"
	"strings"

	"github.com/alimy/yesql"
)

//go:generate go run $GOFILE
func main() {
	log.Println("[Yesql] generate code start")
	yesql.SetDefaultQueryHook(func(query *yesql.Query) (*yesql.Query, error) {
		query.Query = strings.TrimRight(query.Query, ";")
		return query, nil
	})
	if err := yesql.Generate("yesql.sql", "cc", "cc", yesql.SqlxPkgName("github.com/bitbus/sqlx")); err != nil {
		log.Fatalf("generate code occurs error: %s", err)
	}
	log.Println("[Yesql] generate code finish")
}
