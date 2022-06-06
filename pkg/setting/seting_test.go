package setting

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
	features := newFeatures(suites, kv)
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
		if v, ok := features.Cfg(data.key); ok != data.exist || v != data.expect {
			t.Errorf("key: %s expect: %s exist: %t got v: %s ok: %t", data.key, data.expect, data.exist, v, ok)
		}
	}
	for exp, res := range map[string]bool{
		"Sms":           true,
		"Sms = SmsJuhe": true,
		"SmsJuhe":       false,
	} {
		if ok := features.CfgIf(exp); res != ok {
			t.Errorf("CfgIf(%s) want %t got %t", exp, res, ok)
		}
	}
}
