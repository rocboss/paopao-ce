// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cfg

import (
	"testing"
)

func TestUseDefault(t *testing.T) {
	suites := map[string][]string{
		"default": {"Sms", "Alipay", "Zinc", "MySQL", "Redis", "AliOSS", "LogZinc"},
		"develop": {"Zinc", "MySQL", "AliOSS", "LogFile"},
		"slim":    {"Zinc", "MySQL", "Redis", "AliOSS", "LogFile"},
	}
	kv := map[string]string{
		"sms": "SmsJuhe",
	}
	f := NewFeatures(suites, kv)
	for _, data := range []struct {
		key    string
		expect string
		exist  bool
	}{
		{"Sms", "SmsJuhe", true},
		{"Alipay", "", true},
		{"Zinc", "", true},
		{"Redis", "", true},
		{"Database", "", false},
	} {
		if v, ok := f.Cfg(data.key); ok != data.exist || v != data.expect {
			t.Errorf("key: %s expect: %s exist: %t got v: %s ok: %t", data.key, data.expect, data.exist, v, ok)
		}
	}
	for exp, res := range map[string]bool{
		"Sms":           true,
		"Sms = SmsJuhe": true,
		"SmsJuhe":       false,
		"default":       true,
	} {
		if ok := f.CfgIf(exp); res != ok {
			t.Errorf("CfgIf(%s) want %t got %t", exp, res, ok)
		}
	}
}

func TestUse(t *testing.T) {
	suites := map[string][]string{
		"default": {"Sms", "Alipay", "Zinc", "MySQL", "Redis", "AliOSS", "LogZinc"},
		"develop": {"Zinc", "MySQL", "AliOSS", "LogFile"},
		"slim":    {"Zinc", "MySQL", "Redis", "AliOSS", "LogFile"},
	}
	kv := map[string]string{
		"sms": "SmsJuhe",
	}
	f := NewFeatures(suites, kv)

	f.Use([]string{"develop"}, true)
	for _, data := range []struct {
		key    string
		expect string
		exist  bool
	}{
		{"Sms", "", false},
		{"Alipay", "", false},
		{"Zinc", "", true},
		{"Redis", "", false},
		{"Database", "", false},
	} {
		if v, ok := f.Cfg(data.key); ok != data.exist || v != data.expect {
			t.Errorf("key: %s expect: %s exist: %t got v: %s ok: %t", data.key, data.expect, data.exist, v, ok)
		}
	}
	for exp, res := range map[string]bool{
		"Sms":           false,
		"Sms = SmsJuhe": false,
		"SmsJuhe":       false,
		"default":       false,
		"develop":       true,
	} {
		if ok := f.CfgIf(exp); res != ok {
			t.Errorf("CfgIf(%s) want %t got %t", exp, res, ok)
		}
	}

	f.UseDefault()
	f.Use([]string{"slim", "", "demo"}, false)
	for _, data := range []struct {
		key    string
		expect string
		exist  bool
	}{
		{"Sms", "SmsJuhe", true},
		{"Alipay", "", true},
		{"Zinc", "", true},
		{"Redis", "", true},
		{"Database", "", false},
		{"demo", "", true},
	} {
		if v, ok := f.Cfg(data.key); ok != data.exist || v != data.expect {
			t.Errorf("key: %s expect: %s exist: %t got v: %s ok: %t", data.key, data.expect, data.exist, v, ok)
		}
	}
	for exp, res := range map[string]bool{
		"Sms":           true,
		"Sms = SmsJuhe": true,
		"SmsJuhe":       false,
		"default":       true,
		"develop":       false,
		"slim":          true,
		"demo":          true,
	} {
		if ok := f.CfgIf(exp); res != ok {
			t.Errorf("CfgIf(%s) want %t got %t", exp, res, ok)
		}
	}
}
