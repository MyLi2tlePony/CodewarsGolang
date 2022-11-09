package string_ends_with

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(in string) string {
	max, min := 0, 0

	for index, strNum := range strings.Split(in, " ") {
		num, _ := strconv.Atoi(strNum)

		if index == 0 {
			max = num
			min = num
		}

		if max < num {
			max = num
		}

		if min > num {
			min = num
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
