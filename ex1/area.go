package ex1

import "fmt"

func Area() {
var b int
var h int

fmt.Println("Digite o número da base do triangulo")
fmt.Scan(&b)
fmt.Println("Digite o número da altura do triangulo")
fmt.Scan(&h)

var s = (b* h) / 2

fmt.Println("A área do triangulo é de ", s)
}