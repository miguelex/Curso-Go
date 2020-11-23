package main

// Si la variable empieza en minusculas es una variable local.. Si la priemra letra es mayusculas se pueden exportar

import "fmt"

var numero1 int
var texto string
var status bool

func main() {
	var numero2, numero4 int
	numero3 := 4

	numero2, numero4, texto := 2, 5, "Este es mi texto"

	var moneda int64 = 0

	var moneda2 int
	moneda2 = int(moneda)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(moneda2)
}
