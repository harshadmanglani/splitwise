package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/harshadmanglani/splitwise/models"
	"github.com/jmoiron/sqlx"
	"github.com/knadh/goyesql/v2"
	goyesqlx "github.com/knadh/goyesql/v2/sqlx"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type App struct {
	db      *sqlx.DB
	queries models.Queries
}

func initDb() *sqlx.DB {
	ko := koanf.New(".")
	ko.Load(file.Provider("config.yaml"), yaml.Parser())

	var c struct {
		Host        string        `koanf:"host"`
		Port        int           `koanf:"port"`
		User        string        `koanf:"user"`
		Password    string        `koanf:"password"`
		DBName      string        `koanf:"database"`
		SSLMode     string        `koanf:"ssl_mode"`
		Params      string        `koanf:"params"`
		MaxOpen     int           `koanf:"max_open"`
		MaxIdle     int           `koanf:"max_idle"`
		MaxLifetime time.Duration `koanf:"max_lifetime"`
	}

	if err := ko.Unmarshal("db", &c); err != nil {
		fmt.Printf("error loading db config: %v", err)
	}

	fmt.Printf(c.Host)
	fmt.Printf("connecting to db: %s:%d/%s", c.Host, c.Port, c.DBName)
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s %s", c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode, c.Params))
	if err != nil {
		fmt.Printf("error connecting to DB: %v", err)
	}

	db.SetMaxOpenConns(c.MaxOpen)
	db.SetMaxIdleConns(c.MaxIdle)
	db.SetConnMaxLifetime(c.MaxLifetime)

	return db
}

func main() {
	qMap := goyesql.MustParseFile("sql/queries.sql")
	var app *App = new(App)
	app.db = initDb()
	if err := goyesqlx.ScanToStruct(&app.queries, qMap, app.db.Unsafe()); err != nil {
		fmt.Errorf("Query preparation failed")
		panic(err)
	}
	// app.db.MustExec(app.queries.InsertUser())
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, there!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
