
public class ListaEnlazada {
    private Nodo cabeza;
    private int longitud;

    public ListaEnlazada(){
        cabeza = null;
        longitud = 0;
    }
   
    public void insertarJugador(String nombre, int numeroCamiseta){
         //Se crea un nuevo nodo con el nombre y el numero de camiseta
        Nodo nuevoNodo = new Nodo(nombre, numeroCamiseta);
        //Si la lista esta vacia o el numero de camiseta es menor a que el primero en la lista,
        //el nuevo jugador se convierte en el primero 
        if (cabeza == null || numeroCamiseta < cabeza.getNumeroCamiseta()){
            nuevoNodo.setSiguiente(cabeza);
            cabeza = nuevoNodo;
        } else {
            Nodo actual = cabeza;

            while (actual.getSiguiente() != null && actual.getSiguiente().getNumeroCamiseta() < numeroCamiseta){
                actual = actual.getSiguiente();
            }
            nuevoNodo.setSiguiente(actual.getSiguiente());
            actual.setSiguiente(nuevoNodo);

        }

    }

    public void imprimirLista(){
        Nodo actual = cabeza;
        while (actual != null){
            System.out.println(actual.toString());
            actual = actual.getSiguiente();
        }
    }
}

