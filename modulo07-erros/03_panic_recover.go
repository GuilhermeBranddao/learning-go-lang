package main

import "fmt"

/*
PANIC E RECOVER

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    try:
        raise Exception("erro!")
    except Exception as e:
        print(f"Capturado: {e}")

Go:
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Capturado:", r)
        }
    }()
    panic("erro!")

IMPORTANTE:
- Panic é para erros IRRECUPERÁVEIS
- Use erro normal (return error) sempre que possível
- Panic para bugs, não para erros esperados
*/

func main() {
	fmt.Println("=== PANIC BÁSICO ===\n")

	// Panic simples (encerra o programa!)
	// Descomente para ver:
	// panic("algo muito errado aconteceu!")

	fmt.Println("=== RECOVER ===\n")

	// Recover captura panic
	exemploRecover()
	fmt.Println("Programa continuou após recover!")

	fmt.Println("\n=== DEFER COM RECOVER ===\n")

	funcaoQuePodePanicar()
	fmt.Println("Programa ainda está rodando!")

	fmt.Println("\n=== MÚLTIPLOS PANICS ===\n")

	multiplasOperacoes()

	fmt.Println("\n=== QUANDO USAR PANIC ===\n")

	exemploQuandoUsarPanic()
}

// ========================================
// RECOVER BÁSICO
// ========================================

func exemploRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado do panic:", r)
		}
	}()

	fmt.Println("Antes do panic")
	panic("ops!")
	fmt.Println("Depois do panic (nunca executa)")
}

// ========================================
// FUNÇÃO QUE PODE PANICAR
// ========================================

func funcaoQuePodePanicar() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic capturado: %v\n", r)
			fmt.Println("Limpeza de recursos...")
		}
	}()

	fmt.Println("Iniciando operação arriscada...")
	causarPanic(true)
	fmt.Println("Operação concluída (não executa se panic)")
}

func causarPanic(devePanicar bool) {
	if devePanicar {
		panic("operação não suportada!")
	}
	fmt.Println("Operação bem-sucedida")
}

// ========================================
// MÚLTIPLAS OPERAÇÕES COM RECOVER
// ========================================

func operacaoSegura(id int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Operação %d: panic recuperado - %v\n", id, r)
		}
	}()

	fmt.Printf("  Operação %d: iniciando\n", id)

	if id == 2 {
		panic("erro na operação 2")
	}

	fmt.Printf("  Operação %d: concluída\n", id)
}

func multiplasOperacoes() {
	for i := 1; i <= 4; i++ {
		operacaoSegura(i)
	}
	fmt.Println("Todas as operações processadas!")
}

// ========================================
// QUANDO USAR PANIC VS ERROR
// ========================================

func exemploQuandoUsarPanic() {
	// ✅ USAR ERROR (preferido)
	resultado1, err := dividirComErro(10, 0)
	if err != nil {
		fmt.Println("Erro esperado:", err)
	} else {
		fmt.Println("Resultado:", resultado1)
	}

	// ❌ USAR PANIC (apenas para erros irrecuperáveis)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic capturado:", r)
		}
	}()

	// Isso causaria panic (descomente para testar):
	// resultado2 := dividirComPanic(10, 0)
	// fmt.Println("Resultado:", resultado2)

	fmt.Println("\n=== CASOS DE USO VÁLIDOS PARA PANIC ===\n")

	// 1. Inicialização falhou
	exemploInicializacao()

	// 2. Erro de programação (bug)
	exemploErroProgramacao()

	// 3. Situação impossível
	exemploSituacaoImpossivel()
}

func dividirComErro(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func dividirComPanic(a, b float64) float64 {
	if b == 0 {
		panic("divisão por zero!")
	}
	return a / b
}

// ========================================
// CASOS VÁLIDOS PARA PANIC
// ========================================

// 1. Falha na inicialização
func exemploInicializacao() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Falha na inicialização:", r)
		}
	}()

	// Em init() ou startup, panic é aceitável
	config := carregarConfiguracao()
	if config == nil {
		panic("configuração não pode ser carregada")
	}
}

func carregarConfiguracao() interface{} {
	// Simular falha
	return nil
}

// 2. Erro de programação (bug)
func exemploErroProgramacao() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Bug detectado:", r)
		}
	}()

	// Array out of bounds é um bug
	array := []int{1, 2, 3}
	// Isso causaria panic:
	// valor := array[10]
	_ = array[0] // OK
	fmt.Println("  Array acessado corretamente")
}

// 3. Situação impossível
func exemploSituacaoImpossivel() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Situação impossível:", r)
		}
	}()

	tipo := "conhecido"
	switch tipo {
	case "conhecido":
		fmt.Println("  Tipo conhecido processado")
	default:
		// Isso nunca deveria acontecer
		panic("tipo desconhecido: " + tipo)
	}
}

/*
RESUMO DE PANIC E RECOVER:

PANIC:
    panic(valor)
    - Para erros IRRECUPERÁVEIS
    - Interrompe execução normal
    - Executa defers antes de propagar

RECOVER:
    defer func() {
        if r := recover(); r != nil {
            // tratar panic
        }
    }()
    - DEVE estar em defer
    - Captura panic da goroutine atual
    - Retorna nil se não houve panic

QUANDO USAR PANIC:
✓ Erro de programação (bug)
✓ Falha na inicialização
✓ Situação impossível
✓ Violação de invariante

QUANDO NÃO USAR PANIC:
✗ Erros esperados/previsíveis
✗ Validação de entrada
✗ Falhas de I/O
✗ Erros de rede
✗ Erros de banco de dados

PREFERIR ERROR QUANDO:
✓ Erro é esperado
✓ Chamador pode tratar
✓ Operação pode falhar normalmente
✓ Precisa de controle fino

DIFERENÇAS COM PYTHON:

Python:
    - Exceções para controle de fluxo
    - Try/except muito comum
    - Exceções podem ser ignoradas

Go:
    - Panic raramente usado
    - Erros são valores
    - Impossível ignorar erros

REGRA DE OURO:
"Don't panic!" - Use return error

Execute com:
    go run 03_panic_recover.go
*/
