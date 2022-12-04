package third_angle_of_a_triangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(30, 60), 90)
		require.Equal(t, solution(60, 60), 60)
		require.Equal(t, solution(43, 78), 59)
		require.Equal(t, solution(10, 20), 150)
	})
}
