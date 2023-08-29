package main

import "fmt"

func recuperarExecucao() {
	//funcçãp recover que tenta reucuperar o codigo
	if r := recover(); r != nil {
		fmt.Println("tentativa de recuperar execução. Sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//lança um erro e mata o programa caso vc não use o "recover"
	panic("A MÉDIA É 6!!!!")

}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("fim execução!")
}
