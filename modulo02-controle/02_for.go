package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python tem: for, while
Go tem apenas: FOR (mas funciona de várias formas!)

Python:
    for i in range(5):
        print(i)

    while condicao:
        codigo

Go:
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    for condicao {
        codigo
    }
*/

func main() {
	fmt.Println("=== FOR TRADICIONAL (estilo C) ===\n")

	// For completo: inicialização; condição; pós-iteração
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteração %d\n", i)
	}

	fmt.Println("\n=== FOR COMO WHILE ===\n")

	// For com apenas condição (como while em Python)
	contador := 0
	for contador < 3 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}

	fmt.Println("\n=== FOR INFINITO ===\n")

	// Loop infinito (precisa de break)
	i := 0
	for {
		fmt.Printf("Loop infinito: %d\n", i)
		i++
		if i >= 3 {
			break // sai do loop
		}
	}

	fmt.Println("\n=== BREAK E CONTINUE ===\n")

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("  (pulando 5)")
			continue // pula para próxima iteração
		}
		if i == 8 {
			fmt.Println("  (parando em 8)")
			break // sai do loop
		}
		fmt.Printf("Número: %d\n", i)
	}

	fmt.Println("\n=== FOR COM RANGE (Arrays/Slices) ===\n")

	// Range em slice (similar ao enumerate do Python)
	numeros := []int{10, 20, 30, 40, 50}

	// Com índice e valor
	for indice, valor := range numeros {
		fmt.Printf("numeros[%d] = %d\n", indice, valor)
	}

	// Apenas valores (ignorar índice com _)
	fmt.Println("\nApenas valores:")
	for _, valor := range numeros {
		fmt.Printf("Valor: %d\n", valor)
	}

	// Apenas índices
	fmt.Println("\nApenas índices:")
	for indice := range numeros {
		fmt.Printf("Índice: %d\n", indice)
	}

	fmt.Println("\n=== RANGE COM STRINGS ===\n")

	// Range em string itera sobre RUNES (caracteres Unicode)
	texto := "Olá, 世界"
	for indice, char := range texto {
		fmt.Printf("Posição %d: %c (código: %d)\n", indice, char, char)
	}

	fmt.Println("\n=== RANGE COM MAPS ===\n")

	// Maps (dicionários)
	capitais := map[string]string{
		"Brasil":    "Brasília",
		"Argentina": "Buenos Aires",
		"Chile":     "Santiago",
	}

	for pais, capital := range capitais {
		fmt.Printf("%s -> %s\n", pais, capital)
	}

	fmt.Println("\n=== LOOPS ANINHADOS ===\n")

	// Tabuada
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}

	fmt.Println("=== EXEMPLO PRÁTICO: SOMA DE NÚMEROS ===\n")

	// Soma dos números de 1 a 100
	soma := 0
	for i := 1; i <= 100; i++ {
		soma += i
	}
	fmt.Printf("Soma de 1 a 100: %d\n", soma)

	fmt.Println("\n=== EXEMPLO PRÁTICO: FATORIAL ===\n")

	// Fatorial de 5
	n := 5
	fatorial := 1
	for i := 1; i <= n; i++ {
		fatorial *= i
	}
	fmt.Printf("%d! = %d\n", n, fatorial)

	fmt.Println("\n=== EXEMPLO PRÁTICO: FIBONACCI ===\n")

	// Primeiros 10 números de Fibonacci
	a, b := 0, 1
	fmt.Print("Fibonacci: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b // atribuição múltipla!
	}
	fmt.Println()
}

/*
RESUMO: FOR EM GO

1. For completo:
   for init; condicao; pos { }

2. For como while:
   for condicao { }

3. For infinito:
   for { }

4. For com range:
   for indice, valor := range colecao { }

IMPORTANTE:
- Go tem apenas FOR, mas ele substitui while e do-while
- Use _ para ignorar valores no range
- Range em strings itera sobre RUNES, não bytes
- break sai do loop, continue pula para próxima iteração

Execute com:
    go run 02_for.go
*/
