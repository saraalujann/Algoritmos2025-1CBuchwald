package diccionario

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
)

type estadoCelda int

const (
	_ESTADO_VACIO estadoCelda = iota
	_ESTADO_OCUPADO
	_ESTADO_BORRADO
	_CAPACIDAD_INICIAL = 16
	_FACTOR_CARGA_MAX  = 0.65
	_FACTOR_BORRADOS   = 0.5
	_FACTOR_AUMENTO    = 2
	_FACTOR_REDUCCION  = 0.5
)

type celdaHash[K comparable, V any] struct {
	clave  K
	valor  V
	estado estadoCelda
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	tam      int
	cantidad int
	borrados int
	hasher   hash.Hash64
}

type iterCerrado[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	posicion int
}

/* crearTabla: Inicializa la tabla de hash con un tamaño especifico y reinicia los contadores.

PreCond. nuevoTam debe ser una potencia de 2.
PostCond. Se crea una nueva tabla vacía de tamaño 'nuevoTam'. */

func (h *hashCerrado[K, V]) crearTabla(nuevoTam int) {
	h.tabla = make([]celdaHash[K, V], nuevoTam)
	h.tam = nuevoTam
	h.cantidad = 0
	h.borrados = 0
}

/* CrearHash: Inicializa la tabla de hash con un tamaño especifico y reinicia los contadores.

PreCond. nuevoTam debe ser una potencia de 2.
PostCond. Se crea una tabla vacia de tamaño 'nuevoTam'. */

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	h := &hashCerrado[K, V]{}
	h.hasher = fnv.New64a()
	h.crearTabla(_CAPACIDAD_INICIAL)
	return h
}

/* convertirABytes: Convierte una clave de tipo generico a un slice de bytes.

PreCond. 'clave' tiene que ser un tipo que se pueda imprimir como string.
PostCond. Devuelve la representacion en bytes. */

func convertirABytes[K comparable](clave K) []byte {
	if str, ok := any(clave).(string); ok {
		return []byte(str)
	}
	return []byte(fmt.Sprintf("%v", clave))
}

/* fnv1aHash: Calcula el hash de una secuencia de bytes usando el algoritmo FNV-1a.

PreCond. 'data' debe ser un slice de bytes válido.
PostCond. Devuelve un uint64 con el hash de la entrada. */

func (h *hashCerrado[K, V]) fnv1aHash(data []byte) uint64 {
	h.hasher.Reset()
	h.hasher.Write(data)
	return h.hasher.Sum64()
}

/* posInicial: Calcula la posición inicial de una clave en la tabla de hash.

PreCond. La tabla tiene que estar inicializada.
PostCond. Devuelve una posición valida en la tabla.  */

func (h *hashCerrado[K, V]) posInicial(clave K) int {
	return int(h.fnv1aHash(convertirABytes(clave))) & (h.tam - 1)
}

/* buscarIndice: Busca el índice de una clave o la mejor posición para insertarla.

PreCond. La tabla tiene que estar inicializada.
PostCond. Devuelve el indice donde se encuentra o se puede insertar la clave, y su estado. */

func (h *hashCerrado[K, V]) buscarIndice(clave K) (int, estadoCelda) {
	idx := h.posInicial(clave)
	primerBorrado := -1
	for i := 0; i < h.tam; i++ {
		c := &h.tabla[idx]
		switch c.estado {
		case _ESTADO_VACIO:
			if primerBorrado != -1 {
				return primerBorrado, _ESTADO_BORRADO
			}
			return idx, _ESTADO_VACIO
		case _ESTADO_OCUPADO:
			if c.clave == clave {
				return idx, _ESTADO_OCUPADO
			}
		case _ESTADO_BORRADO:
			if primerBorrado == -1 {
				primerBorrado = idx
			}
		}
		idx = (idx + 1) % h.tam
	}
	return primerBorrado, _ESTADO_BORRADO
}

/* insertarInterno: Inserta una clave y valor sin validar factor de carga ni redimensionar.

PreCond. La tabla tiene que estar inicializada.
PostCond. Inserta la clave y el valor, o actualiza el valor si ya existe. */

func (h *hashCerrado[K, V]) insertarInterno(clave K, valor V) {
	idx, est := h.buscarIndice(clave)
	if est == _ESTADO_OCUPADO {
		h.tabla[idx].valor = valor
	} else {
		h.tabla[idx] = celdaHash[K, V]{clave, valor, _ESTADO_OCUPADO}
		h.cantidad++
		if est == _ESTADO_BORRADO {
			h.borrados--
		}
	}
}

/*Guardar: guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado.

PreCond. La tabla tiene que estar inicializada.
PostCond. Inserta o actualiza la clave. La tabla puede haber sido redimensionada. */

