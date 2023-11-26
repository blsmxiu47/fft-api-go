package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Listener interface {
	Listen(context.Context) error
}

type App struct {
	Router* mux.Router
	DB* sql.DB
}

type OnStart func(context.Context, *App) ([]Listener, error)

func start(onStart OnStart) {
	// a := App{}
	// TODO: handle app startup
	// TODO: gracefully handle shutdown
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString :=
        fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":8010", a.Router))
}
