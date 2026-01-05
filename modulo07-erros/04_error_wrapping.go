package main

import (
	"errors"
	"fmt"
)

/*
ERROR WRAPPING (Go 1.13+)

Permite "embrulhar" erros para adicionar contexto sem perder o erro original.

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    try:
        processar()
    except Exception as e:
        raise Exception(f"Falha no processo: {e}") from e

Go:
    if err := processar(); err != nil {
        return fmt.Errorf("falha no processo: %w", err)
    }

NOVIDADES GO 1.13+:
- %w para wrapping
- errors.Is() para comparar
- errors.As() para type assertion
- errors.Unwrap() para desembrulhar
*/

func main() {
	fmt.Println("=== ERROR WRAPPING ===\n")

	// Exemplo básico
	err := operacaoCompleta()
	if err != nil {
		fmt.Println("Erro:", err)
		fmt.Println()
	}

	fmt.Println("=== errors.Is (Comparação) ===\n")

	exemploErrorsIs()

	fmt.Println("\n=== errors.As (Type Assertion) ===\n")

	exemploErrorsAs()

	fmt.Println("\n=== errors.Unwrap ===\n")

	exemploUnwrap()

	fmt.Println("\n=== EXEMPLO REAL: PROCESSAMENTO DE ARQUIVO ===\n")

	exemploReal()
}

// ========================================
// WRAPPING BÁSICO
// ========================================

var (
	ErrConexao = errors.New("erro de conexão")
	ErrTimeout = errors.New("timeout")
)

func conectarBanco() error {
	// Simular erro de conexão
	return ErrConexao
}

func buscarDados() error {
	if err := conectarBanco(); err != nil {
		// Wrapping com %w
		return fmt.Errorf("buscar dados falhou: %w", err)
	}
	return nil
}

func processarDados() error {
	if err := buscarDados(); err != nil {
		// Outro nível de wrapping
		return fmt.Errorf("processar dados falhou: %w", err)
	}
	return nil
}

func operacaoCompleta() error {
	if err := processarDados(); err != nil {
		return fmt.Errorf("operação completa falhou: %w", err)
	}
	return nil
}

// ========================================
// errors.Is - Comparar Erros
// ========================================

func exemploErrorsIs() {
	err := processarDados()

	// Verificar se erro contém ErrConexao (em qualquer nível)
	if errors.Is(err, ErrConexao) {
		fmt.Println("✓ Detectado erro de conexão na cadeia")
		fmt.Println("  Erro completo:", err)
	}

	// Verificar outro erro
	if errors.Is(err, ErrTimeout) {
		fmt.Println("✓ Detectado timeout")
	} else {
		fmt.Println("✗ Não é erro de timeout")
	}
}

// ========================================
// errors.As - Type Assertion
// ========================================

type ErroHTTP struct {
	StatusCode int
	URL        string
}

func (e *ErroHTTP) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.URL)
}

func fazerRequisicao() error {
	return &ErroHTTP{
		StatusCode: 404,
		URL:        "http://exemplo.com",
	}
}

func chamarAPI() error {
	if err := fazerRequisicao(); err != nil {
		return fmt.Errorf("chamar API falhou: %w", err)
	}
	return nil
}

func exemploErrorsAs() {
	err := chamarAPI()
	if err != nil {
		fmt.Println("Erro:", err)

		// Type assertion com errors.As
		var httpErr *ErroHTTP
		if errors.As(err, &httpErr) {
			fmt.Printf("✓ Status Code: %d\n", httpErr.StatusCode)
			fmt.Printf("✓ URL: %s\n", httpErr.URL)
		}
	}
}

// ========================================
// errors.Unwrap
// ========================================

func exemploUnwrap() {
	err := processarDados()

	fmt.Println("Erro wrapped:", err)

	// Desembrulhar camada por camada
	nivel := 1
	for err != nil {
		fmt.Printf("Nível %d: %v\n", nivel, err)
		err = errors.Unwrap(err)
		nivel++
	}
}

// ========================================
// EXEMPLO REAL: PROCESSAMENTO DE ARQUIVO
// ========================================

