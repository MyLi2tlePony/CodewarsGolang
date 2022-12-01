package character_with_longest_consecutive_repetition

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution("aaaabb"), Result{'a', 4})
	require.Equal(t, solution("bbbaaabaaaa"), Result{'a', 4})
	require.Equal(t, solution("cbdeuuu900"), Result{'u', 3})
	require.Equal(t, solution("abbbbb"), Result{'b', 5})
	require.Equal(t, solution("aabb"), Result{'a', 2})
	require.Equal(t, solution("ba"), Result{'b', 1})
	require.Equal(t, solution(""), Result{})
}
