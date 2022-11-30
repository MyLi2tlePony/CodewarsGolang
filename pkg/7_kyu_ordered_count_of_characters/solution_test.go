package ordered_count_of_characters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution("abracadabra"), []Tuple{{'a', 5}, {'b', 2}, {'r', 2}, {'c', 1}, {'d', 1}})
	require.Equal(t, solution("Code Wars"), []Tuple{{'C', 1}, {'o', 1}, {'d', 1}, {'e', 1}, {' ', 1}, {'W', 1}, {'a', 1}, {'r', 1}, {'s', 1}})
	require.Equal(t, solution(""), []Tuple{})
}
