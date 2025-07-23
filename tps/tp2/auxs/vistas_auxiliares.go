package auxs

import (
	TDAColaPrioridad "tdas/cola_prioridad"
	TDALista "tdas/lista"
	"time"
	"tp2/aeropuerto"
)

/* CompararFechas compara dos valores time.Time y devuelve:
-1 si a < b,
1 si a > b,
0 si son iguales. */

func CompararFechas(a, b time.Time) int {
	switch {
	case a.Before(b):
		return -1
	case a.After(b):
		return 1
	default:
		return 0
	}
}

/*ListaVuelosAscendentes crea y devuelve una lista enlazada de vuelos ordenados ascendentemente
por fecha y número de vuelo, obtenidos del diccionario en el rango desde-hasta.
Si uno de los límites es nil, devuelve lista vacía. */

func ListaVuelosAscendentes(almacen *aeropuerto.AlmacenVuelos, desde, hasta *time.Time) TDALista.Lista[*aeropuerto.Vuelo] {
	lista := TDALista.CrearListaEnlazada[*aeropuerto.Vuelo]()
	if desde == nil || hasta == nil {
		return lista
	}
	for iter := almacen.Fecha.IteradorRango(desde, hasta); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelos := iter.VerActual()
		if len(vuelos) > 1 {
			heap := TDAColaPrioridad.CrearHeap(func(a, b *aeropuerto.Vuelo) int {
				return a.CompararAscendentePorFechaYNumero(b)
			})
			for _, vuelo := range vuelos {
				heap.Encolar(vuelo)
			}
			for !heap.EstaVacia() {
				lista.InsertarUltimo(heap.Desencolar())
			}
		} else {
			lista.InsertarUltimo(vuelos[0])
		}
	}
	return lista
}

/*InvertirLista devuelve una lista enlazada con los vuelos de la lista descendente.*/

func InvertirLista(almacen *aeropuerto.AlmacenVuelos, desde, hasta *time.Time) TDALista.Lista[*aeropuerto.Vuelo] {
	ascendente := ListaVuelosAscendentes(almacen, desde, hasta)
	listaInvertida := TDALista.CrearListaEnlazada[*aeropuerto.Vuelo]()
	iter := ascendente.Iterador()
	for iter.HaySiguiente() {
		listaInvertida.InsertarPrimero(iter.VerActual())
		iter.Siguiente()
	}
	return listaInvertida
}

/*compararPrioridades compara dos vuelos según prioridad y número de vuelo.*/
func compararPrioridades(a, b *aeropuerto.Vuelo) int {
	return a.CompararPorPrioridad(b)
}

/*extraerTopN extrae los primeros n elementos del heap y los devuelve en un slice.
Si n es menor o igual a cero, devuelve slice vacío.*/

func extraerTopNVuelosPrioritarios(heap TDAColaPrioridad.ColaPrioridad[*aeropuerto.Vuelo], n int) []*aeropuerto.Vuelo {
	if n <= 0 {
		return []*aeropuerto.Vuelo{}
	}
	resultado := make([]*aeropuerto.Vuelo, 0, min(n, heap.Cantidad()))
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		resultado = append(resultado, heap.Desencolar())
	}
	for i, j := 0, len(resultado)-1; i < j; i, j = i+1, j-1 {
		resultado[i], resultado[j] = resultado[j], resultado[i]
	}
	return resultado
}

/*ObtenerTopPrioritarios devuelve un slice con los vuelos de mayor prioridad
hasta la cantidad solicitada, usando un heap para ordenarlos.*/

func ObtenerTopPrioritarios(almacen *aeropuerto.AlmacenVuelos, cantidad int) []*aeropuerto.Vuelo {
	if cantidad <= 0 {
		return []*aeropuerto.Vuelo{}
	}
	heap := TDAColaPrioridad.CrearHeapArr([]*aeropuerto.Vuelo{}, func(a, b *aeropuerto.Vuelo) int {
		return compararPrioridades(b, a)
	})
	iter := almacen.Codigo.Iterador()
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		if heap.Cantidad() < cantidad {
			heap.Encolar(vuelo)
		} else if compararPrioridades(vuelo, heap.VerMax()) > 0 {
			heap.Desencolar()
			heap.Encolar(vuelo)
		}
		iter.Siguiente()
	}
	return extraerTopNVuelosPrioritarios(heap, cantidad)
}
