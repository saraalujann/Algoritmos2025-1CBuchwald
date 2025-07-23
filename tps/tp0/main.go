package main

import (
	"fmt"
	"tp0/ejercicios"
)

func main() {
	arreglo1 := ejercicios.LeerArchivo("archivo1.in")
	arreglo2 := ejercicios.LeerArchivo("archivo2.in")
	arregloMayor := ejercicios.ObtenerArregloMayor(arreglo1, arreglo2)
	ejercicios.Seleccion(arregloMayor)
	for _, numero := range arregloMayor {
		fmt.Println(numero)
	}
}
