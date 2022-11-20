package not_very_secure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(".*?"), false)
		require.Equal(t, solution("a"), true)
		require.Equal(t, solution("Mazinkaiser"), true)
		require.Equal(t, solution("hello world_"), false)
		require.Equal(t, solution("PassW0rd"), true)
		require.Equal(t, solution("     "), false)
		require.Equal(t, solution(""), false)
		require.Equal(t, solution("\n\t\n"), false)
		require.Equal(t, solution("ciao\n$$_"), false)
		require.Equal(t, solution("__ * __"), false)
		require.Equal(t, solution("&))((("), false)
		require.Equal(t, solution("43534h56jmTHHF3k"), true)
	})
}
