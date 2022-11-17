package beginner_reduce_but_grow

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution([]int{1, 2, 3}), 6)
		require.Equal(t, solution([]int{4, 1, 1, 1, 4}), 16)
		require.Equal(t, solution([]int{2, 2, 2, 2, 2, 2}), 64)
	})
}
