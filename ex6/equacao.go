package ex6

import (
	"fmt"
	"math"
)

func EquacaoSegundoGrau (a,b,c float64)(float64, float64){

	fmt.Println("Entre com o valor de A")
	fmt.Scan(&a)
	fmt.Println("Entre com o valor de B")
	fmt.Scan(&b)
	fmt.Println("Entre com o valor de C")
	fmt.Scan(&c)

	 

		if a == 0 {
			fmt.Println("Não é uma equação de segundo grau")
			
		} 
			delta := (b*b) - 4*a*c
			
			
			if delta < 0 {
				fmt.Println("Não existem raízes reais")
			} 
				r1 := (-b + math.Sqrt(delta)) / (2*a)
				r2 := (-b - math.Sqrt(delta)) / (2*a)
				fmt.Println("O valor de r1 é de ", r1, " o valor de r2 é de ", r2)
			
				return r1, r2
			
		} 



