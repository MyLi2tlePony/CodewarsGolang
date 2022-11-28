package find_the_smallest_power_higher_than_a_given_a_value

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(12385, 3), 13824)
	require.Equal(t, solution(1245678, 5), 1419857)
	require.Equal(t, solution(1245678, 6), 1771561)
}
