package task

import (
	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/user"
)

func Entity2PO(task *entity.Task) *Task {
	return &Task{
		Uid:       task.Uid,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func PO2Entity(t *Task, u *user.User) *entity.Task {
	return &entity.Task{
		Uid:       t.Uid,
		UserName:  u.UserName,
		Title:     t.Title,
		Status:    t.Status,
		Content:   t.Content,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
	}
}
