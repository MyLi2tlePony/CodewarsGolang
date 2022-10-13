package string_ends_with

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("", ""), true)
		require.Equal(t, solution(" ", ""), true)
		require.Equal(t, solution("abc", "c"), true)
		require.Equal(t, solution("banana", "ana"), true)
		require.Equal(t, solution("a", "z"), false)
		require.Equal(t, solution("", "t"), false)
		require.Equal(t, solution("$a = $b + 1", "+1"), false)
		require.Equal(t, solution("    ", "   "), true)
		require.Equal(t, solution("1oo", "100"), false)
	})
}
