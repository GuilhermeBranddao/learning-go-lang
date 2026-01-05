# 🚀 Módulo 15 - Tópicos Avançados

## 🎯 Objetivos
Recursos avançados de Go: Reflection e Generics.

## 🔑 Generics (Go 1.18+)

```go
// Função genérica
func Maximo[T int | float64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Type constraint
type Number interface {
    int | int64 | float64
}

func Somar[T Number](a, b T) T {
    return a + b
}
```

---

## 🪞 Reflection

```go
import "reflect"

// Obter tipo
t := reflect.TypeOf(obj)

// Obter valor
v := reflect.ValueOf(obj)

// Modificar valor (se ponteiro)
v.Elem().SetInt(42)
```

---

## 📋 Tópicos

1. **Generics** (tipo parametrizado)
2. **Type constraints**
3. **Reflection**
4. **Type assertions**
5. **Unsafe**

---

**Anterior:** [Módulo 14 - Context](../modulo14-context/)  
**Início:** [README Principal](../README.md)
