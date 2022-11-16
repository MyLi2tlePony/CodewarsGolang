package sum_of_a_sequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(2, 6, 2), 12)
		require.Equal(t, solution(1, 5, 1), 15)
		require.Equal(t, solution(1, 5, 3), 5)
		require.Equal(t, solution(0, 15, 3), 45)
		require.Equal(t, solution(16, 15, 3), 0)
		require.Equal(t, solution(2, 24, 22), 26)
		require.Equal(t, solution(2, 2, 2), 2)
		require.Equal(t, solution(2, 2, 1), 2)
		require.Equal(t, solution(1, 15, 3), 35)
		require.Equal(t, solution(15, 1, 3), 0)
	})
}
