package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
JSON EM GO

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    import json

    # Para JSON
    json_str = json.dumps({"nome": "João", "idade": 30})

    # De JSON
    obj = json.loads(json_str)

Go:
    import "encoding/json"

    // Para JSON (Marshal)
    type Pessoa struct {
        Nome  string `json:"nome"`
        Idade int    `json:"idade"`
    }
    jsonData, _ := json.Marshal(pessoa)

    // De JSON (Unmarshal)
    var p Pessoa
    json.Unmarshal(jsonData, &p)

STRUCT TAGS:
    `json:"nome_no_json"`
    `json:"-"`              // ignorar
    `json:",omitempty"`     // omitir se vazio
*/

// ========================================
// STRUCT COM TAGS
// ========================================

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Email string `json:"email"`
}

type Usuario struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Senha    string `json:"-"`               // Nunca serializar
	Admin    bool   `json:"admin,omitempty"` // Omitir se false
	Ativo    bool   `json:"ativo"`
}

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Preco     float64 `json:"preco"`
	Estoque   int     `json:"estoque,omitempty"`
	Descricao string  `json:"descricao,omitempty"`
}

func main() {
	fmt.Println("=== MARSHAL (GO → JSON) ===\n")

	// Exemplo básico
	p1 := Pessoa{
		Nome:  "João Silva",
		Idade: 30,
		Email: "joao@email.com",
	}

	// Marshal: struct → JSON
	jsonData, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON compacto:")
	fmt.Println(string(jsonData))

	// Marshal com indentação
	jsonPretty, err := json.MarshalIndent(p1, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nJSON bonito:")
	fmt.Println(string(jsonPretty))

	fmt.Println("\n=== UNMARSHAL (JSON → GO) ===\n")

	// JSON como string
	jsonString := `{"nome":"Maria","idade":25,"email":"maria@email.com"}`

	// Unmarshal: JSON → struct
	var p2 Pessoa
	err = json.Unmarshal([]byte(jsonString), &p2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Struct a partir de JSON:")
	fmt.Printf("Nome: %s, Idade: %d, Email: %s\n", p2.Nome, p2.Idade, p2.Email)

	fmt.Println("\n=== STRUCT TAGS ===\n")

	// Usuário com campo ignorado (senha)
	u1 := Usuario{
		ID:       1,
		Username: "joao123",
		Email:    "joao@email.com",
		Senha:    "senha-secreta", // Não aparece no JSON
		Admin:    false,           // Omitido (omitempty)
		Ativo:    true,
	}

	jsonUser, _ := json.MarshalIndent(u1, "", "  ")
	fmt.Println("Usuário (senha ignorada, admin omitido):")
	fmt.Println(string(jsonUser))

	fmt.Println("\n=== SLICES E ARRAYS ===\n")

	pessoas := []Pessoa{
		{Nome: "João", Idade: 30, Email: "joao@email.com"},
		{Nome: "Maria", Idade: 25, Email: "maria@email.com"},
		{Nome: "Pedro", Idade: 35, Email: "pedro@email.com"},
	}

	jsonPessoas, _ := json.MarshalIndent(pessoas, "", "  ")
	fmt.Println("Array de pessoas:")
	fmt.Println(string(jsonPessoas))

	// Unmarshal de array
	jsonArray := `[
		{"nome":"Ana","idade":28,"email":"ana@email.com"},
		{"nome":"Carlos","idade":32,"email":"carlos@email.com"}
	]`

	var pessoasArray []Pessoa
	json.Unmarshal([]byte(jsonArray), &pessoasArray)

	fmt.Println("\nPessoas do JSON:")
	for i, p := range pessoasArray {
		fmt.Printf("%d. %s (%d anos)\n", i+1, p.Nome, p.Idade)
	}

	fmt.Println("\n=== MAPS ===\n")

	// JSON para map (quando não conhece estrutura)
	jsonMap := `{
		"nome": "Produto X",
		"preco": 99.90,
		"disponivel": true,
		"quantidade": 42
	}`

	var dados map[string]interface{}
	json.Unmarshal([]byte(jsonMap), &dados)

	fmt.Println("Dados do map:")
	for key, value := range dados {
		fmt.Printf("%s: %v (tipo: %T)\n", key, value, value)
	}

	fmt.Println("\n=== STRUCTS ANINHADAS ===\n")

	type Endereco struct {
		Rua    string `json:"rua"`
		Numero int    `json:"numero"`
		Cidade string `json:"cidade"`
	}

	type Cliente struct {
		Nome     string   `json:"nome"`
		Email    string   `json:"email"`
		Endereco Endereco `json:"endereco"`
	}

	cliente := Cliente{
		Nome:  "João",
		Email: "joao@email.com",
		Endereco: Endereco{
			Rua:    "Rua A",
			Numero: 123,
			Cidade: "São Paulo",
		},
	}

	jsonCliente, _ := json.MarshalIndent(cliente, "", "  ")
	fmt.Println("Cliente com endereço:")
	fmt.Println(string(jsonCliente))

	fmt.Println("\n=== EXEMPLO: PRODUTOS ===\n")

	produtos := []Produto{
		{ID: 1, Nome: "Notebook", Preco: 2500.00, Estoque: 10},
		{ID: 2, Nome: "Mouse", Preco: 50.00}, // Estoque omitido
		{ID: 3, Nome: "Teclado", Preco: 150.00, Estoque: 5, Descricao: "Mecânico RGB"},
	}

	jsonProdutos, _ := json.MarshalIndent(produtos, "", "  ")
	fmt.Println("Lista de produtos:")
	fmt.Println(string(jsonProdutos))
}

/*
RESUMO DE JSON:

MARSHAL (Go → JSON):
    data, err := json.Marshal(obj)
    data, err := json.MarshalIndent(obj, "", "  ")

UNMARSHAL (JSON → Go):
    err := json.Unmarshal(jsonData, &obj)

STRUCT TAGS:
    `json:"nome"`          - Nome do campo
    `json:"-"`             - Ignorar
    `json:",omitempty"`    - Omitir se zero
    `json:",string"`       - Forçar string

TIPOS SUPORTADOS:
    ✓ bool, int, float64, string
    ✓ struct, slice, array
    ✓ map[string]tipo
    ✓ ponteiros
    ✓ interface{}

DIFERENÇAS COM PYTHON:

Python:
✓ Dicionários nativos
✓ json.dumps/loads
✓ Dinâmico
✓ Tipos mesclados fácil

Go:
✓ Structs tipados
✓ Marshal/Unmarshal
✓ Type-safe
✓ Struct tags
✓ Melhor performance

VANTAGENS GO:
✓ Type safety
✓ Validação em compile time
✓ Performance
✓ Documentação via structs

Execute com:
    go run 01_json_basico.go
*/
