package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	posicion_maximo := 0
	for i := 1; i < len(vector); i++ {
		if vector[i] > vector[posicion_maximo] {
			posicion_maximo = i
		}
	}
	return posicion_maximo
}

// Compara elemento a elemento para saber que devolver dependiendo
// su comparacion.
func compararElementos(primerElem, segundoElem int) int {
	if primerElem < segundoElem {
		return -1
	} else if primerElem > segundoElem {
		return 1
	} else {
		return 0
	}
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for i := 0; i < len(vector1) && i < len(vector2); i++ {
		resultado := compararElementos(vector1[i], vector2[i])
		if resultado != 0 {
			return resultado
		}
	}
	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	} else {
		return 0
	}
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i > 0; i-- {
		indice_maximo := Maximo(vector[:i+1])
		Swap(&vector[i], &vector[indice_maximo])
	}
}

// Se llama las veces necesarias hasta que llegue al final del arreglo
func sumaRecursiva(vector []int, indice int) int {
	if indice >= len(vector) {
		return 0
	}
	return vector[indice] + sumaRecursiva(vector, indice+1)
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return sumaRecursiva(vector, 0)
}

// Se llama las veces necesarias hasta que llegue a la mitad o mas del arreglo para saber si es capicua.
func capicuaRecursiva(cadena string, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	if cadena[inicio] != cadena[fin] {
		return false
	}
	return capicuaRecursiva(cadena, inicio+1, fin-1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return capicuaRecursiva(cadena, 0, len(cadena)-1)
}

// Lee un archivo pasado por parametro y devuelve un arreglo de enteros.
func LeerArchivo(nombreArchivo string) []int {
  archivo, _ := os.Open(nombreArchivo)
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	var arreglo []int
	for s.Scan() {
		numero, _ := strconv.Atoi(s.Text())
		arreglo = append(arreglo, numero)
	}
	err := s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return arreglo
}

// Al pasarle dos arreglos saber cual es el mayor usando la funcion de "Comparar"
func ObtenerArregloMayor(arreglo1, arreglo2 []int) []int {
	comparacion := Comparar(arreglo1, arreglo2)
	var arregloMayor []int
	if comparacion < 0 {
		arregloMayor = arreglo2
	} else {
		arregloMayor = arreglo1
	}
	return arregloMayor
}
