package http

import (
	"github.com/fredy-bambang/golearn/app"
	"github.com/uptrace/bun"
)

// Server hosts endpoints to manage tasks.
type Server struct {
	db  *bun.DB
	svc app.Service
}

// ServerConfig contains everything needed by a Server.
type ServerConfig struct {
	DB  *bun.DB
	Svc app.Service
}

// NewServer creates a new Server.
func NewServer(cfg *ServerConfig) *Server {
	return &Server{
		db:  cfg.DB,
		svc: cfg.Svc,
	}
}
