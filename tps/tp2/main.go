package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/aeropuerto"
)

const (
	AGREGAR_ARCHIVO  = "agregar_archivo"
	VER_TABLERO      = "ver_tablero"
	INFO_VUELO       = "info_vuelo"
	PRIORIDAD_VUELOS = "prioridad_vuelos"
	SIGUIENTE_VUELO  = "siguiente_vuelo"
	BORRAR           = "borrar"
	MENSAJE_OK       = "OK"
	FORMATO_VACIO    = ""
)

func ejecutarComando(campos []string, almacen *aeropuerto.AlmacenVuelos) error {
	if len(campos) == 0 {
		return nil
	}
	switch comando := campos[0]; comando {
	case AGREGAR_ARCHIVO:
		return ejecutarAgregarArchivo(campos, almacen)
	case VER_TABLERO:
		return ejecutarVerTablero(campos, almacen)
	case INFO_VUELO:
		return ejecutarInfoVuelo(campos, almacen)
	case PRIORIDAD_VUELOS:
		return ejecutarPrioridadVuelos(campos, almacen)
	case SIGUIENTE_VUELO:
		return ejecutarSiguienteVuelo(campos, almacen)
	case BORRAR:
		return ejecutarBorrar(campos, almacen)
	default:
		return fmt.Errorf("comando desconocido: %s", comando)
	}
}

func leerComando() {
	scanner := bufio.NewScanner(os.Stdin)
	almacen := aeropuerto.NuevoAlmacenVuelos()
	for scanner.Scan() {
		linea := scanner.Text()
		if linea == FORMATO_VACIO {
			continue
		}
		campos := strings.Fields(linea)
		err := ejecutarComando(campos, almacen)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		} else {
			fmt.Println(MENSAJE_OK)
		}
	}
}

func main() {
	leerComando()
}
