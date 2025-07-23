package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	CAPACIDAD_INICIAL  = 1
	CAPACIDAD_MINIMA   = 1
	FACTOR_CRECIMIENTO = 2
	FACTOR_REDUCCION   = 4
)

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, 0, CAPACIDAD_INICIAL),
		cantidad: 0,
	}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	if nuevaCapacidad < CAPACIDAD_INICIAL {
		nuevaCapacidad = CAPACIDAD_INICIAL
	}
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, pila.datos[:pila.cantidad])
	pila.datos = nuevosDatos
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == len(pila.datos) {
		nuevaCapacidad := cap(pila.datos) * FACTOR_CRECIMIENTO
		pila.redimensionar(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.cantidad--
	elemento := pila.datos[pila.cantidad]
	if pila.cantidad*FACTOR_REDUCCION <= cap(pila.datos) && cap(pila.datos) > CAPACIDAD_MINIMA {
		nuevaCapacidad := cap(pila.datos) / 2
		pila.redimensionar(nuevaCapacidad)
	}
	return elemento
}
