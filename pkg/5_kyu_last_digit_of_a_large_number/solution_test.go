package last_digit_of_a_large_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("4", "1"), 4)
		require.Equal(t, solution("4", "2"), 6)
		require.Equal(t, solution("9", "7"), 9)
		require.Equal(t, solution("10", "10000000000"), 0)
		require.Equal(t, solution("1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376"), 6)
		require.Equal(t, solution("3715290469715693021198967285016729344580685479654510946723", "68819615221552997273737174557165657483427362207517952651"), 7)
	})
}
