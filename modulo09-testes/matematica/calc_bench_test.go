package matematica

import "testing"

/*
BENCHMARKS

Medir performance de funções.

COMPARAÇÃO COM PYTHON:
--------------------
Python (timeit):
    import timeit
    timeit.timeit('somar(2, 3)', number=1000000)

Go (testing):
    func BenchmarkSomar(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Somar(2, 3)
        }
    }

EXECUTAR:
    go test -bench .
    go test -bench . -benchmem    (com memória)
    go test -bench BenchmarkNome  (específico)
*/

// ========================================
// BENCHMARKS BÁSICOS
// ========================================

func BenchmarkSomar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Somar(10, 20)
	}
}

func BenchmarkMultiplicar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplicar(10, 20)
	}
}

func BenchmarkDividir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dividir(10, 2)
	}
}

func BenchmarkFatorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fatorial(10)
	}
}

func BenchmarkEhPrimo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EhPrimo(97)
	}
}

// ========================================
// BENCHMARKS COM SUB-BENCHMARKS
// ========================================

func BenchmarkFatorial_Varios(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"pequeno", 5},
		{"medio", 10},
		{"grande", 15},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fatorial(bm.n)
			}
		})
	}
}

func BenchmarkEhPrimo_Varios(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"pequeno_primo", 7},
		{"grande_primo", 997},
		{"pequeno_composto", 100},
		{"grande_composto", 1000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EhPrimo(bm.n)
			}
		})
	}
}

/*
RESUMO DE BENCHMARKS:

ESTRUTURA:
    func BenchmarkNome(b *testing.B) {
        for i := 0; i < b.N; i++ {
            FuncaoParaTestar()
        }
    }

EXECUTAR:
    go test -bench .
    go test -bench BenchmarkNome
    go test -bench . -benchmem       (memória)
    go test -bench . -benchtime=10s  (duração)

SUB-BENCHMARKS:
    b.Run(nome, func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            // código
        }
    })

MÉTODOS ÚTEIS:
    b.ResetTimer()     - Reinicia timer
    b.StopTimer()      - Para timer
    b.StartTimer()     - Inicia timer
    b.ReportAllocs()   - Reporta alocações

SAÍDA:
    BenchmarkSomar-8    500000000    3.21 ns/op

    - BenchmarkSomar: nome
    - 8: GOMAXPROCS
    - 500000000: iterações (b.N)
    - 3.21 ns/op: tempo por operação

Execute:
    go test -bench .
    go test -bench . -benchmem
*/
