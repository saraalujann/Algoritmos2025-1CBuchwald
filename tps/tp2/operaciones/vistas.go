package operaciones

import (
	"fmt"
	TDALista "tdas/lista"
	"time"
	"tp2/aeropuerto"
	"tp2/auxs"
)

const (
	DESCENDENTE = "desc"
)

/*Imprime una lista de vuelos dentro del rango de fechas extraídos de un ABB, con
un límite en la cantidad de vuelos a mostrar y en el orden pedido (ascendente o descendente).*/

func VerTablero(almacen *aeropuerto.AlmacenVuelos, cantidad int, modo string, desde, hasta time.Time) {
	var lista TDALista.Lista[*aeropuerto.Vuelo]
	if modo == DESCENDENTE {
		lista = auxs.InvertirLista(almacen, &desde, &hasta)
	} else {
		lista = auxs.ListaVuelosAscendentes(almacen, &desde, &hasta)
	}
	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente() && contador < cantidad {
		vuelo := iter.VerActual()
		fmt.Println(vuelo.ImprimirTablero())
		iter.Siguiente()
		contador++
	}
}

/*Busca un vuelo por su código en un diccionario
hash y muestra la información detallada del vuelo.*/

func InfoVuelo(almacen *aeropuerto.AlmacenVuelos, codigo string) error {
	if !almacen.Codigo.Pertenece(codigo) {
		return fmt.Errorf("vuelo con código %s no encontrado\n", codigo)
	}
	vuelo := almacen.Codigo.Obtener(codigo)
	fmt.Print(vuelo.ImprimirDetallado())
	return nil
}

/*Imprime los vuelos con mayor prioridad, limitados por la cantidad indicada.*/

func PrioridadVuelos(almacen *aeropuerto.AlmacenVuelos, cantidad int) {
	topVuelos := auxs.ObtenerTopPrioritarios(almacen, cantidad)
	for _, vuelo := range topVuelos {
		fmt.Print(vuelo.FormatoPrioridad())
	}
}

/*Busca el siguiente vuelo desde
origen hacia destino con fecha igual o posterior a fecha.*/

func SiguienteVuelo(almacen *aeropuerto.AlmacenVuelos, origen, destino string, fecha time.Time) {
	iter := almacen.Fecha.IteradorRango(&fecha, nil)
	for iter.HaySiguiente() {
		_, vuelos := iter.VerActual()
		for _, vuelo := range vuelos {
			if vuelo.ConectaVuelo(origen, destino) {
				fmt.Print(vuelo.ImprimirDetallado())
				return
			}
		}
		iter.Siguiente()
	}
	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n",
		origen, destino, fecha.Format(FORMATO_FECHA))
}
