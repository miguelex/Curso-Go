package main

import "fmt"

func main() {

	paises := make(map[string]string) // make(map[string]string), 5) Crearia 5 elementos
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"]) //Muestra D.F.
	fmt.Println(paises)           // Muestra map[Argentina:Buenos Aires Mexico:D.F.]

	campeonato := map[string]int{
		"Madrid":    40,
		"Barcelona": 38,
		"Chivas":    37,
		"Boca":      30}

	fmt.Println(campeonato) // Muesrra map[Barcelona:38 Boca:30 Chivas:37 Madrid:40] Ordenado alfabeticamente

	campeonato["River"] = 25

	fmt.Println(campeonato)

	delete(campeonato, "Madrid")
	fmt.Println(campeonato)

	for equipo, puntos := range campeonato {
		fmt.Printf("El equipo %s tiene %d puntos\n", equipo, puntos)
	}

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t\n", puntaje, existe)

}
