package take_a_number_and_sum_its_digits_raised_to_the_consecutive_powers

import (
	"math"
)

func solution(a, b uint64) []uint64 {
	var result []uint64

	for i := a; i <= b; i++ {
		if isEureka(i) {
			result = append(result, i)
		}
	}

	return result
}

func isEureka(num uint64) bool {
	digits := getDigits(num)
	var result uint64 = 0

	for i, digit := range digits {
		result += pow(uint64(digit), i+1)
	}

	return result == num
}

func getDigits(num uint64) []uint8 {
	lastDigitIndex := int8(math.Log10(float64(num)))
	digits := make([]uint8, 0)

	for i := lastDigitIndex; i >= 0; i-- {
		tens := uint64(math.Pow(float64(10), float64(i)))
		digits = append(digits, uint8(num/tens))
		num -= num / tens * tens
	}

	return digits
}

func pow(num uint64, degree int) uint64 {
	var result uint64 = 1

	for i := 0; i < degree; i++ {
		result *= num
	}

	return result
}
