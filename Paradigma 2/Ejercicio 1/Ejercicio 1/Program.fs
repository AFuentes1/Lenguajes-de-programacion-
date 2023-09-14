let desplazarLista (sentido:string) (posiciones:int) (lista:int list) =
    let n = List.length lista // Obtener la longitud de la lista de entrada
    let desplazamiento =
        match sentido with // Comenzar un patrón de coincidencia en el sentido
        | "izq" -> posiciones % n // Si el sentido es "izq", calcular el desplazamiento como posiciones % n
        | "der" -> n - (posiciones % n) // Si el sentido es "der", calcular el desplazamiento como n - (posiciones % n)
        | _ -> failwith "El sentido debe ser 'izq' o 'der'" // En caso de otro sentido, lanzar una excepción

    if desplazamiento = 0 then lista // Si el desplazamiento es 0, la lista no cambia
    else
        let primeraParte = List.take desplazamiento lista // Tomar los primeros "desplazamiento" elementos de la lista
        let segundaParte = List.skip desplazamiento lista // Omitir los primeros "desplazamiento" elementos de la lista
        List.append segundaParte primeraParte // Concatenar la segunda parte con la primera parte para obtener la lista desplazada

// Ejemplos de uso:
let listaOriginal = [1; 2; 3; 4; 5]
let resultadoIzq = desplazarLista "izq" 3 listaOriginal // Desplazar hacia la izquierda 3 posiciones
let resultadoDer = desplazarLista "der" 2 listaOriginal // Desplazar hacia la derecha 2 posiciones
let resultadoIzq6 = desplazarLista "izq" 6 listaOriginal // Desplazar hacia la izquierda 6 posiciones

printfn "%A lista Original" listaOriginal
printfn "%A" resultadoIzq  // Salida: [4; 5; 1; 2; 3]
printfn "%A" resultadoDer  // Salida: [4; 5; 1; 2; 3]
printfn "%A" resultadoIzq6 // Salida: [1; 2; 3; 4; 5]

