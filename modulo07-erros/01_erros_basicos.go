package main

import (
	"errors"
	"fmt"
)

/*
TRATAMENTO DE ERROS EM GO

COMPARAÇÃO COM PYTHON:
--------------------
Python (Exceções):
    try:
        resultado = dividir(10, 0)
    except ZeroDivisionError as e:
        print(f"Erro: {e}")

Go (Valores de Erro):
    resultado, err := dividir(10, 0)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

FILOSOFIA GO:
- Erros são VALORES, não exceções
- Tratamento EXPLÍCITO
- Impossível ignorar sem intenção
*/

func main() {
	fmt.Println("=== ERROS BÁSICOS ===\n")

	// Exemplo 1: Divisão
	resultado1, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", resultado1)
	}

	// Exemplo 2: Divisão por zero
	resultado2, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", resultado2)
	}

	fmt.Println("\n=== MÚLTIPLOS RETORNOS COM ERRO ===\n")

	// Padrão comum: (resultado, erro)
	usuario, err := buscarUsuario(1)
	if err != nil {
		fmt.Println("Erro ao buscar usuário:", err)
	} else {
		fmt.Printf("Usuário encontrado: %s\n", usuario)
	}

	// Usuário não existe
	_, err = buscarUsuario(999)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	fmt.Println("\n=== CRIANDO ERROS ===\n")

	// errors.New
	err1 := errors.New("erro simples")
	fmt.Println("err1:", err1)

	// fmt.Errorf (com formatação)
	codigo := 404
	err2 := fmt.Errorf("erro %d: recurso não encontrado", codigo)
	fmt.Println("err2:", err2)

	fmt.Println("\n=== PATTERN: EARLY RETURN ===\n")

	// Padrão idiomático: verificar erro e retornar cedo
	if err := processar(true); err != nil {
		fmt.Println("Erro no processamento:", err)
	}

	fmt.Println("\n=== EXEMPLO: VALIDAÇÃO ===\n")

	exemploValidacao()

	fmt.Println("\n=== EXEMPLO: CÁLCULO COM ERROS ===\n")

	exemploCalculadora()
}

// Função que retorna erro
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// Função que retorna resultado e erro
func buscarUsuario(id int) (string, error) {
	usuarios := map[int]string{
		1: "João",
		2: "Maria",
		3: "Pedro",
	}

	usuario, existe := usuarios[id]
	if !existe {
		return "", fmt.Errorf("usuário %d não encontrado", id)
	}

	return usuario, nil
}

// Função para demonstrar early return
func processar(comErro bool) error {
	// Primeiro passo
	if comErro {
		return errors.New("erro no primeiro passo")
	}

	// Segundo passo (só executa se não houve erro)
	fmt.Println("  Processamento bem-sucedido!")
	return nil
}

// ========================================
// EXEMPLO: VALIDAÇÃO
// ========================================

type Usuario struct {
	Nome  string
	Email string
	Idade int
}

func (u Usuario) Validar() error {
	if u.Nome == "" {
		return errors.New("nome é obrigatório")
	}

	if u.Email == "" {
		return errors.New("email é obrigatório")
	}

	if u.Idade < 0 {
		return errors.New("idade não pode ser negativa")
	}

	if u.Idade < 18 {
		return errors.New("usuário deve ser maior de idade")
	}

	return nil // Sem erros
}

func exemploValidacao() {
	// Usuário válido
	u1 := Usuario{Nome: "João", Email: "joao@email.com", Idade: 25}
	if err := u1.Validar(); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Usuário válido:", u1.Nome)
	}

	// Usuário inválido
	u2 := Usuario{Nome: "", Email: "maria@email.com", Idade: 30}
	if err := u2.Validar(); err != nil {
		fmt.Println("Erro:", err)
	}

	// Menor de idade
	u3 := Usuario{Nome: "Pedro", Email: "pedro@email.com", Idade: 15}
	if err := u3.Validar(); err != nil {
		fmt.Println("Erro:", err)
	}
}

// ========================================
// EXEMPLO: CALCULADORA COM ERROS
// ========================================

func calcular(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("divisão por zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("operação inválida: %s", op)
	}
}

func exemploCalculadora() {
	operacoes := []struct {
		a, b float64
		op   string
	}{
		{10, 5, "+"},
		{10, 5, "-"},
		{10, 5, "*"},
		{10, 5, "/"},
		{10, 0, "/"}, // Erro!
		{10, 5, "%"}, // Erro!
	}

	for _, calc := range operacoes {
		resultado, err := calcular(calc.a, calc.b, calc.op)
		if err != nil {
			fmt.Printf("%.0f %s %.0f = ERRO: %v\n",
				calc.a, calc.op, calc.b, err)
		} else {
			fmt.Printf("%.0f %s %.0f = %.2f\n",
				calc.a, calc.op, calc.b, resultado)
		}
	}
}

/*
RESUMO DE TRATAMENTO DE ERROS:

PADRÃO BÁSICO:
    resultado, err := funcao()
    if err != nil {
        // tratar erro
        return
    }
    // usar resultado

CRIAR ERROS:
    errors.New("mensagem")
    fmt.Errorf("formato %d", valor)

MÚLTIPLOS RETORNOS:
    func nome() (tipo, error) {
        if problema {
            return valorZero, errors.New("erro")
        }
        return resultado, nil
    }

CONVENÇÕES:
✓ Erro é sempre o último retorno
✓ nil significa sucesso
✓ Verificar erro imediatamente
✓ Não ignorar erros
✓ Retornar erros para cima (propagação)

DIFERENÇAS COM PYTHON:

Python:
    ✓ Exceções podem ser não tratadas
    ✓ Try/except pode estar longe do erro
    ✓ Fluxo de controle por exceções
    ✓ Pode ignorar exceções silenciosamente

Go:
    ✓ Erros DEVEM ser verificados
    ✓ Tratamento local e explícito
    ✓ Fluxo de controle normal
    ✓ Impossível ignorar sem intenção

VANTAGENS DO APPROACH GO:
✓ Código mais previsível
✓ Impossível esquecer de tratar erro
✓ Fluxo de controle claro
✓ Menos surpresas em runtime

Execute com:
    go run 01_erros_basicos.go
*/
