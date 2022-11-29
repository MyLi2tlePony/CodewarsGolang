package find_the_odd_int

func solution(sequence []int) (result int) {
	frequency := make(map[int]int)

	for _, num := range sequence {
		frequency[num]++
	}

	for num, freq := range frequency {
		if freq%2 == 1 {
			result = num
			break
		}
	}

	return
}
