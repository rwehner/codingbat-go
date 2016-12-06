package list2

import (
	"sort"
)

// Return the number of even ints in the given array. Note: the % "mod" operator
// computes the remainder, e.g. 5 % 2 is 1.
func CountEvens(a []int) int {
	count := 0
	for _, n := range a {
		if n%2 == 0 {
			count++
		}
	}
	return count
}

// Return the sum of the numbers in the array, returning 0 for an empty array.
// Except the number 13 is very unlucky, so it does not count and numbers that
// come immediately after a 13 also do not count.
func Sum13(a []int) int {
	sum := 0
	var clean []int
	for i := 0; i <= len(a)-1; i++ {
		if a[i] == 13 {
			i++
		} else {
			clean = append(clean, a[i])
		}
	}
	for _, i := range clean {
		sum += i
	}
	return sum
}

// Given an array length 1 or more of ints, return the difference between the
// largest and smallest values in the array.
func BigDiff(a []int) int {
	min, max := a[0], a[0]
	for _, i := range a {
		if i <= min {
			min = i
		}
		if i >= max {
			max = i
		}
	}
	return max - min
}

// Return the sum of the numbers in the array, except ignore sections of numbers
//starting with a 6 and extending to the next 7 (every 6 will be followed by at
// least one 7). Return 0 for no numbers.
func Sum67(a []int) int {
	sum := 0
	blocked := false
	for _, i := range a {
		switch {
		case blocked == true && i == 7:
			blocked = false
		case blocked == false && i == 6:
			blocked = true
		case blocked == false:
			sum += i
		}
	}
	return sum
}

// Return the "centered" average of an array of ints, which we'll say is the mean
// average of the values, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and
// likewise for the largest value. Use int division to produce the final average.
// You may assume that the array is length 3 or more.
func CenteredAverage(a []int) int {
	sort.Ints(a)
	trimmed := a[1 : len(a)-1]

	sum := 0
	for _, i := range trimmed {
		sum += i
	}
	return sum / len(trimmed)
}

// Given an array of ints, return true if the array contains a 2 next to a 2 somewhere.
func Has22(a []int) bool {
	for i, n := range a {
		if i+1 <= len(a)-1 { // don't look out of index range
			if n == 2 && a[i+1] == 2 {
				return true
			}
		}
	}
	return false
}
