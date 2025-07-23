package aeropuerto

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
)

type Vuelo struct {
	flightNumber   string
	airline        string
	origin         string
	destination    string
	tailNumber     string
	priority       int
	date           time.Time
	departureDelay int
	airTime        int
	cancelled      int
}

type AlmacenVuelos struct {
	Fecha  TDADiccionario.DiccionarioOrdenado[time.Time, []*Vuelo]
	Codigo TDADiccionario.Diccionario[string, *Vuelo]
}

const (
	FORMATO_FECHA = "2006-01-02T15:04:05"
)

func NuevoVuelo(flightNumber string, airline string, origin string,
	destination string, tailNumber string, priority int, date time.Time, departureDelay int,
	airTime int, cancelled int,
) *Vuelo {
	return &Vuelo{
		flightNumber:   flightNumber,
		airline:        airline,
		origin:         origin,
		destination:    destination,
		tailNumber:     tailNumber,
		priority:       priority,
		date:           date,
		departureDelay: departureDelay,
		airTime:        airTime,
		cancelled:      cancelled,
	}
}

func NuevoAlmacenVuelos() *AlmacenVuelos {
	return &AlmacenVuelos{
		Fecha:  TDADiccionario.CrearABB[time.Time, []*Vuelo](CompararFechas),
		Codigo: TDADiccionario.CrearHash[string, *Vuelo](),
	}
}

func CompararFechas(a, b time.Time) int {
	if a.Before(b) {
		return -1
	}
	if a.After(b) {
		return 1
	}
	return 0
}

func (v *Vuelo) ImprimirTablero() string {
	return fmt.Sprintf("%s - %s",
		v.date.Format(FORMATO_FECHA),
		v.flightNumber,
	)
}

func (v *Vuelo) ImprimirDetallado() string {
	return fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d\n",
		v.flightNumber,
		v.airline,
		v.origin,
		v.destination,
		v.tailNumber,
		v.priority,
		v.date.Format(FORMATO_FECHA),
		v.departureDelay,
		v.airTime,
		v.cancelled,
	)
}

func (v *Vuelo) ClaveHash() string {
	return v.flightNumber
}

func (v *Vuelo) Fecha() time.Time {
	return v.date
}

func (v *Vuelo) FormatoPrioridad() string {
	return fmt.Sprintf("%d - %s\n", v.priority, v.flightNumber)
}

func (v *Vuelo) ConectaVuelo(origen, destino string) bool {
	return v.origin == origen && v.destination == destino
}

func (v *Vuelo) CompararAscendentePorFechaYNumero(otroVuelo *Vuelo) int {
	if v.date.Before(otroVuelo.date) {
		return -1
	}
	if v.date.After(otroVuelo.date) {
		return 1
	}
	res := strings.Compare(v.flightNumber, otroVuelo.flightNumber)
	if res > 0 {
		return -1
	} else if res < 0 {
		return 1
	}
	return 0
}

func (v *Vuelo) CompararPorPrioridad(otroVuelo *Vuelo) int {
	if v.priority != otroVuelo.priority {
		if v.priority > otroVuelo.priority {
			return 1
		}
		return -1
	}
	if v.flightNumber > otroVuelo.flightNumber {
		return -1
	} else if v.flightNumber < otroVuelo.flightNumber {
		return 1
	}
	return 0
}
