package tuples_test

import (
	"testing"

	"github.com/lekoller/goexp/tuples"
	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	tuple := tuples.NewTuple(1, 3, 7, 8, 7, 5, 4, 6, 8, 5)
	x := tuple.Count(7)

	require.Equal(t, 2, x)
}
