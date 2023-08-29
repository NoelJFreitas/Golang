package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000
	fmt.Println(numero)

	var numero2 uint32 = 100000000
	fmt.Println(numero2)

	//RUNE = INT32
	var numero3 rune = 100000000
	fmt.Println(numero3)

	//BYTE = INT8
	var numero4 byte = 232
	fmt.Println(numero4)

	var numero5 float64 = 12345.54
	fmt.Println(numero5)

	var error error = errors.New("Erro interno")
	fmt.Println(error)
}
