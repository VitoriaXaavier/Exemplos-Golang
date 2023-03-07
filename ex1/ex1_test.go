package ex1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {

	area := Area(20, 10)
	areaEsperada := 100

	assert.Equal(t, area, areaEsperada)

	area2 := Area(2, 4)
	areaEsperada = 4

	assert.Equal(t, area2, areaEsperada)
}
