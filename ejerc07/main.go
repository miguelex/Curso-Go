package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Printf("Suma 5 + 7= %d\n", Calculo(5, 7))

	//Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restar 5 - 7= %d\n", Calculo(5, 7))

	//Dividir
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividimos 12 / 3= %d\n", Calculo(12, 3))

	Operaciones()

	// CLOUSURES
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
