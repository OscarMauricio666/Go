package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nombre1 string //declaracion de variables tipadas
	edad := 30
	edad2 := "99"
	nombre1 = "óscar"             // asignacion
	nombre := "Mauricio"          // declaración dinamica
	edadStr := strconv.Itoa(edad) // convierte un entero a string
	fmt.Println("hello world" + edadStr)
	fmt.Println(nombre)
	fmt.Println(nombre1, edadStr)
	edad2Int, _ := strconv.Atoi(edad2) //convierte a entero operador _ desecha el 2do valor
	fmt.Println(edad + edad2Int)

	// Aquí van Unas ccondicionales Simples pero poderosas
	if edad == 29 {
		fmt.Println("Hola Mauricio")
	} else if edad == 30 {
		fmt.Println("Esta imprimiendo 30")
	} else {
		fmt.Println("En ninguna de las condicones")
	}

	//___________-Cliclos  un soo ciclo
	tres := 22

	for i := 0; i < 10; i++ {
		dos := "Mauricio"
		fmt.Println(dos)
		fmt.Println(tres)
		tres = +1
	}
	j := 1
	for j < 10 {
		fmt.Println(j * 10)
		j++
		if j > 7 {
			break
		}
	}
	// ARREGLOS
	var arreglo [10]int

	fmt.Println(arreglo)

	var arreglo2 [66]int {6,6,6}

}
