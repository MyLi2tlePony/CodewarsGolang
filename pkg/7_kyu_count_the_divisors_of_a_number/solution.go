package count_the_divisors_of_a_number

func solution(n int) (divisorsNumber int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisorsNumber++
		}
	}

	return
}
