package bouncing_balls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(3, 0.66, 1.5), 3)
		require.Equal(t, solution(40, 0.4, 10), 3)
		require.Equal(t, solution(10, 0.6, 10), -1)
		require.Equal(t, solution(40, 1, 10), -1)
		require.Equal(t, solution(5, -1, 1.5), -1)
	})
}
