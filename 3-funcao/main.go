package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	fmt.Println("Sum 1", sum(20, 30))

	fmt.Println("............")

	var1, var2 := sum2(20, 30)
	fmt.Println("Sum 2", var1, var2)

	fmt.Println("............")

	fmt.Println(sum3())

}

func sum(a, b int) int {
	return a + b
}

func sum2(a, b int) (soma, subtracao int) {
	soma = a + b
	subtracao = a - b
	return
}

func sum3(numeros ...int) int {
	total := 0

	for _, n := range numeros {
		total += n
	}

	return total
}
