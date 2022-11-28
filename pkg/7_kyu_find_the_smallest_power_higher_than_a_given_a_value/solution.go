package find_the_smallest_power_higher_than_a_given_a_value

import "math"

func solution(val, pow int) int {
	result := 0

	for i := 0.0; result <= val; i++ {
		result = int(math.Pow(i, float64(pow)))
	}

	return result
}
