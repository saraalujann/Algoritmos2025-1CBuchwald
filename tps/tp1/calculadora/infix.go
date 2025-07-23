package infix

import (
	"errors"
	"tdas/cola"
	"tdas/pila"
)

// procesarToken procesa un token de la expresion infija y lo maneja segun su tipo.
//
// PreCond. Token es un string valido, salida y operadores son estructuras ya inicializadas.
// PostCond. Modifica "salida" y/o "operadores" segun el token. Devuelve error si los parentesis estan mal balanceados.

func procesarToken(token string, salida cola.Cola[string], operadores pila.Pila[string]) error {
	switch {
	case EsNumero(token):
		salida.Encolar(token)
		return nil
	case token == PAR_ABRE:
		operadores.Apilar(token)
		return nil
	case token == PAR_CIERRA:
		for !operadores.EstaVacia() && operadores.VerTope() != PAR_ABRE {
			salida.Encolar(operadores.Desapilar())
		}
		if operadores.EstaVacia() {
			return errors.New("paréntesis no balanceados")
		} else {
			operadores.Desapilar()
			return nil
		}
	case EsOperador(token):
		return ManejarOperador(token, salida, operadores)
	}
	return nil
}

// InfixAPosFix es una funcion que se exporta, convierte una expresion tokenizada en
// notacion infija a su equivalencia en posfija.
//
// PreCond. "Tokens" es una lista de strings ya verificados.
// PostCond. Devuelve el arreglo de strings en orden posfijo.

func InfixAPosfix(tokens []string) ([]string, error) {
	salida := cola.CrearColaEnlazada[string]()
	operadores := pila.CrearPilaDinamica[string]()

	for _, token := range tokens {
		err := procesarToken(token, salida, operadores)
		if err != nil {
			return nil, err
		}
	}
	err := VaciarPila(salida, operadores)
	if err != nil {
		return nil, err
	}
	var resultado []string
	for !salida.EstaVacia() {
		resultado = append(resultado, salida.Desencolar())
	}
	return resultado, nil
}
