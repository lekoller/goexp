package lists_test

import (
	"testing"

	"github.com/lekoller/goexp/lists"
	"github.com/stretchr/testify/require"
)

func TestExtend(t *testing.T) {
	x := lists.NewList("apple", "banana", "cherry")
	y := lists.NewList("orange", "mango", "grapes")

	x.Extend(y)

	require.Equal(t, 6, len(x))
	require.Contains(t, x, "apple")
	require.Contains(t, x, "banana")
	require.Contains(t, x, "cherry")
	require.Contains(t, x, "orange")
	require.Contains(t, x, "mango")
	require.Contains(t, x, "grapes")
}
