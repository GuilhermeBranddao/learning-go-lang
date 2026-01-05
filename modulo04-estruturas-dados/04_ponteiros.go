package main

import "fmt"

/*
PONTEIROS

COMPARAÇÃO COM PYTHON:
--------------------
Python NÃO tem ponteiros explícitos:
    x = 10
    y = x  # y é outra referência ao mesmo objeto (para mutáveis)

Go TEM ponteiros explícitos:
    x := 10
    p := &x    // p é um ponteiro para x
    *p = 20    // modifica x através do ponteiro
    fmt.Println(x)  // 20

OPERADORES:
& - endereço de memória (address-of)
* - desreferência (dereference)
*/

func main() {
	fmt.Println("=== PONTEIROS BÁSICOS ===\n")

	// Variável normal
	x := 10
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Endereço de x: %p\n", &x)

	// Ponteiro para x
	p := &x // & obtém o endereço
	fmt.Printf("p = %p (ponteiro para x)\n", p)
	fmt.Printf("*p = %d (valor apontado)\n", *p)

	// Modificar através do ponteiro
	*p = 20 // * desreferencia o ponteiro
	fmt.Printf("Após *p = 20:\n")
	fmt.Printf("  x = %d\n", x)
	fmt.Printf("  *p = %d\n\n", *p)

	fmt.Println("=== PONTEIROS NULOS ===\n")

	var ptr *int
	fmt.Printf("Ponteiro nulo: %v\n", ptr)
	fmt.Printf("É nil? %t\n\n", ptr == nil)

	// CUIDADO: desreferenciar ponteiro nil causa panic!
	// fmt.Println(*ptr)  // PANIC!

	fmt.Println("=== PONTEIROS E FUNÇÕES ===\n")

	a := 10
	fmt.Printf("Antes: a = %d\n", a)

	// Passar por valor (copia)
	dobrarValor(a)
	fmt.Printf("Após dobrarValor: a = %d (não mudou)\n", a)

	// Passar por referência (ponteiro)
	dobrarPonteiro(&a)
	fmt.Printf("Após dobrarPonteiro: a = %d (mudou!)\n\n", a)

	fmt.Println("=== PONTEIROS E STRUCTS ===\n")

	// Struct normal
	pessoa1 := Pessoa{Nome: "João", Idade: 25}
	fmt.Printf("pessoa1: %+v\n", pessoa1)

	// Ponteiro para struct
	pessoa2 := &Pessoa{Nome: "Maria", Idade: 30}
	fmt.Printf("pessoa2 (ponteiro): %+v\n", pessoa2)

	// Go permite acessar campos sem *
	pessoa2.Idade = 31 // poderia ser (*pessoa2).Idade
	fmt.Printf("pessoa2 após modificar: %+v\n\n", pessoa2)

	fmt.Println("=== EXEMPLO: PASSAGEM POR VALOR vs REFERÊNCIA ===\n")

	pessoa := Pessoa{Nome: "Carlos", Idade: 35}
	fmt.Printf("Original: %+v\n", pessoa)

	// Método com receiver por valor (NÃO modifica)
	pessoa.envelhecerValor()
	fmt.Printf("Após envelhecerValor: %+v (não mudou)\n", pessoa)

	// Método com receiver por ponteiro (MODIFICA)
	pessoa.envelhecerPonteiro()
	fmt.Printf("Após envelhecerPonteiro: %+v (mudou!)\n\n", pessoa)

	fmt.Println("=== NEW vs & ===\n")

	// new: aloca e retorna ponteiro (zero value)
	p1 := new(Pessoa)
	fmt.Printf("new(Pessoa): %+v\n", p1)

	// &: ponteiro para struct literal
	p2 := &Pessoa{Nome: "Ana", Idade: 28}
	fmt.Printf("&Pessoa{...}: %+v\n\n", p2)

	fmt.Println("=== EXEMPLO PRÁTICO: LISTA LIGADA ===\n")

	// Criar lista
	lista := &No{Valor: 1}
	lista.Proximo = &No{Valor: 2}
	lista.Proximo.Proximo = &No{Valor: 3}

	// Imprimir lista
	fmt.Print("Lista: ")
	atual := lista
	for atual != nil {
		fmt.Printf("%d -> ", atual.Valor)
		atual = atual.Proximo
	}
	fmt.Println("nil")

	fmt.Println("\n=== QUANDO USAR PONTEIROS ===\n")

	// 1. Modificar valores em funções
	numero := 100
	incrementar(&numero)
	fmt.Printf("Número incrementado: %d\n", numero)

	// 2. Evitar cópias de structs grandes
	grande := StructGrande{Data: [1000]int{}}
	processarGrande(&grande) // Eficiente (não copia 1000 ints)

	// 3. Permitir valores opcionais (nil)
	var opcionalStr *string
	if opcionalStr == nil {
		fmt.Println("String não fornecida")
	}

	texto := "valor"
	opcionalStr = &texto
	fmt.Printf("String fornecida: %s\n", *opcionalStr)
}

// Passar por valor (não modifica)
func dobrarValor(n int) {
	n = n * 2
	fmt.Printf("  Dentro de dobrarValor: n = %d\n", n)
}

// Passar por ponteiro (modifica)
func dobrarPonteiro(n *int) {
	*n = *n * 2
	fmt.Printf("  Dentro de dobrarPonteiro: *n = %d\n", *n)
}

// Struct de exemplo
type Pessoa struct {
	Nome  string
	Idade int
}

// Receiver por valor (NÃO modifica)
func (p Pessoa) envelhecerValor() {
	p.Idade++
	fmt.Printf("  Dentro de envelhecerValor: %+v\n", p)
}

// Receiver por ponteiro (MODIFICA)
func (p *Pessoa) envelhecerPonteiro() {
	p.Idade++
	fmt.Printf("  Dentro de envelhecerPonteiro: %+v\n", p)
}

// Lista ligada (usa ponteiros)
type No struct {
	Valor   int
	Proximo *No
}

// Função que usa ponteiro
func incrementar(n *int) {
	*n++
}

// Struct grande
type StructGrande struct {
	Data [1000]int
}

func processarGrande(s *StructGrande) {
	// Processa sem copiar 1000 inteiros
	_ = s.Data[0]
}

/*
RESUMO DE PONTEIROS:

SINTAXE:
    var p *tipo        // declarar ponteiro
    p = &variavel      // obter endereço
    valor := *p        // obter valor
    *p = novoValor     // modificar valor

DIFERENÇAS COM PYTHON:
- Python: tudo é referência (implícito)
- Go: ponteiros são explícitos

QUANDO USAR PONTEIROS:
1. ✓ Modificar variáveis em funções
2. ✓ Evitar cópias de structs grandes
3. ✓ Valores opcionais (nil)
4. ✓ Compartilhar dados entre funções

QUANDO NÃO USAR:
1. ✗ Tipos pequenos (int, float, bool)
2. ✗ Quando não precisa modificar
3. ✗ Slices, maps, channels (já são referências)

IMPORTANTE:
- nil é o zero value de ponteiros
- Desreferenciar ponteiro nil causa PANIC
- Go permite acessar campos de struct sem * (açúcar sintático)
- new(T) retorna *T com zero value
- &T{} retorna ponteiro para literal

Execute com:
    go run 04_ponteiros.go
*/
