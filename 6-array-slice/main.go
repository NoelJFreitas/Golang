package main

import "fmt"

func main() {
	//arrays
	var array1 [5]int
	fmt.Println(array1)
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	fmt.Println("--------------------------------")

	var sliceVazio []int
	fmt.Println("sliceVazio: ", sliceVazio)
	fmt.Println("sliceVazio cap: ", cap(sliceVazio))
	fmt.Println("sliceVazio len: ", len(sliceVazio))
	sliceVazio = append(sliceVazio, 1)
	fmt.Println("sliceVazio adicionado: ", sliceVazio)
	fmt.Println("sliceVazio cap: ", cap(sliceVazio))
	fmt.Println("sliceVazio len: ", len(sliceVazio))

	fmt.Println("--------------------------------")
	//slices
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println("slice antes do append", slice)
	slice = append(slice, "f")
	fmt.Println("slice apos o append", slice)

	fmt.Println("--------------------------------")
	//apartir da posição 0 pega os proximos 5 itens
	slice2 := slice[0:5]
	fmt.Println("slice2", slice2)

	fmt.Println("--------------------------------")
	//array interno
	slice3 := make([]float32, 10, 15)
	fmt.Println("slice3: ", slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
