package cats_and_shelves

func solution(start, finish int) int {
	jumpNumber := 0

	for jumpNumber = 0; finish-start > 0; jumpNumber++ {
		if finish-start > 2 {
			start += 3
		} else {
			start += 1
		}
	}

	return jumpNumber
}
