package count_ip_addresses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("10.0.0.0", "10.0.0.50"), 50)
		require.Equal(t, solution("20.0.0.10", "20.0.1.0"), 246)
	})
}
