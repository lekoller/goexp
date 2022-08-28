package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestIsDisjoint(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	j := sets.NewSet("google", "microsoft")

	answer := x.IsDisjoint(y)

	require.Equal(t, false, answer)

	answer = x.IsDisjoint(j)

	require.Equal(t, true, answer)
}
