package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestUnion(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry")

	union := x.Union(y, z)

	require.Equal(t, 6, len(union))
}
