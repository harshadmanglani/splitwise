package main

import (
	"github.com/harshadmanglani/splitwise/jwt"
	"github.com/harshadmanglani/splitwise/models"
	"github.com/jmoiron/sqlx"
	"github.com/knadh/stuffbin"
	_ "github.com/lib/pq"
)

type App struct {
	db      *sqlx.DB
	queries *models.Queries
	fs      stuffbin.FileSystem
	jwt     *jwt.JwtManager
}

func newApp() *App {
	app := &App{
		db:  initDb(),
		fs:  initFs(),
		jwt: initJwt(),
	}
	app.queries = readAndPrepareQueries("./queries.sql", app.db.Unsafe(), app.fs)
	return app
}

func main() {
	app := newApp()
	initHTTPServer(app)
}
