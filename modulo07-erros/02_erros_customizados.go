package main

import (
	"errors"
	"fmt"
)

/*
ERROS CUSTOMIZADOS

Em Go, você pode criar seus próprios tipos de erro
implementando a interface error.

type error interface {
    Error() string
}

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    class ErroCustomizado(Exception):
        def __init__(self, codigo, mensagem):
            self.codigo = codigo
            self.mensagem = mensagem

Go:
    type ErroCustomizado struct {
        Codigo   int
        Mensagem string
    }

    func (e ErroCustomizado) Error() string {
        return fmt.Sprintf("erro %d: %s", e.Codigo, e.Mensagem)
    }
*/

func main() {
	fmt.Println("=== ERROS CUSTOMIZADOS ===\n")

	// Erro simples vs customizado
	err1 := errors.New("erro genérico")
	fmt.Println("Erro genérico:", err1)

	err2 := ErroValidacao{
		Campo:    "email",
		Valor:    "invalido",
		Mensagem: "formato de email inválido",
	}
	fmt.Println("Erro customizado:", err2)

	fmt.Println("\n=== TYPE ASSERTION COM ERROS ===\n")

	// Verificar tipo específico de erro
	resultado, err := validarEmail("teste@exemplo.com")
	if err != nil {
		// Type assertion para erro específico
		if errVal, ok := err.(ErroValidacao); ok {
			fmt.Printf("Erro de validação no campo '%s': %s\n",
				errVal.Campo, errVal.Mensagem)
		} else {
			fmt.Println("Erro genérico:", err)
		}
	} else {
		fmt.Println("Email válido:", resultado)
	}

	// Email inválido
	_, err = validarEmail("invalido")
	if errVal, ok := err.(ErroValidacao); ok {
		fmt.Printf("Campo: %s, Valor: %s, Erro: %s\n",
			errVal.Campo, errVal.Valor, errVal.Mensagem)
	}

	fmt.Println("\n=== EXEMPLO: ERRO HTTP ===\n")

	exemploErroHTTP()

	fmt.Println("\n=== EXEMPLO: ERRO DE BANCO ===\n")

	exemploErroBanco()

	fmt.Println("\n=== EXEMPLO: SENTINEL ERRORS ===\n")

	exemploSentinelErrors()
}

// ========================================
// ERRO CUSTOMIZADO: VALIDAÇÃO
// ========================================

type ErroValidacao struct {
	Campo    string
	Valor    string
	Mensagem string
}

func (e ErroValidacao) Error() string {
	return fmt.Sprintf("validação falhou no campo '%s' (valor: '%s'): %s",
		e.Campo, e.Valor, e.Mensagem)
}

func validarEmail(email string) (string, error) {
	if email == "" {
		return "", ErroValidacao{
			Campo:    "email",
			Valor:    email,
			Mensagem: "email não pode ser vazio",
		}
	}

	// Validação simples
	if len(email) < 5 || !contemCaractere(email, '@') {
		return "", ErroValidacao{
			Campo:    "email",
			Valor:    email,
			Mensagem: "formato de email inválido",
		}
	}

	return email, nil
}

