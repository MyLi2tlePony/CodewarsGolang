package character_with_longest_consecutive_repetition

type Result struct {
	C rune // character
	L int  // count
}

func solution(text string) Result {
	var bestResult, countResult Result

	for _, char := range text {
		if char == countResult.C {
			countResult.L++
		} else {
			countResult.C = char
			countResult.L = 1
		}

		if countResult.L > bestResult.L {
			bestResult = countResult
		}
	}

	return bestResult
}
