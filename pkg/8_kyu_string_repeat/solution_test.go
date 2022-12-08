package string_repeat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	require.Equal(t, solution(4, "a"), "aaaa")
	require.Equal(t, solution(3, "hello "), "hello hello hello ")
	require.Equal(t, solution(2, "abc"), "abcabc")
}
