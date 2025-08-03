package task

import (
	"context"
	"sync"

	"github.com/CocaineCong/todolist-ddd/domain/task/service"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

type Service interface {
	CreateTask(ctx context.Context, req *types.CreateTaskReq) (any, error)
	ListTask(ctx context.Context, req *types.ListTasksReq) (any, error)
	DetailTask(ctx context.Context, req *types.DetailReq) (any, error)
	UpdateTask(ctx context.Context, req *types.UpdateTaskReq) error
	SearchTask(ctx context.Context, req *types.SearchTaskReq) (any, error)
	DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (any, error)
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

func (s *ServiceImpl) CreateTask(ctx context.Context, req *types.CreateTaskReq) (any, error) {
	task, err := CreateReqDTO2Entity(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := s.td.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	resp := Entity2TaskResp(result)
	return resp, nil
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

func (s *ServiceImpl) DetailTask(ctx context.Context, req *types.DetailReq) (any, error) {
	result, err := s.td.FindTaskByTid(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resp := Entity2TaskResp(result)
	return resp, nil
}

func (s *ServiceImpl) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) error {
	task, err := UpdateReqDTO2Entity(ctx, req)
	if err != nil {
		return err
	}
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

func (s *ServiceImpl) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (any, error) {
	err := s.td.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
