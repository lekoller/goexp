package lists_test

import (
	"testing"

	"github.com/lekoller/goexp/lists"
	"github.com/stretchr/testify/require"
)

func TestClear(t *testing.T) {
	x := lists.NewList("apple", "banana", "cherry", "orange")
	x.Clear()

	require.Equal(t, 0, len(x))
}
