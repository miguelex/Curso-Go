package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var z = 0
	for z < 10 {
		fmt.Printf("\n Valor de i: %d", z)
		if z == 5 {
			fmt.Printf("multiplico por 2 \n")
			z = z * 2
			continue
		}
		fmt.Printf("    Paso por aqui")
		z++

	}

	var y int = 0

RUTINA:
	for y < 10 {
		if y == 4 {
			y = y + 2
			fmt.Println("voy a RUTINA sumando 2 a y")
			goto RUTINA
		}
		fmt.Println(" El valor de y: \n", y)
		y++
	}
}
