# 📘 Módulo 07 - Tratamento de Erros

## 🎯 Objetivos
Dominar o tratamento de erros em Go - completamente diferente do Python!

## 🔑 Diferença Fundamental

### Python (Exceções)
```python
try:
    resultado = dividir(10, 0)
    print(resultado)
except ZeroDivisionError as e:
    print(f"Erro: {e}")
except Exception as e:
    print(f"Erro inesperado: {e}")
finally:
    print("Sempre executa")
```

### Go (Valores de Erro)
```go
resultado, err := dividir(10, 0)
if err != nil {
    fmt.Println("Erro:", err)
    return
}
fmt.Println(resultado)
```

**🌟 Filosofia do Go:**
- **Erros são valores**, não exceções
- Tratamento **explícito** (não há try/catch)
- **Impossível ignorar** erros sem intenção
- Código mais **previsível**

---

## 📝 Tópicos

1. **Error interface**
2. **Criando erros**
3. **Erros customizados**
4. **Error wrapping (Go 1.13+)**
5. **Panic e Recover**
6. **Quando usar panic vs error**

---

## Error Interface

```go
type error interface {
    Error() string
}
```

Qualquer tipo com método `Error() string` é um erro!

---

## Criando Erros

```go
// errors.New
err := errors.New("algo deu errado")

// fmt.Errorf (com formatação)
err := fmt.Errorf("valor %d inválido", valor)

// Erro customizado
type MeuErro struct {
    Codigo int
    Msg    string
}

func (e MeuErro) Error() string {
    return fmt.Sprintf("erro %d: %s", e.Codigo, e.Msg)
}
```

---

## Error Wrapping (Go 1.13+)

```go
// Embrulhar erro
err := fmt.Errorf("falha ao processar: %w", erroOriginal)

// Desembrulhar
errors.Unwrap(err)
errors.Is(err, erroOriginal)
errors.As(err, &target)
```

---

## Panic e Recover

```go
// Panic: para erros irrecuperáveis
panic("algo muito errado!")

// Recover: capturar panic
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recuperado:", r)
    }
}()
```

**⚠️ Use panic raramente!** Prefira retornar erros.

---

**Anterior:** [Módulo 06 - Interfaces](../modulo06-interfaces/)  
**Próximo:** [Módulo 08 - Packages](../modulo08-packages/)
