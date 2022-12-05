package incrementer

func solution(elements []int) []int {
	for i, element := range elements {
		elements[i] = (element + i + 1) % 10
	}

	return elements
}
