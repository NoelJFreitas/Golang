package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u *usuario) salvar() {
	fmt.Println("Salvando...")
	fmt.Printf("usuario: %s, salvo com sucesso!", u.nome)
}

func main() {
	usuario1 := usuario{nome: "Noel", idade: 22}
	usuario1.salvar()

}
