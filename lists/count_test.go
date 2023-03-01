package lists_test

import (
	"testing"

	"github.com/lekoller/goexp/lists"
	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	list := lists.NewList(1, 3, 7, 8, 7, 5, 4, 6, 8, 5)
	x := list.Count(7)

	require.Equal(t, 2, x)
}
