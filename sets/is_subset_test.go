package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestIsSubset(t *testing.T) {
	x := sets.NewSet("a", "b", "c")
	y := sets.NewSet("f", "e", "d", "c", "b", "a")
	j := sets.NewSet("f", "e", "d", "c", "b", "b")

	answer := x.IsSubset(y)

	require.Equal(t, true, answer)

	answer = x.IsSubset(j)

	require.Equal(t, false, answer)

	// reverse
	answer = y.IsSubset(x)

	require.Equal(t, false, answer)
}
