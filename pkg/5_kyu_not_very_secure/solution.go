package not_very_secure

func solution(str string) bool {
	if str == "" {
		return false
	}

	for _, char := range str {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'a',
			'B', 'b', 'C', 'c', 'D', 'd', 'E', 'e', 'F', 'f', 'G', 'g', 'H', 'h', 'I', 'i', 'J', 'j',
			'K', 'k', 'L', 'l', 'M', 'm', 'N', 'n', 'O', 'o', 'P', 'p', 'Q', 'q', 'R', 'r', 'S', 's',
			'T', 't', 'U', 'u', 'V', 'v', 'W', 'w', 'X', 'x', 'Y', 'y', 'Z', 'z':
		default:
			return false

		}
	}

	return true
}
