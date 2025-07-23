package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	VOLUMEN_1 = 10000
	VOLUMEN_2 = 100000
)

func TestInsertarEnPrincipioIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(100)
	require.False(t, lista.EstaVacia(), "La lista tiene un elemento")
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 100, lista.VerPrimero())
	require.Equal(t, 100, lista.VerUltimo())
	require.True(t, iterador.HaySiguiente())
	require.Equal(t, 100, iterador.VerActual())
}

func TestInsertarMultiplePrincipioIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iterador := lista.Iterador()
	iterador.Insertar("Hola")
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "Hola", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
	require.Equal(t, "Hola", iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Insertar("Algoritmos")
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, "Algoritmos", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
	require.Equal(t, "Algoritmos", iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
	iterador.Insertar("Programacion")
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, "Programacion", lista.VerPrimero())
	require.Equal(t, "Hola", lista.VerUltimo())
	require.Equal(t, "Programacion", iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
}

func TestInsertarElementoAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.Largo())
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	iter.Insertar(3)
	require.Equal(t, 3, lista.Largo())
	iter2 := lista.Iterador()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 2, iter2.VerActual())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 3, iter2.VerActual())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func TestInsertarElementoEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	require.Equal(t, 2, lista.Largo())
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	iter2 := lista.Iterador()
	require.Equal(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.Equal(t, 2, iter2.VerActual())
	iter2.Siguiente()
	require.Equal(t, 3, iter2.VerActual())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func TestRemueveAlCrearIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	iter := lista.Iterador()
	elem := iter.Borrar()
	require.Equal(t, 10, elem, "El primer elemento debería ser 10")
	require.Equal(t, 2, lista.Largo(), "La lista debe tener 2 elementos después de borrar uno")
	require.Equal(t, 20, lista.VerPrimero(), "El primer elemento debería ser 20 ahora")
	require.Equal(t, 30, lista.VerUltimo(), "El último elemento debe seguir siendo 30")
	require.True(t, iter.HaySiguiente(), "El iterador debe tener un siguiente después de borrar el primer elemento")
	require.Equal(t, 20, iter.VerActual(), "El iterador debería estar ahora en el valor 20")
	iter.Siguiente()
	require.Equal(t, 30, iter.VerActual(), "El iterador debe estar ahora en el valor 30")
	require.True(t, iter.HaySiguiente(), "El iterador no debe tener un siguiente después del valor 30")
	require.Equal(t, 2, lista.Largo(), "La lista no debe haber cambiado de tamaño")
	require.Equal(t, 20, lista.VerPrimero(), "El primer elemento de la lista debe seguir siendo 20")
	require.Equal(t, 30, lista.VerUltimo(), "El último elemento de la lista debe seguir siendo 30")
}

func TestRemoverElementoConIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for _, v := range []int{10, 20, 30} {
		lista.InsertarUltimo(v)
	}
	iter := lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 20, iter.Borrar())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 30, lista.VerUltimo())
	require.Equal(t, 30, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 30, lista.VerUltimo())
}

func TestVerificarEliminadoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	iter.Siguiente()
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	iter2 := lista.Iterador()
	require.Equal(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.Equal(t, 3, iter2.VerActual())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func TestInsertaEnCasoBordeIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter1 := lista.Iterador()
	iter1.Insertar(1)
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 1, lista.VerPrimero())
	lista.InsertarUltimo(3)
	iter2 := lista.Iterador()
	iter2.Insertar(0)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	iter3 := lista.Iterador()
	for iter3.HaySiguiente() {
		iter3.Siguiente()
	}
	iter3.Insertar(4)
	require.Equal(t, 4, lista.Largo())
	require.Equal(t, 4, lista.VerUltimo())
	iter4 := lista.Iterador()
	require.Equal(t, 0, iter4.VerActual())
	iter4.Siguiente()
	require.Equal(t, 1, iter4.VerActual())
	iter4.Siguiente()
	require.Equal(t, 3, iter4.VerActual())
	iter4.Siguiente()
	require.Equal(t, 4, iter4.VerActual())
	iter4.Siguiente()
	require.False(t, iter4.HaySiguiente())
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
}

func TestBorraSecuencialmenteIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 3, iter.Borrar())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.False(t, iter.HaySiguiente())
}

func TestIterarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cond := false
	lista.Iterar(func(valor int) bool {
		cond = true
		return true
	})
	require.False(t, cond)
}

func TestInsertaEliminaAlternadosIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{1, 3, 5}
	for _, v := range valores {
		lista.InsertarUltimo(v)
	}
	iter1 := lista.Iterador()
	iter1.Insertar(0)
	require.True(t, iter1.HaySiguiente())
	iter1.Siguiente()
	require.True(t, iter1.HaySiguiente())
	require.Equal(t, 1, iter1.VerActual())
	iter1.Borrar()
	require.Equal(t, 3, lista.Largo())
	iter1.Siguiente()
	require.True(t, iter1.HaySiguiente())
	require.Equal(t, 5, iter1.VerActual())
	iter1.Insertar(2)
	require.Equal(t, 2, iter1.VerActual())
	iter2 := lista.Iterador()
	require.True(t, iter2.HaySiguiente())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 2, iter2.VerActual())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 5, iter2.VerActual())
	iter2.Siguiente()
	require.False(t, iter2.HaySiguiente())
}

func TestMultiplesIteradores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter1 := lista.Iterador()
	iter2 := lista.Iterador()
	require.True(t, iter1.HaySiguiente())
	require.Equal(t, 1, iter1.VerActual())
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 1, iter2.VerActual())
	iter2.Siguiente()
	require.True(t, iter2.HaySiguiente())
	require.Equal(t, 2, iter2.VerActual())
	iter2.Borrar()
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, iter1.VerActual())
	iter3 := lista.Iterador()
	require.Equal(t, 1, iter3.VerActual())
	iter3.Siguiente()
	require.Equal(t, 3, iter3.VerActual())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	cont := 0
	lista.Iterar(func(valor int) bool {
		cont++
		if valor == 10 {
			require.Equal(t, 10, valor)
		} else if valor == 20 {
			require.Equal(t, 20, valor)
		} else if valor == 30 {
			require.Equal(t, 30, valor)
		}
		return true
	})
	require.Equal(t, 3, cont)
}

func TestIteradorExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valores := []int{1, 2, 3, 5, 6, 2, 3}
	for _, valor := range valores {
		lista.InsertarUltimo(valor)
	}
	iterador := lista.Iterador()
	for i := 0; i < len(valores); i++ {
		require.True(t, iterador.HaySiguiente(), "Deberia haber siguiente en posicion %d", i)
		require.Equal(t, valores[i], iterador.VerActual(), "Valor incorrecto en posicion %d", i)
		iterador.Siguiente()
	}
	require.False(t, iterador.HaySiguiente(), "No hay elementos para iterar")
}
func TestListaCreadaVacia(t *testing.T) {
	listaStr := TDALista.CrearListaEnlazada[string]()
	require.True(t, listaStr.EstaVacia(), "La lista string debe estar vacía")
	require.Equal(t, 0, listaStr.Largo(), "El largo debe ser 0")
	require.Panics(t, func() { listaStr.VerPrimero() }, "VerPrimero debería panic en lista vacía")
	require.Panics(t, func() { listaStr.VerUltimo() }, "VerUltimo debería panic en lista vacía")
	listaInt := TDALista.CrearListaEnlazada[int]()
	require.True(t, listaInt.EstaVacia(), "La lista int debe estar vacía")
	require.Equal(t, 0, listaInt.Largo(), "El largo debe ser 0")
	require.Panics(t, func() { listaInt.VerPrimero() }, "VerPrimero debería panic en lista vacía")
	require.Panics(t, func() { listaInt.VerUltimo() }, "VerUltimo debería panic en lista vacía")
}

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista está vacía después de ser creada")
	require.Equal(t, 0, lista.Largo(), "El largo debe ser 0 en una lista vacía")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	lista.InsertarPrimero(4)
	require.False(t, lista.EstaVacia(), "La lista no está vacía después de insertar un elemento")
	require.Equal(t, 1, lista.Largo(), "El largo debe ser 1 después de insertar un elemento")
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "La lista está vacía después de borrar el único elemento")
	require.Equal(t, 0, lista.Largo(), "El largo debe volver a 0")
}

