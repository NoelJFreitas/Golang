package main

import "fmt"

func main() {
	usario := map[string]string{
		"name":      "noel",
		"sobrenome": "Freitas",
	}
	fmt.Println(usario["name"])
}
