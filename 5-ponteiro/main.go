package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	//copia o valor da variavel1 para variavel2
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	fmt.Println("-----")

	//variavel4 aponta para o mesmo valor de variavel3, por isso muda o valor se a variavel3 for alterada
	var variavel3 int = 10
	var ponteiro *int = &variavel1
	fmt.Println(variavel3, *ponteiro)
	variavel3++
	fmt.Println(variavel3, *ponteiro)

}
