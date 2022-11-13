package multiply_word_in_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		require.Equal(t, solution("This is a string", 3, 5),
			"string-string-string-string-string")

		require.Equal(t, solution("Creativity is the process of having original ideas that have value. It is a process; it's not random.", 8, 10),
			"that-that-that-that-that-that-that-that-that-that")

		require.Equal(t, solution("Self-control means wanting to be effective at some random point in the infinite radiations of my spiritual existence", 1, 1),
			"means")

		require.Equal(t, solution("Is sloppiness in code caused by ignorance or apathy? I don't know and I don't care.", 6, 8),
			"ignorance-ignorance-ignorance-ignorance-ignorance-ignorance-ignorance-ignorance")

		require.Equal(t, solution("Everything happening around me is very random. I am enjoying the phase, as the journey is far more enjoyable than the destination.", 2, 5),
			"around-around-around-around-around")
	})
}
