package string2

import "testing"

func TestDoubleChar(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"The", "TThhee"},
		{"AAbb", "AAAAbbbb"},
		{"Hi-There", "HHii--TThheerree"},
		{"Word!", "WWoorrdd!!"},
		{"!!", "!!!!"},
		{"", ""},
		{"a", "aa"},
		{".", ".."},
		{"aa", "aaaa"},
	}
	for _, test := range tests {
		if got := DoubleChar(test.s); got != test.want {
			t.Errorf("DoubleChar(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestCountCode(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"aaacodebbb", 1},
		{"codexxcode", 2},
		{"cozexxcope", 2},
		{"cozfxxcope", 1},
		{"xxcozeyycop", 1},
		{"cozcop", 0},
		{"abcxyz", 0},
		{"code", 1},
		{"ode", 0},
		{"c", 0},
		{"", 0},
		{"AAcodeBBcoleCCccoreDD", 3},
		{"AAcodeBBcoleCCccorfDD", 2},
		{"coAcodeBcoleccoreDD", 3},
	}
	for _, test := range tests {
		if got := CountCode(test.s); got != test.want {
			t.Errorf("CountCode(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestCountHi(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"abc hi ho", 1},
		{"ABChi hi", 2},
		{"hihi", 2},
		{"hiHIhi", 2},
		{"", 0},
		{"h", 0},
		{"hi", 1},
		{"Hi is no HI on ahI", 0},
		{"hiho not HOHIhi", 2},
	}
	for _, test := range tests {
		if got := CountHi(test.s); got != test.want {
			t.Errorf("CountHi(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestEndOther(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"Hiabc", "abc", true},
		{"AbC", "HiaBc", true},
		{"abc", "abXabc", true},
		{"Hiabc", "abcd", false},
		{"Hiabc", "bc", true},
		{"Hiabcx", "bc", false},
		{"abc", "abc", true},
		{"xyz", "12xyz", true},
		{"yz", "12xz", false},
		{"Z", "12xz", true},
		{"12", "12", true},
		{"abcXYZ", "abcDEF", false},
		{"ab", "ab12", false},
		{"ab", "12ab", true},
	}
	for _, test := range tests {
		if got := EndOther(test.s1, test.s2); got != test.want {
			t.Errorf("EndOther(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestCatDog(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"catdog", true},
		{"catcat", false},
		{"1cat1cadodog", true},
		{"catxxdogxxxdog", false},
		{"catxdogxdogxcat", true},
		{"catxdogxdogxca", false},
		{"dogdogcat", false},
		{"dogogcat", true},
		{"dog", false},
		{"cat", false},
		{"ca", true},
		{"c", true},
		{"", true},
	}
	for _, test := range tests {
		if got := CatDog(test.s); got != test.want {
			t.Errorf("CatDog(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestXyzThere(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"abcxyz", true},
		{"abc.xyz", false},
		{"xyz.abc", true},
		{"abcxy", false},
		{"xyz", true},
		{"xy", false},
		{"x", false},
		{"", false},
		{"abc.xyzxyz", true},
		{"abc.xxyz", true},
		{".xyz", false},
		{"12.xyz", false},
		{"12xyz", true},
		{"1.xyz.xyz2.xyz", false},
	}
	for _, test := range tests {
		if got := XyzThere(test.s); got != test.want {
			t.Errorf("XyzThere(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}
