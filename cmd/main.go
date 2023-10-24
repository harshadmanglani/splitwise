package main

import (
	"github.com/harshadmanglani/splitwise/models"
	"github.com/jmoiron/sqlx"
	"github.com/knadh/stuffbin"
	_ "github.com/lib/pq"
)

type App struct {
	db      *sqlx.DB
	queries *models.Queries
	fs      stuffbin.FileSystem
}

var (
	db      *sqlx.DB
	queries *models.Queries
	fs      stuffbin.FileSystem
)

func init() {
	db = initDb()
	fs = initFs()
	queries = readAndPrepareQueries("./queries.sql", db.Unsafe(), fs)
}

func main() {
	app := &App{
		db:      db,
		queries: queries,
		fs:      fs,
	}
	initHTTPServer(app)
}
