package main

import (
	"fmt"
	"time"
)

/*
CHANNELS - Comunicação entre Goroutines

COMPARAÇÃO COM PYTHON:
--------------------
Python (Queue):
    from queue import Queue

    q = Queue()
    q.put(10)
    valor = q.get()

Go (Channels):
    ch := make(chan int)
    ch <- 10          // enviar
    valor := <-ch     // receber

VANTAGENS DOS CHANNELS:
- Nativos da linguagem
- Type-safe
- Sincronização automática
- Podem ser fechados
*/

func main() {
	fmt.Println("=== CHANNELS BÁSICOS ===\n")

	// Criar channel
	ch := make(chan string)

	// Enviar em goroutine (evita deadlock)
	go func() {
		ch <- "Olá do channel!" // Enviar
	}()

	// Receber
	mensagem := <-ch // Receber (bloqueia até receber)
	fmt.Println("Recebido:", mensagem)

	fmt.Println("\n=== CHANNELS COMO PARÂMETROS ===\n")

	messages := make(chan string)
	go enviarMensagem(messages)

	msg := <-messages
	fmt.Println("Recebido:", msg)

	fmt.Println("\n=== MÚLTIPLAS MENSAGENS ===\n")

	numeros := make(chan int)

	// Produtor
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("  Enviando: %d\n", i)
			numeros <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(numeros) // Importante: fechar quando terminar!
	}()

	// Consumidor
	for num := range numeros {
		fmt.Printf("  Recebido: %d\n", num)
	}

	fmt.Println("\n=== BUFFERED CHANNELS ===\n")

	// Channel com buffer (não bloqueia até encher)
	buffered := make(chan int, 3) // capacidade 3

	// Pode enviar 3 valores sem bloquear
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("Enviados 3 valores (não bloqueou)")

	// Receber
	fmt.Println("Recebido:", <-buffered)
	fmt.Println("Recebido:", <-buffered)
	fmt.Println("Recebido:", <-buffered)

	fmt.Println("\n=== CHANNEL COMO SINCRONIZAÇÃO ===\n")

	done := make(chan bool)

	go trabalhar(done)

	<-done // Espera sinal de conclusão
	fmt.Println("Trabalho concluído!")

	fmt.Println("\n=== DIREÇÃO DOS CHANNELS ===\n")

	// Channel bidirecional
	ch1 := make(chan string)

	// Apenas enviar
	go enviarApenas(ch1)

	// Apenas receber
	msg1 := <-ch1
	fmt.Println("Recebido:", msg1)

	fmt.Println("\n=== EXEMPLO: PIPELINE ===\n")

	// Pipeline: gerador → quadrado → imprimir
	naturais := gerador(1, 2, 3, 4, 5)
	quadrados := quadrado(naturais)

	for resultado := range quadrados {
		fmt.Printf("Resultado: %d\n", resultado)
	}

	fmt.Println("\n=== EXEMPLO: WORKER POOL ===\n")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Criar 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Enviar 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Coletar resultados
	for r := 1; r <= 5; r++ {
		resultado := <-results
		fmt.Printf("Resultado: %d\n", resultado)
	}
}

func enviarMensagem(ch chan string) {
	ch <- "Mensagem importante!"
}

func trabalhar(done chan bool) {
	fmt.Println("  Trabalhando...")
	time.Sleep(1 * time.Second)
	fmt.Println("  Trabalho finalizado")
	done <- true // Sinalizar conclusão
}

// Channel apenas para enviar
func enviarApenas(ch chan<- string) {
	ch <- "Só posso enviar!"
}

// Channel apenas para receber
func receberApenas(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}

// Pipeline: gerador
func gerador(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Pipeline: quadrado
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

// Worker
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("  Worker %d processando job %d\n", id, j)
		time.Sleep(500 * time.Millisecond)
		results <- j * 2
	}
}

/*
RESUMO DE CHANNELS:

CRIAR:
    ch := make(chan tipo)           // unbuffered
    ch := make(chan tipo, N)        // buffered (cap N)

OPERAÇÕES:
    ch <- valor                     // enviar
    valor := <-ch                   // receber
    valor, ok := <-ch               // receber com verificação
    close(ch)                       // fechar

ITERAR:
    for valor := range ch {
        // processa valor
    }

DIREÇÃO:
    chan<- tipo                     // apenas enviar
    <-chan tipo                     // apenas receber

CARACTERÍSTICAS:
- Unbuffered: bloqueiam até receber/enviar
- Buffered: bloqueiam apenas quando cheio/vazio
- Fechar: indica que não haverá mais valores
- Receber de channel fechado: retorna zero value

PADRÕES COMUNS:
1. Pipeline: ch1 → ch2 → ch3
2. Fan-out: 1 produtor → N consumidores
3. Fan-in: N produtores → 1 consumidor
4. Worker Pool: jobs → workers → results

CUIDADOS:
⚠️ Fechar channel já fechado: PANIC
⚠️ Enviar para channel fechado: PANIC
⚠️ Receber de channel não fechado sem goroutine: DEADLOCK
⚠️ Apenas o sender deve fechar o channel

Execute com:
    go run 02_channels.go
*/
