package ex4

import "testing"

func TestMultiplicacao ( t * testing.T) {

	multiplicação := CalculadoraMultiplicacao(2, 3)

	if multiplicação != 6 {
		t.Error("Erro na multiplicação de 2 * 3")
	}


	mult := CalculadoraMultiplicacao(2,2)
	 
	if mult != 4 {
		t.Error("Erro na multiplicação de 2 * 2")
	}

}

func TestDivisao ( t * testing.T) {

	divisao, _ := CalculadoraDivisao(10,2)

	if divisao != 5 {
		t.Error("Erro na divisão de 10 / 2")
	}

	div, _ := CalculadoraDivisao(30, 3)

	if div != 10 {
		t.Error("Erro na divisão de 30 / 3")
	}

}
	func TestDiv0(t * testing.T) {
		
		div, err := CalculadoraDivisao(0,2)

		if div != 5 {
			t.Error("Erro! Não pode haver divisor ou dividendos = 0", err)
		
		}
		
		
	}
	