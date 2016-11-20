package warmup2

import "testing"

func TestArray123(t *testing.T) {
	var tests = []struct {
		s    []int
		want bool
	}{
		{[]int{1, 1, 2, 3, 1}, true},
		{[]int{1, 1, 2, 4, 1}, false},
		{[]int{1, 1, 2, 1, 2, 3}, true},
		{[]int{1, 1, 2, 1, 2, 1}, false},
		{[]int{1, 2, 3, 1, 2, 3}, true},
		{[]int{1, 2, 3}, true},
		{[]int{1, 1, 1}, false},
		{[]int{1, 2}, false},
		{[]int{1}, false},
		{[]int{}, false},
	}
	for _, test := range tests {
		if got := Array123(test.s); got != test.want {
			t.Errorf("Array123(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestArrayCount9(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 9}, 1},
		{[]int{1, 9, 9}, 2},
		{[]int{1, 9, 9, 3, 9}, 3},
		{[]int{1, 2, 3}, 0},
		{[]int{}, 0},
		{[]int{4, 2, 4, 3, 1}, 0},
		{[]int{9, 2, 4, 3, 1}, 1},
	}
	for _, test := range tests {
		if got := ArrayCount9(test.s); got != test.want {
			t.Errorf("ArrayCount9(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestArrayFront(t *testing.T) {
	var tests = []struct {
		s    []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 9}, false},
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{9, 2, 3}, true},
		{[]int{1, 9, 9}, true},
		{[]int{1, 2, 3}, false},
		{[]int{1, 9}, true},
		{[]int{5, 5}, false},
		{[]int{2}, false},
		{[]int{9}, true},
		{[]int{}, false},
		{[]int{3, 9, 2, 3, 3}, true},
	}
	for _, test := range tests {
		if got := ArrayFront(test.s); got != test.want {
			t.Errorf("ArrayFront(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestFrontTimes(t *testing.T) {
	var tests = []struct {
		s    string
		i    int
		want string
	}{
		{"Chocolate", 2, "ChoCho"},
		{"Chocolate", 3, "ChoChoCho"},
		{"Abc", 3, "AbcAbcAbc"},
		{"Ab", 4, "AbAbAbAb"},
		{"A", 4, "AAAA"},
		{"", 4, ""},
		{"Abc", 0, ""},
	}
	for _, test := range tests {
		if got := FrontTimes(test.s, test.i); got != test.want {
			t.Errorf("FrontTimes(%v, %v) = %v [wanted: %v]", test.s, test.i, got, test.want)
		}
	}
}

func TestLast2(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"hixxhi", 1},
		{"xaxxaxaxx", 1},
		{"axxxaaxx", 2},
		{"xxaxxaxxaxx", 3},
		{"xaxaxaxx", 0},
		{"xxxx", 2},
		{"13121312", 1},
		{"11212", 1},
		{"13121311", 0},
		{"1717171", 2},
		{"hi", 0},
		{"h", 0},
		{"", 0},
	}
	for _, test := range tests {
		if got := Last2(test.s); got != test.want {
			t.Errorf("Last2(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestStringBits(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Hello", "Hlo"},
		{"Hi", "H"},
		{"Heeololeo", "Hello"},
		{"HiHiHi", "HHH"},
		{"", ""},
		{"Greetings", "Getns"},
		{"Chocoate", "Coot"},
		{"pi", "p"},
		{"Hello Kitten", "HloKte"},
		{"hxaxpxpxy", "happy"},
	}
	for _, test := range tests {
		if got := StringBits(test.s); got != test.want {
			t.Errorf("StringBits(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestStringMatch(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want int
	}{
		{"xxcaazz", "xxbaaz", 3},
		{"abc", "abc", 2},
		{"abc", "axc", 0},
		{"hello", "he", 1},
		{"he", "hello", 1},
		{"h", "hello", 0},
		{"", "hello", 0},
		{"aabbccdd", "abbbxxd", 1},
		{"aaxxaaxx", "iaxxai", 3},
		{"iaxxai", "aaxxaaxx", 3},
	}
	for _, test := range tests {
		if got := StringMatch(test.s1, test.s2); got != test.want {
			t.Errorf("StringMatch(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestStringSplosion(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Code", "CCoCodCode"},
		{"abc", "aababc"},
		{"ab", "aab"},
		{"x", "x"},
		{"fade", "ffafadfade"},
		{"There", "TThTheTherThere"},
		{"Kitten", "KKiKitKittKitteKitten"},
		{"Bye", "BByBye"},
		{"Good", "GGoGooGood"},
		{"Bad", "BBaBad"},
	}
	for _, test := range tests {
		if got := StringSplosion(test.s); got != test.want {
			t.Errorf("StringSplosion(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestStringTimes(t *testing.T) {
	var tests = []struct {
		s    string
		i    int
		want string
	}{
		{"Hi", 2, "HiHi"},
		{"Hi", 3, "HiHiHi"},
		{"Hi", 1, "Hi"},
		{"Hi", 0, ""},
		{"Hi", 5, "HiHiHiHiHi"},
		{"Oh Boy!", 2, "Oh Boy!Oh Boy!"},
		{"x", 4, "xxxx"},
		{"", 4, ""},
		{"code", 2, "codecode"},
		{"code", 3, "codecodecode"},
	}
	for _, test := range tests {
		if got := StringTimes(test.s, test.i); got != test.want {
			t.Errorf("StringTimes(%v, %v) = %v [wanted: %v]", test.s, test.i, got, test.want)
		}
	}
}
