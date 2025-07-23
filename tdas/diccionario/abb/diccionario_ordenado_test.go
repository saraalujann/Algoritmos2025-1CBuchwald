package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	VOLUMEN_1 = 10000
	VOLUMEN_2 = 5000
)

func cmpInt(a, b int) int {
	return a - b
}

func cmpStr(a, b string) int {
	return strings.Compare(a, b)
}

func TestClavesInt(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(2, "C")
	abb.Guardar(3, "B")
	require.Equal(t, "A", abb.Obtener(1))
	require.Equal(t, "B", abb.Obtener(3))
	require.Equal(t, "C", abb.Obtener(2))
}

func TestClavesStr(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](cmpStr)
	abb.Guardar("a", "A")
	abb.Guardar("b", "C")
	abb.Guardar("c", "B")
	require.Equal(t, "A", abb.Obtener("a"))
	require.Equal(t, "B", abb.Obtener("c"))
	require.Equal(t, "C", abb.Obtener("b"))
}

func TestReemplazoDatoABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(2, "B")
	require.Equal(t, "A", abb.Obtener(1))
	require.Equal(t, "B", abb.Obtener(2))
	require.Equal(t, "C", abb.Obtener(3))
	abb.Guardar(1, "D")
	require.Equal(t, "D", abb.Obtener(1))
	require.NotEqual(t, "A", abb.Obtener(1))
}

func TestObtenerClaveExistenteABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(2, "B")
	require.Equal(t, "B", abb.Obtener(2))
	require.True(t, abb.Pertenece(2))
}

func TestABBVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	require.Panics(t, func() { abb.Obtener(2) })
	require.False(t, abb.Pertenece(2))
}

func TestObtenerVacioABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "")
	require.Equal(t, "", abb.Obtener(1))
	require.True(t, abb.Pertenece(1))
}

func TestBorrarABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(2, "B")
	abb.Borrar(3)
	require.True(t, abb.Pertenece(1))
	require.True(t, abb.Pertenece(2))
	require.False(t, abb.Pertenece(3))
	abb.Borrar(1)
	abb.Borrar(2)
	require.False(t, abb.Pertenece(1))
	require.False(t, abb.Pertenece(2))
}

func TestBorrarClaveNoExistente(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(2, "B")
	abb.Guardar(3, "C")
	require.Panics(t, func() { abb.Borrar(5) })
}

func TestBorrarTodasLasClaves(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(2, "B")
	abb.Guardar(3, "C")
	abb.Borrar(3)
	require.False(t, abb.Pertenece(3))
	abb.Borrar(2)
	require.False(t, abb.Pertenece(2))
	abb.Borrar(1)
	require.False(t, abb.Pertenece(1))
}

func TestBorrarSinHijosABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "A")
	abb.Guardar(2, "B")
	abb.Guardar(3, "C")
	abb.Borrar(2)
	require.False(t, abb.Pertenece(2))
	require.True(t, abb.Pertenece(1))
	require.True(t, abb.Pertenece(3))
}

func TestBorrarHijoIzqABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(2, "B")
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(4, "D")
	abb.Borrar(1)
	require.False(t, abb.Pertenece(1))
	require.True(t, abb.Pertenece(2))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(4))
}

func TestBorrarHijoDerABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(2, "B")
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(4, "D")
	abb.Borrar(4)
	require.False(t, abb.Pertenece(4))
	require.True(t, abb.Pertenece(2))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(1))
}

func TestBorrarConDosHijosABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(5, "B")
	abb.Guardar(3, "A")
	abb.Guardar(7, "C")
	abb.Guardar(9, "D")
	abb.Guardar(1, "E")
	abb.Borrar(1)
	require.False(t, abb.Pertenece(1))
	require.True(t, abb.Pertenece(5))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(7))
	require.True(t, abb.Pertenece(9))
}

func TestBorrarRaiz(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(2, "B")
	abb.Guardar(1, "A")
	abb.Guardar(3, "C")
	abb.Guardar(4, "D")
	abb.Borrar(2)
	require.False(t, abb.Pertenece(2))
	require.True(t, abb.Pertenece(1))
	require.True(t, abb.Pertenece(3))
	require.True(t, abb.Pertenece(4))
}

func TestGuardarYBorrarIntercaladoABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(50, "A")
	abb.Guardar(8, "B")
	require.True(t, abb.Pertenece(50))
	require.True(t, abb.Pertenece(8))
	abb.Borrar(8)
	require.False(t, abb.Pertenece(8))
	abb.Guardar(20, "C")
	abb.Guardar(68, "D")
	require.True(t, abb.Pertenece(20))
	require.True(t, abb.Pertenece(68))
	abb.Borrar(50)
	require.False(t, abb.Pertenece(50))
	abb.Guardar(16, "E")
	abb.Guardar(75, "F")
	require.True(t, abb.Pertenece(16))
	require.True(t, abb.Pertenece(75))
	abb.Borrar(68)
	require.False(t, abb.Pertenece(68))
	require.False(t, abb.Pertenece(50))
	require.False(t, abb.Pertenece(8))
	require.True(t, abb.Pertenece(20))
	require.True(t, abb.Pertenece(16))
	require.True(t, abb.Pertenece(75))
}

func TestInOrderIterarCompleto(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(10, "L")
	abb.Guardar(4, "D")
	abb.Guardar(1, "F")
	abb.Guardar(5, "A")
	abb.Guardar(2, "S")
	abb.Guardar(15, "P")
	abb.Guardar(13, "W")
	abb.Guardar(12, "R")
	var claves []int
	var valores []string
	abb.Iterar(func(clave int, valor string) bool {
		claves = append(claves, clave)
		valores = append(valores, valor)
		return true
	})
	require.Equal(t, []int{1, 2, 4, 5, 10, 12, 13, 15}, claves)
	require.Equal(t, []string{"F", "S", "D", "A", "L", "R", "W", "P"}, valores)
}

func TestInOrderIterarRangoInterna(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](cmpStr)
	abb.Guardar("J", 10)
	abb.Guardar("N", 3)
	abb.Guardar("Z", 98)
	abb.Guardar("A", 27)
	abb.Guardar("I", 5)
	abb.Guardar("S", 100)
	abb.Guardar("L", 15)
	abb.Guardar("O", 55)
	abb.Guardar("C", 12)
	desde := "J"
	hasta := "S"
	var claves []string
	var valores []int
	abb.IterarRango(&desde, &hasta, func(clave string, valor int) bool {
		claves = append(claves, clave)
		valores = append(valores, valor)
		return true
	})
	require.Equal(t, []string{"J", "L", "N", "O", "S"}, claves)
	require.Equal(t, []int{10, 15, 3, 55, 100}, valores)
}

func TestInOrderIteradorRangoExterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(92, "G")
	abb.Guardar(8, "P")
	abb.Guardar(124, "Z")
	abb.Guardar(16, "A")
	abb.Guardar(27, "L")
	abb.Guardar(15, "J")
	abb.Guardar(100, "B")
	abb.Guardar(60, "E")
	desde := 20
	hasta := 70
	var claves []int
	var valores []string
	iter := abb.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		claves = append(claves, clave)
		valores = append(valores, valor)
		iter.Siguiente()
	}
	require.Equal(t, []int{27, 60}, claves)
	require.Equal(t, []string{"L", "E"}, valores)
}

func TestVolumenABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	for i := 0; i < VOLUMEN_1; i++ {
		abb.Guardar(i, i*i)
	}
	for i := 0; i < VOLUMEN_1; i++ {
		require.True(t, abb.Pertenece(i))
		valor := abb.Obtener(i)
		require.Equal(t, i*i, valor)
	}
	require.Equal(t, VOLUMEN_1, abb.Cantidad())
	for i := 0; i < VOLUMEN_1; i++ {
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(i) })
	}
	require.Equal(t, 0, abb.Cantidad())
}

func TestVolumenIntABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	for i := 0; i < VOLUMEN_2; i++ {
		abb.Guardar(i, fmt.Sprintf("valor-%d", i))
	}
	for i := 0; i < VOLUMEN_2; i++ {
		require.True(t, abb.Pertenece(i))
		valor := abb.Obtener(i)
		require.Equal(t, fmt.Sprintf("valor-%d", i), valor)
	}
	for i := 0; i < VOLUMEN_2; i += 2 {
		abb.Borrar(i)
	}
	for i := 0; i < VOLUMEN_2; i++ {
		if i%2 == 0 {
			require.False(t, abb.Pertenece(i))
			require.PanicsWithValue(t, "La clave no pertenece al diccionario",
				func() { abb.Obtener(i) })
		} else {
			require.True(t, abb.Pertenece(i))
			valor := abb.Obtener(i)
			require.Equal(t, fmt.Sprintf("valor-%d", i), valor)
		}
	}
}

func TestVolumenStrABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](cmpStr)
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)
		abb.Guardar(clave, i*10)
	}
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)
		require.True(t, abb.Pertenece(clave))
		valor := abb.Obtener(clave)
		require.Equal(t, i*10, valor)
	}
	for i := 0; i < VOLUMEN_1; i += 2 {
		clave := fmt.Sprintf("clave-%08d", i)
		abb.Borrar(clave)
	}
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)
		if i%2 == 0 {
			require.False(t, abb.Pertenece(clave))
			require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(clave) })
		} else {
			require.True(t, abb.Pertenece(clave))
			valor := abb.Obtener(clave)
			require.Equal(t, i*10, valor)
		}
	}
}

