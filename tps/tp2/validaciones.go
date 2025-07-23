package main

import (
	"fmt"
	"strconv"
	"time"
	"tp2/aeropuerto"
	"tp2/operaciones"
)

const (
	ASCENDENTE          = "asc"
	DESCENDENTE         = "desc"
	ERROR_CANTIDAD      = "cantidad invalida"
	ERROR_MODO          = "el modo debe ser 'asc' o 'desc'"
	ERROR_FORMATO       = "fecha con formato incorrecto"
	ERROR_FORMATO_DESDE = "fecha 'desde' con formato incorrecto"
	ERROR_FORMATO_HASTA = "fecha 'hasta' con formato incorrecto"
	FORMATO_FECHA       = "2006-01-02T15:04:05"
	AVISO_FECHA         = "'hasta' debe ser igual o posterior a 'desde'"
)

func errorComando(nombre string) error {
	return fmt.Errorf("Error en comando %s", nombre)
}

func procesarVerTablero(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	cantidad, err := strconv.Atoi(campos[1])
	if err != nil {
		return fmt.Errorf(ERROR_CANTIDAD)
	}
	modo := campos[2]
	if modo != ASCENDENTE && modo != DESCENDENTE {
		return fmt.Errorf(ERROR_MODO)
	}
	desde, err := time.Parse(FORMATO_FECHA, campos[3])
	if err != nil {
		return fmt.Errorf(ERROR_FORMATO_DESDE)
	}
	hasta, err := time.Parse(FORMATO_FECHA, campos[4])
	if err != nil {
		return fmt.Errorf(ERROR_FORMATO_HASTA)
	}
	if hasta.Before(desde) {
		return fmt.Errorf(AVISO_FECHA)
	}
	operaciones.VerTablero(almacen, cantidad, modo, desde, hasta)
	return nil
}

func procesarBorrar(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	desde, err := time.Parse(FORMATO_FECHA, campos[1])
	if err != nil {
		return fmt.Errorf(ERROR_FORMATO_DESDE)
	}
	hasta, err := time.Parse(FORMATO_FECHA, campos[2])
	if err != nil {
		return fmt.Errorf(ERROR_FORMATO_HASTA)
	}
	if hasta.Before(desde) {
		return fmt.Errorf(AVISO_FECHA)
	}
	operaciones.Borrar(almacen, desde, hasta)
	return nil
}

func procesarInfoVuelo(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	codigo := campos[1]
	err := operaciones.InfoVuelo(almacen, codigo)
	if err != nil {
		return errorComando(INFO_VUELO)
	}
	return nil
}

func procesarPrioridadVuelos(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	cantidad, err := strconv.Atoi(campos[1])
	if err != nil || cantidad < 0 {
		return fmt.Errorf(ERROR_CANTIDAD)
	}
	operaciones.PrioridadVuelos(almacen, cantidad)
	return nil
}

func procesarSiguienteVuelo(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	fecha, err := time.Parse(FORMATO_FECHA, campos[3])
	if err != nil {
		return fmt.Errorf(ERROR_FORMATO)
	}
	operaciones.SiguienteVuelo(almacen, campos[1], campos[2], fecha)
	return nil
}

func ejecutarAgregarArchivo(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 2 {
		return errorComando(AGREGAR_ARCHIVO)
	}
	if err := operaciones.AgregarArchivo(campos[1], almacen); err != nil {
		return errorComando(AGREGAR_ARCHIVO)
	}
	return nil
}

func ejecutarVerTablero(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 5 {
		return errorComando(VER_TABLERO)
	}
	if err := procesarVerTablero(campos, almacen); err != nil {
		return errorComando(VER_TABLERO)
	}
	return nil
}

func ejecutarInfoVuelo(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 2 {
		return errorComando(INFO_VUELO)
	}
	if err := procesarInfoVuelo(campos, almacen); err != nil {
		return errorComando(INFO_VUELO)
	}
	return nil
}

func ejecutarPrioridadVuelos(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 2 {
		return errorComando(PRIORIDAD_VUELOS)
	}
	if err := procesarPrioridadVuelos(campos, almacen); err != nil {
		return errorComando(PRIORIDAD_VUELOS)
	}
	return nil
}

func ejecutarSiguienteVuelo(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 4 {
		return errorComando(SIGUIENTE_VUELO)
	}
	if err := procesarSiguienteVuelo(campos, almacen); err != nil {
		return errorComando(SIGUIENTE_VUELO)
	}
	return nil
}

func ejecutarBorrar(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) != 3 {
		return errorComando(BORRAR)
	}
	if err := procesarBorrar(campos, almacen); err != nil {
		return errorComando(BORRAR)
	}
	return nil
}
