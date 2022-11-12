package cats_and_shelves

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("  nAa"), false)
		require.Equal(t, solution("abcde"), true)
		require.Equal(t, solution("++-"), false)
		require.Equal(t, solution("AaBbC"), true)
	})
}
