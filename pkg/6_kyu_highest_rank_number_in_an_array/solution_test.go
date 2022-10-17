package highest_rank_number_in_an_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}), 12)
		require.Equal(t, solution([]int{12, 10, 8, 12, 7, 6, 4, 4, 4, 10, 12}), 12)
		require.Equal(t, solution([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10, 10}), 10)
		require.Equal(t, solution([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 13, 13, 13}), 13)
	})
}
