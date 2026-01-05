# 📘 Módulo 02 - Estruturas de Controle

## 🎯 Objetivos
Aprender as estruturas de controle do Go: if/else, loops e switch.

## 📝 Tópicos

1. **If/Else**
2. **For (o único loop do Go!)**
3. **Switch**
4. **Defer**

---

## 1. If/Else

### Python
```python
if idade >= 18:
    print("Maior de idade")
elif idade >= 13:
    print("Adolescente")
else:
    print("Criança")
```

### Go
```go
if idade >= 18 {
    fmt.Println("Maior de idade")
} else if idade >= 13 {
    fmt.Println("Adolescente")
} else {
    fmt.Println("Criança")
}
```

**Diferenças:**
- SEMPRE use chaves `{}`
- NÃO use parênteses `()` na condição (opcional, mas não idiomático)
- Pode inicializar variável no if: `if x := getValue(); x > 0 {}`

---

## 2. For - O Único Loop do Go!

Go tem apenas `for`, mas ele funciona de várias maneiras:

### Python
```python
# Loop básico
for i in range(5):
    print(i)

# While
while condicao:
    # código

# Infinito
while True:
    # código
```

### Go
```go
// For tradicional
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// "While"
for condicao {
    // código
}

// Infinito
for {
    // código
    break
}

// Range (similar ao for...in do Python)
nums := []int{1, 2, 3, 4, 5}
for index, valor := range nums {
    fmt.Printf("Index: %d, Valor: %d\n", index, valor)
}
```

---

## 3. Switch

Muito mais poderoso que em Python!

### Python
```python
# Python usa if/elif (não tem switch nativo até 3.10)
match status:  # Python 3.10+
    case 200:
        print("OK")
    case 404:
        print("Not Found")
    case _:
        print("Outro")
```

### Go
```go
switch status {
case 200:
    fmt.Println("OK")
case 404:
    fmt.Println("Not Found")
default:
    fmt.Println("Outro")
}

// Switch sem break! (não precisa)
// Switch sem expressão
switch {
case idade < 13:
    fmt.Println("Criança")
case idade < 18:
    fmt.Println("Adolescente")
default:
    fmt.Println("Adulto")
}
```

**Vantagens:**
- NÃO precisa de `break`
- Pode ter múltiplos valores: `case 1, 2, 3:`
- Pode usar expressões nas condições

---

## 4. Break e Continue

Igual ao Python!

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // sai do loop
    }
    if i % 2 == 0 {
        continue  // pula para próxima iteração
    }
    fmt.Println(i)
}
```

---

**Próximo:** [Módulo 03 - Funções](../modulo03-funcoes/)
