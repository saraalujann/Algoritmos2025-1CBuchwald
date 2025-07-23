from tp3.biblioteca.analisis_red import calcular_centralidad, generar_nueva_aerolinea, calcular_itinerario, calcular_camino_mejor_peso, calcular_camino_menos_escalas
from tp3.comandos.errores import ERROR_CANTIDAD_ARGUMENTOS

"""Parsea una cadena de argumentos separados por comas, limpia
los espacios y valida que la cantidad coincida con la esperada.
Lanza un error si la cantidad es incorrecta."""
def parsear_argumentos(argumentos, esperados):
    args = [arg.strip() for arg in argumentos.split(",")]
    if len(args) != esperados:
        raise ValueError(ERROR_CANTIDAD_ARGUMENTOS.format(esperados))
    return args

"""Devuelve el mejor camino entre aeropuertos
de las ciudades según el peso indicado (por ejemplo, tiempo o precio)."""
def camino_mejor_peso(grafo, origenes, destinos, funcion_peso):
    return calcular_camino_mejor_peso(grafo, origenes, destinos, funcion_peso)

"""Devuelve el camino entre aeropuertos con menor cantidad de escalas mediante BFS."""
def camino_menos_escalas(grafo, origenes, destinos):
    return calcular_camino_menos_escalas(grafo, origenes, destinos)

"""Devuelve una cadena con los primeros `n` aeropuertos más centrales predefinidos,
separados por una coma y un espacio."""
def centralidad(grafo, n, funcion_peso):
    if n <= 0:
        raise ValueError("El número debe ser positivo")
    return calcular_centralidad(grafo, n, funcion_peso)

"""Crea un archivo con nombre dado en argumentos y escribe en él
un mensaje predeterminado. Retorna "OK" si la operación es exitosa."""
def nueva_aerolinea(grafo, argumentos):
    archivo = parsear_argumentos(argumentos, 1)[0]
    return generar_nueva_aerolinea(grafo, archivo)

"""Lee el archivo de itinerario y genera el orden topológico de ciudades
con restricciones. Luego obtiene los caminos con menos escalas
entre cada par consecutivo de ciudades.
Retorna el orden y la lista de caminos o (None, None) si hay error"""
def itinerario(grafo, mapa_ciudades, argumentos):
    archivo = parsear_argumentos(argumentos, 1)[0]
    orden, caminos = calcular_itinerario(grafo, mapa_ciudades, archivo)
    if orden is None or caminos is None:
        return None, None
    return orden, caminos
