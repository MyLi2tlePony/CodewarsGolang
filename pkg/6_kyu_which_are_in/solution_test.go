package which_are_in

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("tests", func(t *testing.T) {
		var a1 = []string{"live", "arp", "strong"}
		var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
		var r = []string{"arp", "live", "strong"}
		require.Equal(t, solution(a1, a2), r)

		a1 = []string{"cod", "code", "wars", "ewar", "ar"}
		a2 = []string{}
		r = []string{}
		require.Equal(t, solution(a1, a2), r)
	})
}
