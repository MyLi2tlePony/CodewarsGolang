package split_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("abc"), []string{"ab", "c_"})
		require.Equal(t, solution("dawsd"), []string{"da", "ws", "d_"})
		require.Equal(t, solution("awsaws"), []string{"aw", "sa", "ws"})
	})
}
