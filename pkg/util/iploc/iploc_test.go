package iploc

import (
	"testing"
)

func TestFind(t *testing.T) {
	for _, data := range []struct {
		ip      string
		country string
		city    string
	}{
		{ip: "127.0.0.1", country: "本机地址", city: " CZ88.NET"},
		{ip: "180.89.94.9", country: "北京市", city: "鹏博士宽带"},
	} {
		country, city := Find(data.ip)
		t.Logf("ip:%v, country:%v, city:%v", data.ip, country, city)
		if country != data.country {
			t.Errorf("find ip:%s expect country: %s got: %s", data.ip, data.country, country)
		}
		if city != data.city {
			t.Errorf("find ip:%s expect city: %s got: %s", data.ip, data.city, city)
		}
	}
}