func contemCaractere(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

// ========================================
// ERRO CUSTOMIZADO: HTTP
// ========================================

type ErroHTTP struct {
	StatusCode int
	Mensagem   string
	URL        string
}

func (e ErroHTTP) Error() string {
	return fmt.Sprintf("HTTP %d ao acessar %s: %s",
		e.StatusCode, e.URL, e.Mensagem)
}

func fazerRequisicao(url string) error {
	// Simular requisição
	if url == "" {
		return ErroHTTP{
			StatusCode: 400,
			Mensagem:   "URL não pode ser vazia",
			URL:        url,
		}
	}

	if url == "http://naoexiste.com" {
		return ErroHTTP{
			StatusCode: 404,
			Mensagem:   "página não encontrada",
			URL:        url,
		}
	}

	return nil
}

func exemploErroHTTP() {
	urls := []string{
		"http://exemplo.com",
		"",
		"http://naoexiste.com",
	}

	for _, url := range urls {
		if err := fazerRequisicao(url); err != nil {
			fmt.Println("Erro:", err)

			// Type assertion
			if httpErr, ok := err.(ErroHTTP); ok {
				fmt.Printf("  Status Code: %d\n", httpErr.StatusCode)
			}
		} else {
			fmt.Printf("Sucesso ao acessar: %s\n", url)
		}
	}
}

// ========================================
// ERRO CUSTOMIZADO: BANCO DE DADOS
// ========================================

type ErroBanco struct {
	Operacao string
	Tabela   string
	Detalhe  error
}

func (e ErroBanco) Error() string {
	return fmt.Sprintf("erro ao %s na tabela '%s': %v",
		e.Operacao, e.Tabela, e.Detalhe)
}

func buscarUsuarioBanco(id int) (string, error) {
	if id < 0 {
		return "", ErroBanco{
			Operacao: "buscar",
			Tabela:   "usuarios",
			Detalhe:  errors.New("ID inválido"),
		}
	}

	if id > 1000 {
		return "", ErroBanco{
			Operacao: "buscar",
			Tabela:   "usuarios",
			Detalhe:  errors.New("registro não encontrado"),
		}
	}

	return fmt.Sprintf("Usuario_%d", id), nil
}

func exemploErroBanco() {
	ids := []int{10, -1, 9999}

	for _, id := range ids {
		usuario, err := buscarUsuarioBanco(id)
		if err != nil {
			fmt.Println("Erro:", err)

			// Type assertion
			if dbErr, ok := err.(ErroBanco); ok {
				fmt.Printf("  Operação: %s, Tabela: %s\n",
					dbErr.Operacao, dbErr.Tabela)
			}
		} else {
			fmt.Printf("Usuário encontrado: %s\n", usuario)
		}
	}
}

// ========================================
// SENTINEL ERRORS (Erros Predefinidos)
// ========================================

var (
	ErrNaoEncontrado     = errors.New("recurso não encontrado")
	ErrNaoAutorizado     = errors.New("não autorizado")
	ErrParametroInvalido = errors.New("parâmetro inválido")
)

func buscarItem(id int) (string, error) {
	if id < 0 {
		return "", ErrParametroInvalido
	}

	if id > 100 {
		return "", ErrNaoEncontrado
	}

	return fmt.Sprintf("Item_%d", id), nil
}

func exemploSentinelErrors() {
	ids := []int{50, -1, 200}

	for _, id := range ids {
		item, err := buscarItem(id)
		if err != nil {
			// Comparar com sentinel errors usando errors.Is (Go 1.13+)
			switch err {
			case ErrNaoEncontrado:
				fmt.Printf("ID %d: Item não encontrado\n", id)
			case ErrParametroInvalido:
				fmt.Printf("ID %d: Parâmetro inválido\n", id)
			case ErrNaoAutorizado:
				fmt.Printf("ID %d: Não autorizado\n", id)
			default:
				fmt.Printf("ID %d: Erro: %v\n", id, err)
			}
		} else {
			fmt.Printf("ID %d: %s\n", id, item)
		}
	}
}

/*
RESUMO DE ERROS CUSTOMIZADOS:

CRIAR ERRO CUSTOMIZADO:
    type MeuErro struct {
        Campo1 tipo1
        Campo2 tipo2
    }

    func (e MeuErro) Error() string {
        return fmt.Sprintf("erro: %v", e.Campo1)
    }

TYPE ASSERTION:
    if errCustom, ok := err.(MeuErro); ok {
        // usar campos específicos
        fmt.Println(errCustom.Campo1)
    }

SENTINEL ERRORS:
    var (
        ErrNaoEncontrado = errors.New("não encontrado")
        ErrInvalido      = errors.New("inválido")
    )

    if err == ErrNaoEncontrado {
        // tratar erro específico
    }

QUANDO USAR:

Erro Customizado:
✓ Precisa de contexto adicional
✓ Tratamento específico necessário
✓ Informações estruturadas

Sentinel Error:
✓ Erros predefinidos e conhecidos
✓ Comparação direta
✓ Constantes de erro

Error Simples (errors.New):
✓ Erro único/não reutilizado
✓ Não precisa de contexto extra
✓ Mensagem descritiva suficiente

VANTAGENS:
✓ Informações ricas sobre o erro
✓ Type safety
✓ Tratamento específico por tipo
✓ Melhor debugging

Execute com:
    go run 02_erros_customizados.go
*/
