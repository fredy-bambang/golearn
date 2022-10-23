package main

import (
	"context"
	"fmt"
	"net/http"

	apphttp "github.com/fredy-bambang/golearn/app/http"
	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Port          string `env:"PORT,default=8080"`
	PostgresURL   string `env:"POSTGRES_URL,default=postgres://postgres:password@localhost/todo?sslmode=disable"`
	MigrationsDir string `env:"MIGRATIONS_DIR,default=app/database/migrations"`
}

func main() {
	realMain(context.Background())
	// log.Println("listening on http://localhost:9999")
	// log.Println(http.ListenAndServe(":9999", router))
}

func realMain(ctx context.Context) error {
	// 1. Load the config.
	var cfg config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	server := apphttp.NewServer(&apphttp.ServerConfig{})

	fmt.Println("listening on http://localhost:9999")
	err := http.ListenAndServe(":9999", server.Routes())

	if err != nil {
		return err
	}
	return nil
}
