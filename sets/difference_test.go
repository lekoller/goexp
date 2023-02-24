package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestDifference(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry")

	result := x.Difference(y, z)

	require.Equal(t, 1, len(result))
	require.Equal(t, "banana", result[0])
}

func TestDifference_WithDifferentType(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet(2, 4, 5)
	z := sets.NewSet("peach", "cherry")

	result := x.Difference(y, z)

	require.Equal(t, 2, len(result))
	require.Equal(t, "apple", result[0])
	require.Equal(t, "banana", result[1])
}

func TestDifferenceUpdate(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry")

	x.DifferenceUpdate(y, z)

	require.Equal(t, 1, len(x))
	require.Equal(t, "banana", x[0])
}
