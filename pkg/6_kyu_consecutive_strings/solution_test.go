package create_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2),
			"abigailtheta")

		require.Equal(t, solution([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1),
			"oocccffuucccjjjkkkjyyyeehh")

		require.Equal(t, solution([]string{}, 3),
			"")

		require.Equal(t, solution([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2),
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	})
}
