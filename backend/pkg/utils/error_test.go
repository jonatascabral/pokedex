package utils

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicOnError_WithPanic(t *testing.T) {
	err := errors.New("I'm an error")
	assert.Panics(t, func() {
		PanicOnError(err)
	})
}

func TestPanicOnError_WithoutPanic(t *testing.T) {
	var err error
	assert.NotPanics(t, func() {
		PanicOnError(err)
	})
}
