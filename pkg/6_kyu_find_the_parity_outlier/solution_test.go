package find_the_parity_outlier

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution([]int{2, 6, 8, -10, 3}), 3)
		require.Equal(t, solution([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}), 206847684)
		require.Equal(t, solution([]int{math.MaxInt32, 0, 1}), 0)
	})
}
