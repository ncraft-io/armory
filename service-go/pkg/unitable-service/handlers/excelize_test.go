package handlers

import (
	"github.com/zeebo/assert"
	"testing"
)

func TestGetColumnIndex(t *testing.T) {
	i := getColIndex(26)
	assert.Equal(t, "AA", i)
	assert.Equal(t, "B", getColIndex(1))
}
