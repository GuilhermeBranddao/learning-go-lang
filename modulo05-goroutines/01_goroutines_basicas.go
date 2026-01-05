package main

import (
	"fmt"
	"time"
)

/*
GOROUTINES - Concorrência Simples

COMPARAÇÃO COM PYTHON:
--------------------
Python (threading):
    import threading

    def tarefa():
        print("Executando")

    t = threading.Thread(target=tarefa)
    t.start()
    t.join()

Go (goroutines):
    func tarefa() {
        fmt.Println("Executando")
    }

    go tarefa()  // Tão simples quanto isso!

DIFERENÇAS:
- Goroutines são MUITO mais leves que threads
- Python tem GIL (limita paralelismo real)
- Go pode rodar milhares de goroutines simultaneamente
*/

func main() {
	fmt.Println("=== GOROUTINES BÁSICAS ===\n")

	// Função normal - execução sequencial
	fmt.Println("Chamada normal:")
	dizer("Mundo")
	fmt.Println("Continuou após dizer()\n")

	// Goroutine - execução concorrente
	fmt.Println("Com goroutine:")
	go dizer("Mundo") // Não espera terminar!
	fmt.Println("Continuou imediatamente!")

	// PROBLEMA: programa termina antes da goroutine!
	// Vamos esperar um pouco
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== MÚLTIPLAS GOROUTINES ===\n")

	// Lançar várias goroutines
	for i := 0; i < 5; i++ {
		go contarAte(i, 3)
	}

	// Esperar goroutines terminarem
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== GOROUTINE COM FUNÇÃO ANÔNIMA ===\n")

	// Goroutine com closure
	mensagem := "Olá"
	go func() {
		fmt.Println(mensagem, "do closure!")
	}()

	time.Sleep(100 * time.Millisecond)

	// Goroutine com parâmetros
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("Goroutine %d executando\n", n)
		}(i) // Passa i como argumento!
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== EXEMPLO: DOWNLOAD CONCORRENTE ===\n")

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
	}

	fmt.Println("Baixando URLs (simulado)...")
	for _, url := range urls {
		go baixarURL(url)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\n=== PROBLEMA: RACE CONDITION ===\n")

	// Exemplo de problema com dados compartilhados
	contador := 0

	for i := 0; i < 1000; i++ {
		go func() {
			contador++ // RACE CONDITION!
		}()
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Contador: %d (deveria ser 1000)\n", contador)
	fmt.Println("⚠️ Valor incorreto devido a race condition!")
	fmt.Println("Solução: usar Mutex ou Channels (próximos exemplos)")
}

func dizer(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("  %s %d\n", texto, i)
		time.Sleep(50 * time.Millisecond)
	}
}

func contarAte(id, max int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("  Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func baixarURL(url string) {
	fmt.Printf("  Iniciando download: %s\n", url)

	// Simular tempo de download
	time.Sleep(time.Duration(500+len(url)*10) * time.Millisecond)

	fmt.Printf("  ✓ Concluído: %s\n", url)
}

/*
RESUMO DE GOROUTINES:

SINTAXE:
    go funcao()
    go funcao(args)
    go func() { codigo }()

CARACTERÍSTICAS:
- Muito leves (alguns KB de memória)
- Escalonadas pelo Go runtime (não pelo OS)
- Podem rodar milhares simultaneamente
- Comunicação via channels (próximo exemplo)

CUIDADOS:
1. ⚠️ Goroutines não são garbage collected
2. ⚠️ Programa termina se main() terminar
3. ⚠️ Cuidado com closures em loops
4. ⚠️ Dados compartilhados precisam sincronização

EXEMPLO CORRETO em loop:
    // ❌ ERRADO
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)  // i pode mudar!
        }()
    }

    // ✅ CORRETO
    for i := 0; i < 5; i++ {
        go func(n int) {
            fmt.Println(n)  // n é cópia
        }(i)
    }

COMPARAÇÃO DE PERFORMANCE:
- Thread (Python): ~1-2 MB cada
- Goroutine (Go): ~2 KB cada
- Resultado: Go pode ter 1000x mais goroutines!

Execute com:
    go run 01_goroutines_basicas.go

Para detectar race conditions:
    go run -race 01_goroutines_basicas.go
*/
