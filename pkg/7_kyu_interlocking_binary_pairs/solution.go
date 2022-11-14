package interlocking_binary_pairs

func solution(a, b uint64) bool {
	for a > 0 && b > 0 {
		if (a&1 == 1) && (b&1 == 1) {
			return false
		}

		a >>= 1
		b >>= 1
	}

	return true
}
