package main

import (
	"fmt"
	"sync"
	"time"
)

/*
WAITGROUP E MUTEX - Sincronização

COMPARAÇÃO COM PYTHON:
--------------------
Python (Thread.join):
    threads = []
    for i in range(5):
        t = Thread(target=worker)
        t.start()
        threads.append(t)

    for t in threads:
        t.join()

Go (WaitGroup):
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            worker()
        }()
    }

    wg.Wait()
*/

func main() {
	fmt.Println("=== WAITGROUP BÁSICO ===\n")

	var wg sync.WaitGroup

	// Lançar 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementar contador

		go func(id int) {
			defer wg.Done() // Decrementar ao finalizar
			fmt.Printf("  Worker %d iniciando\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("  Worker %d finalizando\n", id)
		}(i)
	}

	fmt.Println("Esperando workers...")
	wg.Wait() // Esperar todas terminarem
	fmt.Println("Todos os workers finalizaram!\n")

	fmt.Println("=== PROBLEMA: RACE CONDITION ===\n")

	// Exemplo de race condition
	contador := 0
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			contador++ // ⚠️ RACE CONDITION!
		}()
	}

	wg2.Wait()
	fmt.Printf("Contador: %d (deveria ser 1000)\n", contador)
	fmt.Println("⚠️ Valor incorreto!\n")

	fmt.Println("=== SOLUÇÃO: MUTEX ===\n")

	// Solução com Mutex
	var (
		contador2 int
		mu        sync.Mutex
		wg3       sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mu.Lock()   // Travar
			contador2++ // Região crítica
			mu.Unlock() // Destravar
		}()
	}

	wg3.Wait()
	fmt.Printf("Contador com Mutex: %d ✓\n\n", contador2)

	fmt.Println("=== RWMUTEX (Read-Write Mutex) ===\n")

	exemploBancoComMutex()

	fmt.Println("\n=== EXEMPLO: CONTADOR CONCORRENTE ===\n")

	counter := &SafeCounter{valores: make(map[string]int)}
	var wg4 sync.WaitGroup

	// Múltiplas goroutines incrementando
	for i := 0; i < 100; i++ {
		wg4.Add(1)
		go func(n int) {
			defer wg4.Done()
			key := fmt.Sprintf("key%d", n%10)
			counter.Inc(key)
		}(i)
	}

	wg4.Wait()

	// Ler valores
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Printf("%s: %d\n", key, counter.Value(key))
	}
}

func exemploBancoComMutex() {
	conta := &ContaBancaria{saldo: 1000}
	var wg sync.WaitGroup

	// 10 depósitos simultâneos
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(valor float64) {
			defer wg.Done()
			conta.Depositar(valor)
		}(100.0)
	}

	// 5 saques simultâneos
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(valor float64) {
			defer wg.Done()
			conta.Sacar(valor)
		}(50.0)
	}

	wg.Wait()
	fmt.Printf("Saldo final: R$ %.2f\n", conta.Saldo())
}

// Conta bancária thread-safe
type ContaBancaria struct {
	mu    sync.RWMutex
	saldo float64
}

func (c *ContaBancaria) Depositar(valor float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saldo += valor
	fmt.Printf("  Depositou: R$ %.2f\n", valor)
}

func (c *ContaBancaria) Sacar(valor float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.saldo >= valor {
		c.saldo -= valor
		fmt.Printf("  Sacou: R$ %.2f\n", valor)
	}
}

func (c *ContaBancaria) Saldo() float64 {
	c.mu.RLock() // Read lock
	defer c.mu.RUnlock()
	return c.saldo
}

// Contador thread-safe
type SafeCounter struct {
	mu      sync.Mutex
	valores map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valores[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valores[key]
}

/*
RESUMO:

WAITGROUP:
    var wg sync.WaitGroup

    wg.Add(1)      // Incrementar contador
    wg.Done()      // Decrementar (geralmente com defer)
    wg.Wait()      // Esperar contador chegar a 0

MUTEX:
    var mu sync.Mutex

    mu.Lock()      // Travar
    // região crítica
    mu.Unlock()    // Destravar

RWMUTEX (melhor para muitas leituras):
    var mu sync.RWMutex

    mu.RLock()     // Travar para leitura (múltiplos permitidos)
    mu.RUnlock()   // Destravar leitura

    mu.Lock()      // Travar para escrita (exclusivo)
    mu.Unlock()    // Destravar escrita

QUANDO USAR:

WAITGROUP:
✓ Esperar goroutines finalizarem
✓ Sincronizar conclusão de tarefas
✓ Coordenar trabalho paralelo

MUTEX:
✓ Proteger dados compartilhados
✓ Prevenir race conditions
✓ Acesso exclusivo a recursos

CHANNELS vs MUTEX:

Use CHANNELS quando:
- Passar ownership de dados
- Distribuir unidades de trabalho
- Comunicar resultados

Use MUTEX quando:
- Proteger estado compartilhado
- Cache/estruturas compartilhadas
- Contadores/estatísticas

"Don't communicate by sharing memory; share memory by communicating."
(Mas às vezes Mutex é mais simples!)

DETECÇÃO DE RACE CONDITIONS:
    go run -race programa.go

Execute com:
    go run 04_waitgroup_mutex.go
    go run -race 04_waitgroup_mutex.go
*/
