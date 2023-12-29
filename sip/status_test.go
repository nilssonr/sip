package sip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	assert.Equal(t, "Trying", StatusText(StatusTrying))
	assert.Equal(t, "", StatusText(99999))
}
