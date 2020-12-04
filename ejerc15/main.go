package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Miguel Delgado")
	fmt.Println("Estoy aqui")

	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
