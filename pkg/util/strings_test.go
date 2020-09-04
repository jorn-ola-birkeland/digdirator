package util_test

import (
	"github.com/nais/digdirator/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	c := "foo"
	t.Run("Empty list should return false", func(t *testing.T) {
		s := make([]string, 0)
		assert.False(t, util.ContainsString(s, c))
	})
	t.Run("List does not contain string should return false", func(t *testing.T) {
		s := []string{"bar"}
		assert.False(t, util.ContainsString(s, c))
	})
	t.Run("List contains string should return true", func(t *testing.T) {
		s := []string{"bar", c}
		assert.True(t, util.ContainsString(s, c))
	})
}

func TestRemoveString(t *testing.T) {
	c := "foo"
	t.Run("Empty list should empty list", func(t *testing.T) {
		s := make([]string, 0)
		assert.Empty(t, util.RemoveString(s, c))
	})
	t.Run("List does not contain string should return same list", func(t *testing.T) {
		s := []string{"bar"}
		assert.Equal(t, s, util.RemoveString(s, c))
	})
	t.Run("List contains string should list without element", func(t *testing.T) {
		s := []string{"bar", c}
		expected := []string{"bar"}
		assert.Equal(t, expected, util.RemoveString(s, c))
	})
}
