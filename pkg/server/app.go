package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	Router *http.ServeMux
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error

	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = http.NewServeMux()
	loadRoutes(app.Router)
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
