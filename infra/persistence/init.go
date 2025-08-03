package persistence

import (
	"gorm.io/gorm"

	tRepo "github.com/CocaineCong/todolist-ddd/domain/task/repository"
	uRepo "github.com/CocaineCong/todolist-ddd/domain/user/repository"
	"github.com/CocaineCong/todolist-ddd/infra/persistence/task"
	"github.com/CocaineCong/todolist-ddd/infra/persistence/user"
)

type Repositories struct {
	User uRepo.User
	Task tRepo.Task
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepository(db),
		Task: task.NewRepository(db),
	}
}
