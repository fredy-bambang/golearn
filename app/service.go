package app

import (
	"context"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
	"gopkg.in/go-playground/validator.v9"
)

// Service represents the interface for app usecases.
type Service interface {
	CreateTask(context.Context, *CreateTaskInput) (*CreateTaskOutput, error)
	CompleteTask(context.Context, *CompleteTaskInput) error
	DeleteTask(context.Context, *DeleteTaskInput) error
	GetTask(context.Context, *GetTaskInput) (*GetTaskOutput, error)
	ListTasks(w http.ResponseWriter, req bunrouter.Request) error
}

type service struct {
	db       *bun.DB
	validate *validator.Validate
}

// ServiceConfig contains everything needed by the service.
type ServiceConfig struct {
	DB       *bun.DB
	Validate *validator.Validate
}

// NewService creates a new task service.
func NewService(cfg *ServiceConfig) Service {
	return &service{
		db:       cfg.DB,
		validate: cfg.Validate,
	}
}
