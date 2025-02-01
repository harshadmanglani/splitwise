package main

import (
<<<<<<< HEAD
	"github.com/harshadmanglani/splitwise/jwt"
	"github.com/harshadmanglani/splitwise/models"
=======
	"github.com/harshadmanglani/whopays/jwt"
	"github.com/harshadmanglani/whopays/models"
>>>>>>> ca4e8dcb09fdce6af5de352209cec842475b85ed
	"github.com/jmoiron/sqlx"
	"github.com/knadh/stuffbin"
	_ "github.com/lib/pq"
)

type App struct {
	db      *sqlx.DB
	queries *models.Queries
	fs      stuffbin.FileSystem
	jwtg    *jwt.JwtGenerator
}

var (
	db      *sqlx.DB
	queries *models.Queries
	fs      stuffbin.FileSystem
	jwtg    *jwt.JwtGenerator
)

func init() {
	db = initDb()
	fs = initFs()
	queries = readAndPrepareQueries("./queries.sql", db.Unsafe(), fs)
	jwtg = initJwt()
}

func main() {
	app := &App{
		db:      db,
		queries: queries,
		fs:      fs,
		jwtg:    jwtg,
	}
	initHTTPServer(app)
}
