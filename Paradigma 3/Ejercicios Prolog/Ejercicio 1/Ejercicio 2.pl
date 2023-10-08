% Caso base: Una lista vacía es un subconjunto de cualquier lista.
subconj([], _).

% Caso recursivo: Verifica si el primer elemento de S1 está en S y luego verifica el resto de S1.
subconj([X|S1], S) :-
    member(X, S),     % Comprueba si X está en S.
    subconj(S1, S).   % Llamada recursiva para el resto de S1.

% Se verifica el primer elemento si esta en la lista S, luego se realiza
% una llamada recursiva para verificar el resto de S1, si todos los
% elementos de S1 estan en S, entonces se considera verdadero
