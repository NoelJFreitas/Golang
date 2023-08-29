package main

import "fmt"

type forma interface {
	area() float64
}

func escreveArea(f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func teste(t interface{}) {

}

func main() {
	r := retangulo{10, 23}

	escreveArea(r)
}
