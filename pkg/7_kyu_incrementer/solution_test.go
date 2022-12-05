package incrementer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution([]int{}), []int{})
	require.Equal(t, solution([]int{1, 2, 3}), []int{2, 4, 6})
	require.Equal(t, solution([]int{4, 6, 7, 1, 3}), []int{5, 8, 0, 5, 8})
	require.Equal(t, solution([]int{3, 6, 9, 8, 9}), []int{4, 8, 2, 2, 4})
	require.Equal(t, solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8}), []int{2, 4, 6, 8, 0, 2, 4, 6, 8, 9, 0, 1, 2, 2})
}
