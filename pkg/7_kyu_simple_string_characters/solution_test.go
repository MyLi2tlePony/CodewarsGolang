package string_ends_with

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("Codewars@codewars123.com"), []int{1, 18, 3, 2})
		require.Equal(t, solution("bgA5<1d-tOwUZTS8yQ"), []int{7, 6, 3, 2})
		require.Equal(t, solution("P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H"), []int{9, 9, 6, 9})
		require.Equal(t, solution("RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD"), []int{15, 8, 6, 9})
		require.Equal(t, solution("$Cnl)Sr<7bBW-&qLHI!mY41ODe"), []int{10, 7, 3, 6})
	})
}
