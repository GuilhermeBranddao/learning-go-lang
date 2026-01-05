# 📘 Módulo 04 - Estruturas de Dados

## 🎯 Objetivos
Dominar as estruturas de dados fundamentais do Go.

## 📝 Tópicos

1. **Arrays**
2. **Slices (arrays dinâmicos)**
3. **Maps (dicionários)**
4. **Structs (estruturas)**
5. **Ponteiros**

---

## 1. Arrays

### Python
```python
numeros = [1, 2, 3, 4, 5]
```

### Go
```go
// Array tem tamanho FIXO
var numeros [5]int = [5]int{1, 2, 3, 4, 5}

// Ou com inferência de tamanho
numeros := [...]int{1, 2, 3, 4, 5}
```

**⚠️ Arrays em Go têm tamanho fixo!**

---

## 2. Slices (Mais Usados!)

### Python
```python
lista = [1, 2, 3]
lista.append(4)  # Adiciona elemento
```

### Go
```go
// Slice - tamanho dinâmico (como listas Python)
numeros := []int{1, 2, 3}
numeros = append(numeros, 4)  // Adiciona elemento
```

**Operações:**
```go
slice[inicio:fim]  // Fatiar
len(slice)         // Tamanho
cap(slice)         // Capacidade
append(slice, val) // Adicionar
```

---

## 3. Maps

### Python
```python
pessoa = {
    "nome": "Maria",
    "idade": 25
}
```

### Go
```go
pessoa := map[string]interface{}{
    "nome": "Maria",
    "idade": 25,
}

// Ou tipado
pessoa := map[string]string{
    "nome": "Maria",
    "cidade": "SP",
}
```

---

## 4. Structs

### Python
```python
class Pessoa:
    def __init__(self, nome, idade):
        self.nome = nome
        self.idade = idade

p = Pessoa("Maria", 25)
```

### Go
```go
type Pessoa struct {
    Nome  string
    Idade int
}

p := Pessoa{Nome: "Maria", Idade: 25}
```

---

## 5. Ponteiros

Go tem ponteiros explícitos (diferente de Python):

```go
x := 10
p := &x    // p é um ponteiro para x
*p = 20    // modifica x através do ponteiro
```

**Uso comum:** passar grandes estruturas por referência.

---

**Próximo:** [Módulo 05 - Goroutines e Concorrência](../modulo05-goroutines/)