func (h *hashCerrado[K, V]) Guardar(clave K, valor V) {
	carga := float64(h.cantidad+h.borrados) / float64(h.tam)
	if carga >= _FACTOR_CARGA_MAX {
		nuevoTam := h.tam * _FACTOR_AUMENTO
		h.redimensionar(nuevoTam)
	}
	h.insertarInterno(clave, valor)
}

/* Obtener: Devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje 'La clave no pertenece al diccionario'

PreCond. La clave debe existir, o no, en el diccionario.
PostCond. Devuelve el valor que corresponde o 'panic' si no encuentra nada. */

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	idx, est := h.buscarIndice(clave)
	if est != _ESTADO_OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	return h.tabla[idx].valor
}

/* Borrar: Borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'

PreCond. La clave debe existir, o no, en el diccionario.
PostCond. Marca la celda como BORRADA y devuelve el valor eliminado. O ´panic´ si no encuentra. */

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	idx, est := h.buscarIndice(clave)
	if est != _ESTADO_OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	val := h.tabla[idx].valor
	h.tabla[idx].estado = _ESTADO_BORRADO
	h.cantidad--
	h.borrados++
	tomb := float64(h.borrados) / float64(h.tam)
	if tomb > _FACTOR_BORRADOS {
		nuevoTam := int(math.Max(float64(_CAPACIDAD_INICIAL), float64(h.tam)*_FACTOR_REDUCCION))
		h.redimensionar(nuevoTam)
	}
	return val
}

/* Pertenece: Determina si una clave ya se encuentra en el diccionario, o no.

PreCond. -
PostCond. Devuelve true si la clave esta presente, false en caso contrario. */

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, est := h.buscarIndice(clave)
	return est == _ESTADO_OCUPADO
}

/* Cantidad: devuelve la cantidad de elementos dentro del diccionario.

PreCond. -
PostCond. Devuelve el numero de claves insertadas y no borradas.  */

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cantidad
}

/*redimensionar: Crea una nueva tabla de tamaño distinto y reinserta los elementos existentes.

PreCond. 'nuevoTam' > 0
PostCond. El diccionario va a contener los mismos elementos, pero sin celdas borradas. */

func (h *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	vieja := h.tabla
	h.crearTabla(nuevoTam)
	for i := range vieja {
		if vieja[i].estado == _ESTADO_OCUPADO {
			h.insertarInterno(vieja[i].clave, vieja[i].valor)
		}
	}
}

/*Iterar: itera internamente el diccionario, aplicando la función pasada por parámetro
a todos los elementos del mismo

PreCond. 'visitar' no tiene que ser nil.
PostCond. Ejecuta visitar(clave, valor) por cada valor y se detiene si devuelve false. */

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	for i := 0; i < h.tam; i++ {
		c := h.tabla[i]
		if c.estado == _ESTADO_OCUPADO {
			if !visitar(c.clave, c.valor) {
				return
			}
		}
	}
}

/* Iterador: Devuelve un IterDiccionario para este Diccionario.

PreCond.-
PostCond. Devuelve un iterador valido. */

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterCerrado[K, V]{hash: h, posicion: -1}
	iter.avanzar()
	return iter
}

/* Avanzar: Avanza el iterador a la siguiente celda ocupada..

PreCond. 'iter' debe estar asociado a un hash válido.
PostCond. 'iter.posicion' apunta a la siguiente clave válida o al final. */

func (iter *iterCerrado[K, V]) avanzar() {
	iter.posicion++
	for iter.posicion < iter.hash.tam && iter.hash.tabla[iter.posicion].estado != _ESTADO_OCUPADO {
		iter.posicion++
	}
}

/* HaySiguiente: devuelve si hay más datos para ver. Esto es,
si en el lugar donde se encuentra parada el iterador hay un elemento.

PreCond. -
PostCond. Devuelve true si hay un siguiente elemento válido.. */

func (iter *iterCerrado[K, V]) HaySiguiente() bool {
	return iter.posicion < iter.hash.tam
}

/* VerActual: devuelve la clave y el dato del elemento actual en el que se encuentra
posicionado el iterador. Si no HaySiguiente, debe entrar en pánico con el mensaje
'El iterador termino de iterar'

PreCond. HaySiguiente() puede o no ser true.
PostCond. Devuelve la clave y el valor actual si hay, caso contrario un mensaje de 'panic'. */

func (iter *iterCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	c := iter.hash.tabla[iter.posicion]
	return c.clave, c.valor
}

/* Siguiente: si HaySiguiente avanza al siguiente elemento en el diccionario.
Si no HaySiguiente, entonces debe entrar en pánico con mensaje
'El iterador termino de iterar'

PreCond. El iterador fue creado correctamente con Iterador(). HaySiguiente puede o no ser true.
PostCond. El iterador avanza a la siguiente posición con una celda OCUPADO, si HaySiguiente
devuelve false, entra al mensaje de panic. */

func (iter *iterCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.avanzar()
}
