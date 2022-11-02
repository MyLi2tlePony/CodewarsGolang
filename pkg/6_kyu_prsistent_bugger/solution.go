package prsistent_bugger

import "math"

func solution(n int) int {
	count := 0

	for int(math.Log10(float64(n))) > 0 {
		sub := 1

		for n > 0 {
			sub *= n % 10
			n /= 10
		}
		
		n = sub
		count++
	}

	return count
}
