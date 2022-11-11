package vowel_count

import "strings"

func solution(str string) bool {
	start := 0
	end := len(str) - 1

	lowerStr := strings.ToLower(str)
	for start < end {
		if lowerStr[start] != lowerStr[end] {
			return false
		}

		start++
		end--
	}

	return true
}
