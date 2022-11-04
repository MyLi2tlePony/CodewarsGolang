package bouncing_balls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		require.Equal(t, solution(a1, a2), true)

		a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
		a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		require.Equal(t, solution(a1, a2), false)

		a1 = nil
		a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
		require.Equal(t, solution(a1, a2), false)
	})
}
