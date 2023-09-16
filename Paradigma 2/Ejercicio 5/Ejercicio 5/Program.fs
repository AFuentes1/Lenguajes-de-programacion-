//Ejercicio 5 
type Point = { X: int; Y: int }

let isWithinBounds (laberinto: int list list) (point: Point) =
    let numRows = List.length laberinto
    if numRows = 0 then
        false
    else
        let numCols = List.length laberinto.[0]
        point.X >= 0 && point.X < numRows && point.Y >= 0 && point.Y < numCols

let findShortestPath (laberinto: int list list) (start: Point) (endPos: Point) =
    let directions = [(-1, 0); (1, 0); (0, -1); (0, 1)]

    let rec dfs currentPos visitedPath =
        if currentPos = endPos then
            Some (List.rev visitedPath)
        else
            let isValidNeighbor (dx, dy) =
                let neighbor = { X = currentPos.X + dx; Y = currentPos.Y + dy }
                isWithinBounds laberinto neighbor && laberinto.[neighbor.X].[neighbor.Y] = 0 && not (List.contains neighbor visitedPath)

            let mutable foundPath = None

            for (dx, dy) in directions do
                let neighbor = { X = currentPos.X + dx; Y = currentPos.Y + dy }
                if isValidNeighbor (dx, dy) then
                    let newPath = neighbor :: visitedPath
                    foundPath <- dfs neighbor newPath

            foundPath

    if not (isWithinBounds laberinto start) || not (isWithinBounds laberinto endPos) then
        None
    elif laberinto.[start.X].[start.Y] = 1 || laberinto.[endPos.X].[endPos.Y] = 1 then
        None
    else
        dfs start [start]

let start = { X = 0; Y = 0 }
let endPos = { X = 5; Y = 5 }

let laberinto = [
    [0; 0; 0; 1; 0; 0];
    [1; 1; 0; 1; 0; 1];
    [0; 0; 0; 0; 0; 0];
    [1; 0; 1; 1; 1; 1];
    [0; 0; 0; 0; 0; 0];
    [0; 1; 1; 0; 1; 1];
]

match findShortestPath laberinto start endPos with
| Some path -> printfn "Ruta más corta encontrada: %A" path
| None -> printfn "No se encontró una ruta válida."