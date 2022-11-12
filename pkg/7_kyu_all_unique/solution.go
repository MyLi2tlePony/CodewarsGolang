package cats_and_shelves

func solution(str string) bool {
	unique := make(map[rune]int)

	for _, s := range str {
		if _, ok := unique[s]; ok {
			return false
		}

		unique[s]++
	}

	return true
}
