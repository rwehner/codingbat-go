package list1

// Given an array of ints, return true if 6 appears as either the first or last
// element in the array. The array will be length 1 or more.
func FirstLast6(i []int) bool {
	if i[0] == 6 || i[len(i)-1] == 6 {
		return true
	}
	return false
}

// Given 2 arrays of ints, a and b, return true if they have the same first
// element or they have the same last element. Both arrays will be length 1 or more.
func CommonEnd(a, b []int) bool {
	if a[0] == b[0] || a[len(a)-1] == b[len(b)-1] {
		return true
	}
	return false
}

// Given an array of ints length 3, return a new array with the elements in
// reverse order, so [3]int{1, 2, 3} becomes [3]int{3, 2, 1}.
func Reverse3(a [3]int) [3]int {
	var reversed [3]int
	for i, v := range a {
		reversed[len(a)-1-i] = v
	}
	return reversed
}

// Given 2 int arrays, a and b, each length 3, return a new array length 2
// containing their middle elements.
func MiddleWay(a, b [3]int) [2]int {
	return [2]int{a[1], b[1]}
}

// Given an array of ints, return true if the array is length 1 or more, and
// the first element and the last element are equal.
func SameFirstLast(i []int) bool {
	if len(i) >= 1 {
		if i[0] == i[len(i)-1] {
			return true
		}
	}
	return false
}

// Given an array of ints length 3, return the sum of all the elements
func Sum3(i [3]int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

// Given an array of ints length 3, figure out which is larger, the first or
// last element in the array, and set all the other elements to be that value.
// Return the changed array.
func MaxEnd3(a [3]int) [3]int {
	var largest int
	var new [3]int
	if a[0] >= a[len(a)-1] {
		largest = a[0]
	} else {
		largest = a[len(a)-1]
	}
	for i := range a {
		new[i] = largest
	}
	return new
}

// Given an array of ints, return a new array length 2 containing the first and
// last elements from the original array. The original array will be length 1 or more.
func MakeEnds(i []int) [2]int {
	return [2]int{i[0], i[len(i)-1]}
}

// Return an int array length 3 containing the first 3 digits of pi, {3, 1, 4}.
func MakePi() []int {
	return []int{3, 1, 4}
}

// Given an array of ints length 3, return an array with the elements
// "rotated left" so [3]int{1, 2, 3} yields [3]int{2, 3, 1}.
func RotateLeft3(a [3]int) [3]int {
	var new [3]int
	new[0] = a[1]
	new[1] = a[2]
	new[2] = a[0]
	return new
}

// Given an array of ints, return the sum of the first 2 elements in the array.
// If the array length is less than 2, just sum up the elements that exist,
// returning 0 if the array is length 0.
func Sum2(a []int) int {
	imax := 1
	sum := 0
	if len(a) < 2 {
		imax = len(a) - 1
	}
	for i := 0; i <= imax; i++ {
		sum += a[i]
	}
	return sum
}

// Given an int array length 2, return true if it contains a 2 or a 3.
func Has23(i [2]int) bool {
	for _, v := range i {
		if v == 2 || v == 3 {
			return true
		}
	}
	return false
}
