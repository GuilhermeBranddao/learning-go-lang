package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python tem tipagem dinâmica:
    numero = 42         # int
    numero = "texto"    # agora é string (OK!)

Go tem tipagem ESTÁTICA:
    numero := 42
    numero = "texto"    // ERRO! Não pode mudar o tipo
*/

func main() {
	fmt.Println("=== TIPOS NUMÉRICOS ===\n")

	// Inteiros
	var idade int = 25    // tamanho depende da arquitetura (32 ou 64 bits)
	var ano int64 = 2024  // sempre 64 bits
	var bytes uint8 = 255 // 0 a 255 (sem sinal)
	fmt.Printf("Idade: %d (tipo: int)\n", idade)
	fmt.Printf("Ano: %d (tipo: int64)\n", ano)
	fmt.Printf("Byte: %d (tipo: uint8)\n\n", bytes)

	// Ponto flutuante
	var altura float64 = 1.75
	var peso float32 = 70.5
	fmt.Printf("Altura: %.2f (tipo: float64)\n", altura)
	fmt.Printf("Peso: %.2f (tipo: float32)\n\n", peso)

	fmt.Println("=== STRINGS E CARACTERES ===\n")

	// Strings (aspas duplas)
	nome := "Gopher"
	frase := "Go é awesome!"
	multiLinha := `Esta é uma string
	que pode ter
	múltiplas linhas`

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Frase: %s\n", frase)
	fmt.Printf("Multi-linha:\n%s\n\n", multiLinha)

	// Rune (caractere - aspas simples)
	var letra rune = 'G'
	var emoji rune = '🚀'
	fmt.Printf("Letra: %c (código: %d)\n", letra, letra)
	fmt.Printf("Emoji: %c (código: %d)\n\n", emoji, emoji)

	fmt.Println("=== BOOLEANOS ===\n")

	// Boolean
	var aprovado bool = true
	var reprovado = false
	fmt.Printf("Aprovado: %t\n", aprovado)
	fmt.Printf("Reprovado: %t\n\n", reprovado)

	fmt.Println("=== CONSTANTES ===\n")

	// Constantes
	const PI = 3.14159
	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)

	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("HTTP OK: %d\n", StatusOK)
	fmt.Printf("HTTP Not Found: %d\n", StatusNotFound)
	fmt.Printf("HTTP Error: %d\n\n", StatusError)

	fmt.Println("=== CONVERSÃO DE TIPOS ===\n")

	// Em Go, conversão precisa ser EXPLÍCITA
	var inteiro int = 42
	var flutuante float64 = float64(inteiro) // conversão explícita
	var outro int = int(flutuante)

	fmt.Printf("Inteiro: %d\n", inteiro)
	fmt.Printf("Flutuante: %.2f\n", flutuante)
	fmt.Printf("De volta para int: %d\n", outro)

	// CUIDADO: conversão pode perder precisão
	var pi float64 = 3.14159
	var piInteiro int = int(pi) // perde a parte decimal
	fmt.Printf("\nPI original: %.5f\n", pi)
	fmt.Printf("PI como int: %d (perdeu a parte decimal!)\n", piInteiro)
}

/*
PLACEHOLDERS ÚTEIS para fmt.Printf():
%d - inteiro
%f - float
%.2f - float com 2 casas decimais
%s - string
%c - caractere (rune)
%t - booleano
%v - valor no formato padrão
%T - tipo da variável

Execute com:
    go run 03_tipos.go
*/
