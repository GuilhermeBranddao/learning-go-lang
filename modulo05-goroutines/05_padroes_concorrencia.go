package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
PADRÕES COMUNS DE CONCORRÊNCIA EM GO

Este arquivo demonstra padrões práticos que você
usará frequentemente em programas Go reais.
*/

func main() {
	fmt.Println("=== PADRÃO 1: WORKER POOL ===\n")
	exemploWorkerPool()

	fmt.Println("\n=== PADRÃO 2: FAN-OUT / FAN-IN ===\n")
	exemploFanOutFanIn()

	fmt.Println("\n=== PADRÃO 3: PIPELINE ===\n")
	exemploPipeline()

	fmt.Println("\n=== PADRÃO 4: TIMEOUT E CANCELAMENTO ===\n")
	exemploTimeout()

	fmt.Println("\n=== PADRÃO 5: RATE LIMITING ===\n")
	exemploRateLimiting()
}

// ========================================
// PADRÃO 1: WORKER POOL
// ========================================
func exemploWorkerPool() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Iniciar workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Enviar jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Coletar resultados
	for r := 1; r <= numJobs; r++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("  Worker %d processando job %d\n", id, j)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		results <- j * 2
	}
}

// ========================================
// PADRÃO 2: FAN-OUT / FAN-IN
// ========================================
func exemploFanOutFanIn() {
	// Entrada
	input := gerarNumeros(1, 2, 3, 4, 5)

	// Fan-out: distribuir para múltiplos workers
	c1 := quadrado(input)
	c2 := quadrado(input)
	c3 := quadrado(input)

	// Fan-in: combinar resultados
	for resultado := range merge(c1, c2, c3) {
		fmt.Printf("  Resultado: %d\n", resultado)
	}
}

func gerarNumeros(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func quadrado(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Função para copiar valores
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Fechar out quando todos terminarem
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// ========================================
// PADRÃO 3: PIPELINE
// ========================================
func exemploPipeline() {
	// Pipeline: gerar → filtrar pares → multiplicar por 2
	numeros := gerarNumeros2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	pares := filtrarPares(numeros)
	dobrados := multiplicar(pares, 2)

	for resultado := range dobrados {
		fmt.Printf("  Resultado: %d\n", resultado)
	}
}

func gerarNumeros2(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func filtrarPares(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func multiplicar(in <-chan int, fator int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * fator
		}
		close(out)
	}()
	return out
}

// ========================================
// PADRÃO 4: TIMEOUT E CANCELAMENTO
// ========================================
func exemploTimeout() {
	resultado := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		resultado <- "Operação concluída"
	}()

	select {
	case res := <-resultado:
		fmt.Printf("  %s\n", res)
	case <-time.After(1 * time.Second):
		fmt.Println("  ⏱️ Timeout após 1 segundo")
	}
}

// ========================================
// PADRÃO 5: RATE LIMITING
// ========================================
func exemploRateLimiting() {
	// Limitar a 1 requisição por 200ms
	limiter := time.Tick(200 * time.Millisecond)

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-limiter // Esperar ticker
		fmt.Printf("  Processando request %d às %s\n",
			req, time.Now().Format("15:04:05.000"))
	}
}

/*
RESUMO DOS PADRÕES:

1. WORKER POOL
   - Número fixo de workers
   - Jobs distribuídos entre workers
   - Uso: processar muitas tarefas com recursos limitados

2. FAN-OUT / FAN-IN
   - Fan-out: distribuir trabalho para múltiplos workers
   - Fan-in: combinar resultados
   - Uso: paralelizar processamento

3. PIPELINE
   - Cadeia de processamento
   - Cada estágio é uma goroutine
   - Uso: transformações em série

4. TIMEOUT
   - Limitar tempo de espera
   - Usar time.After com select
   - Uso: evitar esperas infinitas

5. RATE LIMITING
   - Limitar taxa de execução
   - Usar time.Tick
   - Uso: controlar uso de APIs, recursos

OUTROS PADRÕES IMPORTANTES:

6. CANCELLATION
   - context.Context para cancelar operações
   - Uso: cleanup de recursos

7. SEMAPHORE
   - Limitar número de goroutines
   - Usar buffered channel
   - Uso: controlar concorrência

8. PUBLISH/SUBSCRIBE
   - Múltiplos subscribers para eventos
   - Uso: notificações, eventos

9. FUTURE/PROMISE
   - Valor que será calculado no futuro
   - Usar channel com capacidade 1
   - Uso: operações assíncronas

QUANDO USAR CADA PADRÃO:

Worker Pool:
  ✓ Processar muitos itens
  ✓ Limitar recursos (CPU, memória)
  ✓ Download de arquivos, processamento de imagens

Pipeline:
  ✓ Transformações em série
  ✓ Streaming de dados
  ✓ Processamento de logs

Fan-out/Fan-in:
  ✓ Paralelizar tarefas independentes
  ✓ Agregar resultados
  ✓ Busca paralela, agregação

Execute com:
    go run 05_padroes_concorrencia.go
*/
