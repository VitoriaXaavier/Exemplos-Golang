package ex1

import "fmt"

func Area(b int, h int) int {

fmt.Println("Programa que calcula a área de um triângulo!")

fmt.Println("Digite o número da base do triangulo")
fmt.Scan(&b)
fmt.Println("Digite o número da altura do triangulo")
fmt.Scan(&h)

var s int = (b* h) / 2

fmt.Println("Dado a base do triangulo de ", b, " e altura de ", h, "temos que a área é de ", s)

return s
}