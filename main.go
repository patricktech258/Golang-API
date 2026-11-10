package main

import (
	"example/main/database"
	"example/main/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Banco inicializado!")

	mux := routes.SetupRoutes(db)

	fmt.Println("Rodando a API em Go na porta: 8000")

	if err := http.ListenAndServe(":8000", mux); err != nil {
		fmt.Print("Erro ao iniciar o servidor")
	}
}
