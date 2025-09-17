package task

import (
	"time"

	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

func ListResponse(list []*entity.Task, count int64) types.List[*entity.Task] {
	return types.List[*entity.Task]{
		Items: list,
		Count: count,
	}
}

func UpdateReq2TaskEntity(tid, uid uint, username string, req *types.UpdateTaskReq) *entity.Task {
	return &entity.Task{
		Id:        tid,
		Uid:       uid,
		UserName:  username,
		Title:     req.Title,
		Status:    req.Status,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}
}
