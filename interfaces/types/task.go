package types

import (
	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
)

type List[T any] struct {
	Count int64 `json:"count"`
	Items []T   `json:"items"`
}

var TaskListResp = List[*entity.Task]{}

type DetailReq struct {
	Id uint `json:"id" form:"id"`
}

type DeleteTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type UpdateTaskReq struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type CreateTaskReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

type SearchTaskReq struct {
	Info string `form:"info" json:"info"`
	Pagination
}

type ListTasksReq struct {
	Pagination
}

type TaskResp struct {
	ID        uint   `json:"id,omitempty" example:"1"`       // 任务ID
	Title     string `json:"title,omitempty" example:"吃饭"`   // 题目
	Content   string `json:"content,omitempty" example:"睡觉"` // 内容
	View      uint64 `json:"view,omitempty" example:"32"`    // 浏览量
	Status    int    `json:"status,omitempty" example:"0"`   // 状态(0未完成，1已完成)
	CreatedAt int64  `json:"created_at,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
}
