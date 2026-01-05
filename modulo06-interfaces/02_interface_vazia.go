package main

import "fmt"

/*
INTERFACE VAZIA (interface{} ou any)

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    def processar(valor: Any):  # Type hint
        # Aceita qualquer tipo
        print(valor)

Go:
    func processar(valor interface{}) {
        // Aceita qualquer tipo
        fmt.Println(valor)
    }

    // Go 1.18+
    func processar(valor any) {
        fmt.Println(valor)
    }

interface{} é satisfeita por QUALQUER tipo!
*/

func main() {
	fmt.Println("=== INTERFACE VAZIA ===\n")

	// interface{} pode conter qualquer tipo
	var x interface{}

	x = 42
	fmt.Printf("x = %v (tipo: %T)\n", x, x)

	x = "texto"
	fmt.Printf("x = %v (tipo: %T)\n", x, x)

	x = true
	fmt.Printf("x = %v (tipo: %T)\n", x, x)

	x = []int{1, 2, 3}
	fmt.Printf("x = %v (tipo: %T)\n", x, x)

	fmt.Println("\n=== SLICE DE INTERFACE{} ===\n")

	// Slice com tipos misturados
	items := []interface{}{
		42,
		"Go",
		3.14,
		true,
		[]string{"a", "b"},
		map[string]int{"x": 10},
	}

	for i, item := range items {
		fmt.Printf("[%d] %v (tipo: %T)\n", i, item, item)
	}

	fmt.Println("\n=== TYPE ASSERTIONS ===\n")

	// Type assertion: verificar e converter tipo
	var valor interface{} = "Hello, Go!"

	// Forma 1: pode causar panic se tipo errado
	// texto := valor.(string)

	// Forma 2: segura (recomendada)
	if texto, ok := valor.(string); ok {
		fmt.Printf("É uma string: %s\n", texto)
	} else {
		fmt.Println("Não é uma string")
	}

	// Testar tipo errado
	if numero, ok := valor.(int); ok {
		fmt.Printf("É um int: %d\n", numero)
	} else {
		fmt.Println("Não é um int")
	}

	fmt.Println("\n=== TYPE SWITCH ===\n")

	descreveTipo := func(v interface{}) {
		switch v := v.(type) {
		case int:
			fmt.Printf("Inteiro: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		case float64:
			fmt.Printf("Float: %.2f\n", v)
		case []int:
			fmt.Printf("Slice de ints: %v\n", v)
		default:
			fmt.Printf("Tipo desconhecido: %T = %v\n", v, v)
		}
	}

	descreveTipo(42)
	descreveTipo("Go")
	descreveTipo(true)
	descreveTipo(3.14)
	descreveTipo([]int{1, 2, 3})
	descreveTipo(map[string]int{"a": 1})

	fmt.Println("\n=== EXEMPLO: FUNÇÃO GENÉRICA ===\n")

	Imprimir(42)
	Imprimir("texto")
	Imprimir(3.14)
	Imprimir([]int{1, 2, 3})

	fmt.Println("\n=== EXEMPLO: MAP COM VALORES MISTOS ===\n")

	config := map[string]interface{}{
		"host":    "localhost",
		"port":    8080,
		"ssl":     true,
		"timeout": 30.5,
		"tags":    []string{"web", "api"},
	}

	for chave, valor := range config {
		fmt.Printf("%s = %v (tipo: %T)\n", chave, valor, valor)
	}

	// Acessar e converter
	if porta, ok := config["port"].(int); ok {
		fmt.Printf("\nPorta configurada: %d\n", porta)
	}

	fmt.Println("\n=== EXEMPLO: VALIDAÇÃO DE TIPOS ===\n")

	processar := func(dados interface{}) {
		switch v := dados.(type) {
		case string:
			fmt.Printf("Processando string de tamanho %d: %s\n", len(v), v)
		case int:
			fmt.Printf("Processando número: %d (dobro: %d)\n", v, v*2)
		case []int:
			soma := 0
			for _, n := range v {
				soma += n
			}
			fmt.Printf("Processando slice: soma = %d\n", soma)
		case map[string]int:
			fmt.Printf("Processando map com %d elementos\n", len(v))
		default:
			fmt.Printf("Tipo não suportado: %T\n", v)
		}
	}

	processar("Go Lang")
	processar(42)
	processar([]int{1, 2, 3, 4, 5})
	processar(map[string]int{"a": 1, "b": 2})
	processar(3.14) // Não suportado
}

// Função que aceita qualquer tipo
func Imprimir(v interface{}) {
	fmt.Printf("Valor: %v (Tipo: %T)\n", v, v)
}

/*
RESUMO DE INTERFACE VAZIA:

DECLARAÇÃO:
    var x interface{}    // Go < 1.18
    var x any           // Go 1.18+ (alias para interface{})

TYPE ASSERTION:
    // Segura
    valor, ok := x.(TipoDesejado)
    if ok {
        // usar valor
    }

    // Perigosa (pode causar panic)
    valor := x.(TipoDesejado)

TYPE SWITCH:
    switch v := x.(type) {
    case int:
        // v é int
    case string:
        // v é string
    default:
        // tipo desconhecido
    }

QUANDO USAR:
✓ APIs flexíveis
✓ Parsing de JSON/dados desconhecidos
✓ Containers genéricos
✓ Reflection

QUANDO EVITAR:
✗ Quando tipo é conhecido
✗ Para "generalizar" desnecessariamente
✗ Em código type-safe

ALTERNATIVAS MODERNAS:
- Generics (Go 1.18+) para código type-safe
- Interfaces específicas quando possível

CUIDADOS:
⚠️ Perde type safety
⚠️ Precisa type assertions
⚠️ Pode causar panic em runtime
⚠️ Dificulta debugging

Execute com:
    go run 02_interface_vazia.go
*/
