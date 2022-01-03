package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialFor(t *testing.T) {

	t.Run("factorial for", func(t *testing.T) {
		assert.Equal(t, 1, FactorialFor(0))
		assert.Equal(t, 1, FactorialFor(1))
		assert.Equal(t, 2, FactorialFor(2))
		assert.Equal(t, 6, FactorialFor(3))
		assert.Equal(t, 24, FactorialFor(4))
	})

	t.Run("factorial req", func(t *testing.T) {
		assert.Equal(t, 1, FactorialReqursive(0))
		assert.Equal(t, 1, FactorialReqursive(1))
		assert.Equal(t, 2, FactorialReqursive(2))
		assert.Equal(t, 6, FactorialReqursive(3))
		assert.Equal(t, 24, FactorialReqursive(4))
	})

	t.Run("factorial req tail", func(t *testing.T) {
		assert.Equal(t, 1, FactorialReqursiveTail(1, 0))
		assert.Equal(t, 1, FactorialReqursiveTail(1, 1))
		assert.Equal(t, 2, FactorialReqursiveTail(1, 2))
		assert.Equal(t, 6, FactorialReqursiveTail(1, 3))
		assert.Equal(t, 24, FactorialReqursiveTail(1, 4))
	})
}
