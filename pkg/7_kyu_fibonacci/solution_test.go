package fibonacci

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(1), 1)
		require.Equal(t, solution(2), 1)
		require.Equal(t, solution(3), 2)
		require.Equal(t, solution(4), 3)
		require.Equal(t, solution(5), 5)
	})
}
