package alphabetical_addition

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution([]rune{'a', 'b', 'c'}), 'f')
	require.Equal(t, solution([]rune{'z'}), 'z')
	require.Equal(t, solution([]rune{'a', 'b'}), 'c')
	require.Equal(t, solution([]rune{'c'}), 'c')
	require.Equal(t, solution([]rune{'z', 'a'}), 'a')
	require.Equal(t, solution([]rune{'y', 'c', 'b'}), 'd')
	require.Equal(t, solution([]rune{}), 'z')
}
