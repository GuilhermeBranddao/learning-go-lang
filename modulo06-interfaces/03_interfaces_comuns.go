package main

import (
	"fmt"
	"io"
	"strings"
)

/*
INTERFACES COMUNS DA STANDARD LIBRARY

Go tem várias interfaces que são usadas em toda a biblioteca padrão.
Implementar essas interfaces torna seus tipos compatíveis com muitas funções!

PRINCIPAIS:
- io.Reader
- io.Writer
- fmt.Stringer
- error
- sort.Interface
*/

func main() {
	fmt.Println("=== io.Reader e io.Writer ===\n")

	// io.Reader: type Reader interface { Read(p []byte) (n int, err error) }
	// io.Writer: type Writer interface { Write(p []byte) (n int, err error) }

	// strings.Reader implementa io.Reader
	reader := strings.NewReader("Hello from Reader!")

	// io.Copy usa io.Reader e io.Writer
	// Nota: não vamos escrever em stdout aqui, apenas demonstrar
	dados := make([]byte, 20)
	n, err := reader.Read(dados)
	if err != nil && err != io.EOF {
		fmt.Println("Erro:", err)
	}
	fmt.Printf("Lidos %d bytes: %s\n", n, string(dados[:n]))

	fmt.Println("\n=== fmt.Stringer ===\n")

	// fmt.Stringer: type Stringer interface { String() string }
	// Usado por fmt.Print, fmt.Println, etc.

	pessoa := Pessoa{Nome: "Maria", Idade: 30}
	fmt.Println(pessoa) // Usa String() automaticamente

	fmt.Println("\n=== EXEMPLO: CUSTOM READER ===\n")

	// Criar nosso próprio Reader
	zeroes := ZeroReader{limite: 10}
	buffer := make([]byte, 5)

	for {
		n, err := zeroes.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Lidos %d bytes: %v\n", n, buffer[:n])
	}

	fmt.Println("\n=== EXEMPLO: CONTADOR DE ESCRITAS ===\n")

	contador := &ContadorWriter{}

	fmt.Fprintf(contador, "Primeira linha\n")
	fmt.Fprintf(contador, "Segunda linha\n")
	fmt.Fprintf(contador, "Terceira linha\n")

	fmt.Printf("Total de bytes escritos: %d\n", contador.Total)

	fmt.Println("\n=== COMPOSIÇÃO DE INTERFACES ===\n")

	// io.ReadWriter compõe Reader e Writer
	exemploComposicao()
}

// ========================================
// fmt.Stringer
// ========================================

type Pessoa struct {
	Nome  string
	Idade int
}

// Implementar fmt.Stringer
func (p Pessoa) String() string {
	return fmt.Sprintf("%s (%d anos)", p.Nome, p.Idade)
}

// ========================================
// Custom io.Reader
// ========================================

// ZeroReader retorna apenas zeros
type ZeroReader struct {
	contador int
	limite   int
}

func (z *ZeroReader) Read(p []byte) (n int, err error) {
	if z.contador >= z.limite {
		return 0, io.EOF
	}

	// Preencher com zeros
	for i := range p {
		if z.contador >= z.limite {
			return i, io.EOF
		}
		p[i] = 0
		z.contador++
	}

	return len(p), nil
}

// ========================================
// Custom io.Writer
// ========================================

// ContadorWriter conta bytes escritos
type ContadorWriter struct {
	Total int
}

func (c *ContadorWriter) Write(p []byte) (n int, err error) {
	c.Total += len(p)
	return len(p), nil
}

// ========================================
// Composição de Interfaces
// ========================================

// Interface que compõe outras
type ReadWriter interface {
	io.Reader
	io.Writer
}

// Tipo que implementa ambas
type Buffer struct {
	dados []byte
	pos   int
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.pos >= len(b.dados) {
		return 0, io.EOF
	}

	n = copy(p, b.dados[b.pos:])
	b.pos += n
	return n, nil
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.dados = append(b.dados, p...)
	return len(p), nil
}

func exemploComposicao() {
	buf := &Buffer{}

	// Usar como Writer
	fmt.Fprintf(buf, "Hello, ")
	fmt.Fprintf(buf, "World!")

	// Usar como Reader
	resultado := make([]byte, 100)
	n, _ := buf.Read(resultado)
	fmt.Printf("Buffer contém: %s\n", string(resultado[:n]))
}

/*
INTERFACES COMUNS:

1. io.Reader
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   - Ler dados de qualquer fonte
   - Usado em: arquivos, rede, strings, etc.

2. io.Writer
   type Writer interface {
       Write(p []byte) (n int, err error)
   }
   - Escrever dados para qualquer destino
   - Usado em: arquivos, rede, buffers, etc.

3. fmt.Stringer
   type Stringer interface {
       String() string
   }
   - Conversão customizada para string
   - Usado por: fmt.Print*, log.Print*, etc.

4. error
   type error interface {
       Error() string
   }
   - Representar erros
   - Veremos no próximo módulo!

5. sort.Interface
   type Interface interface {
       Len() int
       Less(i, j int) bool
       Swap(i, j int)
   }
   - Ordenação customizada

VANTAGENS:
✓ Interoperabilidade com stdlib
✓ Código reutilizável
✓ Padrões bem conhecidos
✓ Composição simples

PADRÃO DE NOMENCLATURA:
- 1 método: sufixo "er" (Reader, Writer, Closer)
- 2+ métodos: nome descritivo

Execute com:
    go run 03_interfaces_comuns.go
*/
