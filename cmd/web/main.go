package main

import (
	"log"
	"net/http"
	"sprig/internal/bootstrap"
	"sprig/internal/config"
	"sprig/internal/router"

	_ "sprig/cmd/web/docs"
)

// @title		sprig-api
// @version	1.0
// @BasePath	/
func main() {
	cfg := config.LoadConfig()

	// mysql connection
	db, err := config.OpenConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	container := bootstrap.NewContainer(db)
	mux := router.New(container)

	server := http.Server{
		Addr:    ":8080", // containerization (cukup ambil portnya saja)
		Handler: mux,
	}

	log.Printf("Server berjalan di port %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
