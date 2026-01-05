package matematica

import "testing"

/*
TESTES EM GO

COMPARAÇÃO COM PYTHON:
--------------------
Python (pytest):
    def test_somar():
        assert somar(2, 3) == 5

Go (testing):
    func TestSomar(t *testing.T) {
        resultado := Somar(2, 3)
        if resultado != 5 {
            t.Errorf("Somar(2, 3) = %d; esperado 5", resultado)
        }
    }

CONVENÇÕES:
- Arquivo: *_test.go
- Função: Test* (maiúscula)
- Parâmetro: *testing.T
- Mesmo package do código

EXECUTAR:
    go test
    go test -v           (verbose)
    go test -cover       (cobertura)
    go test -run TestNome (teste específico)
*/

// ========================================
// TESTES BÁSICOS
// ========================================

func TestSomar(t *testing.T) {
	resultado := Somar(2, 3)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Somar(2, 3) = %d; esperado %d", resultado, esperado)
	}
}

func TestSubtrair(t *testing.T) {
	resultado := Subtrair(10, 5)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Subtrair(10, 5) = %d; esperado %d", resultado, esperado)
	}
}

// ========================================
// TABLE-DRIVEN TESTS (Padrão recomendado!)
// ========================================

func TestSomar_TableDriven(t *testing.T) {
	// Slice de casos de teste
	tests := []struct {
		name     string // Nome do caso
		a, b     int    // Inputs
		esperado int    // Output esperado
	}{
		{"positivos", 2, 3, 5},
		{"negativos", -2, -3, -5},
		{"zero", 0, 5, 5},
		{"zero ambos", 0, 0, 0},
		{"positivo e negativo", 10, -5, 5},
	}

	// Executar cada caso
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado := Somar(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("Somar(%d, %d) = %d; esperado %d",
					tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}

func TestMultiplicar_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		esperado int
	}{
		{"positivos", 3, 4, 12},
		{"negativos", -3, -4, 12},
		{"zero", 0, 5, 0},
		{"um", 1, 5, 5},
		{"positivo e negativo", 3, -4, -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado := Multiplicar(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("Multiplicar(%d, %d) = %d; esperado %d",
					tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}

// ========================================
// TESTE COM ERRO
// ========================================

func TestDividir(t *testing.T) {
	// Caso de sucesso
	resultado, err := Dividir(10, 2)
	if err != nil {
		t.Errorf("Dividir(10, 2) retornou erro: %v", err)
	}
	if resultado != 5.0 {
		t.Errorf("Dividir(10, 2) = %.2f; esperado 5.00", resultado)
	}

	// Caso de erro (divisão por zero)
	_, err = Dividir(10, 0)
	if err == nil {
		t.Error("Dividir(10, 0) deveria retornar erro")
	}
}

func TestDividir_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		esperado  float64
		deveErrar bool
	}{
		{"normal", 10, 2, 5, false},
		{"decimal", 10, 4, 2.5, false},
		{"divisão por zero", 10, 0, 0, true},
		{"zero dividido", 0, 5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado, err := Dividir(tt.a, tt.b)

			if tt.deveErrar {
				if err == nil {
					t.Errorf("Dividir(%.2f, %.2f) deveria retornar erro", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Dividir(%.2f, %.2f) erro inesperado: %v", tt.a, tt.b, err)
				}
				if resultado != tt.esperado {
					t.Errorf("Dividir(%.2f, %.2f) = %.2f; esperado %.2f",
						tt.a, tt.b, resultado, tt.esperado)
				}
			}
		})
	}
}

// ========================================
// TESTES MAIS COMPLEXOS
// ========================================

func TestFatorial(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		esperado  int
		deveErrar bool
	}{
		{"0", 0, 1, false},
		{"1", 1, 1, false},
		{"5", 5, 120, false},
		{"negativo", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado, err := Fatorial(tt.n)

			if tt.deveErrar {
				if err == nil {
					t.Errorf("Fatorial(%d) deveria retornar erro", tt.n)
				}
			} else {
				if err != nil {
					t.Errorf("Fatorial(%d) erro inesperado: %v", tt.n, err)
				}
				if resultado != tt.esperado {
					t.Errorf("Fatorial(%d) = %d; esperado %d", tt.n, resultado, tt.esperado)
				}
			}
		})
	}
}

func TestEhPrimo(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		esperado bool
	}{
		{"2 é primo", 2, true},
		{"3 é primo", 3, true},
		{"4 não é primo", 4, false},
		{"17 é primo", 17, true},
		{"20 não é primo", 20, false},
		{"1 não é primo", 1, false},
		{"0 não é primo", 0, false},
		{"negativo não é primo", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado := EhPrimo(tt.n)
			if resultado != tt.esperado {
				t.Errorf("EhPrimo(%d) = %v; esperado %v", tt.n, resultado, tt.esperado)
			}
		})
	}
}

/*
RESUMO DE TESTES:

ESTRUTURA:
    func TestNome(t *testing.T) {
        resultado := Funcao(args)
        if resultado != esperado {
            t.Errorf("mensagem com %v", resultado)
        }
    }

TABLE-DRIVEN TESTS:
    tests := []struct {
        name     string
        input    tipo
        esperado tipo
    }{
        {"caso1", input1, esperado1},
        {"caso2", input2, esperado2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // teste
        })
    }

MÉTODOS:
    t.Errorf()   - Reporta erro, continua
    t.Fatalf()   - Reporta erro, para teste
    t.Skip()     - Pula teste
    t.Run()      - Subteste

COMANDOS:
    go test              - Executar testes
    go test -v           - Verbose
    go test -cover       - Cobertura
    go test -run Nome    - Teste específico

DIFERENÇAS COM PYTHON:

Python (pytest):
✓ assert x == y
✓ Descoberta automática
✓ Fixtures
✓ Plugins abundantes
✓ pytest.mark.parametrize

Go (testing):
✓ Sem dependências externas
✓ Stdlib apenas
✓ Table-driven tests (manual)
✓ Subtestes com t.Run
✓ Rápido e simples

VANTAGENS GO:
✓ Zero dependências
✓ Muito rápido
✓ Integrado à linguagem
✓ Simples e direto

Execute os testes:
    go test
    go test -v
    go test -cover
*/
