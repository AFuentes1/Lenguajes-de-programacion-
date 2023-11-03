public class EditorGraficoDemo {
    public static void main(String[] args) {
        Hoja hoja1 = new Hoja();
        Hoja hoja2 = new Hoja();

        Circulo circulo = new Circulo(100, 100, 50);
        Rectangulo rectangulo = new Rectangulo(200, 200, 80, 40);

        Grupo grupo1 = new Grupo("Grupo1");
        grupo1.agregarObjeto(circulo);
        grupo1.agregarObjeto(rectangulo);

        Grupo grupo2 = new Grupo("Grupo2");
        grupo2.agregarObjeto(grupo1);

        hoja1.agregarObjeto(circulo);
        hoja1.agregarObjeto(rectangulo);
        hoja2.agregarObjeto(grupo2);

        System.out.println("Contenido de la Hoja 1:");
        for (ObjetoRepresentable objeto : hoja1.getObjetos()) {
            objeto.dibujar();
        }

        System.out.println("\nContenido de la Hoja 2:");
        for (ObjetoRepresentable objeto : hoja2.getObjetos()) {
            objeto.dibujar();
        }
    }
}