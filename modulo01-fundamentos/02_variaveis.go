package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Em Python, você simplesmente faz:
    nome = "Maria"
    idade = 25

Em Go, você tem várias formas de declarar variáveis.
*/

func main() {
	fmt.Println("=== DECLARAÇÃO DE VARIÁVEIS ===\n")

	// Forma 1: var com tipo explícito
	var nome string = "Maria"
	var idade int = 25
	fmt.Printf("Forma 1: %s tem %d anos\n\n", nome, idade)

	// Forma 2: var com inferência de tipo
	var cidade = "São Paulo"
	var temperatura = 28.5
	fmt.Printf("Forma 2: Em %s está %.1f°C\n\n", cidade, temperatura)

	// Forma 3: Short declaration := (mais comum)
	pais := "Brasil"
	populacao := 214_000_000 // Você pode usar _ para facilitar leitura
	fmt.Printf("Forma 3: %s tem %d habitantes\n\n", pais, populacao)

	// Declaração múltipla
	var (
		empresa  = "TechCorp"
		fundacao = 2020
		ativa    = true
	)
	fmt.Printf("Múltiplas: %s fundada em %d (ativa: %t)\n\n", empresa, fundacao, ativa)

	// Short declaration múltipla
	x, y, z := 10, 20, 30
	fmt.Printf("x=%d, y=%d, z=%d\n\n", x, y, z)

	// IMPORTANTE: Variáveis não usadas causam ERRO!
	// Descomente a linha abaixo para ver o erro:
	// naoUsada := "isso vai dar erro"

	// Zero values (valores padrão)
	var numeroInteiro int
	var textoVazio string
	var booleano bool
	fmt.Println("=== ZERO VALUES ===")
	fmt.Printf("int não inicializado: %d\n", numeroInteiro)
	fmt.Printf("string não inicializada: '%s'\n", textoVazio)
	fmt.Printf("bool não inicializado: %t\n", booleano)
}

/*
DICAS IMPORTANTES:
1. Use := dentro de funções (mais rápido)
2. Use var fora de funções ou quando precisa especificar o tipo
3. Go NÃO permite variáveis não usadas (economia de recursos)
4. Todas as variáveis têm valores padrão (zero values)

Execute com:
    go run 02_variaveis.go
*/
