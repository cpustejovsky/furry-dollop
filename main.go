package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/cpustejovsky/furry-dollop/models/psql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	errorLog *log.Logger
	infoLog  *log.Logger
)

type Config struct {
	Addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    interface {
		Insert(string, string, string) error
		Get(string) (*models.User, error)
		Update(string, string, string, string) (*models.User, error)
		Delete(string) error
	}
	posts interface {
		Insert(string, string, string) error
		GetAll() (*[]models.Post, error)
		GetById(string) (*models.Post, error)
		GetByUserId(string) (*[]models.Post, error)
	}
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

	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Llongfile)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
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
	infoLog.Println("Successfully connected to database!")

	app := &application{
		users: &psql.UserModel{DB: db},
		posts: &psql.PostModel{DB: db},
	}

	server := &http.Server{
		Handler: app.Routes(),
		Addr:    cfg.Addr,
	}
	infoLog.Printf("starting server on port %s", cfg.Addr)
	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatalf("could not listen on port %s\n%v", cfg.Addr, err)
	}
}
