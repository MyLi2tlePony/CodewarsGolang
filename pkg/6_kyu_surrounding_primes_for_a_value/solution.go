package surrounding_primes_for_a_value

import (
	"math"
)

func solution(n int) (result [2]int) {
	for i := n - 1; ; i-- {
		if IsPrime(i) {
			result[0] = i
			break
		}
	}

	for i := n + 1; ; i++ {
		if IsPrime(i) {
			result[1] = i
			break
		}
	}

	return
}

func IsPrime(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
