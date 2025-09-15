package entity

import (
	"errors"
	"time"

	"github.com/CocaineCong/todolist-ddd/infrastructure/consts"
)

type Task struct {
	Id        uint      `json:"id"`
	Uid       uint      `json:"uid"`
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Status    int       `json:"status"`
	Content   string    `json:"content"`
	StartTime int64     `json:"start_time"`
	EndTime   int64     `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(uid uint, title, content string) (*Task, error) {
	if uid == 0 {
		return nil, errors.New("owner ID cannot be empty")
	}
	if title == "" {
		return nil, errors.New("task title cannot be empty")
	}
	now := time.Now()
	return &Task{
		Uid:       uid,
		Title:     title,
		Status:    consts.TaskStatusEmunInit,
		Content:   content,
		StartTime: now.Unix(),
		CreatedAt: now,
	}, nil
}

func (t *Task) Complete() error {
	now := time.Now()
	t.Status = consts.TaskStatusEmunFinished
	t.UpdatedAt = now
	t.EndTime = now.Unix()
	return nil
}

func (t *Task) AddUserInfo(uid uint, userName string) {
	t.Uid = uid
	t.UserName = userName
}

func (t *Task) BelongsToUser(userID uint) bool {
	return t.Uid == userID
}

func (t *Task) UpdateContent(title, content string) error {
	if title == "" {
		return errors.New("task title cannot be empty")
	}

	t.Title = title
	t.Content = content
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) IsExist() bool {
	return t.Id > 0
}