var (
	ErrArquivoNaoEncontrado = errors.New("arquivo não encontrado")
	ErrPermissaoNegada      = errors.New("permissão negada")
	ErrFormatoInvalido      = errors.New("formato inválido")
)

type ErroProcessamento struct {
	Arquivo string
	Linha   int
	Erro    error
}

func (e *ErroProcessamento) Error() string {
	return fmt.Sprintf("erro ao processar %s (linha %d): %v",
		e.Arquivo, e.Linha, e.Erro)
}

func (e *ErroProcessamento) Unwrap() error {
	return e.Erro
}

func lerArquivo(nome string) error {
	if nome == "" {
		return ErrArquivoNaoEncontrado
	}
	if nome == "protegido.txt" {
		return ErrPermissaoNegada
	}
	return nil
}

func parsearLinha(linha string, numero int) error {
	if linha == "INVALIDO" {
		return &ErroProcessamento{
			Arquivo: "dados.txt",
			Linha:   numero,
			Erro:    ErrFormatoInvalido,
		}
	}
	return nil
}

func processarArquivo(nome string) error {
	// Tentar ler arquivo
	if err := lerArquivo(nome); err != nil {
		return fmt.Errorf("processar arquivo %s: %w", nome, err)
	}

	// Simular processamento de linhas
	linhas := []string{"OK", "OK", "INVALIDO", "OK"}
	for i, linha := range linhas {
		if err := parsearLinha(linha, i+1); err != nil {
			return fmt.Errorf("processar arquivo %s: %w", nome, err)
		}
	}

	return nil
}

func exemploReal() {
	arquivos := []string{
		"dados.txt",
		"",
		"protegido.txt",
	}

	for _, arquivo := range arquivos {
		fmt.Printf("Processando: %s\n", arquivo)

		if err := processarArquivo(arquivo); err != nil {
			fmt.Println("Erro:", err)

			// Verificar tipo específico
			if errors.Is(err, ErrArquivoNaoEncontrado) {
				fmt.Println("  → Arquivo não existe")
			} else if errors.Is(err, ErrPermissaoNegada) {
				fmt.Println("  → Sem permissão de leitura")
			} else if errors.Is(err, ErrFormatoInvalido) {
				fmt.Println("  → Formato de dados inválido")

				// Extrair informações específicas
				var procErr *ErroProcessamento
				if errors.As(err, &procErr) {
					fmt.Printf("  → Linha problemática: %d\n", procErr.Linha)
				}
			}
		} else {
			fmt.Println("  ✓ Sucesso")
		}
		fmt.Println()
	}
}

/*
RESUMO DE ERROR WRAPPING:

WRAPPING:
    fmt.Errorf("contexto: %w", erroOriginal)
    - %w preserva o erro original
    - Adiciona contexto
    - Permite verificação com Is/As

COMPARAR ERROS:
    errors.Is(err, ErrAlvo)
    - Procura em toda a cadeia
    - Retorna true se encontrar
    - Substitui err == ErrAlvo

TYPE ASSERTION:
    var target *TipoErro
    errors.As(err, &target)
    - Procura tipo específico na cadeia
    - Popula target se encontrar
    - Retorna true/false

DESEMBRULHAR:
    errors.Unwrap(err)
    - Remove uma camada de wrapping
    - Retorna erro interno ou nil
    - Útil para debugging

IMPLEMENTAR Unwrap:
    func (e *MeuErro) Unwrap() error {
        return e.ErroInterno
    }

VANTAGENS:
✓ Contexto rico sobre erros
✓ Cadeia de erros preservada
✓ Type-safe error checking
✓ Debugging facilitado

QUANDO USAR:

%w (wrapping):
✓ Adicionar contexto
✓ Preservar erro original
✓ Permitir verificação posterior

%v (sem wrapping):
✓ Não precisa verificar erro original
✓ Apenas mensagem descritiva
✓ Erro pode ser descartado

PADRÃO COMUM:
    func operacao() error {
        if err := subOperacao(); err != nil {
            return fmt.Errorf("operacao falhou: %w", err)
        }
        return nil
    }

Execute com:
    go run 04_error_wrapping.go
*/
