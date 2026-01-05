package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== ESCREVER ARQUIVO ===\n")

	// Escrever arquivo inteiro
	conteudo := []byte("Olá, Go!\nEsta é a segunda linha.\n")
	err := os.WriteFile("exemplo.txt", conteudo, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Arquivo exemplo.txt criado")

	fmt.Println("\n=== LER ARQUIVO ===\n")

	// Ler arquivo inteiro
	data, err := os.ReadFile("exemplo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conteúdo:")
	fmt.Println(string(data))

	fmt.Println("\n=== LER LINHA POR LINHA ===\n")

	// Criar arquivo com múltiplas linhas
	linhas := "Linha 1\nLinha 2\nLinha 3\nLinha 4\n"
	os.WriteFile("linhas.txt", []byte(linhas), 0644)

	// Ler linha por linha com bufio
	file, err := os.Open("linhas.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	contador := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", contador, scanner.Text())
		contador++
	}

	fmt.Println("\n=== TRABALHAR COM DIRETÓRIOS ===\n")

	// Criar diretório
	err = os.Mkdir("meu_diretorio", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	fmt.Println("✓ Diretório criado")

	// Listar arquivos
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nArquivos no diretório atual:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("[DIR]  %s\n", entry.Name())
		} else {
			fmt.Printf("[FILE] %s\n", entry.Name())
		}
	}

	fmt.Println("\n=== CAMINHOS ===\n")

	// Trabalhar com caminhos
	caminho := filepath.Join("meu_diretorio", "arquivo.txt")
	fmt.Println("Caminho:", caminho)
	fmt.Println("Base:", filepath.Base(caminho))
	fmt.Println("Dir:", filepath.Dir(caminho))
	fmt.Println("Extensão:", filepath.Ext(caminho))

	// Limpeza
	os.Remove("exemplo.txt")
	os.Remove("linhas.txt")
	os.RemoveAll("meu_diretorio")
}

/*
Execute com:
    go run 01_arquivos_basico.go
*/
