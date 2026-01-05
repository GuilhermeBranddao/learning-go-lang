# 📘 Módulo 03 - Funções

## 🎯 Objetivos
Dominar funções em Go, incluindo recursos exclusivos da linguagem.

## 📝 Tópicos

1. **Declaração de funções**
2. **Múltiplos retornos**
3. **Funções variádicas**
4. **Funções anônimas e closures**
5. **Defer, Panic e Recover**

---

## 1. Declaração Básica

### Python
```python
def somar(a, b):
    return a + b

resultado = somar(10, 5)
```

### Go
```go
func somar(a int, b int) int {
    return a + b
}

// Ou com tipos agrupados
func somar(a, b int) int {
    return a + b
}

resultado := somar(10, 5)
```

**Diferenças:**
- Tipos vêm DEPOIS do nome
- Tipo de retorno no final
- Use `func` para declarar

---

## 2. Múltiplos Retornos (🌟 Recurso Poderoso!)

### Python
```python
def dividir(a, b):
    return a / b, a % b

quociente, resto = dividir(10, 3)
```

### Go
```go
func dividir(a, b int) (int, int) {
    return a / b, a % b
}

quociente, resto := dividir(10, 3)

// Ignorar retorno
resultado, _ := dividir(10, 3)
```

**Padrão em Go: retornar valor e erro**
```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
```

---

## 3. Named Return Values

```go
func calcular(a, b int) (soma int, mult int) {
    soma = a + b
    mult = a * b
    return  // naked return
}
```

---

## 4. Funções Variádicas

### Python
```python
def somar(*numeros):
    return sum(numeros)

total = somar(1, 2, 3, 4, 5)
```

### Go
```go
func somar(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

total := somar(1, 2, 3, 4, 5)
```

---

## 5. Defer

Adia a execução até o fim da função (muito útil para cleanup!):

```go
func exemplo() {
    defer fmt.Println("Executado por último")
    fmt.Println("Executado primeiro")
}
```

**Uso comum: fechar arquivos**
```go
file, err := os.Open("arquivo.txt")
if err != nil {
    return
}
defer file.Close()  // Garante que o arquivo será fechado
```

---

**Próximo:** [Módulo 04 - Estruturas de Dados](../modulo04-estruturas-dados/)
