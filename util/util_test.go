package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	assert.Panics(t, func() { Check(errors.New("Test")) })
	assert.NotPanics(t, func() { Check(nil) })
}
