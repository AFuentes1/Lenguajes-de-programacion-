sub_cadenas(Subcadena, ListaCadenas, Filtradas) :-
    incluir_cadenas_con_subcadena(Subcadena, ListaCadenas, Filtradas).

% Caso base: Cuando no hay más cadenas que procesar.
incluir_cadenas_con_subcadena(_, [], []).

% Caso recursivo: Incluye la cadena actual si contiene la subcadena.
incluir_cadenas_con_subcadena(Subcadena, [Cadena|Resto], [Cadena|FiltradasResto]) :-
    sub_atom(Cadena, _, _, _, Subcadena), % Verifica si Cadena contiene Subcadena.
    incluir_cadenas_con_subcadena(Subcadena, Resto, FiltradasResto).

% Caso recursivo: Descarta la cadena actual si no contiene la subcadena.
incluir_cadenas_con_subcadena(Subcadena, [_|Resto], FiltradasResto) :-
    incluir_cadenas_con_subcadena(Subcadena, Resto, FiltradasResto).
