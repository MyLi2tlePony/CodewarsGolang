package count_the_divisors_of_a_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(1), 1)
	require.Equal(t, solution(10), 4)
	require.Equal(t, solution(11), 2)
	require.Equal(t, solution(54), 8)
	require.Equal(t, solution(64), 7)
}
