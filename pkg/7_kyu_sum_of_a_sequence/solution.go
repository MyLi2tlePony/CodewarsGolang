package sum_of_a_sequence

func solution(start, end, step int) int {
	result := 0

	for start <= end {
		result += start
		start += step
	}

	return result
}
