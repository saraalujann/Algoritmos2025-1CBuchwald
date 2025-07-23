ERROR_AEROLINEA = "Error al generar nueva aerolínea"
ERROR_CAMINO = "No hay camino para exportar."
ERROR_CIUDADES = "Error: una o ambas ciudades no existen."
ERROR_INVALIDO_CAMINO = "Error: formato inválido para 'camino_mas'. Uso: camino_mas <rapido|barato>, origen, destino"
ERROR_INVALIDO_ESCALA = "Error: formato inválido para 'camino_escalas'. Uso: camino_escalas origen, destino"
ERROR_INVALIDO_CENTRALIDAD = "Error: para centralidad se espera un número entero."
ERROR_INVALIDO_ITINERARIO = "Error: debe especificar la ruta del archivo del itinerario."
ERROR_INVALIDO_KML = "Error: debe especificar la ruta del archivo KML."
ERROR_INVALIDO_RUTA = "Error: debe especificar la ruta del archivo de salida."
ERROR_KML = "Error al exportar KML"
ERROR_PARAMETROS = "No se encontró camino para los parámetros dados."
ERROR_PROCESAR_ITINERARIO = "Error al procesar itinerario."
ERROR_TIPO = "Error: el tipo debe ser 'rapido' o 'barato'."
ERROR_CIUDAD_NO_ENCONTRADA = "Error: ciudad no encontrada: {}"
ERROR_CICLO = "Hay un ciclo en las restricciones"
ERROR_EXPORTAR_KML = "Error al exportar KML: {}"
ERROR_COMANDO_NO_RECONOCIDO = "Comando de camino no reconocido"
ERROR_CAMINO_EXPORTADO = "Error: No hay un camino calculado para exportar."
ERROR_ORIGEN_NO_EXISTE = "El vértice origen '{}' no está en el grafo."
ERROR_CANTIDAD_ARGUMENTOS = "Cantidad incorrecta de argumentos. Esperados: {}"
MENSAJE_OK = "OK"
MENSAJE_SIN_CONEXION = "(no hay conexión {origen} -> {destino})"
TIPO_RAPIDO = "rapido"
TIPO_BARATO = "barato"

"""Verifica que los argumentos del comando camino sean válidos.
Devuelve una tupla (tipo, origen, destino) o None si hay error."""
def verificar_error_argumentos_camino(argumentos):
    partes = [p.strip() for p in argumentos.split(",")]
    if len(partes) != 3:
        print(ERROR_INVALIDO_CAMINO)
        return None
    tipo, origen, destino = partes
    if tipo not in (TIPO_RAPIDO, TIPO_BARATO):
        print(ERROR_TIPO)
        return None
    return tipo, origen, destino

"""Verifica que todas las ciudades del itinerario existan en el mapa.
Devuelve True si todas están presentes, False si hay alguna inválida."""
def verificar_error_ciudades_itinerario(ciudades, mapa_ciudades):
    for ciudad in ciudades:
        if ciudad not in mapa_ciudades:
            print(ERROR_CIUDAD_NO_ENCONTRADA.format(ciudad))
            return False
    return True

"""Verifica que los argumentos del comando sean válidos.
Devuelve una lista con las partes o None si hay error."""
def verificar_error_argumentos_escalas(argumentos):
    partes = [p.strip() for p in argumentos.split(",")]
    if len(partes) != 2:
        print(ERROR_INVALIDO_ESCALA)
        return None
    return partes

"""Verifica si se encontró un camino válido.
Devuelve True si existe, False si no hay camino."""
def verificar_error_camino_encontrado(camino):
    if not camino:
        print(ERROR_PARAMETROS)
        return False
    return True

"""Verifica que la ruta de salida sea válida.
Devuelve True si es válida, False si es inválida."""
def verificar_error_ruta_salida(ruta):
    if not ruta.strip():
        print(ERROR_INVALIDO_RUTA)
        return False
    return True

"""Verifica que origen y destino existan en el mapa de ciudades.
Devuelve True si ambos existen, False si alguno falta."""
def verificar_error_ciudad_origen_destino(mapa_ciudades, origen, destino):
    if origen not in mapa_ciudades or destino not in mapa_ciudades:
        print(ERROR_CIUDADES)
        return False
    return True

"""Verifica que origen y destino existan en el mapa de ciudades.
Devuelve True si ambos existen, False si alguno falta."""
def verificar_error_entero(argumento):
    try:
        return int(argumento.strip())
    except ValueError:
        print(ERROR_INVALIDO_CENTRALIDAD)
        return None

"""Verifica que la ruta del archivo de itinerario sea válida.
Devuelve True si es no vacía, False en caso contrario."""
def verificar_error_ruta_itinerario(ruta_itinerario):
    if not ruta_itinerario.strip():
        print(ERROR_INVALIDO_ITINERARIO)
        return False
    return True

"""Verifica que se haya podido generar correctamente el itinerario.
Devuelve True si los datos son válidos, False si hay error."""
def verificar_error_resultado_itinerario(orden_ciudades, caminos):
    if orden_ciudades is None or caminos is None:
        print(ERROR_PROCESAR_ITINERARIO)
        return False
    return True

"""Verifica que la ruta del archivo KML sea válida.
Devuelve True si es no vacía, False si está vacía."""
def verificar_error_ruta_kml(ruta_kml):
    if not ruta_kml.strip():
        print(ERROR_INVALIDO_KML)
        return False
    return True

"""Verifica que haya un camino previamente generado.
Devuelve True si existe, False si está vacío o es None."""
def verificar_error_ultimo_camino(ultimo_camino):
    if not ultimo_camino:
        print(ERROR_CAMINO)
        return False
    return True
