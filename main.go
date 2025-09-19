package main

import "fmt"

func main() {
	var peso, altura float64

	fmt.Print("Digite o seu peso (em kg): ")
	fmt.Scan(&peso)

	fmt.Print("Digite a sua altura (em metros): ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)
	fmt.Printf("Seu IMC é: %.2f", imc)
}
