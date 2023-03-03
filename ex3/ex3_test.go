package ex3

import "testing"

func Testdivisor (t testing.T) {

	divisor := MaximoDivisor(18,15)

	if divisor != 3 {
		t.Error("Erro no teste")
	}

	divisor2 := MaximoDivisor(20,30)

	if divisor2 != 10 {
		t.Error("Erro no teste 2")
	}


}