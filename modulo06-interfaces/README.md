# 📘 Módulo 06 - Interfaces

## 🎯 Objetivos
Dominar interfaces em Go - um dos conceitos mais importantes e diferentes do Python!

## 🔑 Conceito Principal

### Python (Classes Abstratas)
```python
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def fazer_som(self):
        pass

class Cachorro(Animal):
    def fazer_som(self):  # DEVE implementar
        return "Au au!"
```

### Go (Interfaces Implícitas)
```go
// Interface
type Animal interface {
    FazerSom() string
}

// Struct - NÃO precisa declarar que implementa!
type Cachorro struct{}

func (c Cachorro) FazerSom() string {
    return "Au au!"
}
// Cachorro implementa Animal automaticamente!
```

**🌟 Diferença Fundamental:**
- Python: Herança explícita (`class Cachorro(Animal)`)
- Go: **Implementação implícita** - se tem os métodos, implementa!

---

## 📝 Tópicos

1. **Interfaces básicas**
2. **Interfaces implícitas**
3. **Interface vazia (interface{})**
4. **Type assertions**
5. **Type switches**
6. **Interfaces comuns (io.Reader, io.Writer)**
7. **Composição de interfaces**

---

## Por que Interfaces são Importantes?

1. **Polimorfismo** - mesma interface, comportamentos diferentes
2. **Desacoplamento** - código flexível e testável
3. **Composição** - preferida sobre herança em Go
4. **Contratos** - garantir que tipos têm certos métodos

---

## Interface Vazia

```go
var x interface{}  // ou any (Go 1.18+)
x = 42
x = "texto"
x = true
// Aceita QUALQUER tipo!
```

**Similar a:** `Any` em Python type hints

---

## Interfaces Comuns

- `io.Reader` - ler dados
- `io.Writer` - escrever dados
- `error` - erros
- `fmt.Stringer` - conversão para string
- `sort.Interface` - ordenação

---

**Anterior:** [Módulo 05 - Goroutines](../modulo05-goroutines/)  
**Próximo:** [Módulo 07 - Tratamento de Erros](../modulo07-erros/)
