package take_a_number_and_sum_its_digits_raised_to_the_consecutive_powers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(1, 10), []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		require.Equal(t, solution(1, 100), []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89})
		require.Equal(t, solution(10, 89), []uint64{89})
		require.Equal(t, solution(10, 100), []uint64{89})
		require.Nil(t, solution(90, 100))
		require.Equal(t, solution(89, 135), []uint64{89, 135})
		require.Equal(t, solution(12157692622039623404, 12157692622039625655), []uint64{12157692622039623539})
	})
}
