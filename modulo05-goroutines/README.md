# 📘 Módulo 05 - Goroutines e Concorrência

## 🎯 Objetivos
Dominar concorrência em Go - uma das maiores vantagens da linguagem!

## 🚀 Por que Go é Excelente para Concorrência?

### Python
```python
import threading

def tarefa():
    print("Executando tarefa")

# Threading tem GIL (Global Interpreter Lock)
# Não é paralelismo real em CPython
t = threading.Thread(target=tarefa)
t.start()
```

### Go
```go
func tarefa() {
    fmt.Println("Executando tarefa")
}

// Goroutines são leves e eficientes
// Paralelismo real!
go tarefa()
```

**Vantagens do Go:**
- Goroutines são muito mais leves que threads (kilobytes vs megabytes)
- Go pode rodar milhares de goroutines simultaneamente
- Não tem GIL (Global Interpreter Lock)
- Paralelismo real em múltiplos cores

---

## 📝 Tópicos

1. **Goroutines** - funções executadas concorrentemente
2. **Channels** - comunicação entre goroutines
3. **Select** - multiplexação de channels
4. **WaitGroups** - sincronização
5. **Mutex** - exclusão mútua

---

## 1. Goroutines

```go
// Função normal
funcao()

// Goroutine - executa em paralelo
go funcao()
```

**Simples assim!** Adicione `go` antes de qualquer chamada de função.

---

## 2. Channels

### Python (usando queue)
```python
from queue import Queue

q = Queue()
q.put(10)
valor = q.get()
```

### Go (channels nativos)
```go
// Criar channel
ch := make(chan int)

// Enviar valor
ch <- 10

// Receber valor
valor := <-ch
```

**Tipos de channels:**
- `make(chan tipo)` - unbuffered (bloqueante)
- `make(chan tipo, N)` - buffered (com capacidade)

---

## 3. Select

Permite esperar por múltiplos channels:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Recebido de ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Recebido de ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}
```

---

## 4. WaitGroup

Espera um grupo de goroutines terminar:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println("Tarefa", n)
    }(i)
}

wg.Wait() // Espera todas terminarem
```

---

## 5. Mutex

Proteção de dados compartilhados:

```go
var (
    contador int
    mu       sync.Mutex
)

mu.Lock()
contador++
mu.Unlock()
```

---

## 🎯 Padrões Comuns

### Worker Pool
```go
jobs := make(chan int, 100)
results := make(chan int, 100)

// Workers
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}

// Enviar jobs
for j := 1; j <= 5; j++ {
    jobs <- j
}
close(jobs)
```

### Pipeline
```go
c1 := gerador(nums)
c2 := quadrado(c1)
c3 := imprimir(c2)
```

---

## ⚠️ Cuidados

1. **Sempre feche channels quando terminar de enviar**
2. **Use WaitGroup para sincronizar goroutines**
3. **Evite race conditions com Mutex**
4. **Não compartilhe memória, comunique por channels**

> "Don't communicate by sharing memory; share memory by communicating."

---

**Anterior:** [Módulo 04 - Estruturas de Dados](../modulo04-estruturas-dados/)
