package ex2

import (
	"fmt"
)

// Compraram-se canetasCompradas canetas iguais, que foram pagas com
//uma nota de valorPago reais, obtendo-se troco reais de troco. Quanto custou cada caneta?
func Caneta (valorPago, troco float64, canetasCompradas int) float64 {

	fmt.Println("Entre com a quantidade de canetas compradas")
	fmt.Scan(&canetasCompradas)
	fmt.Println("Qual o valor da nota dada para pagar")
	fmt.Scan(&valorPago)
	fmt.Println("Qual foi o valor do troco recebido?")
	fmt.Scan(&troco)
	nConvertido := float64(canetasCompradas)
	var valorDeCadaCaneta = (valorPago - troco) / nConvertido

	if valorPago > troco && canetasCompradas > 0 && valorPago > 0 && troco >= 0 {
		fmt.Printf("O valor unitário da caneta é de %g ",valorDeCadaCaneta, "\n")
		}else {
			fmt.Println("Valores inválidos")
		}
		return (valorDeCadaCaneta)
}