package the_supermarket_queue

func solution(customers []int, n int) int {
	totalTime := 0

	for !allValuesZero(customers) {
		totalTime++
		freeN := n

		for i := 0; i < len(customers); i++ {
			if customers[i] > 0 && freeN > 0 {
				customers[i]--
				freeN--
			}
		}
	}

	return totalTime
}

func allValuesZero(customers []int) bool {
	for _, c := range customers {
		if c > 0 {
			return false
		}
	}

	return true
}
