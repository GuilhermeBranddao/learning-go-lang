package main

import "fmt"

// ========================================
// GENERICS (Go 1.18+)
// ========================================

// Função genérica com constraint
func Maximo[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Type constraint personalizado
type Number interface {
	int | int64 | float64
}

func Somar[T Number](a, b T) T {
	return a + b
}

// Slice genérico
func Contem[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Stack genérico
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func main() {
	fmt.Println("=== GENERICS ===\n")

	// Função genérica com int
	fmt.Println("Máximo(10, 20):", Maximo(10, 20))

	// Mesma função com float64
	fmt.Println("Máximo(3.14, 2.71):", Maximo(3.14, 2.71))

	// Somar genérico
	fmt.Println("Somar(5, 3):", Somar(5, 3))
	fmt.Println("Somar(2.5, 3.7):", Somar(2.5, 3.7))

	fmt.Println("\n=== GENERICS COM SLICES ===\n")

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Contém 3?", Contem(numeros, 3))
	fmt.Println("Contém 10?", Contem(numeros, 10))

	nomes := []string{"João", "Maria", "Pedro"}
	fmt.Println("Contém 'Maria'?", Contem(nomes, "Maria"))

	fmt.Println("\n=== STACK GENÉRICO ===\n")

	// Stack de inteiros
	stackInt := Stack[int]{}
	stackInt.Push(10)
	stackInt.Push(20)
	stackInt.Push(30)

	for {
		item, ok := stackInt.Pop()
		if !ok {
			break
		}
		fmt.Println("Pop:", item)
	}

	// Stack de strings
	stackStr := Stack[string]{}
	stackStr.Push("Go")
	stackStr.Push("é")
	stackStr.Push("incrível")

	for {
		item, ok := stackStr.Pop()
		if !ok {
			break
		}
		fmt.Println("Pop:", item)
	}
}

/*
RESUMO DE GENERICS:

SINTAXE:
    func Nome[T constraint](param T) T { }

CONSTRAINTS:
    - any: qualquer tipo
    - comparable: tipos comparáveis (==, !=)
    - int | float64: união de tipos
    - interface customizada

EXEMPLOS:
    func Min[T int | float64](a, b T) T
    type Stack[T any] struct { }

QUANDO USAR:
✓ Estruturas de dados genéricas
✓ Algoritmos que funcionam com vários tipos
✓ Evitar interface{} + type assertion

Execute:
    go run 01_generics.go
*/
