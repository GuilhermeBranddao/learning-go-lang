# 🌐 Módulo 12 - HTTP e APIs REST

## 🎯 Objetivos
Criar servidores HTTP e consumir APIs REST em Go.

## 🔑 Comparação com Python

| Aspecto | Python (Flask) | Go (net/http) |
|---------|----------------|---------------|
| **Servidor** | `@app.route()` | `http.HandleFunc()` |
| **Handler** | Função decorada | `func(w, r)` |
| **Router** | Flask, FastAPI | Gorilla Mux, Chi |
| **Performance** | WSGI/ASGI | Nativo, muito rápido |

---

## 📝 Servidor Básico

```go
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Olá, Mundo!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

---

## 🔧 HTTP Client

```go
// GET
resp, err := http.Get("https://api.exemplo.com/dados")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST com JSON
jsonData := []byte(`{"nome":"João"}`)
resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
```

---

## 📋 Tópicos

1. **Servidor HTTP**
2. **Handlers e rotas**
3. **HTTP Client (requisições)**
4. **JSON APIs**
5. **Middleware**
6. **Gorilla Mux (router)**

---

**Anterior:** [Módulo 11 - Arquivos](../modulo11-arquivos/)  
**Próximo:** [Módulo 13 - Database](../modulo13-database/)
