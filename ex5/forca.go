package ex5

import (
	"fmt"
	"math"
)

func ForcaExercida (altura, diametro, gama float64) float64 {

	fmt.Println("Entre com o valor da altura")
	fmt.Scan(&altura)
	fmt.Println("Entre com o valor do diametro")
	fmt.Scan(&diametro)
	fmt.Println("Entre com o valor gama")
	fmt.Scan(&gama)
	
	pi := 3.14159
	exp2 := math.Pow(diametro, 2)

	f := (pi * gama * exp2 * altura) / 4

	fmt.Println("A força exercida pela coluna de um liquído é de ", f)

	return f
}