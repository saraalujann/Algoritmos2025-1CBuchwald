package cola_prioridad_test

import (
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	VOLUMEN_1 = 100000
	VOLUMEN_2 = 10000
)

func cmpInt(a, b int) int {
	return a - b
}

func cmpStr(a, b string) int {
	return strings.Compare(b, a)
}

func TestCrearHeapVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.Equal(t, 0, heap.Cantidad())
}

func TestIntercaladoOperacionesConHeapVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	heap.Encolar(19)
	require.Equal(t, 19, heap.VerMax())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(100)
	heap.Encolar(312)
	require.Equal(t, 312, heap.VerMax())
}

func TestHeapifyArregloVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]string{}, cmpStr)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar("primero")
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "primero", heap.VerMax())
}

func TestHeapifyArregloUnElemento(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]int{7}, cmpInt)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 7, heap.VerMax())
	require.Equal(t, 7, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestVacioDesencolaUnicoElementoInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.Equal(t, 0, heap.Cantidad())
	heap.Encolar(2)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 2, heap.VerMax())
	heap.Desencolar()
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestVacioDesencolaUnicoElementoStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	require.Equal(t, 0, heap.Cantidad())
	heap.Encolar("Unico")
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, "Unico", heap.VerMax())
	heap.Desencolar()
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestEncolaDesencolaInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.Equal(t, 0, heap.Cantidad())
	heap.Encolar(102)
	heap.Encolar(19293)
	heap.Encolar(1)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 19293, heap.VerMax())
	heap.Desencolar()
	require.Equal(t, 102, heap.VerMax())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	heap.Encolar(27)
	require.Equal(t, 27, heap.VerMax())
	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestEncolaDesencolaStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	require.True(t, heap.EstaVacia())
	heap.Encolar("pera")
	heap.Encolar("banana")
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, "banana", heap.VerMax())
	heap.Encolar("durazno")
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, "banana", heap.VerMax())
	heap.Desencolar()
	require.Equal(t, "durazno", heap.VerMax())
	heap.Desencolar()
	require.Equal(t, "pera", heap.VerMax())
	heap.Encolar("ciruela")
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, "ciruela", heap.VerMax())
	heap.Desencolar()
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolaMuchosElementos(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	capacidadInicial := heap.Cantidad()
	elementosAgregar := capacidadInicial*10 + 1
	for i := 0; i < elementosAgregar; i++ {
		heap.Encolar(i)
		require.Equal(t, i+1, heap.Cantidad(), "Cantidad incorrecta despues de encolar")
		require.Equal(t, i, heap.VerMax())
	}
	require.Equal(t, elementosAgregar, heap.Cantidad(), "Cantidad final incorrecta")
	require.Equal(t, elementosAgregar-1, heap.VerMax())
	for i := elementosAgregar - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
		require.Equal(t, i, heap.Cantidad(), "Cantidad incorrecta despues de desencolar")
	}
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad(), "Cantidad deberia ser 0")
}

func TestMaxHeapInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	heap.Encolar(5)
	heap.Encolar(1)
	heap.Encolar(3)
	heap.Encolar(4)
	heap.Encolar(2)
	arreglo_final := []int{5, 4, 3, 2, 1}
	for _, valorDesencolado := range arreglo_final {
		require.Equal(t, valorDesencolado, heap.Desencolar())
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.True(t, heap.EstaVacia())
}

func TestMaxHeapOperacionesIntercaladasStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	heap.Encolar("banana")
	heap.Encolar("anana")
	require.Equal(t, "anana", heap.Desencolar())
	heap.Encolar("cherry")
	heap.Encolar("grape")
	require.Equal(t, "banana", heap.Desencolar())
	heap.Encolar("mango")
	require.Equal(t, "cherry", heap.Desencolar())
	require.Equal(t, "grape", heap.Desencolar())
	require.Equal(t, "mango", heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestMinHeapStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	heap.Encolar("rojo")
	heap.Encolar("azul")
	heap.Encolar("verde")
	heap.Encolar("amarillo")
	heap.Encolar("naranja")
	require.Equal(t, "amarillo", heap.Desencolar())
	require.Equal(t, "azul", heap.Desencolar())
	require.Equal(t, "naranja", heap.Desencolar())
	require.Equal(t, "rojo", heap.Desencolar())
	require.Equal(t, "verde", heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestMinHeapInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.True(t, heap.EstaVacia())
	valores := []int{1000, 1, 50, 200, 9}
	for _, valor := range valores {
		heap.Encolar(valor)
	}
	esperado := []int{1000, 200, 50, 9, 1}
	for _, valor := range esperado {
		require.Equal(t, valor, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapDesdeArreglo(t *testing.T) {
	arr := []int{10, 40, 5, 30, 100}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)
	require.Equal(t, 5, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	resultadoEsperado := []int{100, 40, 30, 10, 5}
	for _, valor := range resultadoEsperado {
		require.Equal(t, valor, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	heap.Encolar(5)
	heap.Encolar(1)
	heap.Encolar(3)
	heap.Encolar(4)
	heap.Encolar(2)
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 4, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	heap.Encolar("h")
	heap.Encolar("c")
	heap.Encolar("a")
	heap.Encolar("z")
	heap.Encolar("v")
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, "c", heap.Desencolar())
	require.Equal(t, "h", heap.Desencolar())
	require.Equal(t, "v", heap.Desencolar())
	require.Equal(t, "z", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestVolumenMaxHeap(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	for i := 0; i <= VOLUMEN_1; i++ {
		heap.Encolar(i)
	}
	for i := VOLUMEN_1; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapReutilizarDespuesDeVaciar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	heap.Encolar(2)
	heap.Encolar(1)
	heap.Encolar(3)
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	heap.Encolar(7)
	require.Equal(t, 7, heap.VerMax())
	heap.Encolar(8)
	require.Equal(t, 8, heap.VerMax())
	heap.Encolar(9)
	require.Equal(t, 9, heap.VerMax())
}

func TestVolumenElementosIgualesStr(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpStr)
	for i := 0; i < VOLUMEN_2; i++ {
		heap.Encolar("b")
		require.Equal(t, "b", heap.VerMax())
	}
	for i := VOLUMEN_2 - 1; i >= 0; i-- {
		require.Equal(t, "b", heap.Desencolar())
	}
}

func TestVolumenElementosIgualesInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	for i := 0; i < VOLUMEN_1; i++ {
		heap.Encolar(10)
		require.Equal(t, 10, heap.VerMax())
	}
	for i := VOLUMEN_1 - 1; i >= 0; i-- {
		require.Equal(t, 10, heap.Desencolar())
	}
}

func TestIntercaladoVolumen(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	for i := 0; i < VOLUMEN_1; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
	}
	for j := VOLUMEN_1; j >= VOLUMEN_2; j-- {
		heap.Desencolar()
	}
	for k := VOLUMEN_2; k <= VOLUMEN_1; k++ {
		heap.Encolar(k)
		require.Equal(t, k, heap.VerMax())
	}
}

func TestVerMaximo(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	heap.Encolar(100)
	require.Equal(t, 100, heap.VerMax())
}

func TestVerMaximoUnElemento(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.Equal(t, 0, heap.Cantidad())
	heap.Encolar(97)
	require.Equal(t, 97, heap.VerMax())
	require.False(t, heap.EstaVacia())
	heap.Desencolar()
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestElementosIgualesInt(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestElementosIgualesStr(t *testing.T) {
	arr := []string{"a", "a", "a", "a", "a"}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpStr)
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestArregloOrdenado(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 4, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapifySubconjuntoArreglo(t *testing.T) {
	arr := []int{100, 80, 60, 40, 20, 0}
	subconjunto := arr[2:5]
	heap := TDAColaPrioridad.CrearHeapArr(subconjunto, cmpInt)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 60, heap.VerMax())
	require.Equal(t, 60, heap.Desencolar())
	require.Equal(t, 40, heap.Desencolar())
	require.Equal(t, 20, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapVacioVerMaximoDesencolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.Panics(t, func() { heap.VerMax() })
}

func TestHeapMinInt(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int {
		return -cmpInt(a, b)
	})
	heap.Encolar(3)
	heap.Encolar(1)
	heap.Encolar(5)
	require.Equal(t, 1, heap.VerMax())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 3, heap.VerMax())
}

func TestHeapifyYOperaciones(t *testing.T) {
	arr := []int{3, 6, 9}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)
	require.Equal(t, 9, heap.VerMax())
	heap.Encolar(15)
	require.Equal(t, 15, heap.VerMax())
}
