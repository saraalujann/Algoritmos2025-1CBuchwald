package lista

type Lista[T any] interface {
	/*Pre: -
	Post: Devuelve true si la lista esta vacia (no tiene elementos) o false en caso contrario.
	*/
	EstaVacia() bool

	/*Pre: -
	Post: Se añade un elemento en la primera posicion, y sus elementos consecuentes
	se desplazan una posicion. La longitud de la lista aumenta 1.
	*/
	InsertarPrimero(T)

	/*Pre: -
	Post: El elemento se añade al final de la lista, los elementos ya colocados
	mantienen su posicion. La longitud de la lista aumenta 1.
	*/
	InsertarUltimo(T)

	/*Pre: La lista no esta vacia.
	Post: Elimina y devuelve el primer elemento de la lista. La longitud de la lista disminuye 1.
	*/
	BorrarPrimero() T

	/*Pre: La lista no esta vacia.
	Post: Devuelve el primer elemento de la lista sin modificarla.
	*/
	VerPrimero() T

	/*Pre: La lista no esta vacia.
	Post: Devuelve el ultimo elemento de la lista sin modificarla.
	*/
	VerUltimo() T

	/*Pre: -
	Post: Devuelve el numero de elementos actual de la lista. Devuelve 0 si esta vacia.
	*/
	Largo() int

	/*Pre: La funcion visitar no debe ser nil.
	Post: Itera sobre cada elemento de la lista aplicando la funcion visitar. Si visitar
	devuelve false, la iteracion se detiene. Recorre todos los elementos si visitar
	devuelve true hasta el ultimo elemento.
	*/
	Iterar(visitar func(T) bool)

	/*Pre: -
	Post: Devuelve un iterador que es valido mientras la lista original no sea modificada,
	si se modifica, el comportamiento es indefinido.
	*/
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	/*Pre: El iterador esta sobre el elemento actual.
	Post: Devuelve el elemento actual sin modificarlo ni avanzar la posicion.
	No altera el estado del iterador ni de la lista
	*/
	VerActual() T

	/*Pre: -
	Post: Devuelve true si existe un elemento siguiente para recorrer, false en caso contrario.
	*/
	HaySiguiente() bool

	/*Pre: HaySiguiente tiene que ser true.
	Post: Avanza la posicion del iterador al siguiente elemento.
	*/
	Siguiente()

	/*Pre: -
	Post: Inserta el elemento en la posicion actual del iterador.
	*/
	Insertar(T)

	/*Pre: El iterador esta posicionado sobre un elemento actual.
	Post: Elimina y devuelve el elemento actual de la lista. El iterador se posiciona
	sobre el siguiente elemento.
	*/
	Borrar() T
}
