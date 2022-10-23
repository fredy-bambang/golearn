package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fredy-bambang/golearn/app"
	apphttp "github.com/fredy-bambang/golearn/app/http"
	"github.com/pressly/goose/v3"
	"github.com/sethvargo/go-envconfig"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"gopkg.in/go-playground/validator.v9"
)

type config struct {
	Port          string `env:"PORT,default=9999"`
	PostgresURL   string `env:"POSTGRES_URL,default=postgres://postgres:password@localhost/todo?sslmode=disable"`
	MigrationsDir string `env:"MIGRATIONS_DIR,default=app/database/migrations"`
}

func main() {
	realMain(context.Background())
	// log.Println("listening on http://localhost:9999")
	// log.Println(http.ListenAndServe(":9999", router))
}

// var embedMigrations embed.FS

func realMain(ctx context.Context) error {
	// 1. Load the config.
	var cfg config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path + "/app/database/migrations")

	// sqldb := sql.OpenDB(sqliteshim.NewConnector(pgdriver.WithDSN(cfg.PostgresURL)))
	// file:dev.s3db?cache=shared
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:app/database/migrations/database.db?cache=shared")
	if err != nil {
		panic(err)
	}

	// doing migration
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}
	if err := goose.Up(sqldb, "app/database/migrations"); err != nil {
		panic(err)
	}

	// connect db with bun
	db := bun.NewDB(sqldb, sqlitedialect.New())

	validate := validator.New()
	svc := app.NewService(&app.ServiceConfig{
		DB:       db,
		Validate: validate,
	})

	server := apphttp.NewServer(&apphttp.ServerConfig{
		DB:  db,
		Svc: svc,
	})

	fmt.Println("listening on http://localhost:" + cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, server.Routes())
	if err != nil {
		return err
	}

	return nil
}
