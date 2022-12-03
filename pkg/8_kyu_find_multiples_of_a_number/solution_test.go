package find_multiples_of_a_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(5, 25), []int{5, 10, 15, 20, 25})
	require.Equal(t, solution(1, 2), []int{1, 2})
}
