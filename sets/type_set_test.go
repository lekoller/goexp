package sets_test

import (
	"testing"

	"github.com/lekoller/goexp/sets"
	"github.com/stretchr/testify/require"
)

func TestNewSet(t *testing.T) {
	set := sets.NewSet("like", "dislike", "like", "subscribe")

	require.Equal(t, 3, len(set))
	require.Equal(t, "like", set[0])
	require.Equal(t, "dislike", set[1])
	require.Equal(t, "subscribe", set[2])
}
