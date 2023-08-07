// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"

	. "github.com/alimy/mir/v4/core"
	. "github.com/alimy/mir/v4/engine"
	"github.com/gin-gonic/gin"

	_ "github.com/rocboss/paopao-ce/mirc/admin/v1"
	_ "github.com/rocboss/paopao-ce/mirc/bot/v1"
	_ "github.com/rocboss/paopao-ce/mirc/localoss/v1"
	_ "github.com/rocboss/paopao-ce/mirc/space/v1"
	_ "github.com/rocboss/paopao-ce/mirc/web/v1"
)

//go:generate go run $GOFILE
func main() {
	log.Println("[Mir] generate code start")
	opts := Options{
		UseGin(),
		SinkPath("../auto"),
		WatchCtxDone(true),
		RunMode(InSerialMode),
		AssertType[*gin.Context](),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("[Mir] generate code finish")
}
