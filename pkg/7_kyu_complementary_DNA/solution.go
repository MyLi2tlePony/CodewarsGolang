package string_ends_with

import "strings"

var replacer = strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")

func solution(dna string) string {
	return replacer.Replace(dna)
}
