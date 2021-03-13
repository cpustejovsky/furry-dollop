package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cpustejovsky/furry-dollop/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Addr string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "furrydollop"
	password = "password"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Flag and Config Setup
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":5000", "HTTP network address")
	flag.Parse()

	// DB Setup
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database!")
	server := &http.Server{
		Handler: routes.Routes(),
		Addr:    ":5000",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
