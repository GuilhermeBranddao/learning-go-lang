# 🎯 Exercícios Práticos

Pratique seus conhecimentos em Go com estes exercícios!

## 📋 Como Usar

1. Leia o enunciado do exercício
2. Tente resolver sozinho primeiro
3. Se precisar, consulte os módulos anteriores
4. Execute com `go run exercicio_XX.go`
5. Compare com a solução (quando disponível)

---

## 🟢 Nível Básico

### Exercício 1: Calculadora
Crie uma calculadora que:
- Recebe dois números do usuário
- Recebe a operação (+, -, *, /)
- Exibe o resultado
- Trata divisão por zero

**Conceitos:** entrada/saída, switch, operadores

---

### Exercício 2: Números Pares
Crie um programa que:
- Leia 10 números do usuário
- Exiba apenas os números pares
- Exiba a soma dos números pares

**Conceitos:** loops, slices, condicionais

---

### Exercício 3: Tabela de Produtos
Crie um programa que:
- Armazene produtos em um map (nome → preço)
- Permita adicionar produtos
- Permita buscar preço de um produto
- Liste todos os produtos

**Conceitos:** maps, funções, loops

---

## 🟡 Nível Intermediário

### Exercício 4: Sistema de Notas
Crie um sistema de gerenciamento de notas que:
- Use uma struct `Aluno` (Nome, Notas []float64)
- Permita adicionar alunos
- Calcule a média de cada aluno
- Liste alunos aprovados (média >= 7.0)

**Conceitos:** structs, slices, funções, métodos

---

### Exercício 5: Agenda de Contatos
Crie uma agenda que:
- Armazene contatos (Nome, Telefone, Email)
- Permita adicionar, buscar e remover contatos
- Liste todos os contatos
- Busque por nome

**Conceitos:** structs, maps, funções

---

### Exercício 6: Estatísticas
Crie funções que calculem:
- Média de um slice de números
- Maior e menor valor
- Mediana
- Desvio padrão (desafio!)

**Conceitos:** funções, slices, matemática

---

## 🔴 Nível Avançado

### Exercício 7: Sistema Bancário
Implemente um sistema bancário com:
- Struct `ContaBancaria` (Número, Titular, Saldo)
- Métodos: Depositar, Sacar, Transferir
- Validações (saldo suficiente, valores positivos)
- Histórico de transações

**Conceitos:** structs, métodos, ponteiros, slices

---

### Exercício 8: Filtro e Map
Implemente funções genéricas:
- `Filtrar(slice, funcao)` - filtra elementos
- `Mapear(slice, funcao)` - transforma elementos
- `Reduzir(slice, funcao, inicial)` - reduz a um valor

**Conceitos:** funções de alta ordem, closures

---

### Exercício 9: Analisador de Texto
Crie um analisador que:
- Conte palavras em um texto
- Encontre a palavra mais frequente
- Liste palavras únicas
- Calcule comprimento médio das palavras

**Conceitos:** strings, maps, slices, funções

---

## 💡 Dicas

- **Comece pelo básico**: não pule exercícios
- **Teste seu código**: use diferentes entradas
- **Refatore**: depois de funcionar, melhore o código
- **Use `go fmt`**: mantenha o código formatado
- **Leia erros**: compilador do Go é muito útil!

---

## 📚 Recursos Adicionais

- [Documentação oficial do Go](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Tour of Go](https://go.dev/tour/)

Bons estudos! 🚀
