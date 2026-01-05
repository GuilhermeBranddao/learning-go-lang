package utils

import "strings"

/*
PACKAGE UTILS

Funções utilitárias diversas.
*/

// ToUpper converte string para maiúsculas
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converte string para minúsculas
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Reverter inverte uma string
func Reverter(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Maximo retorna o maior valor de um slice
func Maximo(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}
	max := numeros[0]
	for _, n := range numeros {
		if n > max {
			max = n
		}
	}
	return max
}

// Minimo retorna o menor valor de um slice
func Minimo(numeros []int) int {
	if len(numeros) == 0 {
		return 0
	}
	min := numeros[0]
	for _, n := range numeros {
		if n < min {
			min = n
		}
	}
	return min
}

// SomaArray retorna a soma de todos elementos
func SomaArray(numeros []int) int {
	soma := 0
	for _, n := range numeros {
		soma += n
	}
	return soma
}
