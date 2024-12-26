package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WhoDoIt/gofinal/payment_service/internal/app"
	"github.com/WhoDoIt/gofinal/payment_service/internal/dummy"
)

func main() {
	payment_port := os.Getenv("PAYMENT_PORT")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &app.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Dummy: &dummy.Dummy{
			Client:  &http.Client{},
			InfoLog: infoLog,
		},
	}

	app.InfoLog.Printf("HTTP server starting listening on :%s\n", payment_port)
	if err := http.ListenAndServe("payment_service:"+payment_port, app.Routes()); err != nil {
		app.ErrorLog.Fatalln(err)
	}
}
