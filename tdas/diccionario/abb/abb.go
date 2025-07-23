package diccionario

import (
	TDAPila "tdas/pila"
)

type comparador[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      comparador[K]
}

type iterAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   comparador[K]
}

//CrearABB: Inicializa un nuevo ABB y reinicia los contadores.
// PreCond. raiz == nil.
// PostCond. Se crea un nuevo ABB con funcion de comparacion y cantidad 0.

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

// Guardar guarda un nuevo par clave-dato en el diccionario.
// Si la clave ya existe, actualiza su valor.
// Pre: -
// Post: la clave está presente en el diccionario y asociada al dato proporcionado. La cantidad de claves puede haber aumentado o no.

func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.guardar(a.raiz, clave, dato)
}

// guardar. funcion llamada en Guardar.
// Pre: necesita la raiz, la clave y el dato.
// Post: devuelve el nodo con la clave y el dato guardado.

func (a *abb[K, V]) guardar(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		a.cantidad++
		return &nodoAbb[K, V]{clave: clave, dato: dato}
	}
	comp := a.cmp(clave, nodo.clave)
	if comp < 0 {
		nodo.izq = a.guardar(nodo.izq, clave, dato)
	} else if comp > 0 {
		nodo.der = a.guardar(nodo.der, clave, dato)
	} else {
		nodo.dato = dato
	}
	return nodo
}

// Pertenece indica si una clave se encuentra en el diccionario.
// Pre: -
// Post: devuelve true si la clave pertenece al diccionario, false si no lo hace.

func (a *abb[K, V]) Pertenece(clave K) bool {
	_, encontrado := a.buscarNodo(a.raiz, clave)
	return encontrado
}

// Pre: necesita un nodo y una clave.
// Post: devuelve el nodo si la clave pertenece al diccionario o nil si no lo hace

func (a *abb[K, V]) buscarNodo(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], bool) {
	if nodo == nil {
		return nil, false
	}
	comp := a.cmp(clave, nodo.clave)
	if comp < 0 {
		return a.buscarNodo(nodo.izq, clave)
	} else if comp > 0 {
		return a.buscarNodo(nodo.der, clave)
	} else {
		return nodo, true
	}
}

// Pre: necesita una clave y una raiz.
// Post: devuelve el dato del resultado de obtenerNodo(), si la clave no pertenece entra en pánico con mensaje 'La clave no pertenece al diccionario'.

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, encontrado := a.buscarNodo(a.raiz, clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

// Borrar
// Pre: necesita la raiz y la clave a borrar
// Pos: devuelve el resultado de borrar()

func (a *abb[K, V]) Borrar(clave K) V {
	nodo, encontrado := a.buscarNodo(a.raiz, clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	dato := nodo.dato
	a.raiz = a.borrar(a.raiz, clave)
	return dato
}

// borrar.
// Pre: necesita un nodo y la clave a borrar.
// Post: devuelve el dato asociado a la clave que se quiere borrar, si la clave no pertenece al diccionario entra en pánico con un mensaje 'La clave no pertenece al diccionario'

func (a *abb[K, V]) borrar(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comp := a.cmp(clave, nodo.clave)
	if comp < 0 {
		nodo.izq = a.borrar(nodo.izq, clave)
	} else if comp > 0 {
		nodo.der = a.borrar(nodo.der, clave)
	} else {
		if nodo.izq == nil && nodo.der == nil {
			a.cantidad--
			return nil
		} else if nodo.izq == nil || nodo.der == nil {
			a.cantidad--
			return a.borrar_con_un_hijo(nodo)
		} else {
			return a.reemplazar(nodo)
		}
	}
	return nodo
}

// Pre: necesita un nodo con hijos
// Post: reemplaza el nodo a eliminar por el max a la izquierda
func (a *abb[K, V]) reemplazar(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	sucesor := buscarMaximo(nodo.izq)
	nodo.clave = sucesor.clave
	nodo.dato = sucesor.dato
	nodo.izq = a.borrar(nodo.izq, sucesor.clave)
	return nodo
}

func (a *abb[K, V]) borrar_con_un_hijo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo.izq != nil {
		return nodo.izq
	}
	return nodo.der
}

// buscarMaximo.
// Pre: necesita un nodo.
// Post: determina si el nodo tiene hijo derecho y lo devuelve.

func buscarMaximo[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	for nodo.der != nil {
		nodo = nodo.der
	}
	return nodo
}

// Cantidad
// Pre: -
// Post: devuelve la cantidad de elementos dentro del diccionario

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

// Iterar aplica la función visitar a cada elemento del diccionario en
// orden, mientras esta devuelva true.
// Pre: visitar no debe ser nil.
// Post: se recorren los elementos en orden hasta que visitar devuelva false o se terminen.

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	a.IterarRango(nil, nil, visitar)
}

