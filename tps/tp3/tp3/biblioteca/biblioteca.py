from collections import deque
from heapq import heappop, heappush
from tp3.tdas.grafo import Grafo
from tp3.tdas.union_find import UnionFind

#CONSTANTES:
INFINITO = float('inf')

"""Reconstruye el camino desde el origen hasta un destino
usando el diccionario de padres generado por Dijkstra o BFS."""
def reconstruir_camino(padres, destino):
    camino = []
    actual = destino
    while padres[actual]:
        camino.append(actual)
        padre = padres[actual]
        actual = padres[actual][0]
    camino.append(actual)
    camino.reverse()
    return camino    

"""Realiza un orden topológico sobre los vertices dados respetando
las aristas de precedencia.
Lanza excepción si hay un ciclo en las aristas."""
def orden_topologico(grafo):
    vertices = grafo.obtener_vertices()
    grados = {v: 0 for v in vertices}
    for origen in vertices:
        for destino in grafo.adyacentes(origen):
            grados[destino] += 1
    cola = deque([v for v in vertices if grados[v] == 0])
    orden = []
    while cola:
        u = cola.popleft()
        orden.append(u)
        for v in grafo.adyacentes(u):
            grados[v] -= 1
            if grados[v] == 0:
                cola.append(v)
    if len(orden) != len(vertices):
        raise ValueError("Hay un ciclo en las aristas")
    return orden

"""Busca el camino de menor cantidad de aristas entre
un conjunto de orígenes y destinos usando BFS."""
def bfs(grafo, origenes, destinos):
    interseccion = set(origenes) & set(destinos)
    if interseccion:
        destino = interseccion.pop()
        padres = {destino: []}
        return reconstruir_camino(padres, destino)
    padres = {}
    cola = deque(origenes)
    for origen in origenes:
        padres[origen] = []
    while cola:
        actual = cola.popleft()
        if actual in destinos:
            return reconstruir_camino(padres, actual)
        for vecino in grafo.adyacentes(actual):
            if vecino not in padres:
                padres[vecino] = [actual]
                cola.append(vecino)
    return None

"""
El algoritmos Dijkstra se encarga de calcular distancias mínimas, cantidad de caminos, padres 
y el orden topológico inverso desde un vértice origen. 
Usa una función para transformar los pesos originales del grafo antes de operar.
"""
def dijkstra(grafo, origen, funcion_peso):
    if origen not in grafo.obtener_vertices():
        raise ValueError(f"El vértice origen '{origen}' no está en el grafo.")
    vertices = grafo.obtener_vertices()
    dist = {v: INFINITO for v in vertices}
    padres = {v: [] for v in vertices}
    cantidad_caminos = {v: 0 for v in vertices}
    pila_topologica_inversa = []
    dist[origen] = 0
    cantidad_caminos[origen] = 1
    heap = [(0, origen)]
    while heap:
        d_actual, v = heappop(heap)
        if d_actual > dist[v]:
            continue
        pila_topologica_inversa.append(v)
        for w in grafo.adyacentes(v):
            peso_total = grafo.peso(v, w)
            peso = funcion_peso(v, w, peso_total)
            if peso is None or peso == 0 or peso == INFINITO:
                continue
            nueva_dist = dist[v] + peso
            if nueva_dist < dist[w]:
                dist[w] = nueva_dist
                padres[w] = [v]
                cantidad_caminos[w] = cantidad_caminos[v]
                heappush(heap, (nueva_dist, w))
            elif nueva_dist == dist[w]:
                padres[w].append(v)
                cantidad_caminos[w] += cantidad_caminos[v]
    return dist, cantidad_caminos, padres, pila_topologica_inversa

"""El algoritmo de Kruskal para construir el árbol de expansión mínima de un grafo."""
def kruskal(grafo, peso_arista):
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if (v, w) not in visitados and (w, v) not in visitados:
                aristas.append((v, w, peso_arista(v, w)))
                visitados.add((v, w))
    aristas.sort(key=lambda x: x[2])
    uf = UnionFind(grafo.obtener_vertices())
    mst = []
    for v, w, peso in aristas:
        if uf.find(v) != uf.find(w):
            uf.union(v, w)
            mst.append((v, w, peso))
    return mst

"""Implementa la fase de acumulacion de Brandes.
Recorre los vértices en orden inverso y
acumula lo acumulado de cada nodo a lo acumulado de
sus predecesores. """

def calcular_acumulados_centralidad(origen, dist, cantidad_caminos, padres, pila, centralidad):
    acumulado = {v: 0.0 for v in dist.keys()}
    while pila:
        w = pila.pop()
        for v in padres[w]:
            if cantidad_caminos[w] != 0:
                acumulado[v] += (cantidad_caminos[v] / cantidad_caminos[w]) * (1 + acumulado[w])
        if w != origen:
            centralidad[w] += acumulado[w]
