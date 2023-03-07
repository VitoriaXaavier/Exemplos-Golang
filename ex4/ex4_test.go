package ex4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiplicacao ( t * testing.T) {

	multiplicação := CalculadoraMultiplicacao(2, 3)
	multiplicaçãoEsperada := 6

	assert.Equal(t, multiplicação, multiplicaçãoEsperada)

	mult := CalculadoraMultiplicacao(2,2)
	multiplicaçãoEsperada = 4

	assert.Equal(t, mult, multiplicaçãoEsperada)
}

func TestDivisao ( t * testing.T) {
	divisao, _ := CalculadoraDivisao(10,2)
	divisaoEsperada := 5

	assert.Equal(t, divisao,divisaoEsperada)

	div, _ := CalculadoraDivisao(30, 3)

	divisaoEsperada = 10
	assert.Equal(t, div, divisaoEsperada)

}
	func TestDiv0(t * testing.T) {
		_ , err := CalculadoraDivisao(0,2)

		assert.NotNil(t, err)
		assert.Equal(t, ZeroOperationError, err)
	}