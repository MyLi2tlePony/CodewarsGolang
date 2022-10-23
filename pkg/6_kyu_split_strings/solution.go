package split_strings

import "regexp"

var regexpSplitByTwoChar = regexp.MustCompile(`(..)|(.)`)

func solution(str string) []string {
	result := regexpSplitByTwoChar.FindAllString(str, -1)

	if len(result) > 0 && len(result[len(result)-1]) < 2 {
		result[len(result)-1] += "_"
	}
	
	return result
}
