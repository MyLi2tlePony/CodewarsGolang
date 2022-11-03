package rgb_to_hex_conversion

import (
	"fmt"
)

var hexNums = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

func solution(r, g, b int) string {
	return hexConvert(r) + hexConvert(g) + hexConvert(b)
}

func hexConvert(num int) string {
	if num < 0 {
		return "00"
	}
	if num > 255 {
		return "FF"
	}
	return fmt.Sprintf("%s%s", hexNums[num/16], hexNums[num%16])
}
