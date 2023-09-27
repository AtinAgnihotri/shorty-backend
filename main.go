package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/AtinAgnihotri/shorty-backend/internal/database"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ServerConf struct {
	PORT string
	DB   *database.Queries
}

func main() {
	serverConf := ServerConf{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Environment variables not found")
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	if len(port) == 0 || len(dbUrl) == 0 {
		log.Fatal("Unable to retrieve db url or port")
	}

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}

	dbQueries := database.New(db)
	serverConf.PORT = ":" + port
	serverConf.DB = dbQueries

	router := chi.NewRouter()

	router.Get("/{"+SHORT_URL_PARAM+"}", serverConf.RedirectHandler)

	server := &http.Server{
		Addr:    serverConf.PORT,
		Handler: router,
	}

	log.Printf("Listenting on port %v", serverConf.PORT)
	log.Fatal(server.ListenAndServe())

}
