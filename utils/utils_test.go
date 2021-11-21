package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSolvable(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsSolvable([]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 14, 0}))
	assert.True(IsSolvable([]uint8{12, 1, 10, 2, 7, 11, 4, 14, 5, 0, 9, 15, 8, 13, 6, 3}))
}
