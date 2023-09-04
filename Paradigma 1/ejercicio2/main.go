package main

import {
	"fmt"
}


func main() {
	tamaño := 5 //El tamaño del rombo

	if tamaño%2 == 0 {
		fmt.Println("El tamaño debe ser impar")
		return
	}

	// Imprime la mitad superior del rombo
	imprimirMitadSuperior(tamaño)
	// Imprime la mitad inferior del rombo
	imprimirMitadInferior(tamaño)
}

func imprimirMitadSuperior(tamaño int) {
	for i := 0; i < tamaño/2+1; i++{
		imprimirEspacios(size/2 - i)
		imprimirAsteriscos(2*i + 1)
		fmt.Println()
	}

}

func imprimirMitadInferior(tamaño int) {
	for i := tamaño/2 - 1; i >= 0; i-- {
		printSpaces(tamaño/2 - i)
		printAsterisks(2*i + 1)
		fmt.Println()
	}
}

func imprimirEspacios(cont int) {
	for i := 0; i < cont; i++ {
		fmt.Print(" ")
	}
}

func imprimirAsteriscos(cont int) {
	for i := 0; i < cont; i++ {
		fmt.Print("*")
	}
}