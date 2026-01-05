package main

import "fmt"

/*
FUNÇÕES ANÔNIMAS E CLOSURES

COMPARAÇÃO COM PYTHON:
--------------------
Python (lambda):
    dobro = lambda x: x * 2
    resultado = dobro(5)

Go (função anônima):
    dobro := func(x int) int {
        return x * 2
    }
    resultado := dobro(5)

Python (closure):
    def contador():
        x = 0
        def incrementar():
            nonlocal x
            x += 1
            return x
        return incrementar

    c = contador()
    print(c())  # 1
    print(c())  # 2

Go (closure):
    func contador() func() int {
        x := 0
        return func() int {
            x++
            return x
        }
    }

    c := contador()
    fmt.Println(c())  // 1
    fmt.Println(c())  // 2
*/

func main() {
	fmt.Println("=== FUNÇÕES ANÔNIMAS BÁSICAS ===\n")

	// Declarar e executar imediatamente
	func() {
		fmt.Println("Função anônima executada!")
	}()

	// Função anônima com parâmetros
	func(nome string) {
		fmt.Printf("Olá, %s!\n", nome)
	}("Maria")

	// Atribuir a uma variável
	dobro := func(x int) int {
		return x * 2
	}
	fmt.Printf("Dobro de 5: %d\n\n", dobro(5))

	fmt.Println("=== CLOSURES ===\n")

	// Closure: função que captura variáveis do escopo externo
	multiplicador := 10
	multiplicar := func(x int) int {
		return x * multiplicador // captura 'multiplicador'
	}

	fmt.Printf("5 * %d = %d\n", multiplicador, multiplicar(5))
	multiplicador = 20
	fmt.Printf("5 * %d = %d\n\n", multiplicador, multiplicar(5))

	fmt.Println("=== CONTADOR COM CLOSURE ===\n")

	// Closure clássico: contador
	contador := criarContador()
	fmt.Println("Contador:", contador()) // 1
	fmt.Println("Contador:", contador()) // 2
	fmt.Println("Contador:", contador()) // 3

	// Cada closure tem sua própria variável
	outroContador := criarContador()
	fmt.Println("Outro contador:", outroContador()) // 1
	fmt.Println("Outro contador:", outroContador()) // 2

	fmt.Println("\n=== FUNÇÃO COMO RETORNO ===\n")

	somar5 := criarSomador(5)
	somar10 := criarSomador(10)

	fmt.Printf("10 + 5 = %d\n", somar5(10))
	fmt.Printf("10 + 10 = %d\n", somar10(10))

	fmt.Println("\n=== FUNÇÕES COMO PARÂMETROS ===\n")

	numeros := []int{1, 2, 3, 4, 5}

	// Função que dobra valores
	dobrados := aplicar(numeros, func(x int) int {
		return x * 2
	})
	fmt.Printf("Dobrados: %v\n", dobrados)

	// Função que eleva ao quadrado
	quadrados := aplicar(numeros, func(x int) int {
		return x * x
	})
	fmt.Printf("Quadrados: %v\n", quadrados)

	fmt.Println("\n=== EXEMPLO PRÁTICO: FILTRO ===\n")

	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filtrar números pares
	pares := filtrar(valores, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("Pares: %v\n", pares)

	// Filtrar números maiores que 5
	maiores := filtrar(valores, func(x int) bool {
		return x > 5
	})
	fmt.Printf("Maiores que 5: %v\n", maiores)

	fmt.Println("\n=== EXEMPLO PRÁTICO: CALCULADORA ===\n")

	calc := criarCalculadora()

	fmt.Println("Operações:")
	calc.somar(10)      // +10
	calc.subtrair(3)    // -3
	calc.multiplicar(2) // *2
	fmt.Printf("Resultado: %.2f\n", calc.resultado())
}

// Função que retorna um closure (contador)
func criarContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// Função que retorna um closure (somador)
func criarSomador(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// Função que recebe função como parâmetro
func aplicar(numeros []int, fn func(int) int) []int {
	resultado := make([]int, len(numeros))
	for i, num := range numeros {
		resultado[i] = fn(num)
	}
	return resultado
}

// Função de filtro
func filtrar(numeros []int, condicao func(int) bool) []int {
	var resultado []int
	for _, num := range numeros {
		if condicao(num) {
			resultado = append(resultado, num)
		}
	}
	return resultado
}

// Exemplo de closure mais complexo: calculadora
type Calculadora struct {
	valor float64
}

func criarCalculadora() *Calculadora {
	calc := &Calculadora{valor: 0}

	// Métodos que modificam o estado
	calc.somar = func(n float64) {
		calc.valor += n
		fmt.Printf("  +%.2f = %.2f\n", n, calc.valor)
	}

	calc.subtrair = func(n float64) {
		calc.valor -= n
		fmt.Printf("  -%.2f = %.2f\n", n, calc.valor)
	}

	calc.multiplicar = func(n float64) {
		calc.valor *= n
		fmt.Printf("  *%.2f = %.2f\n", n, calc.valor)
	}

	calc.resultado = func() float64 {
		return calc.valor
	}

	return calc
}

// Métodos da calculadora (closures)
func (c *Calculadora) somar(n float64)       {}
func (c *Calculadora) subtrair(n float64)    {}
func (c *Calculadora) multiplicar(n float64) {}
func (c *Calculadora) resultado() float64    { return 0 }

/*
RESUMO:

FUNÇÃO ANÔNIMA:
    func(parametros) tipo {
        corpo
    }

EXECUTAR IMEDIATAMENTE:
    func(x int) {
        fmt.Println(x)
    }(10)

ATRIBUIR A VARIÁVEL:
    f := func(x int) int {
        return x * 2
    }

CLOSURE:
    func externa() func() int {
        x := 0
        return func() int {
            x++  // captura x
            return x
        }
    }

FUNÇÃO COMO PARÂMETRO:
    func processar(fn func(int) int) {
        resultado := fn(10)
    }

FUNÇÃO COMO RETORNO:
    func criar() func() int {
        return func() int {
            return 42
        }
    }

Usos comuns:
- Callbacks
- Processamento de listas (map, filter, reduce)
- Encapsulamento de estado (closures)
- Implementação de iteradores

Execute com:
    go run 04_anonimas_closures.go
*/
