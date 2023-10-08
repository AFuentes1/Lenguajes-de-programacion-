sumlist([],0). %La suma de una lista vacia es 0
sumlist([X|Rest],S):-  %Cuando la lista no es vacia, se descompone en el primer elemento X y Rest
    sumlist(Rest, SRest), %llamada recursiva para sumar el resto de la lista
    S is X + SRest. %S es la suma de X y la suma del resto de la lista
