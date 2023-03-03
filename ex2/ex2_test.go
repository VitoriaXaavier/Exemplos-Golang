package ex2

import "testing"

func TestCaneta ( t * testing.T) {
	
	caneta := Caneta(10,5,4)

	if caneta != 1.25 {
		t.Error("Erro no teste")
	}

	caneta2 := Caneta(100,4,6)

	if caneta2 != 16 {
		t.Error("Erro no teste 2")
	}
}