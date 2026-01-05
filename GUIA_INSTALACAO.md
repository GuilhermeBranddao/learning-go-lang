# 🚀 Guia de Instalação do Go

## Passo 1: Baixar o Go

1. Acesse: https://go.dev/dl/
2. Baixe a versão mais recente para Windows (arquivo `.msi`)
3. Execute o instalador e siga as instruções

## Passo 2: Verificar a Instalação

Abra o PowerShell ou CMD e execute:

```bash
go version
```

Você deve ver algo como: `go version go1.21.x windows/amd64`

## Passo 3: Configurar GOPATH (Opcional)

O Go moderno usa módulos, então não precisa configurar GOPATH, mas é bom conhecer:

```bash
go env GOPATH
```

## Passo 4: Primeiro Programa

Crie um arquivo `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Olá, Go!")
}
```

Execute:

```bash
go run hello.go
```

## Comandos Essenciais

- `go run arquivo.go` - Compila e executa
- `go build arquivo.go` - Compila para executável
- `go fmt arquivo.go` - Formata o código (muito importante!)
- `go mod init nome-projeto` - Inicializa um módulo Go

## VS Code - Extensão Recomendada

Instale a extensão oficial do Go para VS Code:
- Nome: **Go** (by Go Team at Google)
- ID: `golang.go`

Ela oferece autocomplete, formatação automática, debugging e muito mais!

---

✅ **Pronto! Agora você está pronto para começar!**
