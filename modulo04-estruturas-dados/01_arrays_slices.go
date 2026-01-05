package main

import "fmt"

/*
ARRAYS vs SLICES

COMPARAÇÃO COM PYTHON:
--------------------
Python tem apenas listas (dinâmicas):
    numeros = [1, 2, 3, 4, 5]
    numeros.append(6)  # OK

Go tem ARRAYS (fixos) e SLICES (dinâmicos):
    array := [5]int{1, 2, 3, 4, 5}  // tamanho fixo
    slice := []int{1, 2, 3, 4, 5}   // tamanho dinâmico

IMPORTANTE: Slices são muito mais usados!
*/

func main() {
	fmt.Println("=== ARRAYS (Tamanho Fixo) ===\n")

	// Declaração de array com tamanho fixo
	var numeros [5]int
	fmt.Printf("Array vazio: %v\n", numeros) // [0 0 0 0 0]

	// Array com valores
	var cores [3]string = [3]string{"vermelho", "verde", "azul"}
	fmt.Printf("Cores: %v\n", cores)

	// Inferência de tamanho com ...
	dias := [...]string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"}
	fmt.Printf("Dias: %v (tamanho: %d)\n", dias, len(dias))

	// Acessar elementos
	fmt.Printf("Primeiro dia: %s\n", dias[0])
	fmt.Printf("Último dia: %s\n", dias[len(dias)-1])

	// Modificar elementos
	dias[0] = "Domingo"
	fmt.Printf("Dias modificado: %v\n\n", dias)

	fmt.Println("=== SLICES (Tamanho Dinâmico) ===\n")

	// Declaração de slice (SEM tamanho)
	var frutas []string
	fmt.Printf("Slice vazio: %v (len: %d)\n", frutas, len(frutas))

	// Slice com valores
	frutas = []string{"Maçã", "Banana", "Laranja"}
	fmt.Printf("Frutas: %v\n", frutas)

	// Adicionar elementos com append
	frutas = append(frutas, "Uva")
	frutas = append(frutas, "Morango", "Manga")
	fmt.Printf("Frutas após append: %v\n\n", frutas)

	fmt.Println("=== FATIAMENTO (Slicing) ===\n")

	numeros2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", numeros2)

	// slice[inicio:fim] - fim é EXCLUSIVO
	fmt.Printf("numeros2[2:5]: %v\n", numeros2[2:5]) // [2 3 4]
	fmt.Printf("numeros2[:3]: %v\n", numeros2[:3])   // [0 1 2]
	fmt.Printf("numeros2[5:]: %v\n", numeros2[5:])   // [5 6 7 8 9]
	fmt.Printf("numeros2[:]: %v\n\n", numeros2[:])   // cópia completa

	fmt.Println("=== LEN e CAP ===\n")

	s := make([]int, 3, 5) // len=3, cap=5
	fmt.Printf("Slice: %v\n", s)
	fmt.Printf("len(s): %d (tamanho atual)\n", len(s))
	fmt.Printf("cap(s): %d (capacidade)\n\n", cap(s))

	// Adicionar elementos
	s = append(s, 10, 20)
	fmt.Printf("Após append: %v\n", s)
	fmt.Printf("len(s): %d, cap(s): %d\n\n", len(s), cap(s))

	fmt.Println("=== ITERAÇÃO ===\n")

	linguagens := []string{"Go", "Python", "JavaScript", "Rust"}

	// For tradicional
	fmt.Println("For tradicional:")
	for i := 0; i < len(linguagens); i++ {
		fmt.Printf("  [%d] %s\n", i, linguagens[i])
	}

	// For com range (recomendado)
	fmt.Println("\nFor com range:")
	for indice, linguagem := range linguagens {
		fmt.Printf("  [%d] %s\n", indice, linguagem)
	}

	// Apenas valores
	fmt.Println("\nApenas valores:")
	for _, linguagem := range linguagens {
		fmt.Printf("  - %s\n", linguagem)
	}

	fmt.Println("\n=== COPIAR SLICES ===\n")

	origem := []int{1, 2, 3}
	destino := make([]int, len(origem))

	// CUIDADO: atribuição NÃO copia!
	errado := origem
	errado[0] = 999
	fmt.Printf("Origem: %v (modificado!)\n", origem)

	// Use copy() para copiar
	origem = []int{1, 2, 3}
	copy(destino, origem)
	destino[0] = 999
	fmt.Printf("Origem: %v (não modificado)\n", origem)
	fmt.Printf("Destino: %v\n\n", destino)

	fmt.Println("=== SLICES MULTIDIMENSIONAIS ===\n")

	// Matriz 3x3
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz 3x3:")
	for i, linha := range matriz {
		fmt.Printf("  Linha %d: %v\n", i, linha)
	}

	fmt.Println("\n=== EXEMPLOS PRÁTICOS ===\n")

	// Filtrar números pares
	todos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var pares []int
	for _, num := range todos {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	fmt.Printf("Números pares: %v\n", pares)

	// Somar elementos
	soma := 0
	for _, num := range todos {
		soma += num
	}
	fmt.Printf("Soma de todos: %d\n", soma)

	// Encontrar máximo
	maximo := todos[0]
	for _, num := range todos {
		if num > maximo {
			maximo = num
		}
	}
	fmt.Printf("Maior número: %d\n", maximo)
}

/*
DIFERENÇAS IMPORTANTES:

ARRAYS:
- Tamanho FIXO
- [5]int é diferente de [10]int
- Passado por valor (cópia)
- Sintaxe: var a [5]int

SLICES:
- Tamanho DINÂMICO
- Mais flexíveis e usados
- Passado por referência
- Sintaxe: var s []int

OPERAÇÕES COM SLICES:
- len(s) - tamanho
- cap(s) - capacidade
- append(s, val) - adicionar
- copy(dst, src) - copiar
- s[inicio:fim] - fatiar

CRIAR SLICES:
1. Literal: s := []int{1, 2, 3}
2. make: s := make([]int, len, cap)
3. De array: s := array[:]

Execute com:
    go run 01_arrays_slices.go
*/
