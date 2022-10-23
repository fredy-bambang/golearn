package app

import (
	"context"
	"net/http"
	"time"

	"github.com/fredy-bambang/golearn/app/database"
	"github.com/fredy-bambang/golearn/app/view"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

type (
	CreateTaskInput struct {
		Title string `json:"title"`
	}

	CreateTaskOutput struct {
		TaskID string `json:"task_id"`
	}

	CompleteTaskInput struct {
		TaskID string `httpurl:"task_id"`
	}

	DeleteTaskInput struct {
		TaskID string `httpurl:"task_id"`
	}

	GetTaskInput struct {
		TaskID string `httpurl:"task_id"`
	}

	GetTaskOutput struct {
		*database.Task
	}

	ListTasksInput struct{}

	ListTasksOutput struct {
		Data []*database.Task `json:"data"`
	}
)

func (s *service) CreateTask(ctx context.Context, input *CreateTaskInput) (*CreateTaskOutput, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	task := &database.Task{
		ID:    uuid.New().String(),
		Title: input.Title,
	}

	if err := database.InsertTask(ctx, s.db, task); err != nil {
		return nil, errors.WithStack(err)
	}

	return &CreateTaskOutput{TaskID: task.ID}, nil
}

func (s *service) CompleteTask(ctx context.Context, input *CompleteTaskInput) error {
	err := s.validate.Struct(input)
	if err != nil {
		return err
	}

	if err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		task, err := database.FindTaskForUpdate(ctx, tx, input.TaskID)
		if err != nil {
			return err
		}

		currTime := time.Now()
		task.Completed = true
		task.CompletedAt = &currTime

		if err := database.UpdateTask(ctx, tx, task); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *service) DeleteTask(ctx context.Context, input *DeleteTaskInput) error {
	err := s.validate.Struct(input)
	if err != nil {
		return err
	}

	if err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		task, err := database.FindTaskForUpdate(ctx, tx, input.TaskID)
		if err != nil {
			return err
		}

		if err := database.DeleteTask(ctx, tx, task); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.WithStack(err)
	}

	return err
}

func (s *service) GetTask(ctx context.Context, input *GetTaskInput) (*GetTaskOutput, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}

	task, err := view.GetTask(ctx, s.db, input.TaskID)
	return &GetTaskOutput{Task: task}, err
}

func (s *service) ListTasks(w http.ResponseWriter, req bunrouter.Request) error {
	// err := s.validate.Struct(input)
	// if err != nil {
	// 	return nil, err
	// }

	tasks, err := view.ListTasks(req.Context(), s.db)
	if err != nil {
		return err
	}

	err = bunrouter.JSON(w, bunrouter.H{
		"tasks": tasks,
	})
	if err != nil {
		return err
	}

	return nil
}
