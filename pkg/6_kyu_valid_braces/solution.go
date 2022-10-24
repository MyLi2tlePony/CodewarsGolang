package valid_braces

var braces = map[rune]rune{')': '(', '}': '{', ']': '['}

func solution(str string) bool {
	bracesCounter := map[rune]int{'(': 0, ')': 0, '[': 0, ']': 0, '{': 0, '}': 0}
	openedBrace := make([]rune, 0)

	for _, brace := range str {
		if brace == '(' || brace == '[' || brace == '{' {
			openedBrace = append(openedBrace, brace)
		} else if len(openedBrace) > 0 {
			if b, ok := braces[brace]; ok && openedBrace[len(openedBrace)-1] != b {
				return false
			} else {
				openedBrace = openedBrace[:len(openedBrace)-1]
			}
		}

		if _, ok := bracesCounter[brace]; ok {
			bracesCounter[brace]++
		}
	}

	return bracesCounter['('] == bracesCounter[')'] && bracesCounter['['] == bracesCounter[']'] && bracesCounter['{'] == bracesCounter['}']
}
