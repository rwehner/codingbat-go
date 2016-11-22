package logic1

import "testing"

func TestCigarParty(t *testing.T) {
	var tests = []struct {
		i    int
		b    bool
		want bool
	}{
		{30, false, false},
		{50, false, true},
		{70, true, true},
		{30, true, false},
		{50, true, true},
		{60, false, true},
		{61, false, false},
		{40, false, true},
		{39, false, false},
		{40, true, true},
		{39, true, false},
	}
	for _, test := range tests {
		if got := CigarParty(test.i, test.b); got != test.want {
			t.Errorf("CigarParty(%v, %v) = %v [wanted: %v]", test.i, test.b, got, test.want)
		}
	}
}

func TestCaughtSpeeding(t *testing.T) {
	var tests = []struct {
		i    int
		b    bool
		want int
	}{
		{60, false, 0},
		{65, false, 1},
		{65, true, 0},
		{80, false, 1},
		{85, false, 2},
		{85, true, 1},
		{70, false, 1},
		{75, false, 1},
		{75, true, 1},
		{40, false, 0},
		{40, true, 0},
		{90, false, 2},
	}
	for _, test := range tests {
		if got := CaughtSpeeding(test.i, test.b); got != test.want {
			t.Errorf("CaughtSpeeding(%v, %v) = %v [wanted: %v]", test.i, test.b, got, test.want)
		}
	}
}

func TestLove6(t *testing.T) {
	var tests = []struct {
		i1   int
		i2   int
		want bool
	}{
		{6, 4, true},
		{4, 5, false},
		{1, 5, true},
		{1, 6, true},
		{1, 8, false},
		{1, 7, true},
		{7, 5, false},
		{8, 2, true},
		{6, 6, true},
		{-6, 2, false},
		{-4, -10, true},
		{-7, 1, false},
		{7, -1, true},
		{-6, 12, true},
		{-2, -4, false},
		{7, 1, true},
		{0, 9, false},
		{8, 3, false},
		{3, 3, true},
		{3, 4, false},
	}
	for _, test := range tests {
		if got := Love6(test.i1, test.i2); got != test.want {
			t.Errorf("Love6(%v, %v) = %v [wanted: %v]", test.i1, test.i2, got, test.want)
		}
	}
}

func TestDateFashion(t *testing.T) {
	var tests = []struct {
		i1   int
		i2   int
		want int
	}{
		{5, 10, 2},
		{5, 2, 0},
		{5, 5, 1},
		{3, 3, 1},
		{10, 2, 0},
		{2, 9, 0},
		{9, 9, 2},
		{10, 5, 2},
		{2, 2, 0},
		{3, 7, 1},
		{2, 7, 0},
		{6, 2, 0},
	}
	for _, test := range tests {
		if got := DateFashion(test.i1, test.i2); got != test.want {
			t.Errorf("DateFashion(%v, %v) = %v [wanted: %v]", test.i1, test.i2, got, test.want)
		}
	}
}

func TestSortaSum(t *testing.T) {
	var tests = []struct {
		i1   int
		i2   int
		want int
	}{
		{3, 4, 7},
		{9, 4, 20},
		{10, 11, 21},
		{12, -3, 9},
		{-3, 12, 9},
		{4, 5, 9},
		{4, 6, 20},
		{14, 7, 21},
		{14, 6, 20},
	}
	for _, test := range tests {
		if got := SortaSum(test.i1, test.i2); got != test.want {
			t.Errorf("SortaSum(%v, %v) = %v [wanted: %v]", test.i1, test.i2, got, test.want)
		}
	}
}

func TestIn1To10(t *testing.T) {
	var tests = []struct {
		i    int
		b    bool
		want bool
	}{
		{5, false, true},
		{11, false, false},
		{11, true, true},
		{10, false, true},
		{10, true, true},
		{9, false, true},
		{9, true, false},
		{1, false, true},
		{1, true, true},
		{0, false, false},
		{0, true, true},
		{-1, false, false},
	}
	for _, test := range tests {
		if got := In1To10(test.i, test.b); got != test.want {
			t.Errorf("In1To10(%v, %v) = %v [wanted: %v]", test.i, test.b, got, test.want)
		}
	}
}

func TestSquirrelPlay(t *testing.T) {
	var tests = []struct {
		i    int
		b    bool
		want bool
	}{
		{70, false, true},
		{95, false, false},
		{95, true, true},
		{90, false, true},
		{90, true, true},
		{50, false, false},
		{50, true, false},
		{100, false, false},
		{100, true, true},
		{105, true, false},
		{59, false, false},
		{59, true, false},
		{60, false, true},
	}
	for _, test := range tests {
		if got := SquirrelPlay(test.i, test.b); got != test.want {
			t.Errorf("SquirrelPlay(%v, %v) = %v [wanted: %v]", test.i, test.b, got, test.want)
		}
	}
}

func TestAlarmClock(t *testing.T) {
	var tests = []struct {
		i    int
		b    bool
		want string
	}{
		{1, false, "7:00"},
		{5, false, "7:00"},
		{0, false, "10:00"},
		{6, false, "10:00"},
		{0, true, "off"},
		{6, true, "off"},
		{1, true, "10:00"},
		{3, true, "10:00"},
		{5, true, "10:00"},
	}
	for _, test := range tests {
		if got := AlarmClock(test.i, test.b); got != test.want {
			t.Errorf("AlarmClock(%v, %v) = %v [wanted: %v]", test.i, test.b, got, test.want)
		}
	}
}

func TestNearTen(t *testing.T) {
	var tests = []struct {
		i    int
		want bool
	}{
		{12, true},
		{17, false},
		{19, true},
		{31, true},
		{6, false},
		{10, true},
		{11, true},
		{21, true},
		{22, true},
		{23, false},
		{54, false},
		{155, false},
		{158, true},
		{3, false},
		{1, true},
	}
	for _, test := range tests {
		if got := NearTen(test.i); got != test.want {
			t.Errorf("NearTen(%v) = %v [wanted: %v]", test.i, got, test.want)
		}
	}
}
