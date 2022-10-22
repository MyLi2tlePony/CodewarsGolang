package int32_to_IPv4

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(2154959208), "128.114.17.104")
		require.Equal(t, solution(2149583361), "128.32.10.1")
		require.Equal(t, solution(0), "0.0.0.0")
		require.Equal(t, solution(math.MaxUint32), "255.255.255.255")
	})
}
