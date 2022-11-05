package multiples_of_3_or_5

func solution(number int) int {
	if number < 0 {
		return 0
	}

	result := 0

	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	return result
}
