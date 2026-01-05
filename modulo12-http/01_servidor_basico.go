package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mensagem struct {
	Texto string `json:"texto"`
	Autor string `json:"autor"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo ao servidor Go!")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	msg := Mensagem{
		Texto: "Olá do servidor Go!",
		Autor: "Sistema",
	}

	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/mensagem", jsonHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
Execute:
    go run 01_servidor_basico.go

Acesse:
    http://localhost:8080
    http://localhost:8080/api/mensagem
*/
