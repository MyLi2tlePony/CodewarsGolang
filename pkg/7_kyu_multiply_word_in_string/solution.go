package multiply_word_in_string

import (
	"strings"
)

func solution(str string, loc, num int) string {
	if len(str) == 0 {
		return ""
	}

	words := strings.Split(str, " ")

	if num == 0 {
		return words[loc]
	}

	similarWords := make([]string, num)

	for i := 0; i < num; i++ {
		similarWords[i] = words[loc]
	}

	return strings.Join(similarWords, "-")
}
