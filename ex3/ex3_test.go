package ex3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisor(t *testing.T) {

	x := 18
	y := 15
	divisor := MaximoDivisor(x, y)
	divisorEsperado := 3

	assert.Equal(t, divisor, divisorEsperado)

	divisor2 := MaximoDivisor(20, 30)
	divisorEsperado = 10

	assert.Equal(t, divisor2, divisorEsperado)

}
