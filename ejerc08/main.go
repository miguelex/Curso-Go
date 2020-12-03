package main

import "fmt"

var tabla [10]int

func main() {

	tabla[0] = 1
	tabla[3] = 5

	fmt.Println(tabla) // Mostrara [1 0 0 5 0 0 0 0 0 0]

	tabla2 := [10]int{3, 6, 1, 8, 45, 6, 7, 90, 1, 19}

	fmt.Println(tabla2) // Mostrara [3 6 1 8 45 6 7 90 1 19]

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	//Matrices

	var matriz [5][7]int
	matriz[3][4] = 1
	fmt.Println(matriz) // Mostrara [[0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 0 0 0] [0 0 0 0 1 0 0] [0 0 0 0 0 0 0]]

	//Vectores dinamicos o slices

	matrix := []int{2, 5, 4}
	fmt.Println(matrix) // Muestra [2 5 4]
	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]  //desde posicion 3 hasta la ultima
	fmt.Println(porcion)      // Muestra [4 5]
	porcion2 := elementos[:4] //desde posicion 0 hasta la 4
	fmt.Println(porcion2)     // Muestra [1 2 3 4]
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

}

func variante4() {
	nums := make([]int, 0, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
