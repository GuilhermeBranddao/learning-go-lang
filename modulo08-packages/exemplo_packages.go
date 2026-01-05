package main

import (
	"fmt"
	"go-course/modulo08-packages/matematica"
	"go-course/modulo08-packages/utils"
)

/*
PACKAGES EM GO

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    # arquivo: matematica.py
    def somar(a, b):
        return a + b

    # arquivo: main.py
    import matematica
    resultado = matematica.somar(1, 2)

Go:
    // arquivo: matematica/calc.go
    package matematica

    func Somar(a, b int) int {
        return a + b
    }

    // arquivo: main.go
    package main
    import "meuapp/matematica"
    resultado := matematica.Somar(1, 2)

ESTRUTURA DE DIRETÓRIOS:
    meuapp/
    ├── go.mod
    ├── main.go
    └── matematica/
        └── calc.go

DIFERENÇAS PRINCIPAIS:
1. Package = DIRETÓRIO (não arquivo)
2. Todos arquivos no diretório = mesmo package
3. Visibilidade por LETRA MAIÚSCULA
4. Import pelo caminho do módulo
*/

func main() {
	fmt.Println("=== USANDO PACKAGES ===\n")

	// Usando package matematica
	soma := matematica.Somar(10, 5)
	fmt.Printf("10 + 5 = %d\n", soma)

	subtracao := matematica.Subtrair(10, 5)
	fmt.Printf("10 - 5 = %d\n", subtracao)

	multiplicacao := matematica.Multiplicar(10, 5)
	fmt.Printf("10 * 5 = %d\n", multiplicacao)

	divisao, err := matematica.Dividir(10, 5)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", divisao)
	}

	fmt.Println("\n=== USANDO PACKAGE UTILS ===\n")

	// Usando package utils
	texto := "Olá, Mundo!"
	fmt.Printf("Texto original: %s\n", texto)
	fmt.Printf("Maiúsculo: %s\n", utils.ToUpper(texto))
	fmt.Printf("Minúsculo: %s\n", utils.ToLower(texto))
	fmt.Printf("Invertido: %s\n", utils.Reverter(texto))

	numeros := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("\nNúmeros: %v\n", numeros)
	fmt.Printf("Máximo: %d\n", utils.Maximo(numeros))
	fmt.Printf("Mínimo: %d\n", utils.Minimo(numeros))
	fmt.Printf("Soma: %d\n", utils.SomaArray(numeros))

	fmt.Println("\n=== CONSTANTES DO PACKAGE ===\n")

	// Constantes exportadas
	fmt.Println("Versão da biblioteca:", matematica.Versao)
	fmt.Println("Autor:", matematica.Autor)
}

/*
RESUMO DE PACKAGES:

DEFINIR PACKAGE:
    package nome
    // Todo código no arquivo pertence a esse package

EXPORTAR (Público):
    func Publica() { }      // Maiúscula = Exportada
    var PublicaVar int      // Acessível de fora
    const Constante = 10

NÃO EXPORTAR (Privado):
    func privada() { }      // Minúscula = Privada
    var privadaVar int      // Apenas no package

IMPORTAR:
    import "path/do/package"
    import (
        "fmt"
        "meuapp/matematica"
    )

USAR:
    matematica.Somar(1, 2)
    fmt.Println("ok")

ESTRUTURA:
    meuapp/
    ├── go.mod           # module meuapp
    ├── main.go          # package main
    ├── matematica/
    │   └── calc.go      # package matematica
    └── utils/
        └── helpers.go   # package utils

CRIAR MÓDULO:
    go mod init github.com/usuario/meuapp

DIFERENÇAS COM PYTHON:

Python:
✓ Arquivo = módulo
✓ import arquivo
✓ _prefixo para privado
✓ pip install
✓ requirements.txt

Go:
✓ Diretório = package
✓ import "path/package"
✓ Letra maiúscula para público
✓ go get
✓ go.mod + go.sum

REGRAS IMPORTANTES:
✓ Um diretório = um package
✓ Todos .go no diretório = mesmo package
✓ package main = executável
✓ outros packages = bibliotecas
✓ Primeiro caractere maiúsculo = público

Execute com:
    go run main.go
*/
