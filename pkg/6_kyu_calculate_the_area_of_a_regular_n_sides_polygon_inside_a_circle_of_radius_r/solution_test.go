package calculate_the_area_of_a_regular_n_sides_polygon_inside_a_circle_of_radius_r

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(3, 3), 11.691)
		require.Equal(t, solution(5.8, 7), 92.053)
		require.Equal(t, solution(4, 5), 38.042)
	})
}
