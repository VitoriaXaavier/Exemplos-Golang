package ex3

import "fmt"

func MaximoDivisor (x, y int)  int {
	fmt.Println("Algoritimo para calcular o máximo divisor comum entre dois números inteiros!")

	println("Entre com o primeiro número")
	fmt.Scan(&x)
	println("Entre com o segundo número")
	fmt.Scan(&y)

	for y != 0 {

		r := x % y 
		x = y
		y = r

	}
	println("O máximo difisor comum de", x ,"e ", y, "é igual a ", x)
return x
}