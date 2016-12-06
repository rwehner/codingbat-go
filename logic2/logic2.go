package logic2

// We want to make a row of bricks that is goal inches long. We have a number of
// small bricks (1 inch each) and big bricks (5 inches each). Return true if it is
// possible to make the goal by choosing from the given bricks. This is a little harder
// than it looks and can be done without any loops.
// http://codingbat.com/doc/practice/makebricks-introduction.html
func MakeBricks(brick1, brick5, goal int) bool {
	requiredBricks5 := goal / 5
	if requiredBricks5 >= brick5 {
		requiredBricks5 = brick5
	}
	if brick1 >= goal-(requiredBricks5*5) {
		return true
	}
	return false
}

// Given 3 int values, a b c, return their sum. However, if any of the values is
// a teen -- in the range 13..19 inclusive -- then that value counts as 0, except 15 and 16
// do not count as a teens. Write a separate helper "def fixteen(n):"that takes in an
// int value and returns that value fixed for the teen rule. In this way, you avoid
// repeating the teen code 3 times (i.e. "decomposition"). Define the helper below
// and at the same indent level as the main NoTeenSum().
func FixTeen(n int) int {
	if n == 15 || n == 16 {
		return n
	}
	if n > 12 && n < 20 {
		return 0
	}
	return n
}
func NoTeenSum(a, b, c int) int {
	nums := []int{a, b, c}
	sum := 0
	for _, i := range nums {
		sum += FixTeen(i)
	}
	return sum
}

// We want make a package of goal kilos of chocolate. We have small bars (1 kilo each)
// and big bars (5 kilos each). Return the number of small bars to use, assuming we
// always use big bars before small bars. Return -1 if it can't be done.
func MakeChocolate(bar1, bar5, goal int) int {
	requiredBar5 := goal / 5
	if requiredBar5 >= bar5 {
		requiredBar5 = bar5
	}
	requiredBar1 := goal - (requiredBar5 * 5)
	if bar1 >= requiredBar1 {
		return requiredBar1
	}
	return -1
}

// Given 3 int values, a b c, return their sum. However, if one of the values is
// the same as another of the values, it does not count towards the sum.
func LoneSum(a, b, c int) int {
	m := make(map[int]int)
	nums := []int{a, b, c}
	sum := 0
	for _, n := range nums {
		m[n]++
	}
	for k, v := range m {
		if v == 1 {
			sum += k
		}
	}
	return sum
}

// For this problem, we'll round an int value up to the next multiple of 10 if its
// rightmost digit is 5 or more, so 15 rounds up to 20. Alternately, round down to
// the previous multiple of 10 if its rightmost digit is less than 5, so 12 rounds
// down to 10. Given 3 ints, a b c, return the sum of their rounded values. To avoid
// code repetition, write a separate helper "def round10(num):" and call it 3 times.
// Write the helper entirely below and at the same indent level as RoundSum().
func Round10(a int) int {
	tens := a / 10
	mod := a % 10
	if mod >= 5 {
		tens++
	}
	return tens * 10
}
func RoundSum(a, b, c int) int {
	return Round10(a) + Round10(b) + Round10(c)
}

// Given 3 int values, a b c, return their sum. However, if one of the values is 13
// then it does not count towards the sum and values to its right do not count. So
// for example, if b is 13, then both b and c do not count.
func LuckySum(a, b, c int) int {
	nums := []int{a, b, c}
	sum := 0
	for _, n := range nums {
		if n == 13 {
			return sum
		}
		sum += n
	}
	return sum
}

// Given three ints, a b c, return true if one of b or c is "close" (differing from
// a by at most 1), while the other is "far", differing from both other values by 2
// or more.
func Abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
func CloseFar(a, b, c int) bool {
	distanceA2B := Abs(a - b)
	distanceA2C := Abs(a - c)
	distanceB2C := Abs(b - c)

	if distanceA2B <= 1 {
		if distanceA2C >= 2 && distanceB2C >= 2 {
			return true
		}
	} else if distanceA2C <= 1 {
		if distanceA2B >= 2 && distanceB2C >= 2 {
			return true
		}
	}
	return false
}
