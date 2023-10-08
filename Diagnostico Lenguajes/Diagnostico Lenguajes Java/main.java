
public class main {
        public static void main(String[] args){
            ListaEnlazada alineacion = new ListaEnlazada();
    
            alineacion.insertarJugador("Jonathan Fuentes", 10);
            alineacion.insertarJugador("Jose Pablo Cascante", 7);
            alineacion.insertarJugador("Carlos Calderon", 9);
            alineacion.insertarJugador("Andrei Mesen", 5);
            alineacion.insertarJugador("Anthony Fuentes", 1);
            alineacion.insertarJugador("Geremy Avendaño", 2); 
    
            alineacion.imprimirLista();
        }
    }