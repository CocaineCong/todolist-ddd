package repository

import (
	"context"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

type Task interface {
	TaskBase
	TaskQuery
}

type TaskBase interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	ListTaskByUid(ctx context.Context, uid uint, p types.Pagination) ([]*entity.Task, int64, error)
	FindTaskByTid(ctx context.Context, tid, uid uint) (*entity.Task, error)
	SearchTask(ctx context.Context, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, uid, tid uint) error
}

type TaskQuery interface {
	// ...
}
