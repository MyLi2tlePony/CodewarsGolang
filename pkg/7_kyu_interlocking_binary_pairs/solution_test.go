package interlocking_binary_pairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(9, 4), true)
		require.Equal(t, solution(3, 6), false)
		require.Equal(t, solution(2, 5), true)
		require.Equal(t, solution(7, 1), false)
		require.Equal(t, solution(0, 8), true)
	})
}
