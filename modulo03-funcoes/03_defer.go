package main

import "fmt"

/*
DEFER - Adia execução até o fim da função

COMPARAÇÃO COM PYTHON:
--------------------
Python usa context managers (with):
    with open("arquivo.txt") as f:
        # usar arquivo
    # arquivo fechado automaticamente

Go usa defer:
    f, _ := os.Open("arquivo.txt")
    defer f.Close()
    // usar arquivo
    // arquivo fechado automaticamente no final da função

DEFER é perfeito para cleanup!
*/

func main() {
	fmt.Println("=== DEFER BÁSICO ===\n")

	exemploBasico()

	fmt.Println("\n=== MÚLTIPLOS DEFERS ===\n")

	multiplosDefers()

	fmt.Println("\n=== DEFER COM LOOPS ===\n")

	deferComLoops()

	fmt.Println("\n=== DEFER COM PARÂMETROS ===\n")

	deferComParametros()

	fmt.Println("\n=== EXEMPLO PRÁTICO ===\n")

	processarDados()
}

func exemploBasico() {
	fmt.Println("Início da função")

	// defer adia a execução até o final da função
	defer fmt.Println("Executado no FINAL (defer)")

	fmt.Println("Meio da função")
	fmt.Println("Fim da função")

	// Ordem de execução:
	// 1. "Início da função"
	// 2. "Meio da função"
	// 3. "Fim da função"
	// 4. "Executado no FINAL (defer)"
}

func multiplosDefers() {
	fmt.Println("Função com múltiplos defers")

	// Defers são executados em ordem LIFO (Last In, First Out)
	// Como uma pilha!
	defer fmt.Println("  Defer 1")
	defer fmt.Println("  Defer 2")
	defer fmt.Println("  Defer 3")

	fmt.Println("Fim da função normal")

	// Ordem de execução:
	// 1. "Função com múltiplos defers"
	// 2. "Fim da função normal"
	// 3. "Defer 3" (último defer é executado primeiro)
	// 4. "Defer 2"
	// 5. "Defer 1"
}

func deferComLoops() {
	// CUIDADO: defer em loop pode causar problemas
	fmt.Println("Exemplo de defer em loop:")

	for i := 1; i <= 3; i++ {
		// Este defer só será executado no FINAL da função
		// NÃO no final de cada iteração!
		defer fmt.Printf("  Defer do loop i=%d\n", i)
		fmt.Printf("  Iteração %d\n", i)
	}

	fmt.Println("Loop terminado")

	// Ordem:
	// Iteração 1, 2, 3
	// "Loop terminado"
	// Defer i=3, i=2, i=1 (ordem inversa!)
}

func deferComParametros() {
	x := 10
	fmt.Printf("x antes do defer: %d\n", x)

	// Os ARGUMENTOS do defer são avaliados IMEDIATAMENTE
	defer fmt.Printf("  Defer: x = %d (valor capturado)\n", x)

	x = 20
	fmt.Printf("x depois de mudar: %d\n", x)

	// O defer imprimirá x=10, não x=20!
	// Porque o valor foi capturado no momento do defer
}

func processarDados() {
	fmt.Println("Iniciando processamento...")

	// Padrão comum: usar defer para cleanup
	defer fmt.Println("  ✓ Finalizando processamento")
	defer fmt.Println("  ✓ Liberando recursos")
	defer fmt.Println("  ✓ Fechando conexões")

	fmt.Println("Processando dados...")
	fmt.Println("Calculando resultados...")

	// Mesmo se houver erro/return/panic, os defers serão executados!
}

// Exemplo prático: simulação de abertura/fechamento de arquivo
func abrirArquivo() {
	fmt.Println("\n=== EXEMPLO: ARQUIVO ===\n")

	fmt.Println("Abrindo arquivo...")
	defer fmt.Println("Fechando arquivo (defer)")

	fmt.Println("Lendo dados...")
	fmt.Println("Processando dados...")

	// defer garante que o arquivo será fechado,
	// mesmo se houver erro no meio do processamento
}

/*
RESUMO DE DEFER:

1. Adia execução até o fim da função
2. Múltiplos defers = ordem LIFO (pilha)
3. Argumentos avaliados imediatamente
4. Executado SEMPRE (mesmo com panic/return)

Usos comuns:
- Fechar arquivos: defer file.Close()
- Fechar conexões: defer conn.Close()
- Unlock mutexes: defer mutex.Unlock()
- Cleanup de recursos

Sintaxe:
    defer funcao()
    defer funcao(argumentos)

ATENÇÃO em loops:
    // ❌ Ruim
    for i := 0; i < 10; i++ {
        f := open(file)
        defer f.Close()  // só fecha no fim da função!
    }

    // ✅ Bom
    for i := 0; i < 10; i++ {
        func() {
            f := open(file)
            defer f.Close()  // fecha no fim desta função
        }()
    }

Execute com:
    go run 03_defer.go
*/
