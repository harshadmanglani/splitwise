package main

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/harshadmanglani/splitwise/jwt"
	"github.com/harshadmanglani/splitwise/models"
	"github.com/jmoiron/sqlx"
	"github.com/knadh/goyesql/v2"
	goyesqlx "github.com/knadh/goyesql/v2/sqlx"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/stuffbin"
	"github.com/labstack/echo/v4"
)

var (
	ko              = koanf.New(".")
	appDir   string = "../"
	appFiles        = []string{
		"./sql/queries.sql:queries.sql",
		"./sql/schema.sql:schema.sql",
	}
)

func init() {
	ko.Load(file.Provider("../config.yaml"), yaml.Parser())
}

func initDb() *sqlx.DB {
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

	fmt.Printf("connecting to db: %s:%d/%s\n", c.Host, c.Port, c.DBName)
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

func readAndPrepareQueries(sqlFile string, db *sqlx.DB, fs stuffbin.FileSystem) *models.Queries {
	// Load SQL queries.
	fmt.Printf("preparing queries from: %s\n", sqlFile)
	qB, err := fs.Read(sqlFile)
	if err != nil {
		fmt.Printf("An error occurred in opening the file at path: %s\n", sqlFile)
		panic(err)
	}

	qMap, err := goyesql.ParseBytes(qB)
	if err != nil {
		fmt.Printf("An error occurred in parsing the file at path: %s\n", sqlFile)
		panic(err)
	}

	var q models.Queries
	if err := goyesqlx.ScanToStruct(&q, qMap, db.Unsafe()); err != nil {
		panic(err)
	}

	return &q
}

func initFs() stuffbin.FileSystem {
	files := []string{}
	files = append(files, joinFSPaths(appDir, appFiles)...)
	fs, err := stuffbin.NewLocalFS("/", files...)
	if err != nil {
		fmt.Printf("failed reading files from disk: %v", err)
	}
	return fs
}

// initHTTPServer sets up and runs the app's main HTTP server and blocks forever.
func initHTTPServer(app *App) *echo.Echo {
	// Initialize the HTTP server.
	var srv = echo.New()

	// Register app (*App) to be injected into all HTTP handlers.
	srv.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("app", app)
			return next(c)
		}
	})

	// Register all HTTP handlers.
	initHTTPHandlers(srv)

	// Start the server.
	if err := srv.Start(ko.String("app.address")); err != nil {
		if strings.Contains(err.Error(), "Server closed") {
			fmt.Println("Server was shut down")
		} else {
			panic(err)
		}
	}

	return srv
}

func joinFSPaths(root string, paths []string) []string {
	out := make([]string, 0, len(paths))
	for _, p := range paths {
		// real_path:stuffbin_alias
		f := strings.Split(p, ":")

		out = append(out, path.Join(root, f[0])+":"+f[1])
	}

	return out
}

func initJwt() *jwt.JwtGenerator {
	return jwt.NewJwtGenerator("", jwt.HMACSHA256)
}
