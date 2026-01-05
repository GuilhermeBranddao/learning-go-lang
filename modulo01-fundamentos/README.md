# 📘 Módulo 01 - Fundamentos do Go

## 🎯 Objetivos
Aprender os conceitos básicos do Go comparando com Python.

## 📝 Tópicos

1. **Hello World e estrutura básica**
2. **Variáveis e constantes**
3. **Tipos de dados**
4. **Operadores**
5. **Entrada e saída**

---

## 1. Hello World

### Python
```python
print("Hello, World!")
```

### Go
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Diferenças importantes:**
- Todo programa Go precisa de um `package` (geralmente `main`)
- Precisa importar pacotes explicitamente
- A função `main()` é o ponto de entrada
- Use chaves `{}` em vez de indentação

---

## 2. Variáveis

### Python
```python
# Tipagem dinâmica
nome = "Maria"
idade = 25
altura = 1.75
ativo = True
```

### Go
```go
// Forma 1: Declaração completa (com tipo)
var nome string = "Maria"
var idade int = 25
var altura float64 = 1.75
var ativo bool = true

// Forma 2: Inferência de tipo
var nome = "Maria"
var idade = 25

// Forma 3: Short declaration (mais comum)
nome := "Maria"
idade := 25
```

**Diferenças importantes:**
- Go é **estaticamente tipado**
- Use `:=` dentro de funções (mais comum)
- Use `var` fora de funções ou quando precisa especificar tipo
- Variáveis não usadas causam erro de compilação!

---

## 3. Tipos de Dados

### Tipos Numéricos

| Python | Go | Descrição |
|--------|-----|-----------|
| `int` | `int`, `int8`, `int16`, `int32`, `int64` | Inteiros |
| `float` | `float32`, `float64` | Ponto flutuante |
| - | `uint`, `uint8`, `uint16`, `uint32`, `uint64` | Inteiros sem sinal |
| `complex` | `complex64`, `complex128` | Números complexos |

### Outros Tipos

| Python | Go | Exemplo |
|--------|-----|---------|
| `str` | `string` | `"texto"` |
| `bool` | `bool` | `true`, `false` |
| `bytes` | `[]byte` | `[]byte("texto")` |
| - | `rune` | `'A'` (caractere Unicode) |

**Importante:**
- Em Go, use `'` para rune (caractere) e `"` para string
- `true` e `false` são minúsculos (não `True`/`False`)

---

## 4. Constantes

### Python
```python
PI = 3.14159  # Convenção (não é realmente constante)
```

### Go
```go
const PI = 3.14159  // Constante real
const (
    StatusOk = 200
    StatusNotFound = 404
)
```

---

## 5. Zero Values (Valores Padrão)

Em Python, variáveis precisam ser inicializadas. Em Go, todas têm um valor padrão:

```go
var numero int       // 0
var texto string     // ""
var ativo bool       // false
var ponteiro *int    // nil
```

---

## 6. Conversão de Tipos (Type Casting)

### Python
```python
idade = int("25")
texto = str(100)
```

### Go
```go
idade := int("25")        // ERRO! Não funciona assim
idade, _ := strconv.Atoi("25")  // Correto
texto := strconv.Itoa(100)
```

Em Go, a conversão é mais explícita e segura.

---

## 7. Operadores

### Aritméticos
```go
a := 10
b := 3

soma := a + b        // 13
sub := a - b         // 7
mult := a * b        // 30
div := a / b         // 3 (divisão inteira!)
mod := a % b         // 1
```

### Comparação
```go
a == b   // igual
a != b   // diferente
a > b    // maior
a < b    // menor
a >= b   // maior ou igual
a <= b   // menor ou igual
```

### Lógicos
```go
true && false   // AND (e)
true || false   // OR (ou)
!true          // NOT (não)
```

---

## 8. Entrada e Saída

### Python
```python
nome = input("Digite seu nome: ")
print(f"Olá, {nome}!")
```

### Go
```go
var nome string
fmt.Print("Digite seu nome: ")
fmt.Scan(&nome)
fmt.Printf("Olá, %s!\n", nome)

// Ou com formatação moderna
fmt.Println("Olá,", nome, "!")
```

**Diferenças:**
- `fmt.Scan()` precisa do `&` (ponteiro)
- `fmt.Printf()` usa placeholders: `%s` (string), `%d` (int), `%f` (float)

---

## 🎯 Exercícios

Veja os exercícios práticos nos arquivos deste módulo!

**Próximo:** [Módulo 02 - Estruturas de Controle](../modulo02-controle/)
