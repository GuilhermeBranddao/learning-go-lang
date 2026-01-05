package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python:
    nome = input("Digite seu nome: ")
    idade = int(input("Digite sua idade: "))

Go é mais verboso, mas também mais explícito e seguro.
*/

func main() {
	fmt.Println("=== ENTRADA E SAÍDA BÁSICA ===\n")

	// Saída simples
	fmt.Println("Olá, mundo!")           // com quebra de linha
	fmt.Print("Sem quebra de linha... ") // sem quebra de linha
	fmt.Println("Continuando!")

	// Formatação
	nome := "Maria"
	idade := 25
	fmt.Printf("Nome: %s, Idade: %d anos\n\n", nome, idade)

	fmt.Println("=== ENTRADA COM fmt.Scan ===\n")

	// Entrada simples (uma palavra)
	var apelido string
	fmt.Print("Digite seu apelido: ")
	fmt.Scan(&apelido) // Precisa do & (endereço de memória)
	fmt.Printf("Seu apelido é: %s\n\n", apelido)

	// Entrada de número
	var anos int
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&anos)
	fmt.Printf("Você tem %d anos\n\n", anos)

	// PROBLEMA: fmt.Scan não lê espaços!
	// Para ler linha completa, use bufio.Scanner

	fmt.Println("=== ENTRADA DE LINHA COMPLETA ===\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite seu nome completo: ")
	scanner.Scan()
	nomeCompleto := scanner.Text()
	fmt.Printf("Nome completo: %s\n\n", nomeCompleto)

	fmt.Println("=== CONVERSÃO DE TIPOS ===\n")

	fmt.Print("Digite um número: ")
	scanner.Scan()
	texto := scanner.Text()

	// Converter string para int
	numero, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Erro: não é um número válido!")
	} else {
		dobro := numero * 2
		fmt.Printf("O dobro de %d é %d\n\n", numero, dobro)
	}

	fmt.Println("=== FORMATAÇÃO AVANÇADA ===\n")

	produto := "Notebook"
	preco := 3499.99
	estoque := 15

	// Alinhamento e padding
	fmt.Printf("%-15s | R$ %8.2f | Estoque: %3d\n", produto, preco, estoque)
	fmt.Printf("%-15s | R$ %8.2f | Estoque: %3d\n", "Mouse", 49.90, 150)

	fmt.Println("\n=== PLACEHOLDERS ÚTEIS ===\n")

	valor := 42
	fmt.Printf("%%d (decimal):    %d\n", valor)
	fmt.Printf("%%b (binário):    %b\n", valor)
	fmt.Printf("%%o (octal):      %o\n", valor)
	fmt.Printf("%%x (hex):        %x\n", valor)
	fmt.Printf("%%v (padrão):     %v\n", valor)
	fmt.Printf("%%T (tipo):       %T\n\n", valor)

	texto2 := "Go"
	fmt.Printf("%%s (string):     %s\n", texto2)
	fmt.Printf("%%q (quoted):     %q\n", texto2)
	fmt.Printf("%%v (padrão):     %v\n", texto2)
	fmt.Printf("%%T (tipo):       %T\n\n", texto2)

	fmt.Println("=== EXEMPLO PRÁTICO ===\n")
	exemploCalculadora()
}

func exemploCalculadora() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Calculadora Simples")
	fmt.Println("------------------")

	// Primeiro número
	fmt.Print("Digite o primeiro número: ")
	scanner.Scan()
	num1Str := scanner.Text()
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Erro: número inválido!")
		return
	}

	// Operação
	fmt.Print("Digite a operação (+, -, *, /): ")
	scanner.Scan()
	operacao := strings.TrimSpace(scanner.Text())

	// Segundo número
	fmt.Print("Digite o segundo número: ")
	scanner.Scan()
	num2Str := scanner.Text()
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Erro: número inválido!")
		return
	}

	// Calcular
	var resultado float64
	switch operacao {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero!")
			return
		}
		resultado = num1 / num2
	default:
		fmt.Println("Operação inválida!")
		return
	}

	fmt.Printf("\nResultado: %.2f %s %.2f = %.2f\n", num1, operacao, num2, resultado)
}

/*
RESUMO DOS PACOTES:

fmt - Formatação e I/O básico
    fmt.Println() - imprime com quebra de linha
    fmt.Print() - imprime sem quebra de linha
    fmt.Printf() - imprime com formatação
    fmt.Scan() - lê entrada (apenas uma palavra)

bufio - I/O bufferizado (para ler linhas completas)
    bufio.NewScanner() - cria scanner
    scanner.Scan() - lê próxima linha
    scanner.Text() - obtém texto lido

strconv - Conversão de tipos
    strconv.Atoi() - string para int
    strconv.Itoa() - int para string
    strconv.ParseFloat() - string para float64

Execute com:
    go run 05_entrada_saida.go
*/
