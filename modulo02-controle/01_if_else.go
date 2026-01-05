package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python:
    if condicao:
        codigo

Go:
    if condicao {
        codigo
    }

Diferenças:
- SEMPRE use chaves {}
- NÃO use parênteses () (opcional, mas não idiomático)
*/

func main() {
	fmt.Println("=== IF/ELSE BÁSICO ===\n")

	idade := 20

	// If simples
	if idade >= 18 {
		fmt.Println("Maior de idade")
	}

	// If/Else
	if idade < 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Maior de idade")
	}

	// If/Else If/Else
	if idade < 13 {
		fmt.Println("Criança")
	} else if idade < 18 {
		fmt.Println("Adolescente")
	} else if idade < 60 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Idoso")
	}

	fmt.Println("\n=== IF COM INICIALIZAÇÃO ===\n")

	// Go permite inicializar variável no if
	// A variável só existe dentro do escopo do if
	if nota := 7.5; nota >= 7.0 {
		fmt.Printf("Aprovado com nota %.1f\n", nota)
	} else {
		fmt.Printf("Reprovado com nota %.1f\n", nota)
	}
	// fmt.Println(nota)  // ERRO! nota não existe aqui

	fmt.Println("\n=== COMPARAÇÕES ===\n")

	a := 10
	b := 20

	if a == b {
		fmt.Println("a é igual a b")
	}

	if a != b {
		fmt.Println("a é diferente de b")
	}

	if a < b {
		fmt.Println("a é menor que b")
	}

	if a > b {
		fmt.Println("a é maior que b")
	}

	fmt.Println("\n=== OPERADORES LÓGICOS ===\n")

	temCarteira := true
	maiorDe18 := true

	// AND (&&)
	if temCarteira && maiorDe18 {
		fmt.Println("Pode dirigir!")
	}

	// OR (||)
	temDesconto := false
	estudante := true

	if temDesconto || estudante {
		fmt.Println("Tem direito a desconto")
	}

	// NOT (!)
	if !temDesconto {
		fmt.Println("Não tem desconto")
	}

	fmt.Println("\n=== EXEMPLO PRÁTICO ===\n")

	// Sistema de autenticação
	usuario := "admin"
	senha := "1234"

	if usuario == "admin" && senha == "1234" {
		fmt.Println("✓ Login realizado com sucesso!")
	} else {
		fmt.Println("✗ Usuário ou senha inválidos")
	}

	// Verificação de intervalo
	temperatura := 25

	if temperatura < 0 {
		fmt.Println("🥶 Muito frio!")
	} else if temperatura >= 0 && temperatura < 15 {
		fmt.Println("🧊 Frio")
	} else if temperatura >= 15 && temperatura < 25 {
		fmt.Println("😊 Agradável")
	} else if temperatura >= 25 && temperatura < 35 {
		fmt.Println("☀️ Quente")
	} else {
		fmt.Println("🔥 Muito quente!")
	}

	fmt.Println("\n=== CUIDADO COM PARÊNTESES ===\n")

	// Funciona, mas não é idiomático em Go
	if idade > 18 {
		fmt.Println("Maior de idade (com parênteses)")
	}

	// Estilo Go - sem parênteses
	if idade > 18 {
		fmt.Println("Maior de idade (sem parênteses - idiomático)")
	}
}

/*
DICAS IMPORTANTES:

1. SEMPRE use chaves {}, mesmo para uma linha
2. Evite parênteses na condição (não é idiomático)
3. Use if com inicialização para limitar escopo
4. Use go fmt para formatar automaticamente

Execute com:
    go run 01_if_else.go
*/
