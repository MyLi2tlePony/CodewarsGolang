package valid_braces

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("(){}[]"), true)
		require.Equal(t, solution("()"), true)
		require.Equal(t, solution("([{}])"), true)
		require.Equal(t, solution("(}"), false)
		require.Equal(t, solution("[(])"), false)
		require.Equal(t, solution("[({)](]"), false)
	})
}
