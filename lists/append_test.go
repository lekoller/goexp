package lists_test

import (
	"reflect"
	"testing"

	"github.com/lekoller/goexp/lists"
	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	list := lists.NewList(1, 2, 3)
	list.Append(4)
	expected := lists.NewList(1, 2, 3, 4)
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, got %v", expected, list)
	}

	intern := lists.List{"asa", 3}

	list.Append(intern)

	require.Contains(t, list, intern)
}
