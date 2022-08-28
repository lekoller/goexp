package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestIntersection_Level1(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")

	result := x.Intersection(y)

	require.Equal(t, 1, len(result.Data))
	require.Equal(t, "apple", result.Data[0])
}

func TestIntersection_Level2(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry", "apple")

	result := x.Intersection(y, z)

	require.Equal(t, 1, len(result.Data))
	require.Equal(t, "apple", result.Data[0])
}

func TestIntersection_Level3(t *testing.T) {
	x := sets.NewSet("apple", "banana", "cherry")
	y := sets.NewSet("google", "microsoft", "apple")
	z := sets.NewSet("peach", "cherry", "apple")
	j := sets.NewSet("samsung", "apple")

	result := x.Intersection(y, z, j)

	require.Equal(t, 1, len(result.Data))
	require.Equal(t, "apple", result.Data[0])
}
