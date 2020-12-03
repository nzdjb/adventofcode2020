package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	assert.Panics(t, func() { Check(errors.New("Test")) })
	assert.NotPanics(t, func() { Check(nil) })
}
