// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cfg

import (
	"strings"

	"github.com/rocboss/paopao-ce/pkg/types"
)

// Features fetures info struct
type Features struct {
	kv       map[string]string
	suites   map[string][]string
	features map[string]string
}

// Actions feature-func map alias type
type Actions map[string]types.Fn

// NewFeatures create new Features instance
func NewFeatures(suites map[string][]string, kv map[string]string) *Features {
	f := newEmptyFeatures()
	for k, v := range suites {
		if len(k) > 0 {
			for i := 0; i < len(v); i++ {
				// ignore empty string
				if len(v[i]) == 0 {
					lastIdx := len(v) - 1
					v[i] = v[lastIdx]
					v = v[:lastIdx]
					i--
				}
			}
			if len(v) > 0 {
				f.suites[k] = v
			}
		}
	}
	for k, v := range kv {
		if len(k) > 0 && len(v) > 0 {
			f.kv[k] = v
		}
	}
	f.UseDefault()
	return f
}

func newEmptyFeatures() *Features {
	return &Features{
		suites:   make(map[string][]string),
		kv:       make(map[string]string),
		features: make(map[string]string),
	}
}

// UseDefault use default suite for features
func (f *Features) UseDefault() {
	f.Use([]string{"default"}, true)
}

// Use use custom suite for features
func (f *Features) Use(suite []string, noDefault bool) {
	if noDefault && len(f.features) != 0 {
		f.features = make(map[string]string)
	}
	features := f.flatFeatures(suite)
	for _, feature := range features {
		if len(feature) == 0 {
			continue
		}
		f.features[feature] = f.kv[feature]
	}
}

func (f *Features) flatFeatures(suite []string) []string {
	features := make([]string, 0, len(suite)+10)
	for s := suite[:]; len(s) > 0; s = s[:len(s)-1] {
		item := strings.TrimSpace(strings.ToLower(s[0]))
		if len(item) > 0 {
			if items, exist := f.suites[item]; exist {
				s = append(s, items...)
			}
			features = append(features, item)
		}
		s[0] = s[len(s)-1]
	}
	return features
}

// Cfg get value by key if exist
func (f *Features) Cfg(key string) (string, bool) {
	key = strings.ToLower(key)
	value, exist := f.features[key]
	return value, exist
}

// CfgIf check expression is true. if expression just have a string like
// `Sms` is mean `Sms` whether define in suite feature settings. expression like
// `Sms = SmsJuhe` is mean whether `Sms` define in suite feature settings and value
// is `SmsJuhe`
func (f *Features) CfgIf(expression string) bool {
	kv := strings.Split(expression, "=")
	key := strings.Trim(strings.ToLower(kv[0]), " ")
	v, ok := f.features[key]
	if len(kv) == 2 && ok && strings.Trim(kv[1], " ") == v {
		return true
	} else if len(kv) == 1 && ok {
		return true
	}
	return false
}

// CfgIn range actions to check item's expression is true then do the handle, defFn will handle
// if all items are not matched,
func (f *Features) CfgIn(actions Actions, defAct ...types.Fn) {
	itemMatched := false
	for expression, handle := range actions {
		if f.CfgIf(expression) && handle != nil {
			handle()
			itemMatched = true
		}
	}
	if !itemMatched && len(defAct) > 0 {
		for _, handle := range defAct {
			if handle != nil {
				handle()
			}
		}
	}
}

// CfgBe check expression is true then do the handle. if expression just have a string like
// `Sms` is mean `Sms` whether defined in suite feature settings. expression like
// `Sms = SmsJuhe` is mean whether `Sms` define in suite feature settings and value
// is `SmsJuhe`
func (f *Features) CfgBe(expression string, handle types.Fn) {
	if f.CfgIf(expression) && handle != nil {
		handle()
	}
}

// CfgNot check expression is not true then do the handle. if expression just have a string like
// `Sms` is mean `Sms` whether defined in suite feature settings. expression like
// `Sms = SmsJuhe` is mean whether `Sms` define in suite feature settings and value
// is `SmsJuhe`
func (f *Features) CfgNot(expression string, handle func()) {
	if !f.CfgIf(expression) {
		handle()
	}
}
