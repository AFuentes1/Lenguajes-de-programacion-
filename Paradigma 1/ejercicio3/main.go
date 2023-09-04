package main

import "fmt"

// Función para rotar el arreglo de caracteres en posiciones
// en la dirección indicada por dirección
func rotateArray(lista []rune, posiciones int, direccion int) {
	tamaño := len(lista)
	if tamaño <= 0 || posiciones <= 0 {
		return // Los parametros no se pueden rotar si no hay elementos o si las posiciones son 0
	}

	// Realizamos la rotación dependiendo de la dirección
	if direccion == 0 { // Se rota a la izquierda
		posiciones = posiciones % tamaño // Manejo del exceso de posiciones
		for i := 0; i < posiciones; i++ {
			temp := lista[0] // Se guarda el primer elemento
			// Se mueve cada elemento del lugar a la izquierda
			for j := 0; j < tamaño-1; j++ {
				lista[j] = lista[j+1]
			}
			lista[tamaño-1] = temp // Se coloca el primer elemento al final
		}
	} else if direccion == 1 { // Se rota a la derecha
		posiciones = posiciones % tamaño // Manejo de exceso de posiciones
		for i := 0; i < posiciones; i++ {
			temp := lista[tamaño-1] // Se guarda el último elemento
			// Se mueve cada elemento un lugar a la derecha
			for j := tamaño - 1; j > 0; j-- {
				lista[j] = lista[j-1]
			}
			lista[0] = temp // Se coloca el último elemento al principio
		}
	}
}

func main() {
	// Se crea una lista de caracteres
	listaOriginal := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	// Se obtiene el tamaño de la lista
	//tamaño := len(listaOriginal)
	// Se imprime la lista original
	fmt.Println("Lista original:")

	for _, char := range listaOriginal {
		fmt.Printf("%c ", char)
	}
	fmt.Println() // Se imprime un salto de línea

	posiciones := 3                                   // Número de posiciones para rotar
	direccion := 0                                    // Dirección de rotación (0 izquierda, 1 derecha)
	rotateArray(listaOriginal, posiciones, direccion) // Se rota la lista

	fmt.Print("Lista rotada: ") // Se imprime la lista rotada
	for _, char := range listaOriginal {
		fmt.Printf("%c ", char) // Se imprime cada caracter de la lista
	}
	fmt.Println() // Se imprime un salto de línea
}
