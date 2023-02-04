// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package naming

import "testing"

func TestSnakeNamingStrategy_Naming(t *testing.T) {
	ns := NewSnakeNamingStrategy()
	for _, cs := range []struct {
		name     string
		expected string
	}{
		{name: "abc", expected: "abc"},
		{name: "Abc", expected: "abc"},
		{name: "HostName", expected: "host_name"},
		{name: "Host_Name", expected: "host_name"},
		{name: "RESTfulAPI", expected: "res_tful_api"},
		{name: "HTTPS_API", expected: "https_api"},
		{name: "PKG_Name", expected: "pkg_name"},
		{name: "API", expected: "api"},
		{name: "HTTP", expected: "http"},
	} {
		result := ns.Naming(cs.name)
		if result != cs.expected {
			t.Errorf("give:%s expected:%s result:%s", cs.name, cs.expected, result)
		}
	}
}

func TestSimpleNamingStrategy_Naming(t *testing.T) {
	ns := NewSimpleNamingStrategy()
	for _, cs := range []struct {
		name     string
		expected string
	}{
		{name: "abc", expected: "abc"},
		{name: "Abc", expected: "abc"},
		{name: "HostName", expected: "host_name"},
		{name: "Host_Name", expected: "host__name"},
		{name: "RESTfulAPI", expected: "r_e_s_tful_a_p_i"},
		{name: "HTTPS_API", expected: "h_t_t_p_s__a_p_i"},
		{name: "PKG_Name", expected: "p_k_g__name"},
		{name: "API", expected: "a_p_i"},
		{name: "HTTP", expected: "h_t_t_p"},
	} {
		result := ns.Naming(cs.name)
		if result != cs.expected {
			t.Errorf("give:%s expected:%s result:%s", cs.name, cs.expected, result)
		}
	}
}
