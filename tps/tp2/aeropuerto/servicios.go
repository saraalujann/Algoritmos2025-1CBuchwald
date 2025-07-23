package aeropuerto

import (
	"time"
)

type Operaciones interface {
	// Procesa de forma completa un archivo de .csv que contiene datos de vuelos.
	AgregarArchivo(archivo string, almacen *AlmacenVuelos)

	// Muestra los K vuelos ordenados por fecha de forma ascendente (asc) o
	// descendente (desc), cuya fecha de despegue esté dentro del intervalo <desde> <hasta> (inclusive).
	VerTablero(almacen *AlmacenVuelos, cantidad int, modo string, desde, hasta time.Time)

	// Muestra toda la información posible sobre el vuelo que tiene el código pasado por parámetro.
	InfoVuelo(almacen *AlmacenVuelos, codigo string)

	// Muestra los códigos de los K vuelos que tienen mayor prioridad.
	PrioridadVuelos(almacen *AlmacenVuelos, cantidad int)

	// Muestra la información del vuelo (tal cual en info_vuelo) del próximo vuelo directo que conecte
	// los aeropuertos de origen y destino, a partir de la fecha indicada (inclusive).
	// Si no hay un siguiente vuelo cargado, imprimir:
	// "No hay vuelo registrado desde <aeropuerto origen> hacia <aeropuerto destino> desde <fecha>"
	SiguienteVuelo(almacen *AlmacenVuelos, origen, destino string, fecha time.Time)

	// Borra todos los vuelos cuya fecha de despegue estén dentro del intervalo <desde> <hasta> (inclusive).
	Borrar(almacen *AlmacenVuelos, desde, hasta time.Time)
}
