package which_are_in

import (
	"sort"
	"strings"
)

func solution(array1 []string, array2 []string) []string {
	substrings := make([]string, 0)

	for _, substring := range array1 {
		if SliceContain(substrings, substring) {
			continue
		}

		for _, str := range array2 {
			if strings.Contains(str, substring) {
				substrings = append(substrings, substring)
				break
			}
		}
	}

	sort.Strings(substrings)

	return substrings
}

func SliceContain(elements []string, str string) bool {
	for _, element := range elements {
		if element == str {
			return true
		}
	}

	return false
}
