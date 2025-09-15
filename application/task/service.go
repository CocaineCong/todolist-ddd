package task

import (
	"context"
	"errors"
	"sync"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/domain/task/service"
	ctl "github.com/CocaineCong/todolist-ddd/infrastructure/common/context"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

type Service interface {
	CreateTask(ctx context.Context, req *types.CreateTaskReq) (*entity.Task, error)
	ListTask(ctx context.Context, req *types.ListTasksReq) (any, error)
	DetailTask(ctx context.Context, req *types.DetailReq) (*entity.Task, error)
	UpdateTask(ctx context.Context, req *types.UpdateTaskReq) error
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

func (s *ServiceImpl) CreateTask(ctx context.Context, req *types.CreateTaskReq) (*entity.Task, error) {
	// 获取用户信息
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	task, err := entity.NewTask(userInfo.Id, userInfo.Name, req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	return s.td.CreateTask(ctx, task)
}

func (s *ServiceImpl) ListTask(ctx context.Context, req *types.ListTasksReq) (any, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	list, count, err := s.td.ListTaskByUid(ctx, userInfo.Id, req.Pagination)
	if err != nil {
		return nil, err
	}
	return ListResponse(list, count), nil
}

func (s *ServiceImpl) DetailTask(ctx context.Context, req *types.DetailReq) (*entity.Task, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	task, err := s.td.FindTaskByTid(ctx, req.Id, userInfo.Id)
	if err != nil {
		return nil, err
	}
	if !task.IsExist() {
		return nil, errors.New("task not exist")
	}
	if !task.BelongsToUser(userInfo.Id) {
		return nil, errors.New("user not exist")
	}
	return task, nil
}

func (s *ServiceImpl) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) error {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	task := UpdateReq2TaskEntity(userInfo.Id, userInfo.Name, req)
	return s.td.UpdateTask(ctx, task)
}

func (s *ServiceImpl) SearchTask(ctx context.Context, req *types.SearchTaskReq) (any, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	list, count, err := s.td.SearchTask(ctx, userInfo.Id, req.Info, req.Pagination)
	if err != nil {
		return nil, err
	}

	return ListResponse(list, count), nil
}

func (s *ServiceImpl) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) error {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	return s.td.DeleteTask(ctx, userInfo.Id, req.Id)
}
