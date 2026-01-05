package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
EXERCÍCIO 1: CALCULADORA

Crie uma calculadora que:
1. Recebe dois números do usuário
2. Recebe a operação (+, -, *, /)
3. Exibe o resultado
4. Trata divisão por zero
5. Permite fazer várias operações (loop)
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== CALCULADORA ===")
	fmt.Println("Digite 'sair' para encerrar\n")

	for {
		// Primeiro número
		fmt.Print("Primeiro número: ")
		scanner.Scan()
		input1 := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input1) == "sair" {
			fmt.Println("Encerrando...")
			break
		}

		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("❌ Número inválido!\n")
			continue
		}

		// Operação
		fmt.Print("Operação (+, -, *, /): ")
		scanner.Scan()
		operacao := strings.TrimSpace(scanner.Text())

		// Segundo número
		fmt.Print("Segundo número: ")
		scanner.Scan()
		input2 := strings.TrimSpace(scanner.Text())

		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("❌ Número inválido!\n")
			continue
		}

		// Calcular e exibir resultado
		resultado, erro := calcular(num1, num2, operacao)
		if erro != "" {
			fmt.Printf("❌ Erro: %s\n\n", erro)
		} else {
			fmt.Printf("✓ %.2f %s %.2f = %.2f\n\n", num1, operacao, num2, resultado)
		}
	}
}

// Função que realiza o cálculo
func calcular(a, b float64, op string) (float64, string) {
	switch op {
	case "+":
		return a + b, ""
	case "-":
		return a - b, ""
	case "*":
		return a * b, ""
	case "/":
		if b == 0 {
			return 0, "Divisão por zero não permitida"
		}
		return a / b, ""
	default:
		return 0, "Operação inválida"
	}
}

/*
EXEMPLO DE EXECUÇÃO:

=== CALCULADORA ===
Digite 'sair' para encerrar

Primeiro número: 10
Operação (+, -, *, /): +
Segundo número: 5
✓ 10.00 + 5.00 = 15.00

Primeiro número: 20
Operação (+, -, *, /): /
Segundo número: 0
❌ Erro: Divisão por zero não permitida

Primeiro número: sair
Encerrando...

PONTOS DE APRENDIZADO:
- Entrada de dados com bufio.Scanner
- Conversão de string para float
- Tratamento de erros
- Switch para múltiplas condições
- Loops infinitos com break

Execute com:
    go run exercicio01_calculadora.go
*/
