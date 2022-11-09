package string_ends_with

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("8 3 -5 42 -1 0 0 -9 4 7 4 -4"), "42 -9")
		require.Equal(t, solution("1 2 3"), "3 1")
	})
}
