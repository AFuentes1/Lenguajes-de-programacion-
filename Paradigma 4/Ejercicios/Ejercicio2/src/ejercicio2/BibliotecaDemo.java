/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package ejercicio2;

import java.util.ArrayList;
import java.util.List;

public class BibliotecaDemo {
    public static void main(String[] args) {
        Libro libro1 = new Libro(1, "El Señor de los Anillos", "J.R.R. Tolkien");
        Libro libro2 = new Libro(2, "Harry Potter y la Piedra Filosofal", "J.K. Rowling");
        Libro libro3 = new Libro(3, "Cien años de soledad", "Gabriel García Márquez");
        Libro libro4 = new Libro(4, "1984", "George Orwell");

        Socio socio1 = new Socio(101, "Juan", "123 Calle Principal");
        Socio socio2 = new Socio(102, "Maria", "456 Calle Secundaria");
        Socio socio3 = new Socio(103, "Pedro", "789 Calle Terciaria");

        // Juan reserva y toma prestados dos libros
        socio1.tomarPrestadoLibro(libro1);
        socio1.tomarPrestadoLibro(libro2);

        // Maria reserva y toma prestados un libro
        socio2.tomarPrestadoLibro(libro1);

        // Pedro toma prestados tres libros
        socio3.tomarPrestadoLibro(libro2);
        socio3.tomarPrestadoLibro(libro3);
        socio3.tomarPrestadoLibro(libro4);

        // Juan devuelve un libro
        socio1.devolverLibro(libro1);

        System.out.println("Estado de los libros:");
        System.out.println("Libro 1 - " + libro1.getTitulo() + " - " + libro1.obtenerLocalizacion());
        System.out.println("Libro 2 - " + libro2.getTitulo() + " - " + libro2.obtenerLocalizacion());
        System.out.println("Libro 3 - " + libro3.getTitulo() + " - " + libro3.obtenerLocalizacion());
        System.out.println("Libro 4 - " + libro4.getTitulo() + " - " + libro4.obtenerLocalizacion());

        // Control de socios con más de 3 libros prestados
        List<Socio> sociosConMasDeTresLibros = new ArrayList<>();
        for (Socio socio : List.of(socio1, socio2, socio3)) {
            if (socio.obtenerCantidadLibrosPrestados() > 3) {
                sociosConMasDeTresLibros.add(socio);
            }
        }

        System.out.println("\nSocios con más de 3 libros prestados:");
        for (Socio socio : sociosConMasDeTresLibros) {
            System.out.println("Número de socio: " + socio.getNumeroSocio() + ", Nombre: " + socio.getNombre());
        }
    }
}