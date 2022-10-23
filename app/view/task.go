package view

import (
	"context"

	"github.com/fredy-bambang/golearn/app/database"
	"github.com/uptrace/bun"
)

func GetTask(ctx context.Context, db *bun.DB, taskID string) (*database.Task, error) {
	task, err := database.FindTask(ctx, db, taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func ListTasks(ctx context.Context, db *bun.DB) ([]*database.Task, error) {
	tasks, err := database.FindAllTasks(ctx, db)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
