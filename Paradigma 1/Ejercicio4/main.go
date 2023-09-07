package main

import "fmt"

// Se define la estructura "calzado" para almacenar la información de los zapatos
type calzado struct {
	marca  string
	precio float64
	talla  int
	stock  int //Cantidad de zapatos disponibles
}

// Funcion para vender un zapato del inventario
func venderZapato(zapato *calzado) bool {
	if zapato.stock > 0 {
		zapato.stock--
		return true
	}
	return false

}
func main() {
	//Se declara un slice para almacenar los zapatos en el inventario
	var inventario []calzado

	//Se agregan los zapatos al inventario
	inventario = append(inventario, calzado{"Nike", 50000, 42, 5})
	inventario = append(inventario, calzado{"Adidas", 45000, 38, 3})
	inventario = append(inventario, calzado{"Puma", 40000, 40, 2})
	inventario = append(inventario, calzado{"gucci", 200000, 41, 1})
	inventario = append(inventario, calzado{"Vans", 30000, 41, 0})

	//Se imprime el inventario inicial
	fmt.Println("Inventario de zapatos: ")
	for i, zapato := range inventario {
		fmt.Println("Zapato %d \n", i+1)
		fmt.Println("Marca: %s \n", zapato.marca)
		fmt.Println("Precio: %.2f \n", zapato.precio)
		fmt.Println("Talla: %d \n", zapato.talla)
		fmt.Println("Stock: %d \n", zapato.stock)
		fmt.Println()

	}

	//Se realiza la venta de los zapatos y se muestra el resultado
	fmt.Println("Venta de zapatos: ")
	if venderZapato(&inventario[0]) {
		fmt.Println("Se vendio un par Nike talla 42")
	} else {
		fmt.Println("No hay sufiente stock de zapatos Nike talla 42")
	}

	if venderZapato(&inventario[1]) {
		fmt.Println("Se vendio un par Adidas talla 38")
	} else {
		fmt.Println("No hay sufiente stock de zapatos Adidas talla 38")
	}

	if venderZapato(&inventario[2]) {
		fmt.Println("Se vendio un par Puma talla 40")
	} else {
		fmt.Println("No hay sufiente stock de zapatos Puma talla 40")
	}

	if venderZapato(&inventario[3]) {
		fmt.Println("Se vendio un par Gucci talla 41")
	} else {
		fmt.Println("No hay sufiente stock de zapatos Gucci talla 41")
	}

	if venderZapato(&inventario[4]) {
		fmt.Println("Se vendio un par Vans talla 41")
	} else {
		fmt.Println("No hay sufiente stock de zapatos Vans talla 41")
	}

	//Se imprime el inventario final, con las ventas actualizadas
	fmt.Println("Inventario de zapatos actualizado: ")
	for i, zapato := range inventario {
		fmt.Println("Zapato %d \n", i+1)
		fmt.Println("Marca: %s \n", zapato.marca)
		fmt.Println("Precio: %.2f \n", zapato.precio)
		fmt.Println("Talla: %d \n", zapato.talla)
		fmt.Println("Stock: %d \n", zapato.stock)
		fmt.Println()

	}

}
