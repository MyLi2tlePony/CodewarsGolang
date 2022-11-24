package sum_of_angles

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(3), 180)
	require.Equal(t, solution(4), 360)
}
