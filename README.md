# 🎓 Aprendendo Go - De Python para Go

Bem-vindo ao curso de Go para desenvolvedores Python! Este curso foi estruturado para aproveitar seu conhecimento em Python e facilitar sua transição para Go.

## 📚 Estrutura do Curso

### [Módulo 01 - Fundamentos](modulo01-fundamentos/)
- Variáveis e tipos de dados
- Operadores
- Entrada e saída
- **Comparação com Python**

### [Módulo 02 - Estruturas de Controle](modulo02-controle/)
- If/Else
- Loops (for)
- Switch
- **Diferenças importantes vs Python**

### [Módulo 03 - Funções](modulo03-funcoes/)
- Declaração de funções
- Múltiplos retornos
- Funções variádicas
- Defer, Panic e Recover

### [Módulo 04 - Estruturas de Dados](modulo04-estruturas-dados/)
- Arrays e Slices
- Maps
- Structs
- Ponteiros

### [Módulo 05 - Goroutines e Concorrência](modulo05-goroutines/)
- Goroutines
- Channels
- Select
- **A grande vantagem do Go!**

### [Módulo 06 - Interfaces](modulo06-interfaces/)
- Interfaces básicas
- Interface vazia (any)
- Interfaces comuns (io.Reader, io.Writer)
- Type assertions

### [Módulo 07 - Tratamento de Erros](modulo07-erros/)
- Error interface
- Erros customizados
- Panic e Recover
- Error wrapping

### [Módulo 08 - Packages e Módulos](modulo08-packages/)
- Estrutura de packages
- go.mod e go.sum
- Visibilidade
- Imports

### [Módulo 09 - Testes](modulo09-testes/)
- testing package
- Table-driven tests
- Benchmarks
- Cobertura

### [Módulo 10 - JSON](modulo10-json/)
- Marshal/Unmarshal
- Struct tags
- Encoder/Decoder

### [Módulo 11 - Arquivos e I/O](modulo11-arquivos/)
- Ler/escrever arquivos
- bufio
- Diretórios

### [Módulo 12 - HTTP e APIs REST](modulo12-http/)
- Servidor HTTP
- HTTP Client
- JSON APIs

### [Módulo 13 - Banco de Dados](modulo13-database/)
- database/sql
- CRUD operations
- Prepared statements

### [Módulo 14 - Context](modulo14-context/)
- Cancelamento
- Timeouts
- Deadlines

### [Módulo 15 - Tópicos Avançados](modulo15-avancado/)
- Generics (Go 1.18+)
- Reflection
- Type constraints

## 🎯 Como Usar Este Curso

1. **Leia a teoria** em cada arquivo README.md
2. **Execute os exemplos** com `go run`
3. **Faça os exercícios** na pasta exercicios/
4. **Compare com Python** - todos os módulos têm comparações

## 🔧 Pré-requisitos

- Go instalado (veja [GUIA_INSTALACAO.md](GUIA_INSTALACAO.md))
- Conhecimento básico de Python
- Vontade de aprender! 🚀

## 💡 Principais Diferenças: Python vs Go

| Característica | Python | Go |
|----------------|--------|-----|
| **Tipagem** | Dinâmica | Estática |
| **Compilação** | Interpretado | Compilado |
| **Concorrência** | Threading/asyncio | Goroutines nativas |
| **Performance** | Mais lenta | Muito rápida |
| **Sintaxe** | Indentação | Chaves {} |
| **Gerenciamento de memória** | Garbage Collector | Garbage Collector |

## 📝 Convenções do Go

- **Formatação**: Use `go fmt` - é obrigatório!
- **Nomenclatura**: CamelCase (não snake_case)
- **Exportação**: Letras maiúsculas são públicas
- **Simplicidade**: Go valoriza código simples e direto

---

**Vamos começar! Abra o [Módulo 01](modulo01-fundamentos/) 👉**