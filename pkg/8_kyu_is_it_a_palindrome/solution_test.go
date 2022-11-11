package vowel_count

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("a"), true)
		require.Equal(t, solution("aba"), true)
		require.Equal(t, solution("Abba"), true)
		require.Equal(t, solution("hello"), false)
	})
}
