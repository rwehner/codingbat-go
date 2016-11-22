package list1

import "testing"

func TestFirstLast6(t *testing.T) {
	var tests = []struct {
		s    []int
		want bool
	}{
		{[]int{1, 2, 6}, true},
		{[]int{6, 1, 2, 3}, true},
		{[]int{13, 6, 1, 2, 3}, false},
		{[]int{13, 6, 1, 2, 6}, true},
		{[]int{3, 2, 1}, false},
		{[]int{3, 6, 1}, false},
		{[]int{3, 6}, true},
		{[]int{6}, true},
		{[]int{3}, false},
		{[]int{5, 6}, true},
		{[]int{5, 5}, false},
		{[]int{1, 2, 3, 4, 6}, true},
		{[]int{1, 2, 3, 4}, false},
	}
	for _, test := range tests {
		if got := FirstLast6(test.s); got != test.want {
			t.Errorf("FirstLast6(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestCommonEnd(t *testing.T) {
	var tests = []struct {
		s1   []int
		s2   []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{7, 3}, true},
		{[]int{1, 2, 3}, []int{7, 3, 2}, false},
		{[]int{1, 2, 3}, []int{1, 3}, true},
		{[]int{1, 2, 3}, []int{1}, true},
		{[]int{1, 2, 3}, []int{2}, false},
	}
	for _, test := range tests {
		if got := CommonEnd(test.s1, test.s2); got != test.want {
			t.Errorf("CommonEnd(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestReverse3(t *testing.T) {
	var tests = []struct {
		s    [3]int
		want [3]int
	}{
		{[3]int{1, 2, 3}, [3]int{3, 2, 1}},
		{[3]int{5, 11, 9}, [3]int{9, 11, 5}},
		{[3]int{7, 0, 0}, [3]int{0, 0, 7}},
		{[3]int{2, 1, 2}, [3]int{2, 1, 2}},
		{[3]int{1, 2, 1}, [3]int{1, 2, 1}},
		{[3]int{2, 11, 3}, [3]int{3, 11, 2}},
		{[3]int{0, 6, 5}, [3]int{5, 6, 0}},
		{[3]int{7, 2, 3}, [3]int{3, 2, 7}},
	}
	for _, test := range tests {
		if got := Reverse3(test.s); got != test.want {
			t.Errorf("Reverse3(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMiddleWay(t *testing.T) {
	var tests = []struct {
		s1   [3]int
		s2   [3]int
		want [2]int
	}{
		{[3]int{1, 2, 3}, [3]int{4, 5, 6}, [2]int{2, 5}},
		{[3]int{7, 7, 7}, [3]int{3, 8, 0}, [2]int{7, 8}},
		{[3]int{5, 2, 9}, [3]int{1, 4, 5}, [2]int{2, 4}},
		{[3]int{1, 9, 7}, [3]int{4, 8, 8}, [2]int{9, 8}},
		{[3]int{1, 2, 3}, [3]int{3, 1, 4}, [2]int{2, 1}},
		{[3]int{1, 2, 3}, [3]int{4, 1, 1}, [2]int{2, 1}},
	}
	for _, test := range tests {
		if got := MiddleWay(test.s1, test.s2); got != test.want {
			t.Errorf("MiddleWay(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestSameFirstLast(t *testing.T) {
	var tests = []struct {
		s    []int
		want bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 1}, true},
		{[]int{7}, true},
		{[]int{}, false},
		{[]int{1, 2, 3, 4, 5, 1}, true},
		{[]int{1, 2, 3, 4, 5, 13}, false},
		{[]int{13, 2, 3, 4, 5, 13}, true},
		{[]int{7, 7}, true},
	}
	for _, test := range tests {
		if got := SameFirstLast(test.s); got != test.want {
			t.Errorf("SameFirstLast(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestSum3(t *testing.T) {
	var tests = []struct {
		s    [3]int
		want int
	}{
		{[3]int{1, 2, 3}, 6},
		{[3]int{5, 11, 2}, 18},
		{[3]int{7, 0, 0}, 7},
		{[3]int{1, 2, 1}, 4},
		{[3]int{1, 1, 1}, 3},
		{[3]int{2, 7, 2}, 11},
	}
	for _, test := range tests {
		if got := Sum3(test.s); got != test.want {
			t.Errorf("Sum3(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMaxEnd3(t *testing.T) {
	var tests = []struct {
		s    [3]int
		want [3]int
	}{
		{[3]int{1, 2, 3}, [3]int{3, 3, 3}},
		{[3]int{11, 5, 9}, [3]int{11, 11, 11}},
		{[3]int{2, 11, 3}, [3]int{3, 3, 3}},
		{[3]int{11, 3, 3}, [3]int{11, 11, 11}},
		{[3]int{3, 11, 11}, [3]int{11, 11, 11}},
		{[3]int{2, 2, 2}, [3]int{2, 2, 2}},
		{[3]int{2, 11, 2}, [3]int{2, 2, 2}},
		{[3]int{0, 0, 1}, [3]int{1, 1, 1}},
	}
	for _, test := range tests {
		if got := MaxEnd3(test.s); got != test.want {
			t.Errorf("MaxEnd3(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMakeEnds(t *testing.T) {
	var tests = []struct {
		s    []int
		want [2]int
	}{
		{[]int{1, 2, 3}, [2]int{1, 3}},
		{[]int{1, 2, 3, 4}, [2]int{1, 4}},
		{[]int{7, 4, 6, 2}, [2]int{7, 2}},
		{[]int{1, 2, 2, 2, 2, 2, 2, 3}, [2]int{1, 3}},
		{[]int{7, 4}, [2]int{7, 4}},
		{[]int{7}, [2]int{7, 7}},
		{[]int{5, 2, 9}, [2]int{5, 9}},
		{[]int{2, 3, 4, 1}, [2]int{2, 1}},
	}
	for _, test := range tests {
		if got := MakeEnds(test.s); got != test.want {
			t.Errorf("MakeEnds(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMakePi(t *testing.T) {
	want := []int{3, 1, 4}
	got := MakePi()
	for i, n := range want {
		if n != got[i] {
			t.Errorf("MakePi() = %v [wanted: %v]", got, want)
		}
	}
}

func TestRotateLeft3(t *testing.T) {
	var tests = []struct {
		s    [3]int
		want [3]int
	}{
		{[3]int{1, 2, 3}, [3]int{2, 3, 1}},
		{[3]int{5, 11, 9}, [3]int{11, 9, 5}},
		{[3]int{7, 0, 0}, [3]int{0, 0, 7}},
		{[3]int{1, 2, 1}, [3]int{2, 1, 1}},
		{[3]int{0, 0, 1}, [3]int{0, 1, 0}},
	}
	for _, test := range tests {
		if got := RotateLeft3(test.s); got != test.want {
			t.Errorf("RotateLeft3(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestSum2(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1}, 2},
		{[]int{1, 1, 1, 1}, 2},
		{[]int{1, 2}, 3},
		{[]int{1}, 1},
		{[]int{}, 0},
		{[]int{4, 5, 6}, 9},
		{[]int{4}, 4},
	}
	for _, test := range tests {
		if got := Sum2(test.s); got != test.want {
			t.Errorf("Sum2(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestHas23(t *testing.T) {
	var tests = []struct {
		s    [2]int
		want bool
	}{
		{[2]int{2, 5}, true},
		{[2]int{4, 3}, true},
		{[2]int{4, 5}, false},
		{[2]int{2, 2}, true},
		{[2]int{3, 2}, true},
		{[2]int{3, 3}, true},
		{[2]int{7, 7}, false},
		{[2]int{3, 9}, true},
		{[2]int{9, 5}, false},
	}
	for _, test := range tests {
		if got := Has23(test.s); got != test.want {
			t.Errorf("Has23(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}
