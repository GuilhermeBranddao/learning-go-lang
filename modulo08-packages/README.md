# 📦 Módulo 08 - Packages e Módulos

## 🎯 Objetivos
Dominar a organização de código em pacotes e módulos Go.

## 🔑 Conceitos Fundamentais

### Comparação com Python

| Conceito | Python | Go |
|----------|--------|-----|
| **Unidade** | Módulo (arquivo .py) | Package (diretório) |
| **Import** | `import modulo` | `import "path/package"` |
| **Visibilidade** | `_prefixo` ou `__prefixo__` | Primeira letra maiúscula |
| **Gerenciador** | pip, conda, poetry | go mod |

---

## 📝 Estrutura de Package

```
meuapp/
├── go.mod                 # Definição do módulo
├── go.sum                 # Lock file
├── main.go               # Ponto de entrada
└── pacotes/
    ├── matematica/
    │   └── calc.go       # package matematica
    └── utils/
        └── helpers.go    # package utils
```

---

## 🔧 go.mod e go.sum

### go.mod
```go
module github.com/usuario/meuapp

go 1.21

require (
    github.com/gorilla/mux v1.8.0
)
```

### Comandos
```bash
go mod init github.com/usuario/meuapp  # Criar módulo
go mod tidy                            # Limpar dependências
go get github.com/pacote@v1.0.0       # Adicionar dependência
```

---

## 👁️ Visibilidade

```go
// Exportado (público) - primeira letra MAIÚSCULA
func Publica() { }
var PublicVar int
type PublicStruct struct { }

// Não exportado (privado) - primeira letra minúscula
func privada() { }
var privateVar int
type privateStruct struct { }
```

**Regra:** Se começa com **maiúscula**, é **público**!

---

## 📂 Organização de Código

### Package main
```go
package main  // Apenas para executáveis

func main() {
    // Ponto de entrada
}
```

### Packages de biblioteca
```go
package matematica  // Para reutilização

func Somar(a, b int) int {
    return a + b
}
```

---

## 📥 Imports

```go
import "fmt"                          // Stdlib
import "github.com/usuario/pacote"   // Externo
import "./local"                     // Local (evitar!)
import (
    "fmt"
    m "matematica"  // Alias
    . "fmt"         // Import tudo (evitar!)
    _ "pacote"      // Side effects apenas
)
```

---

## 📦 internal/

```
meuapp/
└── internal/
    └── segredo/
        └── config.go
```

**Regra:** Pacotes em `internal/` só podem ser importados por código no mesmo módulo!

---

## 🔍 Tópicos Principais

1. **Criar módulo** (go mod init)
2. **Estrutura de pacotes**
3. **Exportar símbolos** (maiúscula/minúscula)
4. **Imports e aliases**
5. **Dependências externas**
6. **internal/ packages**
7. **Versioning**

---

**Anterior:** [Módulo 07 - Erros](../modulo07-erros/)  
**Próximo:** [Módulo 09 - Testes](../modulo09-testes/)
