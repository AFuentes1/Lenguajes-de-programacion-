package Laboratorio1

import (
	"fmt"
	"math/rand"
	"time"
)

type Matrix interface {
	Get(i, j int) rune
}

type RuneMatrix [N][N]rune

func (rm RuneMatrix) Get(i, j int) rune {
	return rm[i][j]
}

// DefiniciÃ³n de la funciÃ³n genÃ©rica map3
func map3[T Matrix, R any](input []T, mapper func(T, int, int, int) []R, x int, y int, n int) [][]R {
	var result [][]R
	for _, item := range input {
		elements := mapper(item, x, y, n)
		result = append(result, elements)
	}
	return result
}

const (
	N           = 3
	MinSize     = 3
	MaxSize     = 8
	AsciiMin    = 33  // CarÃ¡cter ASCII mÃ­nimo imprimible ('!')
	AsciiMax    = 126 // CarÃ¡cter ASCII mÃ¡ximo imprimible
	MaxAttempts = 100
)

// FunciÃ³n para generar una matriz aleatoria
func generateRandomMatrix() RuneMatrix {
	var matrix RuneMatrix
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			matrix[i][j] = rune(rand.Intn(AsciiMax-AsciiMin+1) + AsciiMin)
		}
	}
	return matrix
}

// FunciÃ³n para obtener elementos en posiciÃ³n vertical
func getVertical(matrix Matrix, x int, y int, n int) []rune {
	var result []rune
	for i := 0; i < n; i++ {
		if y+i < N {
			result = append(result, matrix.Get(x, y+i))
		} else {
			result = append(result, 0) // Agregar nil o un valor indicativo
		}
	}
	return result
}

// FunciÃ³n para obtener elementos en posiciÃ³n horizontal
func getHorizontal(matrix Matrix, x int, y int, n int) []rune {
	var result []rune
	for i := 0; i < n; i++ {
		if x+i < N {
			result = append(result, matrix.Get(x+i, y))
		} else {
			result = append(result, 0) // Agregar nil o un valor indicativo
		}
	}
	return result
}

// FunciÃ³n para obtener elementos en posiciÃ³n diagonal
func getDiagonal(matrix Matrix, x int, y int, n int) []rune {
	var result []rune
	for i := 0; i < n; i++ {
		if x+i < N && y+i < N {
			result = append(result, matrix.Get(x+i, y+i))
		} else {
			result = append(result, 0) // Agregar nil o un valor indicativo
		}
	}
	return result
}

func printMatrix(matrix RuneMatrix) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Establecer semilla para la generaciÃ³n aleatoria

	// Llenar la lista con matrices aleatorias
	var matrixList []Matrix
	for i := 0; i < 5; i++ { // Llenar con 5 matrices de ejemplo
		randomMatrix := generateRandomMatrix()
		matrixList = append(matrixList, randomMatrix)
	}

	// ParÃ¡metros para las funciones de cÃ¡lculo
	x := 0
	y := 1
	n := 2

	// Llamada a la funciÃ³n map3 con diferentes funciones de cÃ¡lculo
	verticalResult := map3(matrixList, getVertical, x, y, n)
	horizontalResult := map3(matrixList, getHorizontal, x, y, n)
	diagonalResult := map3(matrixList, getDiagonal, x, y, n)

	// Imprimir matrices de ejemplo
	fmt.Println("Matrices de ejemplo:")
	for _, matrix := range matrixList {
		printMatrix(matrix.(RuneMatrix))
		fmt.Println()
	}

	// Imprimir resultados
	fmt.Println("Resultado vertical:")
	for _, result := range verticalResult {
		fmt.Println(result)
	}
	fmt.Println()

	fmt.Println("Resultado horizontal:")
	for _, result := range horizontalResult {
		fmt.Println(result)
	}
	fmt.Println()

	fmt.Println("Resultado diagonal:")
	for _, result := range diagonalResult {
		fmt.Println(result)
	}
}
