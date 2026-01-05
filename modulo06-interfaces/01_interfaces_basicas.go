package main

import (
	"fmt"
	"math"
)

/*
INTERFACES BÁSICAS

COMPARAÇÃO COM PYTHON:
--------------------
Python (ABC):
    from abc import ABC, abstractmethod

    class Forma(ABC):
        @abstractmethod
        def area(self):
            pass

    class Circulo(Forma):  # Herda explicitamente
        def area(self):
            return math.pi * self.raio ** 2

Go (Interfaces):
    type Forma interface {
        Area() float64
    }

    type Circulo struct { Raio float64 }

    func (c Circulo) Area() float64 {
        return math.Pi * c.Raio * c.Raio
    }
    // Circulo implementa Forma IMPLICITAMENTE!

DIFERENÇA CHAVE: Go não requer herança explícita!
*/

// Interface: define um contrato (conjunto de métodos)
type Forma interface {
	Area() float64
	Perimetro() float64
}

// Retângulo
type Retangulo struct {
	Largura, Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

// Círculo
type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Triângulo
type Triangulo struct {
	Base, Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	// Simplificado - assumindo triângulo isósceles
	lado := math.Sqrt(math.Pow(t.Base/2, 2) + math.Pow(t.Altura, 2))
	return t.Base + 2*lado
}

// Função que aceita qualquer Forma
func imprimirInfo(f Forma) {
	fmt.Printf("Área: %.2f\n", f.Area())
	fmt.Printf("Perímetro: %.2f\n", f.Perimetro())
}

// Função para calcular área total
func areaTotal(formas ...Forma) float64 {
	total := 0.0
	for _, forma := range formas {
		total += forma.Area()
	}
	return total
}

func main() {
	fmt.Println("=== INTERFACES BÁSICAS ===\n")

	// Criar formas
	ret := Retangulo{Largura: 10, Altura: 5}
	circ := Circulo{Raio: 7}
	tri := Triangulo{Base: 6, Altura: 4}

	fmt.Println("Retângulo:")
	imprimirInfo(ret)

	fmt.Println("\nCírculo:")
	imprimirInfo(circ)

	fmt.Println("\nTriângulo:")
	imprimirInfo(tri)

	fmt.Println("\n=== POLIMORFISMO ===\n")

	// Slice de interfaces - pode conter qualquer Forma
	formas := []Forma{ret, circ, tri}

	fmt.Println("Todas as formas:")
	for i, forma := range formas {
		fmt.Printf("\nForma %d:\n", i+1)
		imprimirInfo(forma)
	}

	fmt.Printf("\nÁrea total de todas as formas: %.2f\n", areaTotal(formas...))

	fmt.Println("\n=== EXEMPLO PRÁTICO: ANIMAIS ===\n")

	exemploAnimais()
}

// ========================================
// EXEMPLO: Interface de Animais
// ========================================

type Animal interface {
	FazerSom() string
	Mover() string
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) FazerSom() string {
	return "Au au!"
}

func (c Cachorro) Mover() string {
	return "Correndo"
}

type Gato struct {
	Nome string
}

func (g Gato) FazerSom() string {
	return "Miau!"
}

func (g Gato) Mover() string {
	return "Pulando"
}

type Pato struct {
	Nome string
}

func (p Pato) FazerSom() string {
	return "Quack!"
}

func (p Pato) Mover() string {
	return "Nadando"
}

func fazerBarulho(a Animal) {
	fmt.Printf("%s está %s e fazendo: %s\n",
		getNome(a), a.Mover(), a.FazerSom())
}

func getNome(a Animal) string {
	// Type assertion para pegar o nome
	switch v := a.(type) {
	case Cachorro:
		return v.Nome
	case Gato:
		return v.Nome
	case Pato:
		return v.Nome
	default:
		return "Animal"
	}
}

func exemploAnimais() {
	animais := []Animal{
		Cachorro{Nome: "Rex"},
		Gato{Nome: "Mimi"},
		Pato{Nome: "Donald"},
	}

	for _, animal := range animais {
		fazerBarulho(animal)
	}
}

/*
RESUMO DE INTERFACES:

DEFINIR:
    type NomeInterface interface {
        Metodo1() tipoRetorno
        Metodo2(param tipo) tipoRetorno
    }

IMPLEMENTAR:
    - Basta ter os métodos com assinatura correta
    - Não precisa declarar explicitamente
    - Implementação IMPLÍCITA

USAR:
    func funcao(i MinhaInterface) {
        i.Metodo1()
    }

VANTAGENS:
✓ Desacoplamento
✓ Testabilidade
✓ Flexibilidade
✓ Composição sobre herança

DIFERENÇAS COM PYTHON:
- Python: Herança explícita (ABC)
- Go: Implementação implícita
- Python: Duck typing em runtime
- Go: Type checking em compile time

CONVENÇÕES:
- Interfaces pequenas (1-2 métodos)
- Nome com sufixo "er" para 1 método (Reader, Writer)
- Definir onde usado, não onde implementado

Execute com:
    go run 01_interfaces_basicas.go
*/
