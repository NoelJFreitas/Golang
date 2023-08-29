package main

import "fmt"

func alunoAprovado(n1, n2 float32) bool {
	//DEFER = ADIAR, adia a função até o ultimo momento possivel
	//No caso abaixo a função executa antes de qualquer dos dois retornos
	defer fmt.Println("Média calculada, o resultado sera retornado!")

	media := (n1 + n2) / 2
	if media >= 7 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoAprovado(7, 8))
}
