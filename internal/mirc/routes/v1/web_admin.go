package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(WebAdmin))
}

type WebAdmin struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`
}
