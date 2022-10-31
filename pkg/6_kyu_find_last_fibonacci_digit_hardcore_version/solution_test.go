package find_last_fibonacci_digit_hardcore_version

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(1), 1)
		require.Equal(t, solution(21), 6)
		require.Equal(t, solution(302), 1)
		require.Equal(t, solution(4003), 7)
		require.Equal(t, solution(50004), 8)
		require.Equal(t, solution(600005), 5)
		require.Equal(t, solution(7000006), 3)
		require.Equal(t, solution(80000007), 8)
		require.Equal(t, solution(900000008), 1)
		require.Equal(t, solution(1000000009), 9)
	})
}
