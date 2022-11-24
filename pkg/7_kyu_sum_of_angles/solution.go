package sum_of_angles

import "math"

func solution(n int) int {
	return int(math.Round((float64(180) - float64(360)/float64(n)) * float64(n)))
}
