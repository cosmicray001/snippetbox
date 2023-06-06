package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()
	infoLogFile, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer infoLogFile.Close()
	infoLog := log.New(infoLogFile, "INFO\t", log.Ldate|log.Ltime)
	errorLogFile, err := os.OpenFile("./tmp/error.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer errorLogFile.Close()
	errorLog := log.New(errorLogFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routers(),
	}
	infoLog.Printf("Starting server on : %s\n", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
