package the_supermarket_queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution([]int{}, 1), 0)
	require.Equal(t, solution([]int{1, 2, 3, 4}, 1), 10)
	require.Equal(t, solution([]int{2, 2, 3, 3, 4, 4}, 2), 9)
	require.Equal(t, solution([]int{1, 2, 3, 4, 5}, 100), 5)
}
