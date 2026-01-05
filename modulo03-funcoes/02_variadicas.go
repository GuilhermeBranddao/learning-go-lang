package main

import "fmt"

/*
COMPARAÇÃO COM PYTHON:
--------------------
Python:
    def somar(*numeros):
        return sum(numeros)

    total = somar(1, 2, 3, 4, 5)

Go:
    func somar(numeros ...int) int {
        total := 0
        for _, n := range numeros {
            total += n
        }
        return total
    }

    total := somar(1, 2, 3, 4, 5)

O ... em Go é como o * em Python!
*/

// Função variádica básica
func somar(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// Função variádica com outros parâmetros
// O parâmetro variádico SEMPRE vem por último
func calcularMedia(nome string, notas ...float64) {
	if len(notas) == 0 {
		fmt.Printf("%s: sem notas\n", nome)
		return
	}

	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	media := soma / float64(len(notas))
	fmt.Printf("%s - Média: %.2f\n", nome, media)
}

// Função que aceita qualquer número de strings
func concatenar(separador string, palavras ...string) string {
	if len(palavras) == 0 {
		return ""
	}

	resultado := palavras[0]
	for i := 1; i < len(palavras); i++ {
		resultado += separador + palavras[i]
	}
	return resultado
}

// Função variádica com interface{} (aceita qualquer tipo)
func imprimir(valores ...interface{}) {
	for i, v := range valores {
		fmt.Printf("  [%d] %v (tipo: %T)\n", i, v, v)
	}
}

func main() {
	fmt.Println("=== FUNÇÕES VARIÁDICAS BÁSICAS ===\n")

	// Chamando com diferentes quantidades de argumentos
	fmt.Printf("Soma(): %d\n", somar())
	fmt.Printf("Soma(5): %d\n", somar(5))
	fmt.Printf("Soma(1, 2, 3): %d\n", somar(1, 2, 3))
	fmt.Printf("Soma(10, 20, 30, 40, 50): %d\n\n", somar(10, 20, 30, 40, 50))

	fmt.Println("=== VARIÁDICAS COM OUTROS PARÂMETROS ===\n")

	calcularMedia("João", 7.5, 8.0, 9.5)
	calcularMedia("Maria", 10.0, 9.0)
	calcularMedia("Pedro", 6.0, 7.0, 8.0, 9.0, 10.0)
	calcularMedia("Ana")

	fmt.Println("\n=== CONCATENAÇÃO DE STRINGS ===\n")

	frase1 := concatenar(" ", "Go", "é", "incrível")
	fmt.Println(frase1)

	frase2 := concatenar(", ", "Maçã", "Banana", "Laranja")
	fmt.Println(frase2)

	frase3 := concatenar(" - ", "Um", "Dois", "Três", "Quatro")
	fmt.Println(frase3)

	fmt.Println("\n=== PASSANDO SLICE PARA FUNÇÃO VARIÁDICA ===\n")

	// Use ... para "desempacotar" um slice
	numeros := []int{1, 2, 3, 4, 5}
	total := somar(numeros...) // desempacota o slice
	fmt.Printf("Soma de %v: %d\n", numeros, total)

	fmt.Println("\n=== FUNÇÃO VARIÁDICA COM INTERFACE{} ===\n")

	fmt.Println("Imprimindo valores de tipos diferentes:")
	imprimir(42, "texto", 3.14, true, 'A')

	fmt.Println("\n=== EXEMPLOS PRÁTICOS ===\n")

	// Encontrar o maior número
	maior := max(10, 5, 8, 15, 3, 12)
	fmt.Printf("Maior número: %d\n", maior)

	// Encontrar o menor número
	menor := min(10, 5, 8, 15, 3, 12)
	fmt.Printf("Menor número: %d\n", menor)

	// Criar título formatado
	titulo := criarTitulo("Go", "Python", "JavaScript", "Rust")
	fmt.Println(titulo)
}

// Função para encontrar o maior número
func max(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}

	maior := numeros[0]
	for _, num := range numeros {
		if num > maior {
			maior = num
		}
	}
	return maior
}

// Função para encontrar o menor número
func min(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}

	menor := numeros[0]
	for _, num := range numeros {
		if num < menor {
			menor = num
		}
	}
	return menor
}

// Criar título com várias linguagens
func criarTitulo(linguagens ...string) string {
	if len(linguagens) == 0 {
		return "Linguagens: nenhuma"
	}

	titulo := "Linguagens: "
	for i, lang := range linguagens {
		if i > 0 {
			titulo += ", "
		}
		titulo += lang
	}
	return titulo
}

/*
PONTOS IMPORTANTES:

1. Funções variádicas usam ...tipo
2. O parâmetro variádico SEMPRE vem por último
3. Dentro da função, é tratado como um slice
4. Use ... para desempacotar um slice ao chamar a função
5. interface{} aceita qualquer tipo

Sintaxe:
    func nome(param ...tipo) {
        // param é um []tipo
    }

Chamada:
    nome(valor1, valor2, valor3)

    // Ou com slice
    slice := []tipo{1, 2, 3}
    nome(slice...)

Execute com:
    go run 02_variadicas.go
*/
