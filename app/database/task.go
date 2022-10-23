package database

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"task"`

	ID          string     `bun:"id" json:"id"`
	Title       string     `bun:"title" json:"title"`
	Completed   bool       `bun:"completed" json:"completed"`
	CompletedAt *time.Time `bun:"completed_at" json:"completed_at"`
	CreatedAt   *time.Time `bun:"created_at" json:"-"`
	UpdatedAt   *time.Time `bun:"updated_at" json:"-"`
	DeletedAt   *time.Time `bun:"deleted_at" json:"-"`
}

func FindTask(ctx context.Context, db bun.IDB, taskID string) (*Task, error) {
	var task Task
	err := db.NewSelect().
		Model(&task).
		Where("id = ?", taskID).
		Where("deleted_at is NULL").
		Scan(ctx)
	return &task, err
}

func FindTaskForUpdate(ctx context.Context, db bun.IDB, taskID string) (*Task, error) {
	var task Task
	err := db.NewSelect().
		Model(&task).
		Where("id = ?", taskID).
		Where("deleted_at is NULL").
		For("Update").
		Scan(ctx)
	return &task, err
}

func FindAllTasks(ctx context.Context, db bun.IDB) ([]*Task, error) {
	var tasks []*Task
	err := db.NewSelect().
		Model(&tasks).
		Where("deleted_at is NULL").
		Scan(ctx)
	return tasks, err
}

func InsertTask(ctx context.Context, db bun.IDB, task *Task) error {
	_, err := db.NewInsert().
		Model(task).
		Exec(ctx)
	return err
}

func UpdateTask(ctx context.Context, db bun.IDB, task *Task) error {
	_, err := db.NewUpdate().
		Model(task).
		WherePK().
		Exec(ctx)
	return err
}

func DeleteTask(ctx context.Context, db bun.IDB, task *Task) error {
	now := time.Now()
	task.DeletedAt = &now
	return UpdateTask(ctx, db, task)
}
