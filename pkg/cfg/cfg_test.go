package cfg

import (
	"testing"
)

func TestCfg(t *testing.T) {
	suites := map[string][]string{
		"default": {"Sms", "Alipay", "Zinc", "MySQL", "Redis", "AliOSS", "LogZinc"},
		"develop": {"Zinc", "MySQL", "AliOSS", "LogFile"},
		"slim":    {"Zinc", "MySQL", "Redis", "AliOSS", "LogFile"},
	}
	kv := map[string]string{
		"sms": "SmsJuhe",
	}

	Initialize(suites, kv)
	UseDefault()

	if !If("Sms") {
		t.Error(`want If("Sms") == true but not`)
	}

	if v, exist := As("Sms"); exist && v != "SmsJuhe" {
		t.Errorf(`want As("Sms") == "SmsJuhe", true but got: "%s", "%t"`, v, exist)
	}

	matched := false
	Be("Alipay", func() {
		matched = true
	})
	if !matched {
		t.Error(`want Be("Alipay", ...) matched but not`)
	}

	matched = false
	Not("LogFile", func() {
		matched = true
	})
	if !matched {
		t.Error(`want Not("LogFile", ...) matched but not`)
	}

	var m1, m2, m3, m4 bool
	In(Actions{
		"Sms": func() {
			m1 = true
		},
		"Alipay": func() {
			m2 = true
		},
		"Meili": func() {
			m4 = true
		},
	}, func() {
		m3 = true
	})
	if !m1 || !m2 || m3 || m4 {
		t.Errorf(`In("Sms", "Alipay", "Meili", ...) not correct -> m1: %t m2:%t m3:%t m4:%t`, m1, m2, m3, m4)
	}

	m1 = false
	m2 = false
	m3 = false
	In(Actions{
		"LogFile": func() {
			m1 = true
		},
		"Meili": func() {
			m2 = true
		},
	}, func() {
		m3 = true
	})
	if m1 || m2 || !m3 {
		t.Errorf(`In("Zinc", "MySQL", ...) not correct -> m1: %t m2:%t m3:%t`, m1, m2, m3)
	}
}
