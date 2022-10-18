package last_digit_of_a_large_number

import (
	"math"
	"regexp"
	"strconv"
)

var (
	lastDigit  = regexp.MustCompile(`\d$`)
	lastDigits = regexp.MustCompile(`\d\d$`)
)

func solution(n1, n2 string) int {
	num1, _ := strconv.ParseFloat(lastDigit.FindString(n1), 32)

	var num2 int
	if len(n2) < 2 {
		num2, _ = strconv.Atoi(lastDigit.FindString(n2))
	} else {
		num2, _ = strconv.Atoi(lastDigits.FindString(n2))
	}

	if num2%4 == 0 {
		return int(math.Pow(num1, 4)) % 10
	}

	return int(math.Pow(num1, float64(num2-num2/4*4))) % 10
}
