# ⏱️ Módulo 14 - Context

## 🎯 Objetivos
Entender e usar `context.Context` para controle de execução.

## 🔑 O que é Context?

Context é usado para:
- **Cancelamento** de operações
- **Timeouts**
- **Passar valores** entre goroutines
- **Deadlines**

---

## 📝 Básico

```go
import "context"

// Context com timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Context com cancelamento
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Context com deadline
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

---

## 🎯 Uso Comum

```go
func operacaoLonga(ctx context.Context) error {
    select {
    case <-time.After(2 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()  // Cancelado ou timeout
    }
}
```

---

**Anterior:** [Módulo 13 - Database](../modulo13-database/)  
**Próximo:** [Módulo 15 - Avançado](../modulo15-avancado/)
