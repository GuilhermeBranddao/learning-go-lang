package matematica

import "errors"

/*
PACKAGE MATEMATICA

Este é um package de biblioteca (não main).
Exporta funções matemáticas básicas.

VISIBILIDADE:
- Somar: EXPORTADA (maiúscula)
- somar: não exportada (minúscula) - não existe aqui
*/

// Constantes exportadas
const (
	Versao = "1.0.0"
	Autor  = "Curso Go"
)

// Somar retorna a soma de dois inteiros
func Somar(a, b int) int {
	return a + b
}

// Subtrair retorna a subtração de dois inteiros
func Subtrair(a, b int) int {
	return a - b
}

// Multiplicar retorna a multiplicação de dois inteiros
func Multiplicar(a, b int) int {
	return a * b
}

// Dividir retorna a divisão de dois números
// Retorna erro se divisor for zero
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// função não exportada (privada ao package)
func auxiliar() int {
	return 42
}
