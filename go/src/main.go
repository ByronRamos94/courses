package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}
func returnValue(a int) int {
	return a * a
}
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	const pi float64 = 3.1416
	var age int = 10
	month := 12
	var year int

	fmt.Println(age, month, year)
	fmt.Println(pi)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	x := 10
	y := 20
	result := x + y

	fmt.Println(result)

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	normalFunction("Hola Mundo")
	normalFunction("Hola Mundo modificado")
	value := returnValue(2)
	fmt.Println("Value:", value)

	//double return
	value1, value2 := doubleReturn(2)
	// omitir un valor de retorno
	// value1, _ := doubleReturn(2)
	fmt.Println(value1)
	fmt.Println(value2)

}
