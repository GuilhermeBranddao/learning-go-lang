package matematica

import "errors"

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
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// Fatorial calcula o fatorial de n
func Fatorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("fatorial de número negativo")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado, nil
}

// EhPrimo verifica se um número é primo
func EhPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
