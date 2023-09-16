module RutaCorta

open System.Collections.Generic

// Grafo de prueba con pesos
let grafo = [
    ("i", [("a", 3); ("b", 6)]);
    ("a", [("i", 3); ("c", 5); ("d", 6)]);
    ("b", [("i", 6); ("c", 2); ("d", 7)]);
    ("c", [("a", 5); ("b", 2); ("x", 4)]);
    ("d", [("a", 6); ("b", 7); ("f", 8)]);
    ("f", [("d", 8)]);
    ("x", [("c", 4)])
]

// Función para generar vecinos con pesos
let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | (head, neighbors) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

// Función para extender una ruta con pesos
let extender ruta grafo = 
    List.collect 
        (fun (x, peso) -> 
            let vecinos_con_pesos = vecinos x grafo in
            List.map (fun (n, p) -> (n, p + peso)) vecinos_con_pesos
        ) ruta

// Función principal de búsqueda en profundidad con pesos
let rec prof2 ini fin grafo =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (fst (List.head ruta)) = fin then
            [List.rev ruta] @ prof_aux fin (List.tail ruta) grafo
        else
            prof_aux fin ((extender ruta grafo) @ (List.tail ruta)) grafo
    prof_aux fin [(ini, 0)] grafo

// Función para encontrar el camino más corto con pesos
let camino_mas_corto ini fin grafo =
    let rutas = prof2 ini fin grafo in
    match rutas with
    | [] -> None
    | _ ->
        let rutas_con_pesos = List.map (fun ruta -> (ruta, List.sumBy snd ruta)) rutas in
        let ruta_mas_corta = List.minBy snd rutas_con_pesos in
        Some ruta_mas_corta

// Ejemplo de uso
let resultado = camino_mas_corto "i" "f" grafo
match resultado with
| Some (ruta, peso) ->
    printfn "Camino más corto de 'i' a 'f':"
    List.iter (fun (nodo, p) -> printfn "Nodo: %s, Peso: %d" nodo p) ruta
    printfn "Peso total: %d" peso
| None -> printfn "No se encontró un camino de 'i' a 'f'"