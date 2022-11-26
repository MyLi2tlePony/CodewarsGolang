package CSV_representation_of_array

import (
	"strconv"
	"strings"
)

func solution(array [][]int) string {
	builder := strings.Builder{}

	for rowI, row := range array {
		for colI, col := range row {
			builder.WriteString(strconv.Itoa(col))

			if colI < len(row)-1 {
				builder.WriteString(",")
			}
		}

		if rowI < len(array)-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
