package even_or_odd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(1), "Odd")
	require.Equal(t, solution(2), "Even")
	require.Equal(t, solution(-1), "Odd")
	require.Equal(t, solution(-2), "Even")
}
