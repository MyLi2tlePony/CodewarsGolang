package two_fighters_one_winner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"), "Lew")
		require.Equal(t, solution(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"), "Harry")
		require.Equal(t, solution(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"), "Harald")
		require.Equal(t, solution(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald"), "Harald")
		require.Equal(t, solution(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"), "Harald")
		require.Equal(t, solution(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald"), "Harald")
	})
}
