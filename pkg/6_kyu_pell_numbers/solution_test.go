package pell_numbers

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(1), big.NewInt(1))
	require.Equal(t, solution(2), big.NewInt(2))
	require.Equal(t, solution(3), big.NewInt(5))
	require.Equal(t, solution(4), big.NewInt(12))
}
