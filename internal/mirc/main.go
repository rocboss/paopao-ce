package main

import (
	"log"

	. "github.com/alimy/mir/v3/core"
	. "github.com/alimy/mir/v3/engine"

	_ "github.com/rocboss/paopao-ce/internal/mirc/routes/v1"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialMode),
		GeneratorName(GeneratorGin),
		SinkPath("auto"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
