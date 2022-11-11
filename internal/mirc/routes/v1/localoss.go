package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(LocalOSS))
}

type LocalOSS struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`
}
