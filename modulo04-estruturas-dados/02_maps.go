package main

import "fmt"

/*
MAPS (Dicionários)

COMPARAÇÃO COM PYTHON:
--------------------
Python:
    pessoa = {
        "nome": "Maria",
        "idade": 25,
        "cidade": "SP"
    }

    pessoa["nome"]  # acessar
    pessoa["email"] = "maria@email.com"  # adicionar
    del pessoa["cidade"]  # remover

Go:
    pessoa := map[string]string{
        "nome": "Maria",
        "cidade": "SP",
    }

    pessoa["nome"]  // acessar
    pessoa["email"] = "maria@email.com"  // adicionar
    delete(pessoa, "cidade")  // remover
*/

func main() {
	fmt.Println("=== CRIANDO MAPS ===\n")

	// Forma 1: map literal vazio
	var cores map[string]string
	fmt.Printf("Map vazio: %v (nil: %t)\n", cores, cores == nil)

	// Forma 2: make
	cores = make(map[string]string)
	cores["vermelho"] = "red"
	cores["azul"] = "blue"
	fmt.Printf("Cores: %v\n", cores)

	// Forma 3: map literal com valores
	capitais := map[string]string{
		"Brasil":    "Brasília",
		"Argentina": "Buenos Aires",
		"Chile":     "Santiago",
		"Peru":      "Lima",
	}
	fmt.Printf("Capitais: %v\n\n", capitais)

	fmt.Println("=== ACESSANDO VALORES ===\n")

	// Acessar valor
	fmt.Printf("Capital do Brasil: %s\n", capitais["Brasil"])

	// IMPORTANTE: verificar se chave existe
	capital, existe := capitais["Uruguai"]
	if existe {
		fmt.Printf("Capital do Uruguai: %s\n", capital)
	} else {
		fmt.Println("Uruguai não está no map")
	}

	// Sem verificação (retorna zero value se não existir)
	fmt.Printf("Bolívia: '%s' (string vazia)\n\n", capitais["Bolívia"])

	fmt.Println("=== ADICIONANDO E MODIFICANDO ===\n")

	// Adicionar
	capitais["Uruguai"] = "Montevidéu"
	fmt.Printf("Capitais após adicionar: %v\n", capitais)

	// Modificar
	capitais["Brasil"] = "Rio de Janeiro"
	fmt.Printf("Brasil modificado: %s\n", capitais["Brasil"])

	// Corrigir
	capitais["Brasil"] = "Brasília"
	fmt.Printf("Brasil corrigido: %s\n\n", capitais["Brasil"])

	fmt.Println("=== REMOVENDO ELEMENTOS ===\n")

	delete(capitais, "Peru")
	fmt.Printf("Após remover Peru: %v\n\n", capitais)

	fmt.Println("=== ITERANDO SOBRE MAPS ===\n")

	fmt.Println("Países e capitais:")
	for pais, capital := range capitais {
		fmt.Printf("  %s -> %s\n", pais, capital)
	}

	// ATENÇÃO: ordem não é garantida!
	fmt.Println("\nIterando novamente (ordem pode mudar):")
	for pais, capital := range capitais {
		fmt.Printf("  %s -> %s\n", pais, capital)
	}

	fmt.Println("\n=== MAPS COM DIFERENTES TIPOS ===\n")

	// Map com valores int
	idades := map[string]int{
		"João":  25,
		"Maria": 30,
		"Pedro": 28,
	}
	fmt.Printf("Idades: %v\n", idades)

	// Map com valores bool
	presentes := map[string]bool{
		"João":  true,
		"Maria": false,
		"Pedro": true,
	}
	fmt.Printf("Presentes: %v\n", presentes)

	// Map com valores slice
	grupos := map[string][]string{
		"Backend":  {"Go", "Python", "Java"},
		"Frontend": {"JavaScript", "TypeScript", "React"},
		"Mobile":   {"Flutter", "React Native", "Swift"},
	}
	fmt.Printf("Grupos: %v\n\n", grupos)

	fmt.Println("=== MAPS ANINHADOS ===\n")

	// Map de maps
	pessoas := map[string]map[string]string{
		"pessoa1": {
			"nome":   "João",
			"cidade": "SP",
		},
		"pessoa2": {
			"nome":   "Maria",
			"cidade": "RJ",
		},
	}

	fmt.Println("Pessoas:")
	for id, dados := range pessoas {
		fmt.Printf("  %s: %s de %s\n", id, dados["nome"], dados["cidade"])
	}

	fmt.Println("\n=== LEN E VERIFICAÇÕES ===\n")

	fmt.Printf("Número de capitais: %d\n", len(capitais))
	fmt.Printf("Número de idades: %d\n", len(idades))

	// Verificar se map está vazio
	var vazio map[string]int
	fmt.Printf("Map vazio é nil? %t\n", vazio == nil)
	fmt.Printf("Map vazio len: %d\n", len(vazio))

	inicializado := make(map[string]int)
	fmt.Printf("Map inicializado é nil? %t\n", inicializado == nil)
	fmt.Printf("Map inicializado len: %d\n\n", len(inicializado))

	fmt.Println("=== EXEMPLOS PRÁTICOS ===\n")

	// Contar ocorrências de palavras
	texto := []string{"go", "python", "go", "java", "python", "go"}
	contagem := make(map[string]int)

	for _, palavra := range texto {
		contagem[palavra]++
	}

	fmt.Println("Contagem de palavras:")
	for palavra, count := range contagem {
		fmt.Printf("  %s: %d\n", palavra, count)
	}

	// Agrupar por categoria
	produtos := map[string]string{
		"Notebook": "Eletrônicos",
		"Mouse":    "Eletrônicos",
		"Cadeira":  "Móveis",
		"Mesa":     "Móveis",
		"Teclado":  "Eletrônicos",
	}

	categorias := make(map[string][]string)
	for produto, categoria := range produtos {
		categorias[categoria] = append(categorias[categoria], produto)
	}

	fmt.Println("\nProdutos por categoria:")
	for categoria, prods := range categorias {
		fmt.Printf("  %s: %v\n", categoria, prods)
	}

	// Cache/Memorização
	fmt.Println("\nCache de resultados:")
	cache := make(map[int]int)

	calcularFibonacci := func(n int) int {
		// Verificar cache
		if val, existe := cache[n]; existe {
			fmt.Printf("  Fibonacci(%d) do cache: %d\n", n, val)
			return val
		}

		// Calcular
		var resultado int
		if n <= 1 {
			resultado = n
		} else {
			resultado = calcularFibonacci(n-1) + calcularFibonacci(n-2)
		}

		// Salvar no cache
		cache[n] = resultado
		fmt.Printf("  Fibonacci(%d) calculado: %d\n", n, resultado)
		return resultado
	}

	calcularFibonacci(5)
	calcularFibonacci(6) // Reutiliza cálculos anteriores
}

/*
RESUMO DE MAPS:

CRIAR:
    var m map[K]V              // nil map
    m := make(map[K]V)         // map vazio
    m := map[K]V{              // map com valores
        "chave": valor,
    }

OPERAÇÕES:
    m[chave] = valor           // adicionar/modificar
    val := m[chave]            // acessar
    val, ok := m[chave]        // acessar com verificação
    delete(m, chave)           // remover
    len(m)                     // tamanho

ITERAR:
    for chave, valor := range m {
        // usar chave e valor
    }

IMPORTANTE:
- Maps são REFERENCE TYPES (passados por referência)
- Ordem NÃO é garantida
- Chaves devem ser comparáveis (não pode ser slice, map ou função)
- Valores podem ser de qualquer tipo
- Map nil causa panic ao adicionar (use make)

Execute com:
    go run 02_maps.go
*/
