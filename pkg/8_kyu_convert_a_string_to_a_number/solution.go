package convert_a_string_to_a_number

import "strconv"

func solution(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
