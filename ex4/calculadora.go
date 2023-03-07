package ex4

import (
	"errors"
	"fmt"
)

var ZeroOperationError = errors.New("divisor ou dividendo não podem ser 0")

func CalculadoraMultiplicacao (fator1 int, fator2 int) int {

	fmt.Println("Entre com o primeiro fator")
	fmt.Scan(&fator1)
	fmt.Println("Entre com o segundo fator")
	fmt.Scan(&fator2)

	produto := fator1 * fator2

	fmt.Println("O resultado da multiplicação de ", fator1, " * ", fator2, " = ", produto)
	return produto

}


func CalculadoraDivisao (dividendo, divisor int) (int, error) {

	fmt.Println("Entre com o dividendo")
	fmt.Scan(&dividendo)
	fmt.Println("Entre com o divisor")
	fmt.Scan(&divisor)
	

	var quociente int

	if dividendo == 0 || divisor == 0 {
		return 0, ZeroOperationError
	
	}else { 
		quociente = dividendo / divisor
		fmt.Println("O resultado da divisão de ", dividendo, " / ", divisor, " = ", quociente)
	}

	return quociente, nil

}