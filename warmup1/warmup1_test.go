package warmup1

import "testing"

func TestSleepIn(t *testing.T) {
	var tests = []struct {
		w    bool
		v    bool
		want bool
	}{
		{false, false, true},
		{false, true, true},
		{true, false, false},
		{true, true, true},
	}
	for _, test := range tests {
		if got := SleepIn(test.w, test.v); got != test.want {
			t.Errorf("sleepIn(%v, %v) = %v [wanted: %v]", test.w, test.v, got, test.want)
		}
	}
}

func TestMonkeyTrouble(t *testing.T) {
	var tests = []struct {
		a    bool
		b    bool
		want bool
	}{
		{true, true, true},
		{false, false, true},
		{true, false, false},
		{false, true, false},
	}
	for _, test := range tests {
		if got := MonkeyTrouble(test.a, test.b); got != test.want {
			t.Errorf("sleepIn(%v, %v) = %v [wanted: %v]", test.a, test.b, got, test.want)
		}
	}
}

func TestSumDouble(t *testing.T) {
	var tests = []struct {
		input1 int
		input2 int
		want   int
	}{
		{1, 2, 3},
		{3, 2, 5},
		{2, 2, 8},
		{-1, 0, -1},
		{3, 3, 12},
		{0, 0, 0},
		{0, 1, 1},
		{3, 4, 7},
	}
	for _, test := range tests {
		if got := SumDouble(test.input1, test.input2); got != test.want {
			t.Errorf("SumDouble(%d, %d) = %d [wanted: %d]", test.input1, test.input2, got, test.want)
		}
	}
}

func TestDiff21(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{19, 2},
		{10, 11},
		{21, 0},
		{22, 2},
		{25, 8},
		{30, 18},
		{0, 21},
		{1, 20},
		{2, 19},
		{-1, 22},
		{-2, 23},
		{50, 58},
	}
	for _, test := range tests {
		if got := Diff21(test.input); got != test.want {
			t.Errorf("Diff21(%d) = %d [wanted: %d]", test.input, got, test.want)
		}
	}
}

func TestParrotTrouble(t *testing.T) {
	var tests = []struct {
		input1 bool
		input2 int
		want   bool
	}{
		{true, 6, true},
		{true, 7, false},
		{false, 6, false},
		{true, 21, true},
		{false, 21, false},
		{false, 20, false},
		{true, 23, true},
		{false, 23, false},
		{true, 20, false},
		{false, 12, false},
	}
	for _, test := range tests {
		if got := ParrotTrouble(test.input1, test.input2); got != test.want {
			t.Errorf("ParrotTrouble(%v, %d) = %v [wanted %v]", test.input1, test.input2, got, test.want)
		}
	}
}

func TestMakes10(t *testing.T) {
	var tests = []struct {
		input1 int
		input2 int
		want   bool
	}{
		{9, 10, true},
		{9, 9, false},
		{1, 9, true},
		{10, 1, true},
		{10, 10, true},
		{8, 2, true},
		{8, 3, false},
		{10, 42, true},
		{12, -2, true},
	}
	for _, test := range tests {
		if got := Makes10(test.input1, test.input2); got != test.want {
			t.Errorf("Makes10(%d, %d) = %v [wanted %v]", test.input1, test.input2, got, test.want)
		}
	}
}

func TestNearHundred(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{93, true},
		{90, true},
		{89, false},
		{0, false},
		{-101, false},
		{-209, false},
		{209, true},
		{211, false},
		{191, true},
		{200, true},
		{100, true},
	}
	for _, test := range tests {
		if got := NearHundred(test.input); got != test.want {
			t.Errorf("NearHundred(%d) = %v [wanted: %v]", test.input, got, test.want)
		}
	}
}

func TestPosNeg(t *testing.T) {
	var tests = []struct {
		input1 int
		input2 int
		input3 bool
		want   bool
	}{
		{1, -1, false, true},
		{1, 1, false, false},
		{-1, 1, false, true},
		{-4, -5, true, true},
		{-4, 5, true, false},
		{-4, 5, false, true},
	}
	for _, test := range tests {
		if got := PosNeg(test.input1, test.input2, test.input3); got != test.want {
			t.Errorf("PosNeg(%d, %d, %v) = %v [wanted %v]", test.input1, test.input2, test.input3, got, test.want)
		}
	}
}

func TestNotString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"candy", "not candy"},
		{"x", "not x"},
		{"not bad", "not bad"},
		{"bad", "not bad"},
		{"is not", "not is not"},
		{"no", "not no"},
	}
	for _, test := range tests {
		if got := NotString(test.input); got != test.want {
			t.Errorf("NotString(%q) = %q [wanted: %q]", test.input, got, test.want)
		}
	}
}

func TestMissingChar(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 int
		want   string
	}{
		{"kitten", 1, "ktten"},
		{"kitten", 0, "itten"},
		{"kitten", 4, "kittn"},
		{"Hi", 0, "i"},
		{"Hi", 1, "H"},
		{"code", 0, "ode"},
		{"code", 1, "cde"},
		{"code", 2, "coe"},
		{"code", 3, "cod"},
		{"chocolate", 8, "chocolat"},
	}
	for _, test := range tests {
		if got := MissingChar(test.input1, test.input2); got != test.want {
			t.Errorf("MissingChar(%q, %d) = %q [wanted %q]", test.input1, test.input2, got, test.want)
		}
	}
}

func TestFrontBack(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"code", "eodc"},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"", ""},
		{"Chocolate", "ehocolatC"},
		{"aavJ", "Java"},
		{"hello", "oellh"},
	}
	for _, test := range tests {
		if got := FrontBack(test.input); got != test.want {
			t.Errorf("FrontBack(%q) = %q [wanted %q]", test.input, got, test.want)
		}
	}
}

func TestFront3(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"Java", "JavJavJav"},
		{"Chocolate", "ChoChoCho"},
		{"abc", "abcabcabc"},
		{"abcXYZ", "abcabcabc"},
		{"ab", "ababab"},
		{"a", "aaa"},
		{"", ""},
	}
	for _, test := range tests {
		if got := Front3(test.input); got != test.want {
			t.Errorf("Front3(%q) = %q [wanted %q]", test.input, got, test.want)
		}
	}
}
