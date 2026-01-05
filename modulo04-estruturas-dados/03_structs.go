package main

import "fmt"

/*
STRUCTS (Estruturas)

COMPARAÇÃO COM PYTHON:
--------------------
Python (classes):
    class Pessoa:
        def __init__(self, nome, idade):
            self.nome = nome
            self.idade = idade

        def apresentar(self):
            return f"{self.nome}, {self.idade} anos"

    p = Pessoa("Maria", 25)
    print(p.apresentar())

Go (structs):
    type Pessoa struct {
        Nome  string
        Idade int
    }

    func (p Pessoa) apresentar() string {
        return fmt.Sprintf("%s, %d anos", p.Nome, p.Idade)
    }

    p := Pessoa{Nome: "Maria", Idade: 25}
    fmt.Println(p.apresentar())
*/

// Definir uma struct
type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

// Struct com campos aninhados
type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome     string
	Email    string
	Endereco Endereco
}

// Struct com tags (metadados)
type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func main() {
	fmt.Println("=== CRIANDO STRUCTS ===\n")

	// Forma 1: zero values
	var p1 Pessoa
	fmt.Printf("Pessoa vazia: %+v\n", p1)

	// Forma 2: com valores posicionais
	p2 := Pessoa{"Maria", 25, "São Paulo"}
	fmt.Printf("Pessoa 2: %+v\n", p2)

	// Forma 3: com nomes de campos (recomendado)
	p3 := Pessoa{
		Nome:   "João",
		Idade:  30,
		Cidade: "Rio de Janeiro",
	}
	fmt.Printf("Pessoa 3: %+v\n\n", p3)

	fmt.Println("=== ACESSANDO CAMPOS ===\n")

	fmt.Printf("Nome: %s\n", p3.Nome)
	fmt.Printf("Idade: %d\n", p3.Idade)
	fmt.Printf("Cidade: %s\n\n", p3.Cidade)

	fmt.Println("=== MODIFICANDO CAMPOS ===\n")

	p3.Idade = 31
	p3.Cidade = "Brasília"
	fmt.Printf("Pessoa modificada: %+v\n\n", p3)

	fmt.Println("=== STRUCTS ANINHADAS ===\n")

	cliente := Cliente{
		Nome:  "Pedro Silva",
		Email: "pedro@email.com",
		Endereco: Endereco{
			Rua:    "Av. Paulista",
			Numero: 1000,
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	fmt.Printf("Cliente: %+v\n", cliente)
	fmt.Printf("Endereço: %s, %d - %s/%s\n\n",
		cliente.Endereco.Rua,
		cliente.Endereco.Numero,
		cliente.Endereco.Cidade,
		cliente.Endereco.Estado,
	)

	fmt.Println("=== PONTEIROS PARA STRUCTS ===\n")

	// Criar ponteiro
	p4 := &Pessoa{Nome: "Ana", Idade: 28, Cidade: "Curitiba"}
	fmt.Printf("Ponteiro: %+v\n", p4)

	// Go permite acessar campos sem desreferenciar
	fmt.Printf("Nome: %s\n", p4.Nome) // poderia ser (*p4).Nome
	p4.Idade = 29
	fmt.Printf("Idade modificada: %d\n\n", p4.Idade)

	fmt.Println("=== COMPARAÇÃO DE STRUCTS ===\n")

	pessoa1 := Pessoa{Nome: "João", Idade: 25, Cidade: "SP"}
	pessoa2 := Pessoa{Nome: "João", Idade: 25, Cidade: "SP"}
	pessoa3 := Pessoa{Nome: "Maria", Idade: 30, Cidade: "RJ"}

	fmt.Printf("pessoa1 == pessoa2: %t\n", pessoa1 == pessoa2)
	fmt.Printf("pessoa1 == pessoa3: %t\n\n", pessoa1 == pessoa3)

	fmt.Println("=== STRUCTS ANÔNIMAS ===\n")

	// Struct sem nome (uso pontual)
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	fmt.Printf("Configuração: %+v\n", config)
	fmt.Printf("Servidor: %s:%d\n\n", config.Host, config.Port)

	fmt.Println("=== MÉTODOS ===\n")

	pessoa := Pessoa{Nome: "Carlos", Idade: 35, Cidade: "Salvador"}

	// Chamar métodos
	pessoa.apresentar()
	pessoa.fazerAniversario()
	pessoa.apresentar()

	fmt.Println("\n=== MÉTODOS COM PONTEIROS ===\n")

	conta := ContaBancaria{Titular: "Maria", Saldo: 1000.0}
	conta.exibirSaldo()

	conta.depositar(500.0)
	conta.exibirSaldo()

	conta.sacar(200.0)
	conta.exibirSaldo()

	conta.sacar(2000.0) // Tentativa de saque maior que saldo
	conta.exibirSaldo()

	fmt.Println("\n=== SLICE DE STRUCTS ===\n")

	pessoas := []Pessoa{
		{Nome: "João", Idade: 25, Cidade: "SP"},
		{Nome: "Maria", Idade: 30, Cidade: "RJ"},
		{Nome: "Pedro", Idade: 28, Cidade: "MG"},
	}

	fmt.Println("Lista de pessoas:")
	for i, p := range pessoas {
		fmt.Printf("  [%d] %s, %d anos, de %s\n", i, p.Nome, p.Idade, p.Cidade)
	}

	fmt.Println("\n=== MAP DE STRUCTS ===\n")

	usuarios := map[string]Pessoa{
		"user1": {Nome: "Ana", Idade: 22, Cidade: "Fortaleza"},
		"user2": {Nome: "Bruno", Idade: 27, Cidade: "Recife"},
		"user3": {Nome: "Carla", Idade: 24, Cidade: "Manaus"},
	}

	fmt.Println("Usuários:")
	for id, usuario := range usuarios {
		fmt.Printf("  %s: %s (%d anos)\n", id, usuario.Nome, usuario.Idade)
	}
}

// Método com receiver (como método de classe em Python)
func (p Pessoa) apresentar() {
	fmt.Printf("Olá, meu nome é %s, tenho %d anos e moro em %s\n",
		p.Nome, p.Idade, p.Cidade)
}

// Método que NÃO modifica o original (receiver por valor)
func (p Pessoa) fazerAniversario() {
	p.Idade++
	fmt.Printf("Feliz aniversário! Agora tenho %d anos (dentro do método)\n", p.Idade)
	// p.Idade volta ao valor original fora do método!
}

// Struct para conta bancária
type ContaBancaria struct {
	Titular string
	Saldo   float64
}

// Método com ponteiro (MODIFICA o original)
func (c *ContaBancaria) depositar(valor float64) {
	c.Saldo += valor
	fmt.Printf("Depósito de R$ %.2f realizado\n", valor)
}

func (c *ContaBancaria) sacar(valor float64) {
	if valor > c.Saldo {
		fmt.Printf("Saldo insuficiente para sacar R$ %.2f\n", valor)
		return
	}
	c.Saldo -= valor
	fmt.Printf("Saque de R$ %.2f realizado\n", valor)
}

func (c ContaBancaria) exibirSaldo() {
	fmt.Printf("Saldo de %s: R$ %.2f\n", c.Titular, c.Saldo)
}

/*
RESUMO DE STRUCTS:

DEFINIR:
    type Nome struct {
        Campo1 tipo1
        Campo2 tipo2
    }

CRIAR:
    var s Nome                      // zero values
    s := Nome{val1, val2}           // posicional
    s := Nome{Campo1: val1}         // com nomes (recomendado)
    s := &Nome{Campo1: val1}        // ponteiro

MÉTODOS:
    // Receiver por valor (não modifica)
    func (s Nome) metodo() {
        // s.Campo
    }

    // Receiver por ponteiro (modifica)
    func (s *Nome) metodo() {
        // s.Campo
    }

DIFERENÇAS COM PYTHON:
- Struct ≈ dataclass ou NamedTuple
- Métodos usam receivers em vez de self
- Não tem __init__ (usa struct literals)
- Não tem herança (usa composição)
- Campos com letra maiúscula são públicos

Execute com:
    go run 03_structs.go
*/
