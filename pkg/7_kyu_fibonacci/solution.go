package fibonacci

func solution(n int) int {
	countNum := 1
	previousNum := 0

	for i := 0; i < n; i++ {
		countNum, previousNum = countNum+previousNum, countNum
	}

	return previousNum
}
