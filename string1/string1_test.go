package string1

import "testing"

func TestHelloName(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Bob", "Hello Bob!"},
		{"Alice", "Hello Alice!"},
		{"X", "Hello X!"},
		{"Dolly", "Hello Dolly!"},
		{"Alpha", "Hello Alpha!"},
		{"Omega", "Hello Omega!"},
		{"Goodbye", "Hello Goodbye!"},
		{"ho ho ho", "Hello ho ho ho!"},
		{"xyz!", "Hello xyz!!"},
		{"Hello", "Hello Hello!"},
	}
	for _, test := range tests {
		if got := HelloName(test.s); got != test.want {
			t.Errorf("HelloName(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMakeoutWord(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want string
	}{
		{"<<>>", "Yay", "<<Yay>>"},
		{"<<>>", "WooHoo", "<<WooHoo>>"},
		{"[[]]", "word", "[[word]]"},
		{"HHoo", "Hello", "HHHellooo"},
		{"abyz", "YAY", "abYAYyz"},
	}
	for _, test := range tests {
		if got := MakeoutWord(test.s1, test.s2); got != test.want {
			t.Errorf("MakeoutWord(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestFirstHalf(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"WooHoo", "Woo"},
		{"HelloThere", "Hello"},
		{"abcdef", "abc"},
		{"ab", "a"},
		{"", ""},
		{"0123456789", "01234"},
		{"kitten", "kit"},
	}
	for _, test := range tests {
		if got := FirstHalf(test.s); got != test.want {
			t.Errorf("FirstHalf(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestNonStart(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want string
	}{
		{"Hello", "There", "ellohere"},
		{"java", "code", "avaode"},
		{"shotl", "java", "hotlava"},
		{"ab", "xy", "by"},
		{"ab", "x", "b"},
		{"x", "ac", "c"},
		{"a", "x", ""},
		{"kit", "kat", "itat"},
		{"mart", "dart", "artart"},
	}
	for _, test := range tests {
		if got := NonStart(test.s1, test.s2); got != test.want {
			t.Errorf("NonStart(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestMakeAbba(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want string
	}{
		{"Hi", "Bye", "HiByeByeHi"},
		{"Yo", "Alice", "YoAliceAliceYo"},
		{"What", "Up", "WhatUpUpWhat"},
		{"aaa", "bbb", "aaabbbbbbaaa"},
		{"x", "y", "xyyx"},
		{"x", "", "xx"},
		{"", "y", "yy"},
		{"Bo", "Ya", "BoYaYaBo"},
		{"Ya", "Ya", "YaYaYaYa"},
	}
	for _, test := range tests {
		if got := MakeAbba(test.s1, test.s2); got != test.want {
			t.Errorf("MakeAbba(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestExtraEnd(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Hello", "lololo"},
		{"ab", "ababab"},
		{"Hi", "HiHiHi"},
		{"Candy", "dydydy"},
		{"Code", "dedede"},
	}
	for _, test := range tests {
		if got := ExtraEnd(test.s); got != test.want {
			t.Errorf("ExtraEnd(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestWithoutEnd(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Hello", "ell"},
		{"java", "av"},
		{"coding", "odin"},
		{"code", "od"},
		{"ab", ""},
		{"Chocolate!", "hocolate"},
		{"kitten", "itte"},
		{"woohoo", "ooho"},
	}
	for _, test := range tests {
		if got := WithoutEnd(test.s); got != test.want {
			t.Errorf("WithoutEnd(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestLeft2(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Hello", "lloHe"},
		{"java", "vaja"},
		{"Hi", "Hi"},
		{"code", "deco"},
		{"cat", "tca"},
		{"12345", "34512"},
		{"Chocolate", "ocolateCh"},
		{"bricks", "icksbr"},
	}
	for _, test := range tests {
		if got := Left2(test.s); got != test.want {
			t.Errorf("Left2(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestMakeTags(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want string
	}{
		{"i", "Yay", "<i>Yay</i>"},
		{"i", "Hello", "<i>Hello</i>"},
		{"cite", "Yay", "<cite>Yay</cite>"},
		{"address", "here", "<address>here</address>"},
		{"body", "Heart", "<body>Heart</body>"},
		{"i", "i", "<i>i</i>"},
		{"i", "", "<i></i>"},
	}
	for _, test := range tests {
		if got := MakeTags(test.s1, test.s2); got != test.want {
			t.Errorf("MakeTags(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}

func TestFirstTwo(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Hello", "He"},
		{"abcdefg", "ab"},
		{"ab", "ab"},
		{"a", "a"},
		{"", ""},
		{"Kitten", "Ki"},
		{"hi", "hi"},
		{"hiya", "hi"},
	}
	for _, test := range tests {
		if got := FirstTwo(test.s); got != test.want {
			t.Errorf("FirstTwo(%v) = %v [wanted: %v]", test.s, got, test.want)
		}
	}
}

func TestComboString(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want string
	}{
		{"Hello", "hi", "hiHellohi"},
		{"hi", "Hello", "hiHellohi"},
		{"aaa", "b", "baaab"},
		{"b", "aaa", "baaab"},
		{"aaa", "", "aaa"},
		{"", "bb", "bb"},
		{"aaa", "1234", "aaa1234aaa"},
		{"aaa", "bb", "bbaaabb"},
		{"a", "bb", "abba"},
		{"bb", "a", "abba"},
		{"xyz", "ab", "abxyzab"},
	}
	for _, test := range tests {
		if got := ComboString(test.s1, test.s2); got != test.want {
			t.Errorf("ComboString(%v, %v) = %v [wanted: %v]", test.s1, test.s2, got, test.want)
		}
	}
}
