package lists_test

import (
	"testing"

	"github.com/lekoller/goexp/lists"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	x := lists.NewList("apple", "banana", "cherry")
	y := x.Copy()

	require.Equal(t, 3, len(y))
	require.Contains(t, y, "apple")
	require.Contains(t, y, "banana")
	require.Contains(t, y, "cherry")
}
