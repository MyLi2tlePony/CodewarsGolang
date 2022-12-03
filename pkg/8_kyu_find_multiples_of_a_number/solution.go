package find_multiples_of_a_number

func solution(integer, limit int) []int {
	results := make([]int, 0)
	result := integer

	for result <= limit {
		results = append(results, result)
		result += integer
	}

	return results
}
