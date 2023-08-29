package main

import "fmt"

func closure() func() {
	texto := "dentro da função closure"

	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
