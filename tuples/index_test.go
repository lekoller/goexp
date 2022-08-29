package tuples_test

import (
	"testing"

	"github.com/lekoller/goexp/tuples"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	tuple := tuples.NewTuple(1, 3, 7, 8, 7, 5, 4, 6, 8, 5)
	x := tuple.Index(8)

	require.Equal(t, 3, x)
}
