package infix

import (
	"errors"
	"strconv"
	"tdas/cola"
	"tdas/pila"
)

type Operador struct {
	Simbolo       string
	Precedencia   int
	Asociatividad string
	Aridad        int
}

const (
	OP_SUMA           = "+"
	OP_RESTA          = "-"
	OP_MULTIPLICACION = "*"
	OP_DIVISION       = "/"
	OP_EXPONENCIAL    = "^"
	OP_LOGARITMICO    = "log"
	OP_RAIZ_CUADRADA  = "sqrt"
	OP_TERNARIO       = "?"
	PAR_ABRE          = "("
	PAR_CIERRA        = ")"
)

var operadores = []Operador{
	{OP_SUMA, 1, "izq", 2},
	{OP_RESTA, 1, "izq", 2},
	{OP_MULTIPLICACION, 2, "izq", 2},
	{OP_DIVISION, 2, "izq", 2},
	{OP_EXPONENCIAL, 4, "der", 2},
	{OP_LOGARITMICO, 4, "izq", 2},
	{OP_RAIZ_CUADRADA, 4, "izq", 1},
	{OP_TERNARIO, 5, "izq", 3},
}

// verificaDesapilar verifica si desapilar un operador de la pila segun la precedencia y
// asociatividad del operador actual.
//
// PreCond. Actual y Tope son operadores validos.
// PostCond. Devuelve true si tiene que desapilar "tope" antes de apilar "actual".

func verificaDesapilar(actual, tope Operador) bool {
	if tope.Simbolo == PAR_ABRE {
		return false
	}
	if actual.Asociatividad == "izq" {
		return actual.Precedencia <= tope.Precedencia
	} else {
		return actual.Precedencia < tope.Precedencia
	}
}

// verificaOperador busca un operador en la lista de operadores disponibles.
//
// PreCond. Token es un string valido.
// PostCond. Devuelve el operador que corresponde y true si lo encuentra, sino uno vacio y false.

func verificaOperador(token string) (Operador, bool) {
	for _, op := range operadores {
		if op.Simbolo == token {
			return op, true
		}
	}
	return Operador{}, false
}

// EsOperador es una funcion que se exporta, ve si el token es un operador valido.
//
// PreCond. Token es un string valido.
// PostCond. Devuelve true si es un operador definido, false caso contrario.

func EsOperador(token string) bool {
	_, es_operador := verificaOperador(token)
	return es_operador
}

// EsNumero es una funcion que se exporta, ve si el token es un numero entero valido.
//
// PreCond. Token es un string valido.
// PostCond. Devuelve true si es un numero entero valido.

func EsNumero(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

// ManejarOperador es una funcion que se exporta, sigue la logica de apilado/desapilado
// de operadores durante la conversion de notacion.
//
// PreCond. Token es un string valido, salida y operadores son estructuras ya inicializadas.
// PostCond. Apila/Desapila operadores segun su precedencia.

func ManejarOperador(token string, salida cola.Cola[string], operadores pila.Pila[string]) error {
	opActual, _ := verificaOperador(token)
	for !operadores.EstaVacia() && operadores.VerTope() != PAR_ABRE {
		opTope, _ := verificaOperador(operadores.VerTope())
		if verificaDesapilar(opActual, opTope) {
			salida.Encolar(operadores.Desapilar())
		} else {
			break
		}
	}
	operadores.Apilar(token)
	return nil
}

// VaciarPila es una funcion que se exporta, trae a todos los operadores restantes al final
// de la conversion, verifica que no queden parentesis abiertos.
//
// PreCond. Salida y operadores son estructuras ya inicializadas.
// PostCond. Encola los operadores restantes o "error" si hay un parentesis sin cerrar.

func VaciarPila(salida cola.Cola[string], operadores pila.Pila[string]) error {
	for !operadores.EstaVacia() {
		if operadores.VerTope() == PAR_ABRE {
			return errors.New("parentesis incorrectos")
		} else {
			salida.Encolar(operadores.Desapilar())
		}
	}
	return nil
}
