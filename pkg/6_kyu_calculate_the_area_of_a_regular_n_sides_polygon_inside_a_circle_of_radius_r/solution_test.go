package create_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}), "(123) 456-7890")
		require.Equal(t, solution([10]uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}), "(111) 111-1111")
		require.Equal(t, solution([10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), "(000) 000-0000")
	})
}
