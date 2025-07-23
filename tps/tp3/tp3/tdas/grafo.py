class Grafo:
    def __init__(self, es_dirigido=True, pesado=True):
        self.grafo = {}
        self.dirigido = es_dirigido
        self.pesado = pesado 

    def agregar_vertice(self, v):
        if v not in self.grafo:
            self.grafo[v] = {}

    def agregar_arista(self, v1, v2, pesos=None):
        self.agregar_vertice(v1)
        self.agregar_vertice(v2)
        valor_arista = pesos if self.pesado else True
        self.grafo[v1][v2] = valor_arista
        if not self.dirigido:
            self.grafo[v2][v1] = valor_arista

    def borrar_vertice(self, v):
        if v in self.grafo:
            del self.grafo[v]
        for ady in self.grafo.values():
            ady.pop(v, None)

    def borrar_arista(self, v1, v2):
        if v1 in self.grafo and v2 in self.grafo[v1]:
            del self.grafo[v1][v2]
            if not self.dirigido and v2 in self.grafo and v1 in self.grafo[v2]:
                del self.grafo[v2][v1]

    def adyacentes(self, v):
        return list(self.grafo.get(v, {}).keys())
        
    def peso(self, v1, v2):
        return self.grafo.get(v1, {}).get(v2, None)

    def vertice_existe(self, v):
        return v in self.grafo

    def obtener_vertices(self):
        return list(self.grafo.keys())
