package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo!", canal)
	fmt.Println("depois de começar ")

	mensagem := <-canal
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		fmt.Println("rodou:", i)
		canal <- texto
		time.Sleep(time.Second)
	}
}
