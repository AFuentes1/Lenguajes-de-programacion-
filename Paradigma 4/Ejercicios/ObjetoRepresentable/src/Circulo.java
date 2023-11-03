public class Circulo extends ObjetoRepresentable {
    private int x;
    private int y;
    private int radio;

    public Circulo(int x, int y, int radio) {
        this.x = x;
        this.y = y;
        this.radio = radio;
    }

    @Override
    public void dibujar() {
        System.out.println("Dibujando círculo en (" + x + ", " + y + ") con radio " + radio);
    }

    @Override
    public void mover(int deltaX, int deltaY) {
        x += deltaX;
        y += deltaY;
    }
}