package prsistent_bugger

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(39), 3)
		require.Equal(t, solution(4), 0)
		require.Equal(t, solution(25), 2)
		require.Equal(t, solution(999), 4)
	})
}
