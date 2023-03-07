package ex2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaneta(t *testing.T) {

	caneta := Caneta(10, 5, 4)
	canetaEsperada := 1.25

	// compara se os valores são iguais
	
	assert.Equal(t, caneta, canetaEsperada)

	caneta2 := Caneta(100, 4, 6)
	canetaEsperada = 16

	assert.Equal(t, caneta2, canetaEsperada)
}
