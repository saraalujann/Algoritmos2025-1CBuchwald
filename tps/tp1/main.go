package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	infix "tp1/calculadora"
)

func procExpresion() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		expresion := s.Text()
		tokens, err := infix.Tokenizar(expresion)
		if err != nil {
			continue
		}
		posfijo, _ := infix.InfixAPosfix(tokens)
		fmt.Println(strings.Join(posfijo, " "))
	}
}

func main() {
	procExpresion()
}
