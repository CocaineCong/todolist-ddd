package types

import (
	"context"
	"time"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	ue "github.com/CocaineCong/todolist-ddd/domain/user/entity"
	lctx "github.com/CocaineCong/todolist-ddd/infrastructure/common/context"
)

func CreateReqDTO2Entity(ctx context.Context, task *CreateTaskReq) (*entity.Task, error) {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &entity.Task{
		Uid:       userInfo.Id,
		UserName:  userInfo.Name,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix(),
	}, nil
}

func UpdateReqDTO2Entity(ctx context.Context, task *UpdateTaskReq) (*entity.Task, error) {
	userInfo, err := lctx.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &entity.Task{
		Uid:       userInfo.Id,
		UserName:  userInfo.Name,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix(),
	}, nil
}

func Entity2TaskResp(task *entity.Task) *TaskResp {
	return &TaskResp{
		ID:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		View:      0,
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func UserReq2Entity(user *UserReq) *ue.User {
	return &ue.User{
		Username: user.UserName,
		Password: user.Password,
	}
}
