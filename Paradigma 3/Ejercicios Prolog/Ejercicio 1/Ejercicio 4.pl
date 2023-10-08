% Caso base: cuando la lista de entrada está vacía, no hay elementos para dividir.
partir([], _, [], []).

% Caso recursivo: dividiendo la lista en menores y mayores respecto al umbral.
partir([X|Resto], Umbral, [X|Menores], Mayores) :-
    X =< Umbral,
    partir(Resto, Umbral, Menores, Mayores).

partir([X|Resto], Umbral, Menores, [X|Mayores]) :-
    X > Umbral,
    partir(Resto, Umbral, Menores, Mayores).

% Cuando se entra a la recursion, para cada elemento X en la lista de
% entrada se verifica si X es menor, igual o mayor que ele umbral dado.
% Dependiendo de esta verificacion, X se agregara a la lista de Menores
% o Mayores. Luego se hace una llamada recursiva para procesar el resto
