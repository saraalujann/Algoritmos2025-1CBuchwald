package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	VOLUMEN = 99999
	NUMERO  = 10000
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola esta vacia despues de haber sido creada")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar(4)
	require.False(t, cola.EstaVacia(), "La cola no tiene que estar vacia despues de encolar")
	require.Equal(t, 4, cola.Desencolar(), "La cola debe haberse quedado vacia despues de desencolar el unico elemento")
	require.True(t, cola.EstaVacia())
}

func TestColaFIFOInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(2)
	cola.Encolar(45)
	cola.Encolar(91929)
	cola.Encolar(1)
	require.Equal(t, 2, cola.Desencolar(), "Fallo en el primer desencolado tiene que ser 2")
	require.Equal(t, 45, cola.Desencolar(), "Fallo en el segundo desencolado tiene que ser 45")
	require.Equal(t, 91929, cola.Desencolar(), "Fallo en el tercer desencolado tiene que ser 91929")
	require.Equal(t, 1, cola.Desencolar(), "Fallo en el cuarto desencolado tiene que ser 1")
}

func TestColaFIFOStr(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("primero")
	cola.Encolar("segundo")
	cola.Encolar("tercero")
	require.Equal(t, "primero", cola.Desencolar(), "Fallo en el primer desencolado tiene que ser la palabra tercero")
	require.Equal(t, "segundo", cola.Desencolar(), "Fallo en el segundo desencolado tiene que ser la palabra segundo")
	require.Equal(t, "tercero", cola.Desencolar(), "Fallo en el tercer desencolado tiene que ser la palabra tercero")
}

func TestVolumenColaInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < VOLUMEN; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < VOLUMEN; i++ {
		require.Equal(t, i, cola.Desencolar(), "Error en posicion %d", i)
	}
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
}

func TestComportamientoAlVaciarInt(t *testing.T) {
	colaNueva := TDACola.CrearColaEnlazada[int]()
	colaNueva.Encolar(10000)
	colaNueva.Desencolar()
	require.True(t, colaNueva.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaNueva.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaNueva.VerPrimero() })
}

func TestComportamientoAlVaciarStr(t *testing.T) {
	colaNueva := TDACola.CrearColaEnlazada[string]()
	colaNueva.Encolar("perro")
	colaNueva.Desencolar()
	require.True(t, colaNueva.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaNueva.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaNueva.VerPrimero() })
}

func TestColaCreadaInvalidaciones(t *testing.T) {
	colaCreada1 := TDACola.CrearColaEnlazada[int]()
	colaCreada2 := TDACola.CrearColaEnlazada[string]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCreada1.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCreada2.VerPrimero() })
}

func TestColaCreadaVacia(t *testing.T) {
	colaStr := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaStr.EstaVacia(), "La cola debe devolver True porque esta vacia")
	colaInt := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaInt.EstaVacia(), "La cola debe devolver True porque esta vacia")
}

func TestComportamientoColaVaciadaInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(55)
	cola.Encolar(102)
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola se vacio despues de desencolar todos los elementos encolados")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestComportamientoColaVaciadaStr(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	cola.Encolar("hello")
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola se vacio despues de desencolar todos los elementos encolados")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestEncolarInt(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(102)
	cola.Encolar(33)
	cola.Encolar(68)
	cola.Encolar(122)
	cola.Encolar(101)
	require.Equal(t, 102, cola.VerPrimero(), "El elemento en el tope deberia ser 102")
	cola.Desencolar()
	require.Equal(t, 33, cola.VerPrimero(), "Despues de desencolar, deberia ser 33")
	cola.Desencolar()
	require.Equal(t, 68, cola.VerPrimero(), "Despues de desencolar, deberia ser 68")
	cola.Desencolar()
	require.Equal(t, 122, cola.VerPrimero(), "Despues de desencolar, deberia ser 122")
	cola.Desencolar()
	require.Equal(t, 101, cola.VerPrimero(), "Despues de desencolar, deberia ser 101")
}

func TestEncolarStr(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	cola.Encolar("hello")
	cola.Encolar("algoritmo")
	cola.Encolar("algoritmos")
	cola.Encolar("materias")
	require.Equal(t, "hola", cola.VerPrimero(), "El elemento en el tope deberia ser 102")
	cola.Desencolar()
	require.Equal(t, "hello", cola.VerPrimero(), "Despues de desencolar, deberia ser 33")
	cola.Desencolar()
	require.Equal(t, "algoritmo", cola.VerPrimero(), "Despues de desencolar, deberia ser 68")
	cola.Desencolar()
	require.Equal(t, "algoritmos", cola.VerPrimero(), "Despues de desencolar, deberia ser 122")
	cola.Desencolar()
	require.Equal(t, "materias", cola.VerPrimero(), "Despues de desencolar, deberia ser 101")
}

func TestEncoladoDesencoladoConstante(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < NUMERO; i++ {
		cola.Encolar(i)
		require.Equal(t, i, cola.Desencolar())
		require.True(t, cola.EstaVacia())
	}
}

func TestVerPrimeroSinDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("primero")
	for i := 0; i < NUMERO; i++ {
		require.Equal(t, "primero", cola.VerPrimero())
	}
	require.Equal(t, "primero", cola.Desencolar())
}

func TestDesencolaHastaElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("programacion")
	cola.Encolar("matematica")
	cola.Encolar("matematicas")
	cola.Encolar("probabilidad")
	for !cola.EstaVacia() && cola.VerPrimero() != "matematica" {
		cola.Desencolar()
	}
	require.Equal(t, "matematica", cola.VerPrimero(), "El tope debe ser la palabra matematica")
	require.False(t, cola.EstaVacia(), "La cola no debe haber quedado vacia")
}

func TestEncolarPostPanic(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	cola.Encolar(134)
	require.Equal(t, 134, cola.VerPrimero(), "El tope de la cola es 30")
}

func TestVaciaIntercalada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(77)
	cola.Encolar(129)
	require.False(t, cola.EstaVacia(), "La cola no esta vacia")
	cola.Desencolar()
	require.False(t, cola.EstaVacia(), "La cola no esta vacia")
	cola.Encolar(21)
	require.False(t, cola.EstaVacia(), "La cola no esta vacia")
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
}

func TestChequeoNodos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("adios")
	cola.Encolar("mariposa")
	cola.Encolar("coding")
	require.Equal(t, "adios", cola.Desencolar())
	require.Equal(t, "mariposa", cola.VerPrimero())
	require.Equal(t, "mariposa", cola.Desencolar())
	require.Equal(t, "coding", cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
}

func TestComportamientoUnNodo(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(27)
	require.Equal(t, 27, cola.VerPrimero())
	require.Equal(t, 27, cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestConStructsFIFO(t *testing.T) {
	type Materia struct {
		Nombre string
		nota   int
	}
	cola := TDACola.CrearColaEnlazada[Materia]()
	cola.Encolar(Materia{"algebra", 9})
	cola.Encolar(Materia{"analisis", 10})
	require.Equal(t, Materia{"algebra", 9}, cola.Desencolar())
	require.Equal(t, Materia{"analisis", 10}, cola.VerPrimero())
	require.False(t, cola.EstaVacia(), "La cola no esta vacia")
	require.Equal(t, Materia{"analisis", 10}, cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
}

func TestVerPrimeroMultiplesIteraciones(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(193)
	for i := 0; i < NUMERO; i++ {
		require.Equal(t, 193, cola.VerPrimero())
	}
	require.Equal(t, 193, cola.Desencolar())
	require.True(t, cola.EstaVacia(), "L cola esta vacia")
}
