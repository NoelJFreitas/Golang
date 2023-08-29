package main

import "fmt"

type usuario struct {
	name string
	age  uint8
	sex  string
}

type estudante struct {
	usuario   // forma de "Herança"
	matricula string
	curso     string
}

func main() {
	fmt.Println("about structs")

	var user1 usuario
	user1.name = "maria"
	user1.age = 25
	user1.sex = "female"
	fmt.Println(user1)

	var user2 usuario = usuario{name: "noel", age: 22, sex: "masculine"}
	fmt.Println(user2)

	user3 := usuario{name: "joão", age: 42, sex: "masculine"}
	fmt.Println(user3)

	estudante := estudante{usuario: usuario{name: "pedro", age: 42, sex: "masculine"}, matricula: "14887878", curso: "eng"}
	fmt.Println(estudante)
}
