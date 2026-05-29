package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.CriaUsuario).Methods(http.MethodPost)
	router.HandleFunc("/users", server.BuscaUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.BuscaUsuarioPorId).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.AtualizaUsuario).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeletaUsuario).Methods(http.MethodDelete)

	fmt.Println("Servidor rodando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
