package bouncing_balls

func solution(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	frequency := make(map[int]int)

	for _, element := range array2 {
		frequency[element]++
	}

	for _, element := range array1 {
		val := element * element

		if frequency[val] < 1 {
			return false
		}

		frequency[val]--
	}

	for _, elementFrequency := range frequency {
		if elementFrequency > 0 {
			return false
		}
	}

	return true
}
