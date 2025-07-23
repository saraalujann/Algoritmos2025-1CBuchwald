#!/usr/bin/python3
import sys
import os
import csv
from tp3.tdas.grafo import Grafo
sys.path.append(os.path.join(os.path.dirname(__file__), 'tp3'))
from tp3.comandos.exportador import ExportadorKML
from tp3.comandos.manejos import (manejar_camino_mejor_peso, manejar_camino_menos_escalas, manejar_centralidad, manejar_nueva_aerolinea, manejar_itinerario, manejar_exportar_kml, CAMINO_MAS_LARGO, CAMINO_MAS_ESCALAS)
from tp3.comandos.errores import ERROR_COMANDO_NO_RECONOCIDO, ERROR_CAMINO_EXPORTADO

# CONSTANTES:
ENCODING = 'utf-8'
MENSAJE_USO = "Uso: ./flycombi aeropuertos.csv vuelos.csv"
PRECIO = "precio"
TIEMPO_PROMEDIO = "tiempo_promedio"
FRECUENCIA = "frecuencia"
CENTRALIDAD = "centralidad"
NUEVA_AEROLINEA = "nueva_aerolinea"
ITINERARIO = "itinerario"
EXPORTAR_KML = "exportar_kml"

"""Carga los datos de aeropuertos y vuelos, construyendo un grafo de
aeropuertos y un mapa de ciudades a aeropuertos.
Las aristas incluyen tiempo promedio y precio, siendo pesadas."""
def cargar_datos(archivo_aeropuertos, archivo_vuelos):
    grafo = Grafo(es_dirigido=False, pesado=True)
    mapa_ciudades = {}
    codigos_agregados = set()
    with open(archivo_aeropuertos, encoding=ENCODING) as f:
        for linea in f:
            datos = linea.strip().split(",")
            if len(datos) < 2:
                continue
            ciudad, cod = datos[0], datos[1]
            if ciudad not in mapa_ciudades:
                mapa_ciudades[ciudad] = []
            mapa_ciudades[ciudad].append(cod)
            if cod not in codigos_agregados:
                grafo.agregar_vertice(cod)
                codigos_agregados.add(cod)
    with open(archivo_vuelos, encoding=ENCODING) as f:
        for linea in f:
            datos = linea.strip().split(",")
            if len(datos) < 5:
                continue
            origen, destino = datos[0], datos[1]
            tiempo = int(datos[2])
            precio = float(datos[3])
            frecuencia = int(datos[4])
            pesos = {PRECIO: precio, TIEMPO_PROMEDIO: tiempo, FRECUENCIA: frecuencia}
            grafo.agregar_arista(origen, destino, pesos)
    return grafo, mapa_ciudades
    
"""Lee el archivo CSV de aeropuertos y extrae las coordenadas
asociadas a cada código de aeropuerto, finalmente las devuelve en un diccionario."""
def cargar_coordenadas_desde_csv(path):
    coords = {}
    with open(path, newline='', encoding=ENCODING) as f:
        reader = csv.reader(f)
        for fila in reader:
            try:
                codigo = fila[1].strip()
                lat = float(fila[2])
                lon = float(fila[3])
                coords[codigo] = (lon, lat)
            except Exception:
                continue
    return coords

"""Inicializa el programa: valida argumentos, carga el grafo, las coordenadas de aeropuertos
y crea el exportador en formato KML."""
def inicializar():
    if len(sys.argv) != 3:
        print(MENSAJE_USO)
        sys.exit(1)
    archivo_aeropuertos = sys.argv[1]
    archivo_vuelos = sys.argv[2]
    grafo, mapa_ciudades = cargar_datos(archivo_aeropuertos, archivo_vuelos)
    coordenadas = cargar_coordenadas_desde_csv(archivo_aeropuertos)
    exportador = ExportadorKML(grafo, coordenadas)
    return grafo, mapa_ciudades, exportador

"""Recibe un comando y sus argumentos, y ejecuta la acción correspondiente."""
def procesar_comando(comando, argumentos, grafo, mapa_ciudades, exportador, caminos_historial):
    if comando == CAMINO_MAS_LARGO:
        camino, tipo = manejar_camino_mejor_peso(argumentos, grafo, mapa_ciudades)
        if camino:
            caminos_historial.append((camino, tipo))
    elif comando == CAMINO_MAS_ESCALAS:
        camino, tipo = manejar_camino_menos_escalas(argumentos, grafo, mapa_ciudades)
        if camino:
            caminos_historial.append((camino, tipo))
    elif comando == CENTRALIDAD:
        manejar_centralidad(argumentos, grafo)
    elif comando == NUEVA_AEROLINEA:
        manejar_nueva_aerolinea(argumentos, grafo)
    elif comando == ITINERARIO:
        manejar_itinerario(argumentos, grafo, mapa_ciudades)
    elif comando == EXPORTAR_KML:
        if not caminos_historial:
            print(ERROR_CAMINO_EXPORTADO)
        else:
            camino_actual, _ = caminos_historial[-1]
            manejar_exportar_kml(argumentos, exportador, camino_actual)
    else:
        print(ERROR_COMANDO_NO_RECONOCIDO.format(comando))

"""Lee comandos desde la entrada estándar y los procesa uno por uno 
hasta que se termine la entrada."""
def ciclo_comandos(grafo, mapa_ciudades, exportador):
    caminos_historial = []
    for linea in sys.stdin:
        linea = linea.strip()
        if linea == "":
            continue
        comando, *resto = linea.split(" ", 1)
        argumentos = resto[0] if resto else ""
        procesar_comando(comando, argumentos, grafo, mapa_ciudades, exportador, caminos_historial)

"""Funcion principal"""
def main():
    grafo, mapa_ciudades, exportador = inicializar()
    ciclo_comandos(grafo, mapa_ciudades, exportador)

if __name__ == "__main__":
    main()
