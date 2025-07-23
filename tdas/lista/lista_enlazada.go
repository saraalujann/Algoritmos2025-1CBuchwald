package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterLista[T any] struct {
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
}

func (l *listaEnlazada[T]) borrarNodo(anterior *nodoLista[T]) T {
	var nodo *nodoLista[T]
	if anterior == nil {
		nodo = l.primero
		l.primero = nodo.siguiente
	} else {
		nodo = anterior.siguiente
		anterior.siguiente = nodo.siguiente
	}
	if nodo == l.ultimo {
		l.ultimo = anterior
	}
	l.largo--
	return nodo.dato
}

func nuevoNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func (l *listaEnlazada[T]) insertarNodo(anterior *nodoLista[T], nuevo *nodoLista[T]) {
	if anterior == nil {
		nuevo.siguiente = l.primero
		l.primero = nuevo
	} else {
		nuevo.siguiente = anterior.siguiente
		anterior.siguiente = nuevo
	}
	if l.largo == 0 || nuevo.siguiente == nil {
		l.ultimo = nuevo
	}
	l.largo++
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := nuevoNodo(elemento)
	if l.EstaVacia() {
		l.insertarNodo(nil, nuevo)
	} else {
		nuevo.siguiente = l.primero
		l.primero = nuevo
		l.largo++
	}
}

func (l *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := nuevoNodo(elemento)
	if l.EstaVacia() {
		l.insertarNodo(nil, nuevo)
	} else {
		l.insertarNodo(l.ultimo, nuevo)
	}
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.borrarNodo(nil)
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil && visitar(actual.dato) {
		actual = actual.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterLista[T]{lista: l, actual: l.primero}
}

func (it *iterLista[T]) VerActual() T {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return it.actual.dato
}

func (it *iterLista[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iterLista[T]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.anterior = it.actual
	it.actual = it.actual.siguiente
}

func (it *iterLista[T]) Insertar(elemento T) {
	nuevo := nuevoNodo(elemento)
	nuevo.siguiente = it.actual
	it.lista.insertarNodo(it.anterior, nuevo)
	it.actual = nuevo
}

func (it *iterLista[T]) Borrar() T {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := it.lista.borrarNodo(it.anterior)
	if it.anterior == nil {
		it.actual = it.lista.primero
	} else {
		it.actual = it.anterior.siguiente
	}
	return dato
}
