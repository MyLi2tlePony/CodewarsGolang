package beginner_reduce_but_grow

func solution(arr []int) int {
	result := 0

	if len(arr) > 0 {
		for i, element := range arr {
			if i == 0 {
				result = element
			} else {

				result *= element
			}
		}
	}

	return result
}
