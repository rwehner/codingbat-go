package logic1

// When squirrels get together for a party, they like to have cigars. A squirrel
// party is successful when the number of cigars is between 40 and 60, inclusive.
// Unless it is the weekend, in which case there is no upper bound on the number
// of cigars. Return true if the party with the given values is successful, or false otherwise.
func CigarParty(cigars int, weekend bool) bool {
	return true
}

// You are driving a little too fast, and a police officer stops you. Write code
// to compute the result, encoded as an int value: 0=no ticket, 1=small ticket, 2=big ticket.
// If speed is 60 or less, the result is 0. If speed is between 61 and 80 inclusive,
// the result is 1. If speed is 81 or more, the result is 2. Unless it is your
// birthday -- on that day, your speed can be 5 higher in all cases.
func CaughtSpeeding(speed int, isBirthday bool) int {
	return 0
}

// The number 6 is a truly great number. Given two int values, a and b, return true
// if either one is 6. Or if their sum or difference is 6.
func Love6(a, b int) bool {
	return false
}

// You and your date are trying to get a table at a restaurant. The parameter "you"
// is the stylishness of your clothes, in the range 0..10, and "date" is the stylishness
// of your date's clothes. The result getting the table is encoded as an int value
// with 0=no, 1=maybe, 2=yes. If either of you is very stylish, 8 or more, then the
// result is 2 (yes). With the exception that if either of you has style of 2 or less,
// then the result is 0 (no). Otherwise the result is 1 (maybe).
func DateFashion(you, date int) int {
	return 0
}

// Given 2 ints, a and b, return their sum. However, sums in the range 10..19 inclusive,
// are forbidden, so in that case just return 20.
func SortaSum(a, b int) int {
	return 0
}

// Given a number n, return true if n is in the range 1..10, inclusive. Unless
// "outsideMode" is true, in which case return true if the number is less or equal
// to 1, or greater or equal to 10.
func In1To10(n int, outsideMode bool) bool {
	return false
}

// The squirrels in Palo Alto spend most of the day playing. In particular, they
// play if the temperature is between 60 and 90 (inclusive). Unless it is summer,
// then the upper limit is 100 instead of 90. Given an int temperature and a boolean
// isSummer, return true if the squirrels play and false otherwise.
func SquirrelPlay(temp int, isSummer bool) bool {
	return false
}

// Given a day of the week encoded as 0=Sun, 1=Mon, 2=Tue, ...6=Sat, and a boolean
// indicating if we are on vacation, return a string of the form "7:00" indicating
// when the alarm clock should ring. Weekdays, the alarm should be "7:00" and on the
// weekend it should be "10:00". Unless we are on vacation -- then on weekdays it
// should be "10:00" and weekends it should be "off".
func AlarmClock(day int, isVacation bool) string {
	return ""
}

// Given a non-negative number "num", return true if num is within 2 of a multiple of 10.
// Note: (a % b) is the remainder of dividing a by b, so (7 % 5), is 2. See also:
// Introduction to Mod (http://codingbat.com/doc/practice/mod-introduction.html)
func NearTen(num int) bool {
	return false
}
