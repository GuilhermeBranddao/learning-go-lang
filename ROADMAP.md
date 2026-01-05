# 🗺️ Roadmap de Aprendizado - Go para Desenvolvedores Python

## ✅ Módulos Atuais (Completos)

### 📘 [Módulo 01 - Fundamentos](modulo01-fundamentos/)
- ✓ Hello World e estrutura básica
- ✓ Variáveis e constantes
- ✓ Tipos de dados
- ✓ Operadores
- ✓ Entrada e saída

### 📘 [Módulo 02 - Estruturas de Controle](modulo02-controle/)
- ✓ If/Else
- ✓ Loops (for)
- ✓ Switch

### 📘 [Módulo 03 - Funções](modulo03-funcoes/)
- ✓ Funções básicas e múltiplos retornos
- ✓ Funções variádicas
- ✓ Defer
- ✓ Funções anônimas e closures

### 📘 [Módulo 04 - Estruturas de Dados](modulo04-estruturas-dados/)
- ✓ Arrays e Slices
- ✓ Maps
- ✓ Structs
- ✓ Ponteiros

### 📘 [Módulo 05 - Goroutines e Concorrência](modulo05-goroutines/)
- ✓ Goroutines básicas
- ✓ Channels
- ✓ Select
- ✓ WaitGroup e Mutex
- ✓ Padrões de concorrência

### 🎯 [Exercícios Práticos](exercicios/)
- ✓ Calculadora
- ✓ Sistema de notas
- ✓ Lista de desafios

---

## 🚀 Próximos Módulos (Sugestões)

### 📘 Módulo 06 - Interfaces e Polimorfismo
**Por que importante:** Interfaces são fundamentais em Go
- [ ] O que são interfaces
- [ ] Interfaces implícitas
- [ ] Polimorfismo em Go
- [ ] Interface vazia (interface{})
- [ ] Type assertions e type switches
- [ ] Interfaces comuns (io.Reader, io.Writer, etc.)
- [ ] Comparação com classes abstratas Python

### 📘 Módulo 07 - Tratamento de Erros
**Por que importante:** Go tem abordagem única para erros
- [ ] Error interface
- [ ] Criar erros personalizados
- [ ] Error wrapping (Go 1.13+)
- [ ] Panic e Recover
- [ ] Quando usar panic vs error
- [ ] Best practices de error handling
- [ ] Comparação com try/except Python

### 📘 Módulo 08 - Packages e Módulos
**Por que importante:** Organizar código profissionalmente
- [ ] Estrutura de packages
- [ ] go.mod e go.sum
- [ ] Imports e exports (maiúsculas/minúsculas)
- [ ] Packages internos
- [ ] Versionamento semântico
- [ ] Publicar packages
- [ ] Comparação com módulos Python

### 📘 Módulo 09 - Testes
**Por que importante:** Go tem excelente suporte a testes
- [ ] Package testing
- [ ] Testes unitários
- [ ] Table-driven tests
- [ ] Benchmarks
- [ ] Coverage
- [ ] Mocks e stubs
- [ ] Comparação com pytest/unittest

### 📘 Módulo 10 - JSON e Serialização
**Por que importante:** Essencial para APIs e dados
- [ ] encoding/json
- [ ] Marshal e Unmarshal
- [ ] Struct tags
- [ ] JSON personalizado
- [ ] XML, CSV, outros formatos
- [ ] Comparação com json Python

### 📘 Módulo 11 - Trabalhando com Arquivos
**Por que importante:** I/O é comum em aplicações
- [ ] Ler e escrever arquivos
- [ ] Package os e io
- [ ] bufio para leitura eficiente
- [ ] Caminhos e diretórios
- [ ] Permissões e metadados
- [ ] Comparação com open() Python

### 📘 Módulo 12 - HTTP e APIs REST
**Por que importante:** Go é excelente para web services
- [ ] net/http básico
- [ ] Criar servidor HTTP
- [ ] Handlers e rotas
- [ ] Middlewares
- [ ] Cliente HTTP
- [ ] Frameworks (gin, echo, chi)
- [ ] Comparação com Flask/FastAPI

