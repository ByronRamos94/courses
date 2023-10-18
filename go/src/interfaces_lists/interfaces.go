package main

import (
	"fmt"
)

type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

type figura interface {
	area() float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}
func (r rectangulo) area() float64 {
	return r.base * r.altura
}
func calcular(f figura) {
	fmt.Println("Area: ", f.area())
}
func main() {
	myCuadrado := cuadrado{base: 2}
	myRectanfulo := rectangulo{base: 2, altura: 4}
	calcular(myCuadrado)
	calcular(myRectanfulo)

	/*slices with different value types*/
	myInterface := []interface{}{"Hola", 12, 10, 1}
	fmt.Println(myInterface...)

}
