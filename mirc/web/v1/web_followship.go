package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(WebFollowship))
}

type WebFollowship struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`
}
