package main

import (
	"fmt"
	"time"
)

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python 3.10+:
    match valor:
        case 1:
            print("Um")
        case 2:
            print("Dois")

Go:
    switch valor {
    case 1:
        fmt.Println("Um")
    case 2:
        fmt.Println("Dois")
    }

VANTAGENS DO SWITCH EM GO:
- NÃO precisa de break
- Pode ter múltiplos valores no case
- Pode usar expressões
- Switch sem expressão (como if/else)
*/

func main() {
	fmt.Println("=== SWITCH BÁSICO ===\n")

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido")
	}

	fmt.Println("\n=== MÚLTIPLOS VALORES NO CASE ===\n")

	// Vários valores no mesmo case
	switch dia {
	case 1, 7:
		fmt.Println("Final de semana! 🎉")
	case 2, 3, 4, 5, 6:
		fmt.Println("Dia de semana 😴")
	default:
		fmt.Println("Dia inválido")
	}

	fmt.Println("\n=== SWITCH SEM EXPRESSÃO ===\n")

	// Switch sem expressão (como if/else)
	hora := 14

	switch {
	case hora < 12:
		fmt.Println("Bom dia! ☀️")
	case hora < 18:
		fmt.Println("Boa tarde! 🌤️")
	default:
		fmt.Println("Boa noite! 🌙")
	}

	fmt.Println("\n=== SWITCH COM INICIALIZAÇÃO ===\n")

	// Switch com inicialização de variável
	switch hoje := time.Now().Weekday(); hoje {
	case time.Saturday, time.Sunday:
		fmt.Printf("É %s - Final de semana!\n", hoje)
	default:
		fmt.Printf("É %s - Dia de semana\n", hoje)
	}

	fmt.Println("\n=== SWITCH COM TIPO (Type Switch) ===\n")

	// Type switch - verifica o tipo da variável
	var x interface{} = "texto"

	switch v := x.(type) {
	case int:
		fmt.Printf("É um inteiro: %d\n", v)
	case string:
		fmt.Printf("É uma string: %s\n", v)
	case bool:
		fmt.Printf("É um booleano: %t\n", v)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}

	fmt.Println("\n=== FALLTHROUGH ===\n")

	// fallthrough força a execução do próximo case
	// (comportamento padrão em C/Java, mas RARO em Go)
	numero := 1

	switch numero {
	case 1:
		fmt.Println("Um")
		fallthrough
	case 2:
		fmt.Println("Dois (executado por fallthrough)")
		fallthrough
	case 3:
		fmt.Println("Três (executado por fallthrough)")
	}

	fmt.Println("\n=== EXEMPLO PRÁTICO: NOTAS ===\n")

	nota := 85

	switch {
	case nota >= 90:
		fmt.Printf("Nota %d: A - Excelente! 🌟\n", nota)
	case nota >= 80:
		fmt.Printf("Nota %d: B - Ótimo! ✨\n", nota)
	case nota >= 70:
		fmt.Printf("Nota %d: C - Bom 👍\n", nota)
	case nota >= 60:
		fmt.Printf("Nota %d: D - Regular 😐\n", nota)
	default:
		fmt.Printf("Nota %d: F - Reprovado 😞\n", nota)
	}

	fmt.Println("\n=== EXEMPLO PRÁTICO: HTTP STATUS ===\n")

	status := 404

	switch status {
	case 200:
		fmt.Println("OK - Sucesso")
	case 201:
		fmt.Println("Created - Recurso criado")
	case 400:
		fmt.Println("Bad Request - Requisição inválida")
	case 401:
		fmt.Println("Unauthorized - Não autorizado")
	case 404:
		fmt.Println("Not Found - Não encontrado")
	case 500:
		fmt.Println("Internal Server Error - Erro no servidor")
	default:
		fmt.Printf("Status code: %d\n", status)
	}

	fmt.Println("\n=== EXEMPLO PRÁTICO: CALCULADORA ===\n")

	a := 10
	b := 5
	operacao := "+"

	switch operacao {
	case "+":
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
	case "-":
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
	case "*":
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
	case "/":
		if b != 0 {
			fmt.Printf("%d / %d = %d\n", a, b, a/b)
		} else {
			fmt.Println("Erro: divisão por zero!")
		}
	default:
		fmt.Println("Operação inválida!")
	}
}

/*
DIFERENÇAS IMPORTANTES:

Python (match/case - Python 3.10+):
    match valor:
        case 1:
            print("Um")
        case 2:
            print("Dois")

Go (switch):
    switch valor {
    case 1:
        fmt.Println("Um")
    case 2:
        fmt.Println("Dois")
    }

VANTAGENS DO GO:
1. NÃO precisa de break (não cai no próximo case)
2. Pode ter múltiplos valores: case 1, 2, 3:
3. Switch sem expressão funciona como if/else
4. Type switch para verificar tipos

Execute com:
    go run 03_switch.go
*/
