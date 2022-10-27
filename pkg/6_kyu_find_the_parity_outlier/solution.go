package find_the_parity_outlier

func solution(integers []int) int {
	var remain int

	if integers[0]&1+integers[1]&1+integers[2]&1 < 2 {
		remain = 1
	}

	for _, integer := range integers {
		if integer&1 == remain {
			return integer
		}
	}

	return 0
}
