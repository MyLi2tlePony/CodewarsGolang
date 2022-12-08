package string_repeat

func solution(repetitions int, value string) (result string) {
	for i := 0; i < repetitions; i++ {
		result += value
	}

	return
}
