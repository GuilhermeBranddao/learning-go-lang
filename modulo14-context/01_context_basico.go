package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== CONTEXT COM TIMEOUT ===\n")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go operacaoLonga(ctx, "Operação 1")
	time.Sleep(3 * time.Second)

	fmt.Println("\n=== CONTEXT COM CANCELAMENTO ===\n")

	ctx2, cancel2 := context.WithCancel(context.Background())

	go operacaoLonga(ctx2, "Operação 2")

	time.Sleep(1 * time.Second)
	cancel2() // Cancelar explicitamente

	time.Sleep(1 * time.Second)
}

func operacaoLonga(ctx context.Context, nome string) {
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s cancelada: %v\n", nome, ctx.Err())
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("%s - Passo %d\n", nome, i)
		}
	}
	fmt.Printf("%s concluída!\n", nome)
}

/*
Execute:
    go run 01_context_basico.go
*/
