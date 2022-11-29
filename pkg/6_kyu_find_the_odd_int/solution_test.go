package find_the_odd_int

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}), 5)
	require.Equal(t, solution([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}), -1)
	require.Equal(t, solution([]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}), 5)
	require.Equal(t, solution([]int{10}), 10)
	require.Equal(t, solution([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}), 10)
	require.Equal(t, solution([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}), 1)
}
