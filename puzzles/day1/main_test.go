package main

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	assert.Panics(t, func() { check(errors.New("Test")) })
	assert.NotPanics(t, func() { check(nil) })
}

// func TestDiffFrom2020(t *testing.T) {
// 	val, err := DiffFrom2020(1903)

// 	assert.Equal(117, val)
// }