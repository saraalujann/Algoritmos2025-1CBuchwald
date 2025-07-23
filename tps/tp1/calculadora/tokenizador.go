package infix

import (
	"errors"
	"strings"
)

// separarOperadoresJuntos inserta espacios entre los caracteres para poder separar
// bien a los operadores cuando estan pegados a los numeros o a los parentesis
//
// PreCond. "expresion" es una cadena no vacia que representa una cadena matematica con not. infija
// PostCond. Retorna la misma cadena con espacios para facilitar la tokenizacion

func separarOperadoresJuntos(expresion string) string {
	for i := 0; i < len(expresion); i++ {
		letra := expresion[i]
		if !EsOperador(string(letra)) || i == 0 {
			continue
		}
		if i > 0 {
			prev := expresion[i-1]
			if (prev >= '0' && prev <= '9') || prev == PAR_CIERRA[0] {
				expresion = expresion[:i] + " " + expresion[i:]
				i++
			}
		}
		if i+1 < len(expresion) {
			sig := expresion[i+1]
			if (sig >= '0' && sig <= '9') || sig == PAR_ABRE[0] {
				expresion = expresion[:i+1] + " " + expresion[i+1:]
				i++
			}
		}
	}
	return expresion
}

// Tokenizar es una funcion exportada que recibe una notacion infija y la convierte en un
// arreglo de tokens insertando espacios donde crea necesario.
//
// PreCond."expresion" es una cadena que representa una expresion matematica valida.
// PostCond. Devuelve una lista de tokens en orden o un error si hay un token invalido.

func Tokenizar(expresion string) ([]string, error) {
	for _, op := range operadores {
		if len(op.Simbolo) > 1 {
			expresion = strings.ReplaceAll(expresion, op.Simbolo, " "+op.Simbolo+" ")
		}
	}
	expresion = separarOperadoresJuntos(expresion)
	expresion = strings.ReplaceAll(expresion, PAR_ABRE, " "+PAR_ABRE+" ")
	expresion = strings.ReplaceAll(expresion, PAR_CIERRA, " "+PAR_CIERRA+" ")

	tokens := strings.Fields(expresion)

	for _, token := range tokens {
		if !EsNumero(token) && !EsOperador(token) && token != PAR_ABRE && token != PAR_CIERRA {
			return nil, errors.New("token invalido: " + token)
		}
	}
	return tokens, nil
}
