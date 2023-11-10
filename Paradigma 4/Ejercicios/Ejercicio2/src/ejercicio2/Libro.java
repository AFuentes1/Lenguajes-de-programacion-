/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package ejercicio2;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;

// Clase para representar un libro
class Libro {
    private int codigo;
    private String titulo;
    private String autor;
    private boolean prestado;
    private Socio socioPrestatario;

    public Libro(int codigo, String titulo, String autor) {
        this.codigo = codigo;
        this.titulo = titulo;
        this.autor = autor;
        this.prestado = false;
        this.socioPrestatario = null;
    }

    public int getCodigo() {
        return codigo;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getAutor() {
        return autor;
    }

    public boolean estaPrestado() {
        return prestado;
    }

    public Socio getSocioPrestatario() {
        return socioPrestatario;
    }

    public void prestar(Socio socio) {
        prestado = true;
        socioPrestatario = socio;
    }

    public void devolver(Socio socio) {
        prestado = false;
        socioPrestatario = null;
    }

    public String obtenerLocalizacion() {
        if (estaPrestado()) {
            return "Prestado a: " + socioPrestatario.getNombre();
        } else {
            return "Disponible en la biblioteca";
        }
    }
}
