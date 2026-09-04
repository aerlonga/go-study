// package main = pacote especial: só ele pode virar um programa executável.
// Todo binário Go tem exatamente UM pacote main, com UMA função main().
package main

import (
	"log"
	"net/http"

	"fase-1/internal/config"
	studentHandler "fase-1/internal/http/handlers/student"
	"fase-1/internal/storage"
)

func main() {
	cfg := config.MustLoad()

	store := storage.NewMemoryStorage()

	// http.NewServeMux() cria um "roteador": decide qual função chamar
	// dependendo do caminho E do método da URL.
	router := http.NewServeMux()

	// Desde Go 1.22, o padrão da rota pode incluir o método HTTP ("GET ", "POST " etc)
	// e parâmetros de caminho tipo "{id}" — antes disso só framework (Gin, Echo) tinha isso.
	// router.HandleFunc(padrão, handler): registra handler pra aquele padrão exato.
	router.HandleFunc("POST /api/students", studentHandler.New(store))
	router.HandleFunc("GET /api/students", studentHandler.GetStudents(store))
	router.HandleFunc("GET /api/students/{id}", studentHandler.GetStudentById(store))
	router.HandleFunc("PUT /api/students", studentHandler.UpdateStudent(store))
	router.HandleFunc("DELETE /api/students/{id}", studentHandler.DeleteStudent(store))

	log.Printf("servidor rodando em %s (env=%s)", cfg.HTTPServer.Address, cfg.Env)
	if err := http.ListenAndServe(cfg.HTTPServer.Address, router); err != nil {
		log.Fatal(err)
	}
}
