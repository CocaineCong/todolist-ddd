package persistence

import (
	"gorm.io/gorm"

	"github.com/CocaineCong/todolist-ddd/domain/user/repository"
	"github.com/CocaineCong/todolist-ddd/infra/persistence/user"
)

type Repositories struct {
	User repository.User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewRepository(db),
	}
}
