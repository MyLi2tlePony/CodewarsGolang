package alphabetical_addition

func solution(letters []rune) rune {
	if len(letters) == 0 {
		return 'z'
	}

	var result rune

	for _, let := range letters {
		result += let - 'a' + 1
	}

	return (result-1)%26 + 'a'
}