// Iterador devuelve un iterador que recorre todos los elementos del diccionario en orden.
// Pre: -
// Post: el iterador apunta al primer elemento en orden, si existe.

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

// IterarRango itera sólo incluyendo a los elementos que se encuentren
// comprendidos en el rango indicado, incluyéndolos en caso de encontrarse
// Pre: -
// Post: Llama a la funcion iterarRango para ejecutar.

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarRango(a.raiz, desde, hasta, visitar)
}

// iterarRango es una funcion que es llamada por IterarRango.
// Pre: tienen que haber valores validos para "desde" y para "hasta"
// Post: itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
//incluyéndolos en caso de encontrarse

func (a *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || a.cmp(nodo.clave, *desde) > 0 {
		if !a.iterarRango(nodo.izq, desde, hasta, visitar) {
			return false
		}
	}
	if (desde == nil || a.cmp(nodo.clave, *desde) >= 0) &&
		(hasta == nil || a.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || a.cmp(nodo.clave, *hasta) < 0 {
		if !a.iterarRango(nodo.der, desde, hasta, visitar) {
			return false
		}
	}
	return true
}

// IteradorRango crea un iterador.
// Pre: la pila debe crearse correctamente.
// Post: devuelve un IterDiccionario para este Diccionario

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	it := &iterAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		cmp:   a.cmp,
	}
	it.iteradorRango(a.raiz, desde, hasta)
	return it
}

// iteradorRango es una funcion que es llamada por IteradorRango.
// Pre: tiene que haber valores validos para "desde" y para "hasta"
// Post: Itera sobre el rango correspondiente.

func (it *iterAbb[K, V]) iteradorRango(nodo *nodoAbb[K, V], desde *K, hasta *K) {
	for nodo != nil {
		if (desde == nil || it.cmp(nodo.clave, *desde) >= 0) &&
			(hasta == nil || it.cmp(nodo.clave, *hasta) <= 0) {
			it.pila.Apilar(nodo)
			nodo = nodo.izq
		} else if hasta != nil && it.cmp(nodo.clave, *hasta) > 0 {
			nodo = nodo.izq
		} else {
			nodo = nodo.der
		}
	}
}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
// Pre: -
// Post: Devuelve true si la pila no esta vacia, caso contrario devuelve false.

func (it *iterAbb[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
// Pre: La pila puede, o no, estar vacia.
// Post: Si la pila no esta vacia muestra su tope, caso contrario muestra un mensaje de panic.

func (it *iterAbb[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.VerTope()
	return nodo.clave, nodo.dato
}

// Siguiente si HaySiguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador termino de iterar'
// Pre: El iteradorRango fue creado correctamente. HaySiguiente puede o no ser true.
// Post: El iterador avanza a la siguiente posición, si HaySiguiente devuelve false, entra al mensaje de panic.
// Aparte el iterador toma el valor desapilado y llama iteradorRango.

func (it *iterAbb[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.Desapilar()
	it.iteradorRango(nodo.der, it.desde, it.hasta)
}
