package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cpustejovsky/furry-dollop/models"
	"github.com/cpustejovsky/furry-dollop/models/psql"
	"github.com/golangcollege/sessions"
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

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	users    interface {
		Authenticate(string, string) (string, error)
		Insert(string, string, string, string) error
		Get(string) (*models.User, error)
		Update(string, string, string) (*models.User, error)
		Delete(string) error
	}
	posts interface {
		Insert(string, string, string) error
		GetAll() (*[]models.Post, error)
		GetById(string) (*models.Post, error)
		GetByUserId(string) (*[]models.Post, error)
		Update(string, string, string) (*models.Post, error)
		Delete(string) error
	}
}

func init() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Llongfile)

	if err := godotenv.Load(); err != nil {
		errorLog.Println("No .env file found")
	}
}

func main() {
	// Flag and Config Setup
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":5000", "HTTP network address")
	flag.Parse()

	// Environemntal Variables
	dbport, err := strconv.Atoi(os.Getenv("PSQL_PORT"))
	if err != nil {
		errorLog.Fatal(err)
	}
	var dbhost = os.Getenv("PSQL_HOST")
	var dbuser = os.Getenv("PSQL_USER")
	var dbname = os.Getenv("PSQL_DBNAME")
	var dbpassword = os.Getenv("PSQL_PW")
	var sessionSecret = []byte(os.Getenv("SESSION_SECRET"))

	// DB Setup
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable", dbhost, dbport, dbuser, dbpassword, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err != nil {
		panic(err)
	}
	infoLog.Println("Successfully connected to database!")

	// Session Setup
	session := sessions.New(sessionSecret)
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		session:  session,
		users:    &psql.UserModel{DB: db},
		posts:    &psql.PostModel{DB: db},
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
