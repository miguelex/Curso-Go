package main

import "fmt"

func main() {
	exponente(2)
}

func exponente(num int) {
	if num > 100000000000 {
		return
	}
	fmt.Println(num)
	exponente(num * 2)
}
