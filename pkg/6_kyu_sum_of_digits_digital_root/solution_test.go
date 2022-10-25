package sum_of_digits_digital_root

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(16), 7)
		require.Equal(t, solution(195), 6)
		require.Equal(t, solution(992), 2)
		require.Equal(t, solution(167346), 9)
		require.Equal(t, solution(0), 0)
	})
}
