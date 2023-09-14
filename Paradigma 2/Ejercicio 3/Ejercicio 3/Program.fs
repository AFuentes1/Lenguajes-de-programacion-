//Ejercicio 3

// Definimos la función n_esimo que toma un índice n y una lista como entrada.
let n_esimo n lista =
    // Definimos una función interna llamada obtener_nesimo que lleva un índice i y una lista como entrada.
    let rec obtener_nesimo i = function
        | [] -> None // Si la lista está vacía, devolvemos None.
        | cabeza::cola when i = n -> Some cabeza // Si encontramos el elemento en la posición n, devolvemos Some con ese elemento.
        | _::cola -> obtener_nesimo (i + 1) cola // Si no hemos encontrado el elemento, avanzamos a la siguiente posición de la lista.
    obtener_nesimo 1 lista // Iniciamos la búsqueda desde la posición 1 de la lista.

// Definimos la función imprimirResultado que toma un valor opcional y lo imprime en el formato deseado.
let imprimirResultado n lista =
    match n_esimo n lista with
    | Some valor -> printfn "%d" valor // Si obtenemos Some con un valor, lo imprimimos como un entero.
    | None -> printfn "false" // Si obtenemos None, imprimimos "false".

// Ejemplos de uso
imprimirResultado 2 [1; 2; 3; 4; 5] // Llamamos a imprimirResultado con n = 2 y la lista [1; 2; 3; 4; 5], debería imprimir "3".
imprimirResultado 3 [1; 2; 3; 4; 5] // Llamamos a imprimirResultado con n = 3 y la lista [1; 2; 3; 4; 5], debería imprimir "4".
imprimirResultado 6 [1; 2; 3; 4; 5] // Llamamos a imprimirResultado con n = 6 (fuera del rango) y la lista [1; 2; 3; 4; 5], debería imprimir "false".