package sort_numbers

import "sort"

func solution(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