func TestVolumenBorradosIntercalados(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](cmpStr)
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)
		abb.Guardar(clave, i)
	}
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)
		if i%3 == 0 {
			require.True(t, abb.Pertenece(clave))
			valor := abb.Obtener(clave)
			require.Equal(t, i, valor)

			abb.Borrar(clave)
			require.False(t, abb.Pertenece(clave))
			require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(clave) })
		} else {
			require.True(t, abb.Pertenece(clave))
		}
	}
	for i := 0; i < VOLUMEN_1; i++ {
		clave := fmt.Sprintf("clave-%08d", i)

		if i%3 == 0 {
			require.False(t, abb.Pertenece(clave))
		} else {
			require.True(t, abb.Pertenece(clave))
			require.Equal(t, i, abb.Obtener(clave))
		}
	}
}

func TestVolumenIteradorRangoCompletoABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	valores := make([]string, 0, VOLUMEN_1)
	for i := 0; i < VOLUMEN_1; i++ {
		valor := fmt.Sprintf("valor-%08d", i)
		valores = append(valores, valor)
		abb.Guardar(i, valor)
	}
	iter := abb.IteradorRango(nil, nil)
	i := 0
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.Equal(t, i, clave)
		require.Equal(t, valores[i], valor)
		iter.Siguiente()
		i++
	}
	require.Equal(t, VOLUMEN_1, i)
	require.False(t, iter.HaySiguiente())
}

func TestVolumenIterarRangoABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	valores := make([]string, VOLUMEN_1)
	for i := 0; i < VOLUMEN_1; i++ {
		valores[i] = fmt.Sprintf("valor-%08d", i)
		abb.Guardar(i, valores[i])
	}
	desde := VOLUMEN_1 / 4
	hasta := VOLUMEN_1 * 3 / 4
	iter := abb.IteradorRango(&desde, &hasta)
	for i := desde; iter.HaySiguiente(); i++ {
		clave, valor := iter.VerActual()
		require.Equal(t, i, clave)
		require.Equal(t, valores[i], valor)
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
}

func TestIterarCompletoCorteABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](cmpStr)
	materias := []string{"Algoritmos2", "Algebra", "Programacion", "OrgaComputador", "BDD", "SISOP"}
	for i, materia := range materias {
		abb.Guardar(materia, (i+1)*10)
	}
	var claves []string
	var valores []int
	abb.Iterar(func(clave string, valor int) bool {
		if clave == "Programacion" {
			return false
		}
		claves = append(claves, clave)
		valores = append(valores, valor)
		return true
	})
	require.Equal(t, []string{"Algebra", "Algoritmos2", "BDD", "OrgaComputador"}, claves)
	require.Equal(t, []int{20, 10, 50, 40}, valores)
}

func TestIterarRangoCorteABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	frutas := []string{"Manzana", "Banana", "Pera", "Sandia", "Palta", "Frutilla"}
	for i, fruta := range frutas {
		abb.Guardar((i+1)*5, fruta)
	}
	var claves []int
	var valores []string
	desde := 10
	hasta := 25
	abb.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		if clave == 20 {
			return false
		}
		claves = append(claves, clave)
		valores = append(valores, valor)
		return true
	})
	require.Equal(t, []int{10, 15}, claves)
	require.Equal(t, []string{"Banana", "Pera"}, valores)
}

func TestIteradorRangoCorteABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	frutas := []string{"Bordo", "Marron", "Plata", "Petroleo", "Salmon"}
	for i, color := range frutas {
		abb.Guardar((i+1)*10, color)
	}
	desde := 15
	hasta := 45
	iter := abb.IteradorRango(&desde, &hasta)
	require.True(t, iter.HaySiguiente())
	clave, valor := iter.VerActual()
	require.Equal(t, 20, clave)
	require.Equal(t, "Marron", valor)
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	clave, valor = iter.VerActual()
	require.Equal(t, "Plata", valor)
	require.Equal(t, 30, clave)
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	clave, valor = iter.VerActual()
	require.Equal(t, "Petroleo", valor)
	require.Equal(t, 40, clave)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Siguiente()
	})
}

func TestIterarRangoArbolVacioABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	desde := 10
	hasta := 30
	llamado := false
	abb.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		llamado = true
		return true
	})
	require.False(t, llamado, "No se debería haber llamado a visitar en un ABB vacío")

}

func TestIteradorRangoArbolVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	desde := 0
	hasta := 100
	iter := abb.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
