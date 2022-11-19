package jaden_casing_strings

import (
	"strings"
)

func solution(str string) string {
	if len(str) < 1 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(strings.ToUpper(string(str[0])))

	for i := 1; i < len(str); i++ {
		if str[i-1] == ' ' {
			builder.WriteString(strings.ToUpper(string(str[i])))
		} else {
			builder.WriteByte(str[i])
		}
	}

	return builder.String()
}
