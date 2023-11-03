import java.util.ArrayList;
import java.util.List;

public class Hoja {
    private List<ObjetoRepresentable> objetos;

    public Hoja() {
        objetos = new ArrayList<>();
    }

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetos.add(objeto);
    }

    public List<ObjetoRepresentable> getObjetos() {
        return objetos;
    }
}