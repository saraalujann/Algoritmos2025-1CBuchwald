package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila esta vacia despues de haber sido creada")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	pila.Apilar(4)
	require.False(t, pila.EstaVacia(), "La pila no tiene que estar vacia despues de apilar")
	require.Equal(t, 4, pila.Desapilar(), "La pila debe haberse quedado vacia despues de desapilar el unico elemento")
	require.True(t, pila.EstaVacia())
}

func TestPilaVaciaSlice(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[[]bool]()
	sliceVacio := []bool{}
	pila.Apilar(sliceVacio)
	require.Empty(t, pila.Desapilar(), "El slice desapilado se vacia")
}

func TestPilaLIFOEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	pila.Apilar(15)
	pila.Apilar(1)
	pila.Apilar(20)
	require.Equal(t, 20, pila.Desapilar(), "Fallo en el primer desapilado tiene que ser 20")
	require.Equal(t, 1, pila.Desapilar(), "Fallo en el segundo desapilado tiene que ser 1")
	require.Equal(t, 15, pila.Desapilar(), "Fallo en el tercer desapilado tiene que ser 15")
	require.Equal(t, 3, pila.Desapilar(), "Fallo en el cuarto desapilado tiene que ser 3")
}

func TestPilaLIFOStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("primero")
	pila.Apilar("segundo")
	pila.Apilar("tercero")
	require.Equal(t, "tercero", pila.Desapilar(), "Fallo en el primer desapilado tiene que ser la palabra tercero")
	require.Equal(t, "segundo", pila.Desapilar(), "Fallo en el segundo desapilado tiene que ser la palabra segundo")
	require.Equal(t, "primero", pila.Desapilar(), "Fallo en el tercer desapilado tiene que ser la palabra tercero")
}

func TestPilaLIFOSlices(t *testing.T) {
	pilaInt := TDAPila.CrearPilaDinamica[[]int]()
	slice_enteros := []int{1, 2, 3, 4, 5}
	pilaInt.Apilar(slice_enteros)
	require.Equal(t, slice_enteros, pilaInt.Desapilar(), "Falla el desapilado del slice")

	pilaStr := TDAPila.CrearPilaDinamica[[]string]()
	slice_strings := []string{"s", "d", "f"}
	pilaStr.Apilar(slice_strings)
	require.Equal(t, slice_strings, pilaStr.Desapilar(), "Fallo el desapilado del slice pila2")
}

func TestVolumenPilaEnteros(t *testing.T) {
	const VOLUMEN = 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < VOLUMEN; i++ {
		pila.Apilar(i)
	}
	for i := VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar(), "Error en posicion %d", i)
	}
}

func TestComportamientoAlVaciarInt(t *testing.T) {
	pilaNueva := TDAPila.CrearPilaDinamica[int]()
	pilaNueva.Apilar(1)
	pilaNueva.Desapilar()
	require.True(t, pilaNueva.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.VerTope() })
}

func TestComportamientoAlVaciarStr(t *testing.T) {
	pilaNueva := TDAPila.CrearPilaDinamica[string]()
	pilaNueva.Apilar("hello")
	pilaNueva.Desapilar()
	require.True(t, pilaNueva.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaNueva.VerTope() })
}

func TestPilaCreadaInvalidaciones(t *testing.T) {
	pilaCreada1 := TDAPila.CrearPilaDinamica[int]()
	pilaCreada2 := TDAPila.CrearPilaDinamica[string]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCreada1.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCreada2.VerTope() })
}

func TestPilaCreadaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia(), "La pila debe devolver True porque una recien creada == pila vacia")
	pilaNumeros := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaNumeros.EstaVacia(), "La pila debe devolver True porque una recien creada == pila vacia")
}

func TestComportamientoPilaVaciadaInt(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(3)
	pila.Apilar(45)
	pila.Apilar(22)
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila se vacio despues de desapilar todos los elementos apilados")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestComportamientoPilaVaciadaStr(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("argentina")
	pila.Apilar("camion")
	pila.Apilar("Palabra")
	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila se vacio despues de desapilar todos los elementos apilados")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestApilarEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(233)
	pila.Apilar(44)
	pila.Apilar(27)
	pila.Apilar(4)
	pila.Apilar(150)
	require.Equal(t, 150, pila.VerTope(), "El elemento en el tope deberia ser 150")
	pila.Desapilar()
	require.Equal(t, 4, pila.VerTope(), "Despues de desapilar, deberia ser 4")
	pila.Desapilar()
	require.Equal(t, 27, pila.VerTope(), "Despues de desapilar, deberia ser 27")
	pila.Desapilar()
	require.Equal(t, 44, pila.VerTope(), "Despues de desapilar, deberia ser 44")
	pila.Desapilar()
	require.Equal(t, 233, pila.VerTope(), "Despues de desapilar, deberia ser 233")
}

func TestApilarStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	pila.Apilar("adios")
	pila.Apilar("perro")
	pila.Apilar("Hoho")
	pila.Apilar("tomate")
	require.Equal(t, "tomate", pila.VerTope(), "El elemento en el top deberia ser la palabra Hola")
	pila.Desapilar()
	require.Equal(t, "Hoho", pila.VerTope(), "Despues de desapilar, deberia ser la palabra adios")
	pila.Desapilar()
	require.Equal(t, "perro", pila.VerTope(), "Despues de desapilar, deberia ser la palabra perro")
	pila.Desapilar()
	require.Equal(t, "adios", pila.VerTope(), "Despues de desapilar, deberia ser la palabra Hoho")
	pila.Desapilar()
	require.Equal(t, "Hola", pila.VerTope(), "Despues de desapilar, deberia ser la palabra tomate")
}

func TestVerificarPilaDesapilada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(34)
	pila.Apilar(55)
	pila.Apilar(2)
	pila.Desapilar()
	require.False(t, pila.EstaVacia(), "La pila no esta vacia despues de desapilar")
}

func TestVerificarTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(33)
	pila.Apilar(12)
	pila.Apilar(2)
	require.Equal(t, 2, pila.VerTope(), "El tope de la pila es 2")
	pila.Desapilar()
	require.Equal(t, 12, pila.VerTope(), "EL tope de la pila es 12")
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila deberia haber quedado vacia")
}

func TestDesapilaHastaElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("aguila")
	pila.Apilar("buho")
	pila.Apilar("animal")
	pila.Apilar("dedo")
	for !pila.EstaVacia() && pila.VerTope() != "animal" {
		pila.Desapilar()
	}
	require.Equal(t, "animal", pila.VerTope(), "El tope debe ser la palabra animal")
	require.False(t, pila.EstaVacia(), "La pila no debe haber quedado vacia")
}

func TestApilarPostPanic(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(30)
	require.Equal(t, 30, pila.VerTope(), "El tope de la pila es 30")
}

func TestVaciaIntercalada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(9)
	pila.Apilar(15)
	require.False(t, pila.EstaVacia(), "La pila no esta vacia")
	pila.Desapilar()
	require.False(t, pila.EstaVacia(), "La pila no esta vacia")
	pila.Apilar(34)
	require.False(t, pila.EstaVacia(), "La pila no esta vacia")
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila esta vacia")
}
