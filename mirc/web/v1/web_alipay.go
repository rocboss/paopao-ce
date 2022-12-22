package v1

import (
	. "github.com/alimy/mir/v3"
	. "github.com/alimy/mir/v3/engine"
)

func init() {
	AddEntry(new(WebAlipay))
}

type WebAlipay struct {
	Chain Chain `mir:"-"`
	Group Group `mir:"v1"`
}
