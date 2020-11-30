package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 23, 45, 60))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 46, 17, 25, 26, 90, 17))

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		// _ recojo el valor de la vaibale pero no la voy a usar. Aunque no la use la tengo que recoger
		total = total + num
	}
	return total
}
