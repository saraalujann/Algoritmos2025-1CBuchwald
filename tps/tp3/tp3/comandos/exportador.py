import xml.etree.ElementTree as ET
from tp3.comandos.errores import ERROR_EXPORTAR_KML

#CONSTANTES:
ENCODING = 'utf-8'
KML_NAMESPACE = "http://earth.google.com/kml/2.1"
KML_ROOT = "kml"
KML_DOCUMENT = "Document"
KML_NAME = "name"
KML_DESCRIPTION = "description"
KML_PLACEMARK = "Placemark"
KML_LINESTRING = "LineString"
KML_COORDINATES = "coordinates"
KML_NOMBRE_ARCHIVO = "Rutas de vuelo"
KML_POINT = "Point"
KML_ARCHIVO = "Archivo generado con ExportadorKML"
XML_ROOT = '<?xml version="1.0" encoding="UTF-8"?>\n'

"""Clase que exporta un camino de aeropuertos a un archivo KML."""
class ExportadorKML:
    def __init__(self, grafo, coordenadas):
        self.grafo = grafo
        self.coordenadas = coordenadas
        self.ultimo_camino = None

    def exportar(self, camino, archivo):
        self.ultimo_camino = camino
        try:
            kml_element = self._generar_kml_element(camino)
            xml_declaration = XML_ROOT
            rough_string = ET.tostring(kml_element, encoding=ENCODING).decode(ENCODING)
            with open(archivo, "w", encoding=ENCODING) as f:
                f.write(xml_declaration)
                f.write(rough_string)
            return True
        except (IOError, OSError, ET.ParseError) as e:
            print(ERROR_EXPORTAR_KML.format(e))
            return False

    def _generar_kml_element(self, camino):
        root = ET.Element(KML_ROOT, xmlns=KML_NAMESPACE)
        ET.register_namespace(" ", KML_NAMESPACE)
        doc = ET.SubElement(root, KML_DOCUMENT)
        self._agregar_informacion(doc)
        self._agregar_aeropuertos(doc, camino)
        self._agregar_rutas(doc, camino)
        return root

    def _agregar_informacion(self, doc):
        nombre = ET.SubElement(doc, KML_NAME)
        nombre.text = KML_NOMBRE_ARCHIVO
        desc = ET.SubElement(doc, KML_DESCRIPTION)
        desc.text = KML_ARCHIVO

    def _agregar_aeropuertos(self, doc, camino):
        for aeropuerto in camino:
            if aeropuerto in self.coordenadas:
                placemark = self._crear_placemark(aeropuerto)
                point = ET.SubElement(placemark, KML_POINT)
                coords = ET.SubElement(point, KML_COORDINATES)
                lon, lat = self.coordenadas[aeropuerto]
                coords.text = f"{lon},{lat} "
                doc.append(placemark)

    def _agregar_rutas(self, doc, camino):
        for i in range(len(camino) - 1):
            origen, destino = camino[i], camino[i+1]
            if origen in self.coordenadas and destino in self.coordenadas:
                placemark = self._crear_placemark(f"Ruta {origen}-{destino}")
                linestring = ET.SubElement(placemark, KML_LINESTRING)
                coords = ET.SubElement(linestring, KML_COORDINATES)
                lon_origen, lat_origen = self.coordenadas[origen]
                lon_destino, lat_destino = self.coordenadas[destino]
                coords.text = f"{lon_origen},{lat_origen} {lon_destino},{lat_destino}"
                doc.append(placemark)

    def _crear_placemark(self, nombre):
        placemark = ET.Element(KML_PLACEMARK)
        name_element = ET.SubElement(placemark, KML_NAME)
        name_element.text = nombre
        return placemark
