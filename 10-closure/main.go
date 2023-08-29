package main

import "fmt"

func inverterSinal(numero *int) {
	//fazendo desreferenciação e mudando valor
	*numero = *numero * -1
}

func main() {
	// passando o ponteiro do numero
	numero := 20
	inverterSinal(&numero)

	fmt.Println(numero)
}
