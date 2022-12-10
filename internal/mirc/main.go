// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	. "github.com/alimy/mir/v3/core"
	. "github.com/alimy/mir/v3/engine"

	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/m/v1"
	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/r/v1"
	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/s/v1"
	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/v1"
	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/x/v1"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialDebugMode),
		GeneratorName(GeneratorGin),
		SinkPath("auto"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
