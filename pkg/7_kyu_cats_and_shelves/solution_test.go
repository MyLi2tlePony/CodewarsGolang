package cats_and_shelves

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(1, 5), 2)
		require.Equal(t, solution(1, 1), 0)
		require.Equal(t, solution(2, 5), 1)
		require.Equal(t, solution(2, 4), 2)
	})
}
