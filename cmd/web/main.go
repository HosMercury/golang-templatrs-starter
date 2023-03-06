package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"snip/db"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	pool     *pgxpool.Pool
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool := db.Connect()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		pool:     pool,
	}

	r := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  r,
	}
	infoLog.Printf("Starting server on %s", *addr)

	defer pool.Close()

	srv.ListenAndServe()
}
