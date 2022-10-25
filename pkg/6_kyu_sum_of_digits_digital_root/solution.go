package sum_of_digits_digital_root

func solution(n int) int {
	if n < 10 {
		return n
	}

	result := 0

	for n > 0 {
		result += n % 10
		n /= 10
	}
	
	return solution(result)
}
