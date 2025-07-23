from tp3.biblioteca.biblioteca import kruskal, orden_topologico, bfs, dijkstra, calcular_acumulados_centralidad, reconstruir_camino
from tp3.comandos.errores import ERROR_PROCESAR_ITINERARIO, verificar_error_ciudades_itinerario
from tp3.tdas.grafo import Grafo

# CONSTANTES
INFINITO = float('inf')
FRECUENCIA = "frecuencia"
PRECIO = "precio"
ENCODING = 'utf-8'
TIEMPO_PROMEDIO = "tiempo_promedio"
MENSAJE_OK = "OK"
MENSAJE_ERROR = "ERROR"

"""Obtiene los valores de tiempo, precio y frecuencia de la arista (v, w).
Devuelve None si alguno de los valores no está definido."""
def obtener_pesos(grafo, v, w):
    pesos = grafo.peso(v, w)
    if pesos == float('inf') or pesos is None:
        return None
    tiempo = pesos.get(TIEMPO_PROMEDIO, INFINITO)
    precio = pesos.get(PRECIO, INFINITO)
    frecuencia = pesos.get(FRECUENCIA, INFINITO)
    if INFINITO in (tiempo, precio, frecuencia):
        return None
    return tiempo, precio, frecuencia

"""Calcula el mejor camino entre dos ciudades (orígenes y destinos)
según un peso dado, aplicando Dijkstra desde
todos los orígenes y eligiendo el de menor distancia."""
def calcular_camino_mejor_peso(grafo, origenes, destinos, funcion_peso):
    mejor_camino = None
    mejor_distancia = INFINITO
    for origen in origenes:
        if origen not in grafo.obtener_vertices():
            continue
        dist, _, padre, _ = dijkstra(grafo, origen, funcion_peso)
        for destino in destinos:
            if destino not in dist or dist[destino] == INFINITO:
                continue
            if dist[destino] < mejor_distancia:
                camino = reconstruir_camino(padre, destino)
                if camino is not None and len(camino) > 1:
                    mejor_distancia = dist[destino]
                    mejor_camino = camino
    return mejor_camino

"""Calcula el camino con la menor cantidad de escalas entre los aeropuertos
de origen y destino, usando búsqueda en anchura (BFS)."""
def calcular_camino_menos_escalas(grafo, origenes, destinos):
    return bfs(grafo, origenes, destinos)

"""Calcula la centralidad de intermediación de todos los vértices del grafo
usando el algoritmo de Brandes, y devuelve los "n" vértices más centrales."""
def calcular_centralidad(grafo, n, funcion_peso):
    centralidad = {v: 0.0 for v in grafo.obtener_vertices()}
    for s in grafo.obtener_vertices():
        dist, cantidad_caminos, padres, pila = dijkstra(grafo, s, funcion_peso)
        calcular_acumulados_centralidad(s, dist, cantidad_caminos, padres, pila, centralidad)
    return sorted(centralidad, key=lambda v: centralidad[v], reverse=True)[:n]

"""Genera un archivo con la nueva aerolínea basada en el árbol de expansión mínima
calculado con el algoritmo de Kruskal."""
def generar_nueva_aerolinea(grafo, archivo):
    mst = kruskal(grafo, lambda v, w: grafo.peso(v, w)[PRECIO])
    try:
        with open(archivo, "w", encoding=ENCODING) as f:
            for v, w, _ in mst:
                pesos = obtener_pesos(grafo, v, w)
                if pesos is None:
                    continue
                tiempo, precio, frecuencia = pesos
                f.write(f"{v},{w},{int(tiempo)},{int(precio)},{int(frecuencia)}\n")
        return MENSAJE_OK
    except IOError as e:
        print(f"Error al escribir archivo: {e}")
        return MENSAJE_ERROR
        
"""Lee un archivo que contiene una lista de ciudades y restricciones,
devuelve una lista de ciudades y una lista de pares."""
def leer_itinerario(archivo):
    with open(archivo) as f:
        lineas = [linea.strip() for linea in f if linea.strip()]
    ciudades = [c.strip() for c in lineas[0].split(",")]
    restricciones = [tuple(map(str.strip, linea.split(","))) for linea in lineas[1:]]
    return ciudades, restricciones

"""Calcula el itinerario respeta las restricciones entre ciudades.
Devuelve el orden topológico y los caminos entre ciudades consecutivas."""
def calcular_itinerario(grafo, mapa_ciudades, archivo):
    try:
        ciudades, restricciones = leer_itinerario(archivo)
        if not verificar_error_ciudades_itinerario(ciudades, mapa_ciudades):
            return None, None
        grafo_itinerario = Grafo(es_dirigido=True, pesado=False)
        for ciudad in ciudades:
            grafo_itinerario.agregar_vertice(ciudad)
        for origen, destino in restricciones:
            grafo_itinerario.agregar_arista(origen, destino)
        orden = orden_topologico(grafo_itinerario)
        caminos = []
        for i in range(len(orden) - 1):
            origen = orden[i]
            destino = orden[i + 1]
            aeropuertos_origen = mapa_ciudades[origen]
            aeropuertos_destino = mapa_ciudades[destino]
            camino = calcular_camino_menos_escalas(grafo, aeropuertos_origen, aeropuertos_destino)
            caminos.append(camino)
        return orden, caminos
    except Exception as e:
        print(ERROR_PROCESAR_ITINERARIO.format(str(e)))
        return None, None
