package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	// ayuda a realizar assertions
	c := require.New(t)

	result := Add(20, 2)

	expect := 22

	// if result != expect {
	// 	t.Errorf("got %q, expected %d", result, expect)
	// }

	c.Equal(expect, result)
}
