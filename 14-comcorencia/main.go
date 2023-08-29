package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORENCIA NÂO É IGUAL A PARALELISMO
	// PARALELISMO = cada atividade roda em uma thead separada no computado, rodando ao mesmo tempo
	// CONCORENCIA = o nucleo tenta executar as duas tarefas ao mesmo tempo na mesma thread

	// CONCORENCIA utilizando a palavra reservada "go"
	go escrever("Olá Mundo!")
	escrever("programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
