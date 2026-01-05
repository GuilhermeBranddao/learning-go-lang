package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python:
    print("Hello, World!")

Go precisa de:
1. package main - define que é um programa executável
2. import - importa pacotes necessários
3. func main() - ponto de entrada do programa
*/

func main() {
	// fmt.Println adiciona quebra de linha automaticamente
	fmt.Println("Hello, World!")

	// fmt.Print não adiciona quebra de linha
	fmt.Print("Hello, ")
	fmt.Print("Go!\n")

	// fmt.Printf permite formatação (como f-string em Python)
	nome := "Gopher"
	fmt.Printf("Olá, %s!\n", nome)
}

/*
Para executar:
    go run 01_hello.go

Para compilar:
    go build 01_hello.go
    ./01_hello (ou 01_hello.exe no Windows)
*/
