// "Implementar un algoritmo para manejar la asignacion de horarios de clases de la carrera de "
// "Ingenieria en computación"
package main

import "fmt"

type curso struct {
	nombre  string
	horario string
	aula    string
	notas   map[string]float64 //Se crea el mapa para almacenar las notas
}
type administradorDeHorarios struct {
	administradorMap map[string]curso
}

func nuevoAdministradorDeHorarios() *administradorDeHorarios {
	return &administradorDeHorarios{
		administradorMap: make(map[string]curso),
	}
}
func (a *administradorDeHorarios) agregarCurso(c curso) bool {
	if _, exist := a.administradorMap[c.horario]; exist {
		return false //El horario esta ocupado
	}
	a.administradorMap[c.horario] = c
	return true //El horario esta disponible
}

//Funcion para imprimir horarios
func (a *administradorDeHorarios) imprimirHorarios() {
	for horario, curso := range a.administradorMap {
		fmt.Printf("%s -%s (%s) \n", horario, curso.nombre, curso.aula)
	}
}

//Funcion para filtrar estudiantes
func filter(notas map[string]float64, condicion func(float64) bool) map[string]float64 {
	result := make(map[string]float64)
	for nombre, nota := range notas {
		if condicion(nota) {
			result[nombre] = nota
		}
	}
	return result
}

func main() {
	administrador := nuevoAdministradorDeHorarios()

	//Se agregan cursos de ejemplo
	curso1 := curso{nombre: "Analisis de algoritmos", horario: "Lunes 9:30-11:30", aula: "B3-03"}
	curso2 := curso{nombre: "Lenguajes de programacion", horario: "Lunes 9:30-11:30", aula: "B3-03"}
	curso3 := curso{nombre: "Estructuras de datos", horario: "Martes 9:30-11:30", aula: "B3-03"}
	curso4 := curso{nombre: "Sistemas operativos", horario: "Jueves 18:00-19:50", aula: "B3-03"}

	if administrador.agregarCurso(curso1) {
		fmt.Println("Curso agregado, Analisis de algoritmos")
	} else {
		fmt.Println("Horario ocupado, Analisis de algoritmos")
	}

	if administrador.agregarCurso(curso2) {
		fmt.Println("Curso agregado, Lenguajes de programacion")
	} else {
		fmt.Println("Horario ocupado, Lenguajes de programacion")
	}

	if administrador.agregarCurso(curso3) {
		fmt.Println("Curso agregado, Estructuras de datos")
	} else {
		fmt.Println("Horario ocupado, Estructuras de datos")
	}

	if administrador.agregarCurso(curso4) {
		fmt.Println("Curso agregado, Sistemas operativos")
	} else {
		fmt.Println("Horario ocupado, Sistemas operativos")
	}

	//simulación de ingreso de notas
	curso1.notas = map[string]float64{
		"Estudiante1": 100,
		"Estudiante2": 90,
		"Estudiante3": 80,
	}

	curso2.notas = map[string]float64{
		"Estudiante1": 70,
		"Estudiante2": 60,
		"Estudiante3": 50,
	}

	curso3.notas = map[string]float64{
		"Estudiante1": 40,
		"Estudiante2": 30,
		"Estudiante3": 20,
	}

	curso4.notas = map[string]float64{
		"Estudiante1": 100,
		"Estudiante2": 90,
		"Estudiante3": 80,
	}

	//Funcion para filtrar estudiantes aprobados
	estaAprobado := func(nota float64) bool {
		return nota >= 67.5
	}

	//Funcion para filtrar estudiantes reprobados
	estaReprobado := func(nota float64) bool {
		return !estaAprobado(nota)

	}

	//Funcion para calcular el promedio de notas
	promedio := func(notas map[string]float64) float64 {
		total := 0.0
		count := 0
		for _, nota := range notas {
			total += nota
			count++
		}
		if count > 0 {
			return total / float64(count)
		}
		return 0.0
	}
	//Calcula las estadisticas de cada curso
	for _, c := range administrador.administradorMap {
		fmt.Printf("Curso: %s \n", c.nombre)

		aprobados := filter(c.notas, estaAprobado)
		reprobados := filter(c.notas, estaReprobado)

		//se calculan los promedios
		promedioGeneral := promedio(c.notas)
		promedioAprobados := promedio(aprobados)
		promedioReprobados := promedio(reprobados)

		//se imprimen los resultados
		fmt.Printf("Aprobados: %d \n", len(aprobados))
		fmt.Printf("Reprobados: %d \n", len(reprobados))

		fmt.Printf("Promedio general: %.2f \n", promedioGeneral)
		fmt.Printf("Promedio aprobados: %.2f \n", promedioAprobados)
		fmt.Printf("Promedio reprobados: %.2f \n", promedioReprobados)

		fmt.Println("------------------------------------------")
	}
	administrador.imprimirHorarios()

}
