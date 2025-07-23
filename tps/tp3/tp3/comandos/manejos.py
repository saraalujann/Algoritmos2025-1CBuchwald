from tp3.biblioteca.red_vuelos import nueva_aerolinea, itinerario, centralidad, camino_menos_escalas, camino_mejor_peso
from tp3.comandos.errores import verificar_error_argumentos_camino, verificar_error_argumentos_escalas, verificar_error_ciudad_origen_destino, verificar_error_camino_encontrado, verificar_error_ruta_salida, verificar_error_ruta_itinerario, verificar_error_resultado_itinerario, verificar_error_ruta_kml, verificar_error_ultimo_camino, verificar_error_entero, MENSAJE_SIN_CONEXION, ERROR_AEROLINEA, ERROR_KML

# CONSTANTES
CAMINO_MAS_ESCALAS = "camino_escalas"
CAMINO_MAS_LARGO = "camino_mas"
SEPARADOR_FLECHA = " -> "
TIPO_PRECIO = "precio"
FRECUENCIA = "frecuencia"
TIPO_RAPIDO = "rapido"
TIPO_TIEMPO_PROMEDIO = "tiempo_promedio"
MENSAJE_OK = "OK"

"""Devuelve el peso como el inverso de la frecuencia, asi 
prioriza caminos con mayor frecuencia."""
def peso_inverso_frecuencia(v, w, pesos):
    valor = pesos.get(FRECUENCIA, None)
    if valor is None or valor == 0:
        return None
    return 1 / valor

"""Devuelve el peso basado en el tiempo."""
def peso_tiempo_promedio(v, w, peso_original):
    return peso_original.get(TIPO_TIEMPO_PROMEDIO, None)

"""Devuelve el peso basado en el peso."""
def peso_precio(v, w, peso_original):
    return peso_original.get(TIPO_PRECIO, None)

"""Procesa el comando para buscar el camino más rápido o más barato entre dos ciudades.
Recibe argumentos en formato 'tipo, origen, destino', valida y ejecuta la búsqueda."""
def manejar_camino_mejor_peso(argumentos, grafo, mapa_ciudades):
    resultado = verificar_error_argumentos_camino(argumentos)
    if not resultado:
        return None, None
    tipo, origen, destino = resultado
    if not verificar_error_ciudad_origen_destino(mapa_ciudades, origen, destino):
        return None, None
    funcion_peso = peso_tiempo_promedio if tipo == TIPO_RAPIDO else peso_precio
    camino = camino_mejor_peso(grafo, mapa_ciudades[origen], mapa_ciudades[destino], funcion_peso)
    if not verificar_error_camino_encontrado(camino):
        return None, None
    print(SEPARADOR_FLECHA.join(camino))
    return camino, CAMINO_MAS_LARGO

"""Procesa el comando para buscar el camino con menor cantidad de
escalas entre dos ciudades.
Recibe argumentos en formato 'origen, destino', valida y ejecuta la búsqueda."""
def manejar_camino_menos_escalas(argumentos, grafo, mapa_ciudades):
    partes = verificar_error_argumentos_escalas(argumentos)
    if not partes:
        return None, None
    origen, destino = partes
    if not verificar_error_ciudad_origen_destino(mapa_ciudades, origen, destino):
        return None, None
    camino = camino_menos_escalas(grafo, mapa_ciudades[origen], mapa_ciudades[destino])
    if not verificar_error_camino_encontrado(camino):
        return None, None
    print(SEPARADOR_FLECHA.join(camino))
    return camino, CAMINO_MAS_ESCALAS

"""Devuelve los n mas centrales"""
def manejar_centralidad(argumentos, grafo):
    n = verificar_error_entero(argumentos)
    if n is None:
        return
    lista_centralidad = centralidad(grafo, n, peso_inverso_frecuencia)
    print(", ".join(lista_centralidad))

""" Procesa el comando para generar una nueva aerolínea escribiendo un archivo de salida."""
def manejar_nueva_aerolinea(argumentos, grafo):
    ruta_salida = argumentos.strip()
    if not verificar_error_ruta_salida(ruta_salida):
        return
    exito = nueva_aerolinea(grafo, ruta_salida)
    print(MENSAJE_OK if exito else ERROR_AEROLINEA)

"""Procesa el comando para generar un itinerario a partir de un archivo.
Valida ciudades, obtiene caminos entre ellas y muestra los resultados."""
def manejar_itinerario(argumentos, grafo, mapa_ciudades):
    ruta_itinerario = argumentos.strip()
    if not verificar_error_ruta_itinerario(ruta_itinerario):
        return
    orden_ciudades, caminos = itinerario(grafo, mapa_ciudades, ruta_itinerario)
    if not verificar_error_resultado_itinerario(orden_ciudades, caminos):
        return
    codigos_ciudades = {ciudad: aeropuertos[0][:3].upper() for ciudad, aeropuertos in mapa_ciudades.items()}
    print(", ".join(orden_ciudades))
    for i in range(len(orden_ciudades) - 1):
        origen = orden_ciudades[i]
        destino = orden_ciudades[i + 1]
        camino = caminos[i]
        cod_origen = codigos_ciudades.get(origen, origen[:3].upper())
        cod_destino = codigos_ciudades.get(destino, destino[:3].upper())
        if camino and len(camino) > 0:
            print(SEPARADOR_FLECHA.join([aeropuerto[:3].upper() for aeropuerto in camino]))
        else:
            print(MENSAJE_SIN_CONEXION.format(origen=cod_origen, destino=cod_destino))

"""Procesa el comando para exportar un camino a un archivo KML."""
def manejar_exportar_kml(argumentos, exportador, camino_a_exportar):
    ruta_kml = argumentos.strip()
    if not verificar_error_ruta_kml(ruta_kml):
        return
    if not verificar_error_ultimo_camino(camino_a_exportar):
        return
    exito = exportador.exportar(camino_a_exportar, ruta_kml)
    print(MENSAJE_OK if exito else ERROR_KML)
