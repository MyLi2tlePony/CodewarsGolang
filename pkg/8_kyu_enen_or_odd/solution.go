package even_or_odd

func solution(number int) string {
	if number&1 == 1 {
		return "Odd"
	}

	return "Even"
}
