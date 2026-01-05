package main

import (
	"fmt"
	"time"
)

/*
SELECT - Multiplexação de Channels

COMPARAÇÃO COM PYTHON:
--------------------
Python não tem equivalente direto.
Você precisaria usar select.select() ou asyncio.

Go tem SELECT nativo:
    select {
    case msg1 := <-ch1:
        // processar ch1
    case msg2 := <-ch2:
        // processar ch2
    }

SELECT permite esperar por múltiplos channels simultaneamente!
*/

func main() {
	fmt.Println("=== SELECT BÁSICO ===\n")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Mensagem do canal 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Mensagem do canal 2"
	}()

	// Select espera qualquer um dos channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido de ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido de ch2:", msg2)
		}
	}

	fmt.Println("\n=== SELECT COM TIMEOUT ===\n")

	resultado := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		resultado <- "Resultado lento"
	}()

	select {
	case res := <-resultado:
		fmt.Println("Sucesso:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("⏱️ Timeout! Operação demorou demais")
	}

	fmt.Println("\n=== SELECT COM DEFAULT ===\n")

	messages := make(chan string)

	// Não bloqueante com default
	select {
	case msg := <-messages:
		fmt.Println("Recebido:", msg)
	default:
		fmt.Println("Nenhuma mensagem disponível")
	}

	fmt.Println("\n=== EXEMPLO: POLLING ===\n")

	tick := time.NewTicker(500 * time.Millisecond)
	boom := time.After(2 * time.Second)

	for {
		select {
		case <-tick.C:
			fmt.Println("  tick.")
		case <-boom:
			fmt.Println("  BOOM!")
			tick.Stop()
			return
		default:
			fmt.Print("  .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

/*
RESUMO DE SELECT:

SINTAXE:
    select {
    case v := <-ch1:
        // canal 1
    case v := <-ch2:
        // canal 2
    case ch3 <- valor:
        // enviar para ch3
    default:
        // se nenhum estiver pronto
    }

CARACTERÍSTICAS:
- Espera múltiplos channels
- Executa o primeiro case que estiver pronto
- Se vários estiverem prontos, escolhe aleatoriamente
- Default torna não-bloqueante

PADRÕES COMUNS:

1. TIMEOUT:
    select {
    case res := <-ch:
        // processar
    case <-time.After(1 * time.Second):
        // timeout
    }

2. NON-BLOCKING:
    select {
    case msg := <-ch:
        // processar
    default:
        // fazer outra coisa
    }

3. TICKER:
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            // executar periodicamente
        }
    }

4. QUIT CHANNEL:
    select {
    case msg := <-messages:
        // processar
    case <-quit:
        return
    }

DIFERENÇAS COM SWITCH:
- Select é para channels
- Select espera operações de I/O
- Select pode bloquear
- Switch não bloqueia

Execute com:
    go run 03_select.go
*/
