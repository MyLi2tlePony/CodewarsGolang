package vowel_count

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("1234"), 1234)
		require.Equal(t, solution("605"), 605)
		require.Equal(t, solution("1405"), 1405)
		require.Equal(t, solution("-7"), -7)
	})
}
