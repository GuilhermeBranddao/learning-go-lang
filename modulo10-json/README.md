# 📦 Módulo 10 - JSON e Serialização

## 🎯 Objetivos
Trabalhar com JSON em Go - codificação e decodificação de dados.

## 🔑 Comparação com Python

| Operação | Python | Go |
|----------|--------|-----|
| **Package** | `json` | `encoding/json` |
| **Para JSON** | `json.dumps(obj)` | `json.Marshal(obj)` |
| **De JSON** | `json.loads(str)` | `json.Unmarshal(data, &obj)` |
| **Bonito** | `json.dumps(indent=2)` | `json.MarshalIndent(obj, "", "  ")` |

---

## 📝 Básico

```go
import "encoding/json"

// Struct para JSON
type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
}

// Para JSON (Marshal)
pessoa := Pessoa{Nome: "João", Idade: 30}
jsonData, err := json.Marshal(pessoa)

// De JSON (Unmarshal)
var p Pessoa
err := json.Unmarshal(jsonData, &p)
```

---

## 🏷️ Struct Tags

```go
type Usuario struct {
    ID        int    `json:"id"`
    Nome      string `json:"nome"`
    Email     string `json:"email"`
    Senha     string `json:"-"`              // Ignorar
    Admin     bool   `json:"admin,omitempty"` // Omitir se vazio
    CreatedAt string `json:"created_at"`
}
```

### Tags Comuns
- `json:"nome"` - Nome do campo no JSON
- `json:"-"` - Ignorar campo
- `json:",omitempty"` - Omitir se valor zero
- `json:",string"` - Forçar como string

---

## 🎨 Formatação

```go
// Compacto
data, _ := json.Marshal(obj)

// Bonito (indentado)
data, _ := json.MarshalIndent(obj, "", "  ")

// Encoder/Decoder (para streams)
encoder := json.NewEncoder(writer)
encoder.Encode(obj)

decoder := json.NewDecoder(reader)
decoder.Decode(&obj)
```

---

## 📋 Tópicos Principais

1. **Marshal/Unmarshal**
2. **Struct tags**
3. **Encoder/Decoder**
4. **JSON para maps**
5. **Campos aninhados**
6. **Arrays e slices**
7. **Validação**

---

**Anterior:** [Módulo 09 - Testes](../modulo09-testes/)  
**Próximo:** [Módulo 11 - Arquivos](../modulo11-arquivos/)
