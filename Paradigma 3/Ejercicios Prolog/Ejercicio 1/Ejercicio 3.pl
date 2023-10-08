% Caso base: La lista vacía se aplanará como una lista vacía.
aplanar([], []).

% Caso recursivo: Aplana la cabeza de la lista y luego aplana la cola.
aplanar([H|T], L2) :-
    aplanar(H, FlatH),   % Aplana la cabeza de la lista.
    aplanar(T, FlatT),   % Aplana la cola de la lista.
    append(FlatH, FlatT, L2).  % Combina la cabeza aplanada y la cola aplanada para obtener la lista final.

% Caso base para aplanar un elemento que no es una lista.
aplanar(X, [X]) :- \+ is_list(X).

% El caso recursivo se toma la lista de entrada H = Primer elemento y
% T = al resto de la lista. Luego se realiza una llamada recursiva para
% aplanar H y T finalmente se usa append/3 para combinar H y T en la
% lista L2
% Hay un caso extra para un elemento X que no sea una lista
