package sort_numbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution([]int{1, 2, 10, 50, 5}), []int{1, 2, 5, 10, 50})
	require.Equal(t, solution([]int{}), []int{})
}
