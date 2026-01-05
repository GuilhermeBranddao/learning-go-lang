package main

import (
	"fmt"
	"reflect"
)

type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

func main() {
	fmt.Println("=== REFLECTION BÁSICO ===\n")

	p := Pessoa{
		Nome:  "João",
		Idade: 30,
		Email: "joao@email.com",
	}

	// Obter tipo
	t := reflect.TypeOf(p)
	fmt.Println("Tipo:", t)
	fmt.Println("Nome do tipo:", t.Name())
	fmt.Println("Tipo de tipo:", t.Kind())

	// Obter valor
	v := reflect.ValueOf(p)
	fmt.Println("\nValor:", v)

	fmt.Println("\n=== INSPECIONAR CAMPOS ===\n")

	// Iterar sobre campos
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("%s: %v (tipo: %s)\n",
			field.Name, value, field.Type)
	}

	fmt.Println("\n=== MODIFICAR VALORES ===\n")

	// Para modificar, precisa de ponteiro
	pPtr := &p
	v = reflect.ValueOf(pPtr).Elem()

	// Modificar campo
	campoNome := v.FieldByName("Nome")
	if campoNome.CanSet() {
		campoNome.SetString("Maria")
		fmt.Println("Nome modificado para:", p.Nome)
	}

	campoIdade := v.FieldByName("Idade")
	if campoIdade.CanSet() {
		campoIdade.SetInt(25)
		fmt.Println("Idade modificada para:", p.Idade)
	}

	fmt.Println("\n=== REFLECTION COM INTERFACE{} ===\n")

	inspecionar(42)
	inspecionar("texto")
	inspecionar(3.14)
	inspecionar([]int{1, 2, 3})
	inspecionar(map[string]int{"a": 1})
}

func inspecionar(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("Tipo: %v, Kind: %v, Valor: %v\n", t, t.Kind(), val)
}

/*
RESUMO DE REFLECTION:

OBTER INFORMAÇÕES:
    reflect.TypeOf(obj)     - Tipo
    reflect.ValueOf(obj)    - Valor

TYPE METHODS:
    .Name()        - Nome do tipo
    .Kind()        - Tipo base (struct, slice, etc)
    .NumField()    - Número de campos (struct)
    .Field(i)      - Campo na posição i

VALUE METHODS:
    .Interface()   - Converter para interface{}
    .Int()         - Obter como int
    .String()      - Obter como string
    .CanSet()      - Pode modificar?
    .Set*()        - Modificar valor

QUANDO USAR:
⚠️ Reflection é lento e complexo
✓ Serialização/deserialização
✓ Frameworks (ORM, validação)
✓ Debugging/inspeção
✗ Código normal (prefira generics!)

Execute:
    go run 02_reflection.go
*/
