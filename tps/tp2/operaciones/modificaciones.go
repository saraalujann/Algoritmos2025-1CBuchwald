package operaciones

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"tp2/aeropuerto"
	"tp2/auxs"
)

const (
	CAMPOS_ESPERADOS = 10
	FORMATO_FECHA    = "2006-01-02T15:04:05"
)

/*AgregarArchivo es una funcion exportable que lee el archivo, con formato CSV
y guarda sus lineas en las estructuras adecuadas.*/

func AgregarArchivo(archivo string, almacen *aeropuerto.AlmacenVuelos) error {
	file, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %v", err)
	}
	defer file.Close()
	lector := csv.NewReader(file)
	for {
		campos, err := lector.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo CSV: %v", err)
		}
		if len(campos) != CAMPOS_ESPERADOS {
			return fmt.Errorf("registro con cantidad incorrecta de campos: %v", campos)
		}
		nuevoVuelo, err := parsearVuelo(campos)
		if err != nil {
			return fmt.Errorf("error parseando vuelo: %v", err)
		}
		procesarVuelo(nuevoVuelo, almacen)
	}
	return nil
}

/*Borrar es una funcion exportable que dado un rango de fechas, elimina los vuelos
mencionados en el rango. SI se meustra un "hasta" menor a "desde" muestra
mensaje de error*/

func Borrar(almacen *aeropuerto.AlmacenVuelos, desde, hasta time.Time) {
	listaVuelos := auxs.ListaVuelosAscendentes(almacen, &desde, &hasta)
	iter := listaVuelos.Iterador()
	for iter.HaySiguiente() {
		vuelo := iter.VerActual()
		fmt.Print(vuelo.ImprimirDetallado())
		almacen.Codigo.Borrar(vuelo.ClaveHash())
		eliminarVueloDeFecha(almacen, vuelo)
		iter.Siguiente()
	}
}

/*parsearVuelo convierte un slice de strings que representa un registro CSV
en un puntero a un aeropuerto.Vuelo, valida y parsea cada campo.
Devuelve error si algun campo no se puede convertir correctamente. */

func parsearVuelo(campos []string) (*aeropuerto.Vuelo, error) {
	priority, err := strconv.Atoi(campos[5])
	if err != nil {
		return nil, err
	}
	date, err := time.Parse(FORMATO_FECHA, campos[6])
	if err != nil {
		return nil, err
	}
	departureDelay, err := strconv.Atoi(campos[7])
	if err != nil {
		return nil, err
	}
	airTime, err := strconv.Atoi(campos[8])
	if err != nil {
		return nil, err
	}
	cancelled, err := strconv.Atoi(campos[9])
	if err != nil {
		return nil, err
	}
	return aeropuerto.NuevoVuelo(
		campos[0],
		campos[1],
		campos[2],
		campos[3],
		campos[4],
		priority,
		date,
		departureDelay,
		airTime,
		cancelled,
	), nil
}

/*procesarVuelo agrega/actualiza un vuelo en las estructuras abb y hash.
Si el vuelo ya existía, elimina la referencia anterior evitando duplicados.*/

func procesarVuelo(vuelo *aeropuerto.Vuelo, almacen *aeropuerto.AlmacenVuelos) {
	numeroVuelo := vuelo.ClaveHash()
	if almacen.Codigo.Pertenece(numeroVuelo) {
		vueloExistente := almacen.Codigo.Obtener(numeroVuelo)
		eliminarVueloDeFecha(almacen, vueloExistente)
	}
	almacen.Codigo.Guardar(numeroVuelo, vuelo)
	fecha := vuelo.Fecha()
	if almacen.Fecha.Pertenece(fecha) {
		lista := almacen.Fecha.Obtener(fecha)
		lista = append(lista, vuelo)
		almacen.Fecha.Guardar(fecha, lista)
	} else {
		almacen.Fecha.Guardar(fecha, []*aeropuerto.Vuelo{vuelo})
	}
}

/*eliminarVueloDeAbb elimina un vuelo específico del abb según su fecha y
número de vuelo. Si la lista queda vacia, borra la clave fecha completa.*/

func eliminarVueloDeFecha(almacen *aeropuerto.AlmacenVuelos, vuelo *aeropuerto.Vuelo) {
	fecha := vuelo.Fecha()
	if almacen.Fecha.Pertenece(fecha) {
		lista := almacen.Fecha.Obtener(fecha)
		lista = borrarVueloDeLista(lista, vuelo)
		if len(lista) == 0 {
			almacen.Fecha.Borrar(fecha)
		} else {
			almacen.Fecha.Guardar(fecha, lista)
		}
	}
}

/* borrarVueloDeLista devuelve una nueva lista de vuelos que no
tiene en cuenta el vuelo especificado según su número de vuelo.*/

func borrarVueloDeLista(lista []*aeropuerto.Vuelo, vuelo *aeropuerto.Vuelo) []*aeropuerto.Vuelo {
	result := make([]*aeropuerto.Vuelo, 0, len(lista))
	for _, v := range lista {
		if v.ClaveHash() != vuelo.ClaveHash() {
			result = append(result, v)
		}
	}
	return result
}
