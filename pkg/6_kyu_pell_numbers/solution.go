package pell_numbers

import "math/big"

var pell = map[int]*big.Int{0: big.NewInt(0), 1: big.NewInt(1)}

func solution(n int) *big.Int {
	if p, ok := pell[n]; ok {
		return p
	}

	if _, ok := pell[n-1]; !ok {
		pell[n-1] = solution(n - 1)
	}

	if _, ok := pell[n-2]; !ok {
		pell[n-2] = solution(n - 2)
	}

	b := big.NewInt(0)
	return b.Add(b.Mul(big.NewInt(int64(2)), solution(n-1)), solution(n-2))
}
