//Realizar en Java el código fuente necesario para implementar una lista enlazada que permita 
//almacenar alineaciones de jugadores de futbol para un equipo cualquiera (nombre y número de camiseta) 
//en un orden cualquiera, para posteriormente imprimirlos en orden secuencial según su número de camiseta.

public class Nodo{
    private String nombre; 
    private int numeroCamiseta; 
    private Nodo siguiente; 

    //constructor
    public Nodo(String nombre, int numeroCamiseta){
        this.nombre = nombre; 
        this.numeroCamiseta = numeroCamiseta;
        this.siguiente = null;
    }

    //getters y setters
    public String getNombre(){
        return nombre;
    }

    public int getNumeroCamiseta(){
        return numeroCamiseta;
    }

    public Nodo getSiguiente(){
        return siguiente;
    }

    public void setNombre(String nombre){
        this.nombre = nombre;
    }

    public void setNumeroCamiseta(int numeroCamiseta){
        this.numeroCamiseta = numeroCamiseta;
    }

    public void setSiguiente(Nodo siguiente){
        this.siguiente = siguiente;
    }

    @Override
    public String toString(){
        return "Nombre: " + nombre + " Numero de camiseta: " + numeroCamiseta;
    }

    
    
}