### 📘 Módulo 13 - Banco de Dados
**Por que importante:** Persistência de dados
- [ ] database/sql
- [ ] Drivers (PostgreSQL, MySQL)
- [ ] CRUD operations
- [ ] Prepared statements
- [ ] Transactions
- [ ] ORMs (GORM, sqlx)
- [ ] Comparação com SQLAlchemy

### 📘 Módulo 14 - Context
**Por que importante:** Controle de requisições e timeouts
- [ ] context.Context
- [ ] Cancelamento
- [ ] Timeouts
- [ ] Propagação de valores
- [ ] Best practices

### 📘 Módulo 15 - Reflection e Generics
**Por que importante:** Recursos avançados
- [ ] Package reflect
- [ ] Type introspection
- [ ] Generics (Go 1.18+)
- [ ] Type constraints
- [ ] Quando usar/evitar

---

## 🎓 Projetos Práticos Sugeridos

### 🟢 Nível Iniciante
1. **CLI Calculator** - Calculadora de linha de comando
2. **File Organizer** - Organizador de arquivos por extensão
3. **Todo List** - Lista de tarefas com persistência
4. **URL Shortener** - Encurtador de URLs simples
5. **Log Parser** - Analisador de arquivos de log

### 🟡 Nível Intermediário
6. **REST API** - API RESTful completa com CRUD
7. **Web Scraper** - Coletor de dados web com goroutines
8. **Chat Server** - Servidor de chat com websockets
9. **Task Queue** - Sistema de filas com workers
10. **Microservice** - Microsserviço com banco de dados

### 🔴 Nível Avançado
11. **Load Balancer** - Balanceador de carga HTTP
12. **Distributed Cache** - Cache distribuído
13. **Search Engine** - Motor de busca simples
14. **API Gateway** - Gateway para microsserviços
15. **Monitoring System** - Sistema de monitoramento

---

## 📚 Recursos Complementares (Criar)

### Cheat Sheets
- [ ] Sintaxe básica Go vs Python
- [ ] Comandos Go CLI
- [ ] Padrões de concorrência
- [ ] Error handling patterns
- [ ] Performance tips

### Guias de Referência
- [ ] Standard Library overview
- [ ] Pacotes mais usados
- [ ] Design patterns em Go
- [ ] Code review checklist
- [ ] Debugging em Go

### Comparações Detalhadas
- [ ] Go vs Python: Performance
- [ ] Go vs Python: Casos de uso
- [ ] Go vs Python: Ecossistema
- [ ] Quando usar Go vs Python
- [ ] Migração de Python para Go

---

## 🎯 Tópicos Especiais

### Performance e Otimização
- [ ] Profiling (pprof)
- [ ] Benchmarking
- [ ] Memory optimization
- [ ] Garbage collector
- [ ] Build optimization

### Ferramentas e Ecosystem
- [ ] go fmt, go vet, go lint
- [ ] Delve (debugger)
- [ ] VS Code setup
- [ ] CI/CD com Go
- [ ] Docker e Go

### Best Practices
- [ ] Effective Go review
- [ ] Code organization
- [ ] Error handling patterns
- [ ] Concurrency patterns
- [ ] Security best practices

---

## 📊 Ordem Sugerida de Estudo

### Fase 1: Fundamentos (Atual - Completa ✓)
1. Módulos 01-05
2. Exercícios básicos

### Fase 2: Intermediário (Próxima)
3. Módulo 06 - Interfaces
4. Módulo 07 - Erros
5. Módulo 08 - Packages
6. Módulo 09 - Testes
7. Projetos nível iniciante

### Fase 3: Prático
8. Módulo 10 - JSON
9. Módulo 11 - Arquivos
10. Módulo 12 - HTTP/APIs
11. Módulo 13 - Banco de Dados
12. Projetos nível intermediário

### Fase 4: Avançado
13. Módulo 14 - Context
14. Módulo 15 - Reflection/Generics
15. Tópicos especiais
16. Projetos nível avançado

---

## 💡 Como Contribuir com Ideias

Quer sugerir mais conteúdo? Considere:
- ✅ Tópicos que você gostaria de aprender
- ✅ Projetos práticos interessantes
- ✅ Comparações específicas com Python
- ✅ Dúvidas comuns na transição
- ✅ Casos de uso reais

---

**Última atualização:** Janeiro 2026
