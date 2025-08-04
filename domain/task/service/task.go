package service

import (
	"context"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/domain/task/repository"
	lctx "github.com/CocaineCong/todolist-ddd/infrastructure/common/context"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

type TaskDomain interface {
	CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error)
	FindTaskByTid(ctx context.Context, taskId uint) (*entity.Task, error)
	ListTaskByUid(ctx context.Context, p types.Pagination) ([]*entity.Task, int64, error)
	UpdateTask(ctx context.Context, in *entity.Task) error
	SearchTask(ctx context.Context, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, tid uint) error
}

type TaskDomainImpl struct {
	repo repository.Task
}

func NewTaskDomainImpl(repo repository.Task) TaskDomain {
	return &TaskDomainImpl{repo: repo}
}

func (t *TaskDomainImpl) CreateTask(ctx context.Context, in *entity.Task) (*entity.Task, error) {
	return t.repo.CreateTask(ctx, in)
}

func (t *TaskDomainImpl) FindTaskByTid(ctx context.Context, taskId uint) (*entity.Task, error) {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return t.repo.FindTaskByTid(ctx, taskId, userInfo.Id)
}

func (t *TaskDomainImpl) ListTaskByUid(ctx context.Context, p types.Pagination) ([]*entity.Task, int64, error) {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return nil, 0, err
	}
	return t.repo.ListTaskByUid(ctx, userInfo.Id, p)
}

func (t *TaskDomainImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	task.Uid = userInfo.Id
	task.UserName = userInfo.Name
	return t.repo.UpdateTask(ctx, task)
}

func (t *TaskDomainImpl) SearchTask(ctx context.Context, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return nil, 0, err
	}
	return t.repo.SearchTask(ctx, userInfo.Id, keyword, p)
}

func (t *TaskDomainImpl) DeleteTask(ctx context.Context, tid uint) error {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	return t.repo.DeleteTask(ctx, tid, userInfo.Id)
}
