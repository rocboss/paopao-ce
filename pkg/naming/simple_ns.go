// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package naming

import (
	"strings"
	"unicode"
)

type simpleNamingStrategy struct {
	// just empty
}

func (s *simpleNamingStrategy) Naming(name string) string {
	b := &strings.Builder{}
	notFirst := false
	b.Grow(len(name) + 3)
	for _, r := range name {
		if unicode.IsUpper(r) {
			if notFirst {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
		notFirst = true
	}
	return b.String()
}

// NewSimpleNamingStrategy return simple naming strategy instance
func NewSimpleNamingStrategy() NamingStrategy {
	return &simpleNamingStrategy{}
}
