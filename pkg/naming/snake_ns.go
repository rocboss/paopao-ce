// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package naming

import (
	"bytes"
	"strings"
)

// snakeNamingStrategy snake naming strategy implement
// This implement is inspiration from https://github.com/jingzhu/gorm's naming logic.
type snakeNamingStrategy struct {
	initialismsReplacer *strings.Replacer
}

func (s *snakeNamingStrategy) Naming(name string) string {
	if name == "" {
		return ""
	}

	var lastCase, currCase, nextCase, nextNumber bool
	value := s.initialismsReplacer.Replace(name)
	buf := bytes.NewBufferString("")
	lower, upper := false, true
	for i, v := range value[:len(value)-1] {
		nextCase = value[i+1] >= 'A' && value[i+1] <= 'Z'
		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'
		if i > 0 {
			if currCase == upper {
				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase, currCase = currCase, nextCase
	}

	buf.WriteByte(value[len(value)-1])
	res := strings.ToLower(buf.String())
	return res
}

// NewSnakeNamingStrategy return snake naming strategy instance
func NewSnakeNamingStrategy() NamingStrategy {
	// Copied from golint
	initialisms := []string{
		"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP",
		"HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA",
		"SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8",
		"VM", "XML", "XSRF", "XSS",
	}

	var oldnews []string
	for _, initialism := range initialisms {
		oldnews = append(oldnews, initialism, strings.Title(strings.ToLower(initialism)))
	}
	replacer := strings.NewReplacer(oldnews...)

	return &snakeNamingStrategy{
		initialismsReplacer: replacer,
	}
}
