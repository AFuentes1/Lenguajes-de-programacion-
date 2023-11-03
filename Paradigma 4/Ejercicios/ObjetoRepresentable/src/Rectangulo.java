public class Rectangulo extends ObjetoRepresentable {
    private int x;
    private int y;
    private int ancho;
    private int alto;

    public Rectangulo(int x, int y, int ancho, int alto) {
        this.x = x;
        this.y = y;
        this.ancho = ancho;
        this.alto = alto;
    }

    @Override
    public void dibujar() {
        System.out.println("Dibujando rectángulo en (" + x + ", " + y + ") con ancho " + ancho + " y alto " + alto);
    }

    @Override
    public void mover(int deltaX, int deltaY) {
        x += deltaX;
        y += deltaY;
    }
}