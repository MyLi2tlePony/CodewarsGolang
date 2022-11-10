package string_ends_with

import (
	"regexp"
)

var (
	regexGetUppercase = regexp.MustCompile(`[A-Z]`)
	regexGetLowercase = regexp.MustCompile(`[a-z]`)
	regexGetNumbers   = regexp.MustCompile(`\d`)
)

func solution(s string) []int {
	uppercase := len(regexGetUppercase.FindAllString(s, -1))
	lowercase := len(regexGetLowercase.FindAllString(s, -1))
	numbers := len(regexGetNumbers.FindAllString(s, -1))

	return []int{uppercase, lowercase, numbers, len(s) - uppercase - lowercase - numbers}
}
