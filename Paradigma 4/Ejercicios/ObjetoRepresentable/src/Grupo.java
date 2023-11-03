import java.util.ArrayList;
import java.util.List;

public class Grupo extends ObjetoRepresentable {
    private String nombre;
    private List<ObjetoRepresentable> objetos;

    public Grupo(String nombre) {
        this.nombre = nombre;
        objetos = new ArrayList<>();
    }

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetos.add(objeto);
    }

    public String getNombre() {
        return nombre;
    }

    @Override
    public void dibujar() {
        System.out.println("Dibujando grupo: " + nombre);
        for (ObjetoRepresentable objeto : objetos) {
            objeto.dibujar();
        }
    }

    @Override
    public void mover(int deltaX, int deltaY) {
        for (ObjetoRepresentable objeto : objetos) {
            objeto.mover(deltaX, deltaY);
        }
    }
}