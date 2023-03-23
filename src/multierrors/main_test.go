package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errFirst  = errors.New("first-error")
	errSecond = errors.New("second-error")
	errThird  = errors.New("third-error")
)

func TestFirstCase_FAIL(t *testing.T) {
	err := fmt.Errorf("%w, %v, %v", errFirst, errSecond, errThird)

	assert.ErrorIs(t, err, errFirst)
	assert.ErrorIs(t, err, errSecond)
	assert.ErrorIs(t, err, errThird)

	expectedStrRepr := "first-error, second-error, third-error"
	assert.Equal(t, expectedStrRepr, err.Error())
}

func TestSecondCase_SUCCESS(t *testing.T) {
	err1 := errFirst
	err2 := fmt.Errorf("%w", err1)
	err3 := fmt.Errorf("%w", err2)

	unwrappedErr3 := errors.Unwrap(err3)
	assert.ErrorIs(t, err2, unwrappedErr3)

	unwrappedErr2 := errors.Unwrap(unwrappedErr3)
	assert.ErrorIs(t, err1, unwrappedErr2)
}

func TestThirdCase_SUCCESS(t *testing.T) {
	err := fmt.Errorf("%w, %w, %w", errFirst, errSecond, errThird)

	assert.ErrorIs(t, err, errFirst)
	assert.ErrorIs(t, err, errSecond)
	assert.ErrorIs(t, err, errThird)

	expectedStrRepr := "first-error, second-error, third-error"
	assert.Equal(t, expectedStrRepr, err.Error())
}

func TestForthCase_SUCCESS(t *testing.T) {
	err := errors.Join(errFirst, errSecond, errThird)

	assert.ErrorIs(t, err, errFirst)
	assert.ErrorIs(t, err, errSecond)
	assert.ErrorIs(t, err, errThird)

	expectedStrRepr := `first-error
second-error
third-error`
	assert.Equal(t, expectedStrRepr, err.Error())
}
