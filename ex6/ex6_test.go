package ex6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquacaoSegundoGrau(t * testing.T) {


// teste 1
	a := 1.0
	b := -5.0
	c := 6.0
	r1Esperado := 3.0
	r2Esperado := 2.0

	r1,r2 := EquacaoSegundoGrau(a,b,c)

	assert.Equal(t, r1, r1Esperado, r2, r2Esperado)
		
	

// teste 2 

	a = 2.0
	b = 4.0
	c = 2.0
	r1Esperado = -1.0
	r2Esperado = -1.0

	r1,r2 = EquacaoSegundoGrau(a,b,c)

	assert.Equal(t, r1, r1Esperado, r2, r2Esperado)

}

