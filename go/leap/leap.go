package leap

// IsLeapYear returns true if the year parameter is a leap year.
func IsLeapYear(year int) bool {
	var (
		isEvenlyDivisibleBy4   = year%4 == 0
		isEvenlyDivisibleBy100 = year%100 == 0
		isEvenlyDivisibleBy400 = year%400 == 0
	)
	return isEvenlyDivisibleBy4 && !isEvenlyDivisibleBy100 || isEvenlyDivisibleBy400
}