func TestListaVaciaMultiple(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	lista.InsertarPrimero(4)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarPrimero(183)
	lista.InsertarUltimo(172)
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, 183, lista.VerPrimero())
	require.Equal(t, 172, lista.VerUltimo())
	iterador := lista.Iterador()
	iterador.Borrar()
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 4, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertaEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(923)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 923, lista.VerPrimero())
	require.Equal(t, 923, lista.VerUltimo())
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 923, iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarMultiplesVeces(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	require.True(t, lista.EstaVacia(), "La lista está vacía")
	lista.InsertarPrimero("Algoritmos")
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "Algoritmos", lista.VerPrimero())
	require.Equal(t, "Algoritmos", lista.VerUltimo())
	lista.InsertarUltimo("Algebra")
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, "Algoritmos", lista.VerPrimero())
	require.Equal(t, "Algebra", lista.VerUltimo())
	iterador := lista.Iterador()
	iterador.Insertar("Sapito")
	require.Equal(t, 3, lista.Largo())
	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente())
	require.Equal(t, "Sapito", iter.VerActual())
	iter.Siguiente()
	require.Equal(t, "Algoritmos", iter.VerActual())
	iter.Siguiente()
	require.Equal(t, "Algebra", iter.VerActual())
	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "No debería haber más elementos para iterar")
}

func TestInsertaPrincipioFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	valoresIniciales := []int{1, 5, 2, 7, 12, 45}
	for _, v := range valoresIniciales {
		lista.InsertarUltimo(v)
	}
	require.Equal(t, 6, lista.Largo(), "Largo inicial incorrecto")
	iterInicio := lista.Iterador()
	iterInicio.Insertar(0)
	require.Equal(t, 7, lista.Largo(), "Largo después de insertar al inicio")
	require.Equal(t, 0, lista.VerPrimero(), "0 debería ser el nuevo primer elemento")
	require.Equal(t, 45, lista.VerUltimo(), "45 debería seguir siendo el último")
	require.Equal(t, 0, iterInicio.VerActual(), "El iterador debería estar apuntando al 0 después de la inserción")
	iterFinal := lista.Iterador()
	for i := 0; i < lista.Largo()-2; i++ {
		iterFinal.Siguiente()
	}
	iterFinal.Siguiente()
	require.Equal(t, 45, iterFinal.VerActual(), "Debería estar posicionado en 45")

	valorBorrado := iterFinal.Borrar()
	require.Equal(t, 45, valorBorrado, "Debería haber borrado 45")
	require.Equal(t, 6, lista.Largo(), "Largo después de borrar")
	require.Equal(t, 12, lista.VerUltimo(), "12 debería ser el nuevo último")

	require.Equal(t, 0, lista.VerPrimero(), "El primer elemento debe seguir siendo 0")

	iterFinal.Insertar(10)
	iterFinal.Siguiente()
}

func TestBorraHastaVaciar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	valores := []string{"Analisis", "Algoritmos", "Programacion", "Algebra", "TDA"}
	for _, valor := range valores {
		lista.InsertarUltimo(valor)
	}
	require.Equal(t, len(valores), lista.Largo(), "El largo inicial es incorrecto")

	for i := 0; i < len(valores); i++ {
		require.Equal(t, valores[i], lista.VerPrimero(), "El primer valor debería ser el esperado")
		require.Equal(t, valores[len(valores)-1], lista.VerUltimo(), "El último valor debería ser el esperado")

		lista.BorrarPrimero()

		require.Equal(t, len(valores)-(i+1), lista.Largo(), "El largo de la lista después de borrar debería ser el esperado")
		iter := lista.Iterador()
		if i+1 < len(valores) {
			require.Equal(t, valores[i+1], iter.VerActual(), "El iterador debe apuntar al siguiente valor correcto")
		}
	}
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.Equal(t, 0, lista.Largo(), "El largo de la lista debe ser 0 después de borrar todos los elementos")
}

func TestInsertaAvanza(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(100)
	require.Equal(t, 1, lista.Largo(), "El largo debe ser 1 después de insertar 100")
	require.Equal(t, 100, lista.VerPrimero(), "El primer elemento debe ser 100")
	require.Equal(t, 100, lista.VerUltimo(), "El último elemento debe ser 100")
	iterador.Siguiente()
	iterador.Insertar(300)
	require.Equal(t, 2, lista.Largo(), "El largo debe ser 2 después de insertar 300")
	require.Equal(t, 100, lista.VerPrimero(), "El primer elemento debe seguir siendo 100")
	require.Equal(t, 300, lista.VerUltimo(), "El último elemento debe ser 300")
	iterador2 := lista.Iterador()
	iterador2.Siguiente()
	iterador2.Insertar(200)
	require.Equal(t, 3, lista.Largo(), "El largo debe ser 3 después de insertar 200")
	require.Equal(t, 100, lista.VerPrimero(), "El primer elemento debe seguir siendo 100")
	require.Equal(t, 300, lista.VerUltimo(), "El último elemento debe seguir siendo 300")
	iter := lista.Iterador()
	require.Equal(t, 100, iter.VerActual(), "El primer valor del iterador debe ser 100")
	iter.Siguiente()
	require.Equal(t, 200, iter.VerActual(), "El segundo valor del iterador debe ser 200")
	iter.Siguiente()
	require.Equal(t, 300, iter.VerActual(), "El tercer valor del iterador debe ser 300")
}

func TestVerActualListaCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.True(t, lista.EstaVacia(), "La lista esta vacia")
	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia ser 0")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.False(t, iterador.HaySiguiente(), "El iterador no deberia tener un siguiente elemento en una lista vacia")
	lista.InsertarUltimo(10)
	nuevoIter := lista.Iterador()
	require.Equal(t, 10, nuevoIter.VerActual())
	require.True(t, nuevoIter.HaySiguiente())
	nuevoIter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { nuevoIter.VerActual() })
	require.Equal(t, 1, lista.Largo())
}

func TestVolumenPrimitivasInsertarYBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista debe comenzar vacia")
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, VOLUMEN_1, lista.Largo())
	require.Equal(t, 0, lista.VerPrimero(), "El primer elemento debe ser 0")
	require.Equal(t, VOLUMEN_1-1, lista.VerUltimo(), "El ultimo elemento debe ser %d", VOLUMEN_1-1)
	for i := 0; i < VOLUMEN_1; i++ {
		require.Equal(t, i, lista.BorrarPrimero(), "Se deberia borrar el elemento %d", i)
	}
	require.True(t, lista.EstaVacia(), "La lista deberia quedar vacia")
}

func TestVolumenPrimitivasInsertarPrimeroYVer(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_2; i++ {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, VOLUMEN_2, lista.Largo())
	require.Equal(t, VOLUMEN_2-1, lista.VerPrimero(), "El primer elemento deberia ser %d", VOLUMEN_2-1)
	require.Equal(t, 0, lista.VerUltimo(), "El ultimo elemento deberia ser 0")
}

func TestVolumenPrimitivasInsertarAlternado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		if i%2 == 0 {
			lista.InsertarPrimero(i)
		} else {
			lista.InsertarUltimo(i)
		}
	}
	require.Equal(t, VOLUMEN_1, lista.Largo(), "La lista deberia tener %d elementos", VOLUMEN_1)
}

func TestVolumenIteradorInternoSuma(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_2; i++ {
		lista.InsertarUltimo(i)
	}
	suma := 0
	lista.Iterar(func(valor int) bool {
		suma += valor
		return true
	})
	require.Equal(t, (VOLUMEN_2*(VOLUMEN_2-1))/2, suma, "La suma total de los valores debe ser la esperada")
}

func TestVolumenIteradorInternoCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	lista.Iterar(func(valor int) bool {
		contador++
		return valor < VOLUMEN_1/2
	})
	require.True(t, contador > 0, "El iterador deberia haber recorrido al menos un elemento")
}

func TestVolumenIteradorInternoTodosRecorridos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_2; i++ {
		lista.InsertarUltimo(1)
	}
	suma := 0
	lista.Iterar(func(valor int) bool {
		suma += valor
		return true
	})
	require.Equal(t, VOLUMEN_2, suma, "La suma deberia ser %d porque todos los elementos son 1", VOLUMEN_2)
}

func TestVolumenIteradorExternoRecorreTodo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		require.Equal(t, contador, iter.VerActual(), "El valor actual deberia ser %d", contador)
		iter.Siguiente()
		contador++
	}
	require.Equal(t, VOLUMEN_1, contador, "Se deberian haber recorrido %d elementos", VOLUMEN_1)
}

func TestVolumenIteradorExternoInsertaDuranteIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	iter.Insertar(999999)
	require.Equal(t, 999999, iter.VerActual(), "El elemento insertado debe ser 999999")
	require.Equal(t, VOLUMEN_1+1, lista.Largo(), "El largo debe haber aumentado en 1")
}

func TestVolumenIteradorExternoBorraTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Borrar()
	}
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todos")
	require.Equal(t, 0, lista.Largo(), "El largo deberia ser 0 despues de borrar todos")
}

func TestVolumenInsertarRecorrerYBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < VOLUMEN_1; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	total := 0
	for iter.HaySiguiente() {
		total += iter.VerActual()
		iter.Siguiente()
	}
	require.Equal(t, (VOLUMEN_1*(VOLUMEN_1-1))/2, total, "La suma debe ser la esperada")

	iter = lista.Iterador()
	for iter.HaySiguiente() {
		iter.Borrar()
	}
	require.True(t, lista.EstaVacia(), "La lista debe estar vacia luego de borrar todos los elementos")
}
