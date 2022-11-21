package surrounding_primes_for_a_value

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(100), [2]int{97, 101})
	require.Equal(t, solution(97), [2]int{89, 101})
	require.Equal(t, solution(101), [2]int{97, 103})
	require.Equal(t, solution(120), [2]int{113, 127})
	require.Equal(t, solution(130), [2]int{127, 131})
}
