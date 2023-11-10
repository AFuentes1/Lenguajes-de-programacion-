/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package ejercicio2;

import java.util.ArrayList;
import java.util.List;

// Clase para representar un socio de la biblioteca
class Socio {
    private int numeroSocio;
    private String nombre;
    private String direccion;
    private List<Libro> librosPrestados;

    public Socio(int numeroSocio, String nombre, String direccion) {
        this.numeroSocio = numeroSocio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.librosPrestados = new ArrayList<>();
    }

    public int getNumeroSocio() {
        return numeroSocio;
    }

    public String getNombre() {
        return nombre;
    }

    public String getDireccion() {
        return direccion;
    }

    public List<Libro> getLibrosPrestados() {
        return librosPrestados;
    }

    public void tomarPrestadoLibro(Libro libro) {
        if (!libro.estaPrestado()) {
            librosPrestados.add(libro);
            libro.prestar(this);
        } else {
            System.out.println("El libro " + libro.getTitulo() + " ya está prestado.");
        }
    }

    public void devolverLibro(Libro libro) {
        if (librosPrestados.contains(libro)) {
            librosPrestados.remove(libro);
            libro.devolver(this);
        } else {
            System.out.println("No puedes devolver un libro que no has tomado prestado.");
        }
    }

    public int obtenerCantidadLibrosPrestados() {
        return librosPrestados.size();
    }
}
