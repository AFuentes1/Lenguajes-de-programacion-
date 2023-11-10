/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package ejercicio2;

import java.util.Date;

class Prestamo {
    private int codigoLibro;
    private int numeroSocio;
    private Date fechaPrestamo;

    public Prestamo(int codigoLibro, int numeroSocio) {
        this.codigoLibro = codigoLibro;
        this.numeroSocio = numeroSocio;
        this.fechaPrestamo = new Date();
    }

    public int getCodigoLibro() {
        return codigoLibro;
    }

    public int getNumeroSocio() {
        return numeroSocio;
    }

    public Date getFechaPrestamo() {
        return fechaPrestamo;
    }
}