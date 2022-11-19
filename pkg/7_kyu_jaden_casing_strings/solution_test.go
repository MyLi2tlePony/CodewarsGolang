package jaden_casing_strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("most trees are blue"), "Most Trees Are Blue")
		require.Equal(t, solution("All the rules in this world were made by someone no smarter than you. So make your own."), "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own.")
		require.Equal(t, solution("When I die. then you will realize"), "When I Die. Then You Will Realize")
		require.Equal(t, solution("Jonah Hill is a genius"), "Jonah Hill Is A Genius")
		require.Equal(t, solution("Dying is mainstream"), "Dying Is Mainstream")
	})
}
