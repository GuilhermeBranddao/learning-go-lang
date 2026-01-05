package main

import "fmt"

/*
EXERCÍCIO 4: SISTEMA DE NOTAS

Crie um sistema de gerenciamento de notas com:
1. Struct Aluno (Nome, Notas []float64)
2. Funções para adicionar alunos
3. Calcular média de cada aluno
4. Listar alunos aprovados (média >= 7.0)
5. Encontrar aluno com maior média
*/

// Struct para representar um aluno
type Aluno struct {
	Nome  string
	Notas []float64
}

// Método para calcular média
func (a Aluno) Media() float64 {
	if len(a.Notas) == 0 {
		return 0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}
	return soma / float64(len(a.Notas))
}

// Método para verificar se está aprovado
func (a Aluno) Aprovado() bool {
	return a.Media() >= 7.0
}

// Método para exibir informações
func (a Aluno) Exibir() {
	status := "❌ Reprovado"
	if a.Aprovado() {
		status = "✓ Aprovado"
	}
	fmt.Printf("  %s - Média: %.2f - %s\n", a.Nome, a.Media(), status)
}

// Sistema de gerenciamento
type SistemaNotas struct {
	Alunos []Aluno
}

// Adicionar aluno
func (s *SistemaNotas) AdicionarAluno(nome string, notas ...float64) {
	aluno := Aluno{
		Nome:  nome,
		Notas: notas,
	}
	s.Alunos = append(s.Alunos, aluno)
	fmt.Printf("✓ Aluno %s adicionado\n", nome)
}

// Listar todos os alunos
func (s SistemaNotas) ListarAlunos() {
	fmt.Println("\n=== TODOS OS ALUNOS ===")
	for _, aluno := range s.Alunos {
		aluno.Exibir()
	}
}

// Listar apenas aprovados
func (s SistemaNotas) ListarAprovados() {
	fmt.Println("\n=== ALUNOS APROVADOS ===")
	aprovados := 0
	for _, aluno := range s.Alunos {
		if aluno.Aprovado() {
			aluno.Exibir()
			aprovados++
		}
	}
	if aprovados == 0 {
		fmt.Println("  Nenhum aluno aprovado")
	}
}

// Encontrar aluno com maior média
func (s SistemaNotas) MelhorAluno() *Aluno {
	if len(s.Alunos) == 0 {
		return nil
	}

	melhor := &s.Alunos[0]
	for i := range s.Alunos {
		if s.Alunos[i].Media() > melhor.Media() {
			melhor = &s.Alunos[i]
		}
	}
	return melhor
}

// Calcular média geral da turma
func (s SistemaNotas) MediaGeral() float64 {
	if len(s.Alunos) == 0 {
		return 0
	}

	soma := 0.0
	for _, aluno := range s.Alunos {
		soma += aluno.Media()
	}
	return soma / float64(len(s.Alunos))
}

func main() {
	fmt.Println("=== SISTEMA DE NOTAS ===\n")

	// Criar sistema
	sistema := SistemaNotas{}

	// Adicionar alunos
	sistema.AdicionarAluno("João Silva", 8.5, 9.0, 7.5, 8.0)
	sistema.AdicionarAluno("Maria Santos", 9.5, 10.0, 9.0, 9.5)
	sistema.AdicionarAluno("Pedro Oliveira", 6.0, 5.5, 6.5, 5.0)
	sistema.AdicionarAluno("Ana Costa", 7.0, 7.5, 7.0, 8.0)
	sistema.AdicionarAluno("Carlos Souza", 5.0, 6.0, 4.5, 5.5)

	// Listar todos
	sistema.ListarAlunos()

	// Listar aprovados
	sistema.ListarAprovados()

	// Melhor aluno
	fmt.Println("\n=== MELHOR ALUNO ===")
	melhor := sistema.MelhorAluno()
	if melhor != nil {
		melhor.Exibir()
	}

	// Média geral
	fmt.Printf("\n=== ESTATÍSTICAS ===\n")
	fmt.Printf("Média geral da turma: %.2f\n", sistema.MediaGeral())
	fmt.Printf("Total de alunos: %d\n", len(sistema.Alunos))

	// Contar aprovados e reprovados
	aprovados := 0
	for _, aluno := range sistema.Alunos {
		if aluno.Aprovado() {
			aprovados++
		}
	}
	fmt.Printf("Aprovados: %d (%.1f%%)\n", aprovados,
		float64(aprovados)/float64(len(sistema.Alunos))*100)
	fmt.Printf("Reprovados: %d (%.1f%%)\n", len(sistema.Alunos)-aprovados,
		float64(len(sistema.Alunos)-aprovados)/float64(len(sistema.Alunos))*100)
}

/*
EXEMPLO DE SAÍDA:

=== SISTEMA DE NOTAS ===

✓ Aluno João Silva adicionado
✓ Aluno Maria Santos adicionado
✓ Aluno Pedro Oliveira adicionado
✓ Aluno Ana Costa adicionado
✓ Aluno Carlos Souza adicionado

=== TODOS OS ALUNOS ===
  João Silva - Média: 8.25 - ✓ Aprovado
  Maria Santos - Média: 9.50 - ✓ Aprovado
  Pedro Oliveira - Média: 5.75 - ❌ Reprovado
  Ana Costa - Média: 7.38 - ✓ Aprovado
  Carlos Souza - Média: 5.25 - ❌ Reprovado

=== ALUNOS APROVADOS ===
  João Silva - Média: 8.25 - ✓ Aprovado
  Maria Santos - Média: 9.50 - ✓ Aprovado
  Ana Costa - Média: 7.38 - ✓ Aprovado

=== MELHOR ALUNO ===
  Maria Santos - Média: 9.50 - ✓ Aprovado

=== ESTATÍSTICAS ===
Média geral da turma: 7.23
Total de alunos: 5
Aprovados: 3 (60.0%)
Reprovados: 2 (40.0%)

PONTOS DE APRENDIZADO:
- Structs para modelar dados
- Métodos com receivers
- Slices de structs
- Ponteiros para structs
- Funções variádicas
- Cálculos estatísticos

Execute com:
    go run exercicio04_notas.go
*/
