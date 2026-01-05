package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
A maioria dos operadores é similar, mas com algumas diferenças:
- Go não tem operador ** para potência (use math.Pow)
- Go não tem operadores += para strings (use + ou fmt.Sprintf)
- Divisão de inteiros sempre resulta em inteiro
*/

func main() {
	fmt.Println("=== OPERADORES ARITMÉTICOS ===\n")

	a := 10
	b := 3

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)    // 13
	fmt.Printf("a - b = %d\n", a-b)    // 7
	fmt.Printf("a * b = %d\n", a*b)    // 30
	fmt.Printf("a / b = %d\n", a/b)    // 3 (divisão inteira!)
	fmt.Printf("a %% b = %d\n\n", a%b) // 1 (módulo/resto)

	// ATENÇÃO: Divisão de inteiros
	fmt.Println("CUIDADO COM DIVISÃO:")
	fmt.Printf("10 / 3 = %d (inteiros)\n", 10/3)
	fmt.Printf("10.0 / 3.0 = %.2f (floats)\n\n", 10.0/3.0)

	// Incremento e decremento
	contador := 0
	contador++ // incrementa (não retorna valor!)
	fmt.Printf("contador++: %d\n", contador)
	contador-- // decrementa
	fmt.Printf("contador--: %d\n\n", contador)

	// IMPORTANTE: em Go, ++ e -- são STATEMENTS, não expressões
	// x := i++  // ERRO! Não funciona
	// if i++ > 0 {}  // ERRO! Não funciona

	fmt.Println("=== OPERADORES DE COMPARAÇÃO ===\n")

	x := 10
	y := 20

	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("x == y: %t\n", x == y)   // false
	fmt.Printf("x != y: %t\n", x != y)   // true
	fmt.Printf("x > y: %t\n", x > y)     // false
	fmt.Printf("x < y: %t\n", x < y)     // true
	fmt.Printf("x >= y: %t\n", x >= y)   // false
	fmt.Printf("x <= y: %t\n\n", x <= y) // true

	fmt.Println("=== OPERADORES LÓGICOS ===\n")

	aprovado := true
	presenca := true

	fmt.Printf("Aprovado: %t, Presença: %t\n", aprovado, presenca)
	fmt.Printf("aprovado && presenca: %t\n", aprovado && presenca) // AND
	fmt.Printf("aprovado || presenca: %t\n", aprovado || presenca) // OR
	fmt.Printf("!aprovado: %t\n\n", !aprovado)                     // NOT

	fmt.Println("=== OPERADORES DE ATRIBUIÇÃO ===\n")

	numero := 10
	fmt.Printf("Inicial: %d\n", numero)

	numero += 5 // numero = numero + 5
	fmt.Printf("Após += 5: %d\n", numero)

	numero -= 3 // numero = numero - 3
	fmt.Printf("Após -= 3: %d\n", numero)

	numero *= 2 // numero = numero * 2
	fmt.Printf("Após *= 2: %d\n", numero)

	numero /= 4 // numero = numero / 4
	fmt.Printf("Após /= 4: %d\n", numero)

	numero %= 3 // numero = numero % 3
	fmt.Printf("Após %%= 3: %d\n\n", numero)

	fmt.Println("=== OPERADORES BIT A BIT ===\n")

	p := 12 // 1100 em binário
	q := 10 // 1010 em binário

	fmt.Printf("p = %d (binário: %04b)\n", p, p)
	fmt.Printf("q = %d (binário: %04b)\n", q, q)
	fmt.Printf("p & q = %d (binário: %04b)\n", p&q, p&q)    // AND
	fmt.Printf("p | q = %d (binário: %04b)\n", p|q, p|q)    // OR
	fmt.Printf("p ^ q = %d (binário: %04b)\n", p^q, p^q)    // XOR
	fmt.Printf("p << 1 = %d (binário: %04b)\n", p<<1, p<<1) // Shift left
	fmt.Printf("p >> 1 = %d (binário: %04b)\n", p>>1, p>>1) // Shift right
}

/*
DIFERENÇAS IMPORTANTES EM RELAÇÃO AO PYTHON:

1. ++ e -- são statements, não expressões:
   Python: x = i++  ✓
   Go: x = i++      ✗

2. Divisão de inteiros SEMPRE retorna inteiro:
   Python: 10 / 3 = 3.333... (float)
   Go: 10 / 3 = 3 (int)

3. Não existe operador ** para potência:
   Python: 2 ** 3 = 8
   Go: math.Pow(2, 3) = 8.0

Execute com:
    go run 04_operadores.go
*/
