package task

import (
	"context"
	"sync"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/domain/task/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

type Service interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	ListTask(ctx context.Context, req *types.ListTasksReq) (any, error)
	DetailTask(ctx context.Context, req *types.DetailReq) (*entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	SearchTask(ctx context.Context, req *types.SearchTaskReq) (any, error)
	DeleteTask(ctx context.Context, req *types.DeleteTaskReq) error
}

type ServiceImpl struct {
	td service.TaskDomain
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func GetServiceImpl(srv service.TaskDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{td: srv}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return s.td.CreateTask(ctx, task)
}

func (s *ServiceImpl) ListTask(ctx context.Context, req *types.ListTasksReq) (any, error) {
	list, count, err := s.td.ListTaskByUid(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}
	var resp = types.TaskListResp
	resp.Items = list
	resp.Count = count
	return resp, nil
}

func (s *ServiceImpl) DetailTask(ctx context.Context, req *types.DetailReq) (*entity.Task, error) {
	return s.td.FindTaskByTid(ctx, req.Id)
}

func (s *ServiceImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	return s.td.UpdateTask(ctx, task)
}

func (s *ServiceImpl) SearchTask(ctx context.Context, req *types.SearchTaskReq) (any, error) {
	list, count, err := s.td.SearchTask(ctx, req.Info, req.Pagination)
	if err != nil {
		return nil, err
	}
	var resp = types.TaskListResp
	resp.Items = list
	resp.Count = count
	return resp, nil
}

func (s *ServiceImpl) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) error {
	return s.td.DeleteTask(ctx, req.Id)
}
