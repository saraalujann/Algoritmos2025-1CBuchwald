package cola_prioridad

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

const (
	CAPACIDAD_INICIAL  = 1
	CAPACIDAD_MINIMA   = 1
	FACTOR_CRECIMIENTO = 2
	FACTOR_REDUCCION   = 4
)

/*CrearHeap. Inicializa un nuevo heap y reinicia todas sus condiciones.
PreCond: funcion_cmp no debe ser nil y debe establecer un orden.
PostCond: Devuelve un heap vacio y con su funcion de comparacion para las operaciones.
*/

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{cantidad: 0, cmp: funcion_cmp}
}

/*CrearHeapArr. Crea un heap a partir de un arreglo usando Heapify.
PreCond: El arreglo no puede estar vacio.
PostCond: Si el arreglo tiene elementos, crea una copia y la transforma en un heap valido.
*/

func CrearHeapArr[T any](arr []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	aux := make([]T, len(arr))
	copy(aux, arr)
	heapify(aux, funcion_cmp)
	return &heap[T]{datos: aux, cantidad: len(aux), cmp: funcion_cmp}
}

/*HeapSort. Ordena un arreglo in-place usando el algoritmo HeapSort.
PreCond: -.
PostCond: El arreglo quedara ordenado segun su funcion_cmp. Se modifica
el arreglo original.
*/

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	n := len(elementos)
	for i := n - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, 0, i, funcion_cmp)
	}
}

/*EstaVacia. Es una funcion que verifica si el heap esta vacio o no.
PreCond: -.
PostCond: Devuelve true si esta vacia, false en caso contrario.
*/

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

/*Encolar. Es una funcion que encola un elemento al heap.
PreCond: El heap debe estar inicializado.
PostCond: El heap va a tener un elemento mas, en el caso de que el heap estuviese
vacio, este elemento se convertiria en la raiz. Se aplica upHeap para
reordenarlo.
*/

func (heap *heap[T]) Encolar(dato T) {
	if heap.cantidad == cap(heap.datos) {
		heap.redimensiona(cap(heap.datos) * FACTOR_CRECIMIENTO)
	}
	heap.datos[heap.cantidad] = dato
	upHeap(heap.datos, heap.cantidad, heap.cmp)
	heap.cantidad++
}

/*VerMax. Es una funcion que devuelve el elemento con mayor prioridad.
PreCond: El heap debe estar inicializado y no debe estar vacio.
PostCond: Si el heap tenia solo un elemento ese va a ser el de mayor prioridad, si tiene
mas se usa la funcion de comparacion, y si no tiene entra en panic. No modifica el heap.
*/

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

/*Desencolar. Es una funcion que desencola un elemento del heap y devuelve el
elemento desencolado manteniendo la propiedad del heap.
PreCond: El heap debe estar inicializado y no debe estar vacio.
PostCond: Si el heap tenia solo un elemento quedara vacio, caso contrario, se
elimina la raiz y se aplica downHeap para reordenarlo. Si esta vacio entra en panic.
*/

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := heap.datos[0]
	heap.datos[0] = heap.datos[heap.cantidad-1]
	heap.cantidad--
	downHeap(heap.datos[:heap.cantidad], 0, heap.cantidad, heap.cmp)
	if len(heap.datos) > CAPACIDAD_INICIAL && heap.cantidad <= len(heap.datos)/FACTOR_REDUCCION {
		nuevaCapacidad := len(heap.datos) / FACTOR_CRECIMIENTO
		if nuevaCapacidad < CAPACIDAD_INICIAL {
			nuevaCapacidad = CAPACIDAD_INICIAL
		}
		if nuevaCapacidad != len(heap.datos) {
			heap.redimensiona(nuevaCapacidad)
		}
	}
	return dato
}

/*Cantidad. Es una funcion que devuelve la cantidad de elementos en la cola de prioridad
PreCond: -.
PostCond: Devuelve la cantidad de elementos validos en la cola de prioridad.
*/

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

/*downHeap. Es una funcion que cumple con la propiedad para reordenar en el tipo
de heap Max-Heap.
PreCond: Tiene que tener un arreglo para reordenar y su funcion de comparacion
para reacomodar por prioridad.
PostCond: Devuelve el heap reordenado con la propiedad downHeap.
*/

func downHeap[T any](arr []T, posicion int, cantidad int, funcion_cmp func(T, T) int) {
	hijo_izq := 2*posicion + 1
	hijo_der := 2*posicion + 2
	posicionMax := posicion
	if hijo_izq < cantidad && funcion_cmp(arr[hijo_izq], arr[posicionMax]) > 0 {
		posicionMax = hijo_izq
	}
	if hijo_der < cantidad && funcion_cmp(arr[hijo_der], arr[posicionMax]) > 0 {
		posicionMax = hijo_der
	}
	if posicionMax != posicion {
		arr[posicion], arr[posicionMax] = arr[posicionMax], arr[posicion]
		downHeap(arr, posicionMax, cantidad, funcion_cmp)
	}
}

/*upHeap. Es una funcion que cumple con la propiedad para reordenar en el tipo
de heap Min-Heap.
PreCond: Tiene que tener un arreglo para reordenar y su funcion de comparacion
para reacomodar por prioridad.
PostCond: Devuelve el heap reordenado con la propiedad upHeap.
*/

func upHeap[T any](arr []T, posicion int, funcion_cmp func(T, T) int) {
	for posicion > 0 {
		padre := (posicion - 1) / 2
		if funcion_cmp(arr[posicion], arr[padre]) <= 0 {
			break
		}
		arr[posicion], arr[padre] = arr[padre], arr[posicion]
		posicion = padre
	}
}

/*Heapify. Es una funcion que reordena el arreglo en un heap valido
PreCond: El arreglo no debe ser nil
PostCond: Devuelve el arreglo con la propiedad de Heapify, usando
el tipo max - heap (downHeap)
*/

func heapify[T any](arr []T, cmp func(T, T) int) {
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		downHeap(arr, i, n, cmp)
	}
}

/*Redimensiona. Es una funcion redimensiona el arreglo.
PreCond: El arreglo no debe ser nil
PostCond: Devuelve el arreglo con la propiedad de redimension */

func (heap *heap[T]) redimensiona(nuevaCapacidad int) {
	if nuevaCapacidad < CAPACIDAD_INICIAL {
		nuevaCapacidad = CAPACIDAD_INICIAL
	}
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos[:heap.cantidad])
	heap.datos = nuevosDatos
}
