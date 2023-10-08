#Realizar en python el codigo fuente necesario para implementar una lista enlazada que permita 
#almacenar alineaciones de jugadores de futbol para un equipo cualquiera(nombre y numero de camiseta)
#en un orden cualquiera, para posteriormente imprimirlos en orden secuencial segun su numero de camiseta 

#Importar librerias
#esta libreria permite agregar y eliminar elementos de una lista
from collections import deque

class Jugador: 
    def __init__(self, nombre, numeroCamiseta) -> None:
        self.nombre = nombre
        self.numeroCamiseta = numeroCamiseta
        
    def __repr__(self) -> str:
        return f"{self.nombre} - {self.numeroCamiseta}"
    
#se crea una lista enlazada
alineacion = deque()

#se agregan los jugadores a la lista
alineacion.append(Jugador("Jonathan Fuentes", 10))
alineacion.append(Jugador("Jose Pablo Cascante", 7))
alineacion.append(Jugador("Carlos Calderon", 9))
alineacion.append(Jugador("Andrei Mesen", 5))
alineacion.append(Jugador("Anthony Fuentes", 1))
alineacion.append(Jugador("Geremy Avendaño", 2)) 

#se imprime la lista
print(sorted(alineacion, key=lambda x: x.numeroCamiseta))


