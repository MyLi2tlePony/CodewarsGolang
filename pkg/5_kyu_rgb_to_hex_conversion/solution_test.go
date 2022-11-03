package rgb_to_hex_conversion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(0, 0, 0), "000000")
		require.Equal(t, solution(1, 2, 3), "010203")
		require.Equal(t, solution(255, 255, 255), "FFFFFF")
		require.Equal(t, solution(254, 253, 252), "FEFDFC")
		require.Equal(t, solution(-20, 275, 125), "00FF7D")
	})
}
