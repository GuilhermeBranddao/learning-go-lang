package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python:
    def somar(a, b):
        return a + b

Go:
    func somar(a, b int) int {
        return a + b
    }

Diferenças:
- Palavra-chave: def vs func
- Tipos vêm DEPOIS do nome
- Tipo de retorno no final
*/

// Função básica
func saudacao() {
	fmt.Println("Olá, mundo!")
}

// Função com parâmetros
func saudar(nome string) {
	fmt.Printf("Olá, %s!\n", nome)
}

// Função com retorno
func dobro(n int) int {
	return n * 2
}

// Múltiplos parâmetros do mesmo tipo
func somar(a, b int) int {
	return a + b
}

// Múltiplos parâmetros de tipos diferentes
func apresentar(nome string, idade int) {
	fmt.Printf("%s tem %d anos\n", nome, idade)
}

// Função com múltiplos retornos
func dividir(a, b int) (int, int) {
	quociente := a / b
	resto := a % b
	return quociente, resto
}

// Named return values
func calcular(a, b int) (soma int, multiplicacao int) {
	soma = a + b
	multiplicacao = a * b
	return // naked return - retorna as variáveis nomeadas
}

// Função que retorna erro (padrão Go)
func dividirSeguro(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== FUNÇÕES BÁSICAS ===\n")

	saudacao()
	saudar("Maria")

	resultado := dobro(5)
	fmt.Printf("Dobro de 5: %d\n\n", resultado)

	fmt.Println("=== MÚLTIPLOS PARÂMETROS ===\n")

	soma := somar(10, 5)
	fmt.Printf("10 + 5 = %d\n", soma)

	apresentar("João", 25)

	fmt.Println("\n=== MÚLTIPLOS RETORNOS ===\n")

	quociente, resto := dividir(17, 5)
	fmt.Printf("17 / 5 = %d (resto %d)\n", quociente, resto)

	// Ignorar um retorno com _
	q, _ := dividir(20, 3)
	fmt.Printf("20 / 3 = %d (resto ignorado)\n", q)

	fmt.Println("\n=== NAMED RETURNS ===\n")

	s, m := calcular(5, 3)
	fmt.Printf("5 + 3 = %d\n", s)
	fmt.Printf("5 * 3 = %d\n", m)

	fmt.Println("\n=== TRATAMENTO DE ERROS ===\n")

	// Divisão válida
	resultado1, err1 := dividirSeguro(10, 2)
	if err1 != nil {
		fmt.Println("Erro:", err1)
	} else {
		fmt.Printf("10 / 2 = %d\n", resultado1)
	}

	// Divisão por zero
	resultado2, err2 := dividirSeguro(10, 0)
	if err2 != nil {
		fmt.Println("Erro:", err2)
	} else {
		fmt.Printf("10 / 0 = %d\n", resultado2)
	}

	fmt.Println("\n=== EXEMPLOS PRÁTICOS ===\n")

	// Função para verificar se é par
	fmt.Printf("5 é par? %t\n", ehPar(5))
	fmt.Printf("8 é par? %t\n", ehPar(8))

	// Função para calcular média
	media := calcularMedia(7.5, 8.0, 9.5)
	fmt.Printf("Média: %.2f\n", media)

	// Função para converter temperatura
	celsius := 25.0
	fahrenheit := celsiusParaFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)
}

// Funções auxiliares
func ehPar(n int) bool {
	return n%2 == 0
}

func calcularMedia(notas ...float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	return soma / float64(len(notas))
}

func celsiusParaFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

/*
RESUMO:

Estrutura:
    func nome(parametros) tipoRetorno {
        corpo
        return valor
    }

Múltiplos retornos:
    func nome() (tipo1, tipo2) {
        return valor1, valor2
    }

Named returns:
    func nome() (x int, y int) {
        x = 10
        y = 20
        return
    }

Padrão erro em Go:
    func nome() (resultado, error) {
        if erro {
            return 0, errors.New("mensagem")
        }
        return resultado, nil
    }

Execute com:
    go run 01_funcoes_basicas.go
*/
