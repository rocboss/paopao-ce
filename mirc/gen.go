// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"

	. "github.com/alimy/mir/v5/core"
	. "github.com/alimy/mir/v5/engine"
)

//go:generate go run $GOFILE
func main() {
	log.Println("[Mir] generate code start")
	if err := Generate(
		UseGin(),
		Schema("web", "space", "localoss", "bot", "admin"),
		SinkPath("../auto"),
		WatchCtxDone(true),
		RunMode(InSerialMode),
	); err != nil {
		log.Fatal(err)
	}
	log.Println("[Mir] generate code finish")
}
