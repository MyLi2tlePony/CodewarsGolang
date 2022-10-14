package bouncing_balls

func solution(h, bounce, window float64) int {
	if window >= h || h <= 0 || bounce >= 1 || bounce <= 0 {
		return -1
	}

	sawNumber := 1
	h *= bounce

	for h > window {
		sawNumber += 2
		h *= bounce
	}

	return sawNumber
}
