//Ejercicio 2
let rec filtrarPorSubcadena subcadena listaCadenas =
    match listaCadenas with
    | [] -> []  // Si la lista de cadenas está vacía, no hay más elementos que verificar, así que devolvemos una lista vacía como resultado.
    | cabeza::cola ->
        // Comprobamos si la cabeza de la lista contiene la subcadena.
        if contieneSubcadena subcadena cabeza then
            // Si la cabeza contiene la subcadena, la agregamos a la lista de resultados y luego llamamos recursivamente a la función con la cola restante.
            cabeza :: filtrarPorSubcadena subcadena cola  
        else
            // Si la cabeza no contiene la subcadena, omitimos la cabeza y llamamos recursivamente a la función con la cola restante.
            filtrarPorSubcadena subcadena cola  

and contieneSubcadena subcadena cadena =
    match subcadena, cadena with
    // Si la subcadena es una cadena vacía, consideramos que está presente en cualquier cadena, por lo que devolvemos verdadero.
    | "", _ -> true  
    // Si la cadena está vacía y la subcadena no lo está, la subcadena no puede estar contenida, así que devolvemos falso.
    | _, "" -> false  
    | sub, cad ->
        // Comprobamos si la cadena comienza con la subcadena utilizando la función StartsWith de cad. Si es así, devolvemos verdadero.
        if cad.StartsWith(sub) then true  
        // Si no comienza con la subcadena, llamamos recursivamente a la función con la cadena recortada en un carácter y verificamos de nuevo.
        else contieneSubcadena sub (cad.Substring(1))  