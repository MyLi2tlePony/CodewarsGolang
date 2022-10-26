package encrypt_this

import (
	"fmt"
	"strings"
)

func solution(text string) string {
	if text == "" {
		return text
	}

	builder := strings.Builder{}

	words := strings.Split(text, " ")
	for i, word := range words {
		if i != 0 {
			builder.Write([]byte(" "))
		}
		builder.Write(EncryptWord(word))
	}

	return builder.String()
}

func EncryptWord(word string) []byte {
	if word == " " {
		return []byte(word)
	}

	if len(word) == 1 {
		return []byte(fmt.Sprint(word[0]))
	}

	if len(word) == 2 {
		return []byte(fmt.Sprint(word[0], string(word[1])))
	}

	str := []string{
		fmt.Sprint(word[0]),
		string(word[len(word)-1]),
		word[2 : len(word)-1],
		string(word[1]),
	}
	return []byte(strings.Join(str, ""))
}
