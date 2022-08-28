package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestIsSuperset(t *testing.T) {
	x := sets.NewSet("a", "b", "c")
	y := sets.NewSet("f", "e", "d", "c", "b", "a")

	answer := x.IsSuperset(y)

	require.Equal(t, false, answer)

	// reverse
	answer = y.IsSuperset(x)

	require.Equal(t, true, answer)
}
