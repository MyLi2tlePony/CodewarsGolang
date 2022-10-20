package string_ends_with

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("AAAA"), "TTTT")
		require.Equal(t, solution("ATTGC"), "TAACG")
		require.Equal(t, solution("GTAT"), "CATA")
	})
}
