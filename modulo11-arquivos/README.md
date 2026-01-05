# 📂 Módulo 11 - Arquivos e I/O

## 🎯 Objetivos
Trabalhar com arquivos e sistema de arquivos em Go.

## 🔑 Comparação com Python

| Operação | Python | Go |
|----------|--------|-----|
| **Ler arquivo** | `open('file.txt')` | `os.ReadFile('file.txt')` |
| **Escrever** | `f.write()` | `os.WriteFile()` |
| **Criar** | `open('f', 'w')` | `os.Create()` |
| **Package** | `os, io` | `os, io, ioutil` |

---

## 📝 Operações Básicas

```go
import "os"

// Ler arquivo inteiro
data, err := os.ReadFile("arquivo.txt")

// Escrever arquivo
err := os.WriteFile("arquivo.txt", []byte("conteúdo"), 0644)

// Criar/abrir arquivo
file, err := os.Create("novo.txt")
defer file.Close()

// Ler linha por linha (bufio)
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    linha := scanner.Text()
}
```

---

## 📋 Tópicos

1. **Ler/escrever arquivos**
2. **bufio (buffered I/O)**
3. **os package**
4. **Caminhos (filepath)**
5. **Diretórios**
6. **Permissões**

---

**Anterior:** [Módulo 10 - JSON](../modulo10-json/)  
**Próximo:** [Módulo 12 - HTTP](../modulo12-http/)
