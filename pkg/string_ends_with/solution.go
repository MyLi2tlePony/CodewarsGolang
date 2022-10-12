package string_ends_with

func solution(text, ending string) bool {
	textIndex := len(text) - 1
	endingIndex := len(ending) - 1

	for textIndex >= 0 && endingIndex >= 0 {
		if text[textIndex] != ending[endingIndex] {
			return false
		}

		textIndex--
		endingIndex--
	}

	return len(text) >= len(ending)
}
