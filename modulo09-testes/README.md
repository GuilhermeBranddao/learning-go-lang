# 🧪 Módulo 09 - Testes

## 🎯 Objetivos
Dominar testes automatizados em Go com o package `testing`.

## 🔑 Comparação com Python

| Aspecto | Python | Go |
|---------|--------|-----|
| **Framework** | pytest, unittest | testing (stdlib) |
| **Arquivo** | `test_*.py` | `*_test.go` |
| **Função** | `def test_*():` | `func Test*(t *testing.T)` |
| **Asserções** | `assert x == y` | `if x != y { t.Error() }` |
| **Executar** | `pytest` | `go test` |

---

## 📝 Estrutura de Teste

```go
// calc_test.go
package matematica

import "testing"

func TestSomar(t *testing.T) {
    resultado := Somar(2, 3)
    esperado := 5
    
    if resultado != esperado {
        t.Errorf("Somar(2, 3) = %d; esperado %d", resultado, esperado)
    }
}
```

---

## 🔧 Comandos

```bash
go test                    # Executar testes
go test -v                 # Verbose
go test -cover             # Cobertura
go test -run TestNome      # Teste específico
go test ./...              # Todos packages
go test -bench .           # Benchmarks
```

---

## 📊 Table-Driven Tests

```go
func TestSomar(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        esperado int
    }{
        {"positivos", 2, 3, 5},
        {"negativos", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resultado := Somar(tt.a, tt.b)
            if resultado != tt.esperado {
                t.Errorf("Somar(%d, %d) = %d; esperado %d",
                    tt.a, tt.b, resultado, tt.esperado)
            }
        })
    }
}
```

---

## 🏆 Benchmarks

```go
func BenchmarkSomar(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Somar(10, 20)
    }
}
```

Executar:
```bash
go test -bench .
go test -bench . -benchmem  # Com estatísticas de memória
```

---

## 📋 Métodos de Testing.T

```go
t.Error()      // Reporta erro, continua
t.Errorf()     // Com formatação
t.Fatal()      // Reporta erro, para teste
t.Fatalf()     // Com formatação
t.Skip()       // Pula teste
t.Skipf()      // Com formatação
t.Run()        // Subteste
t.Parallel()   // Executa em paralelo
```

---

## 🎯 Tópicos Principais

1. **Funções de teste** (Test*)
2. **Table-driven tests**
3. **Subtestes** (t.Run)
4. **Benchmarks** (Benchmark*)
5. **Cobertura de código**
6. **Testes paralelos**
7. **Helper functions**

---

## 💡 Boas Práticas

```go
// ✅ BOM
func TestSomar_PositiveNumbers(t *testing.T) {
    t.Parallel()  // Roda em paralelo
    
    tests := []struct {
        name string
        a, b int
        want int
    }{
        // casos de teste
    }
    
    for _, tt := range tests {
        tt := tt  // Capturar variável
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := Somar(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

---

**Anterior:** [Módulo 08 - Packages](../modulo08-packages/)  
**Próximo:** [Módulo 10 - JSON](../modulo10-json/)
