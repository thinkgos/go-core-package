// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package regex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmail(t *testing.T) {
	emails := map[string]bool{
		`test@example.com`:             true,
		`single-character@b.org`:       true,
		`uncommon_address@test.museum`: true,
		`local@sld.UPPER`:              true,
		`@missing.org`:                 false,
		`missing@.com`:                 false,
		`missing@qq.`:                  false,
		`wrong-ip@127.1.1.1.26`:        false,
	}
	for e, r := range emails {
		require.Equal(t, r, IsEmail(e))
	}
}

func TestIsEmailRFC(t *testing.T) {
	require.True(t, IsEmailRFC("test@example.com"))
}

func TestIsUrl(t *testing.T) {
	urls := map[string]bool{
		"http://www.example.com":                     true,
		"http://example.com":                         true,
		"http://example.com?user=test&password=test": true,
		"http://example.com?user=test#login":         true,
		"ftp://example.com":                          true,
		"https://example.com":                        true,
		"htp://example.com":                          false,
		"http//example.com":                          false,
		"http://example":                             true,
	}
	for u, r := range urls {
		require.Equal(t, r, IsURL(u))
	}
}

func BenchmarkIsEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("test@example.com")
	}
}

func BenchmarkIsUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEmail("http://example.com")
	}
}

func TestIsRGBColor(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"", "", false},
		{"", "rgb(0,31,255)", true},
		{"", "rgb(1,349,275)", false},
		{"", "rgb(01,31,255)", false},
		{"", "rgb(0.6,31,255)", false},
		{"", "rgba(0,31,255)", false},
		{"", "rgb(0,  31, 255)", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRGBColor(tt.s); got != tt.want {
				t.Errorf("IsRGBColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		want  string
		want1 bool
	}{
		{"", "", "", false},
		{"", "#ff", "#ff", false},
		{"", "fff0", "fff0", false},
		{"", "#ff12FG", "#ff12FG", false},
		{"", "CCccCC", "#CCCCCC", true},
		{"", "fff", "#FFF", true},
		{"", "#f00", "#F00", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsHexColor(tt.s)
			if got != tt.want {
				t.Errorf("IsHexColor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsHexColor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
