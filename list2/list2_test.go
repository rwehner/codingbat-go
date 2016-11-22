package list2

import "testing"

// Return the number of even ints in the given array. Note: the % "mod" operator computes the remainder, e.g. 5 % 2 is 1.
func TestCountEvens(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{2, 1, 2, 3, 4}, 3},
		{[]int{2, 2, 0}, 3},
		{[]int{1, 3, 5}, 0},
		{[]int{}, 0},
		{[]int{11, 9, 0, 1}, 1},
		{[]int{2, 11, 9, 0}, 2},
		{[]int{2}, 1},
		{[]int{2, 5, 12}, 2},
	}
	for _, test := range tests {
		if got := CountEvens(test.s); got != test.want {
			t.Errorf("CountEvens(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestSum13(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 2, 1}, 6},
		{[]int{1, 1}, 2},
		{[]int{1, 2, 2, 1, 13}, 6},
		{[]int{1, 2, 13, 2, 1, 13}, 4},
		{[]int{13, 1, 2, 13, 2, 1, 13}, 3},
		{[]int{}, 0},
		{[]int{13}, 0},
		{[]int{13, 13}, 0},
		{[]int{13, 0, 13}, 0},
		{[]int{13, 1, 13}, 0},
		{[]int{5, 7, 2}, 14},
		{[]int{5, 13, 2}, 5},
		{[]int{0}, 0},
		{[]int{13, 0}, 0},
	}
	for _, test := range tests {
		if got := Sum13(test.s); got != test.want {
			t.Errorf("Sum13(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestBigDiff(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{10, 3, 5, 6}, 7},
		{[]int{7, 2, 10, 9}, 8},
		{[]int{2, 10, 7, 2}, 8},
		{[]int{2, 10}, 8},
		{[]int{10, 2}, 8},
		{[]int{10, 0}, 10},
		{[]int{2, 3}, 1},
		{[]int{2, 2}, 0},
		{[]int{2}, 0},
		{[]int{5, 1, 6, 1, 9, 9}, 8},
		{[]int{7, 6, 8, 5}, 3},
		{[]int{7, 7, 6, 8, 5, 5, 6}, 3},
	}
	for _, test := range tests {
		if got := BigDiff(test.s); got != test.want {
			t.Errorf("BigDiff(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestSum67(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 2}, 5},
		{[]int{1, 2, 2, 6, 99, 99, 7}, 5},
		{[]int{1, 1, 6, 7, 2}, 4},
		{[]int{1, 6, 2, 2, 7, 1, 6, 99, 99, 7}, 2},
		{[]int{1, 6, 2, 6, 2, 7, 1, 6, 99, 99, 7}, 2},
		{[]int{2, 7, 6, 2, 6, 7, 2, 7}, 18},
		{[]int{2, 7, 6, 2, 6, 2, 7}, 9},
		{[]int{1, 6, 7, 7}, 8},
		{[]int{6, 7, 1, 6, 7, 7}, 8},
		{[]int{6, 8, 1, 6, 7}, 0},
		{[]int{}, 0},
		{[]int{6, 7, 11}, 11},
		{[]int{11, 6, 7, 11}, 22},
		{[]int{2, 2, 6, 7, 7}, 11},
	}
	for _, test := range tests {
		if got := Sum67(test.s); got != test.want {
			t.Errorf("Sum67(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestCenteredAverage(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{1, 2, 3, 4, 100}, 3},
		{[]int{1, 1, 5, 5, 10, 8, 7}, 5},
		{[]int{-10, -4, -2, -4, -2, 0}, -3},
		{[]int{5, 3, 4, 6, 2}, 4},
		{[]int{5, 3, 4, 0, 100}, 4},
		{[]int{100, 0, 5, 3, 4}, 4},
		{[]int{4, 0, 100}, 4},
		{[]int{0, 2, 3, 4, 100}, 3},
		{[]int{1, 1, 100}, 1},
		{[]int{7, 7, 7}, 7},
		{[]int{1, 7, 8}, 7},
		{[]int{1, 1, 99, 99}, 50},
		{[]int{1000, 0, 1, 99}, 50},
		{[]int{4, 4, 4, 4, 5}, 4},
		{[]int{4, 4, 4, 1, 5}, 4},
		{[]int{6, 4, 8, 12, 3}, 6},
	}
	for _, test := range tests {
		if got := CenteredAverage(test.s); got != test.want {
			t.Errorf("CenteredAverage(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestHas22(t *testing.T) {
	var tests = []struct {
		s    []int
		want bool
	}{
		{[]int{1, 2, 2}, true},
		{[]int{1, 2, 1, 2}, false},
		{[]int{2, 1, 2}, false},
		{[]int{2, 2, 1, 2}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 3, 2, 2}, true},
		{[]int{2, 3, 2, 2}, true},
		{[]int{4, 2, 4, 2, 2, 5}, true},
		{[]int{1, 2}, false},
		{[]int{2, 2}, true},
		{[]int{2}, false},
		{[]int{}, false},
		{[]int{3, 3, 2, 2}, true},
		{[]int{5, 2, 5, 2}, false},
	}
	for _, test := range tests {
		if got := Has22(test.s); got != test.want {
			t.Errorf("Has22(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}
