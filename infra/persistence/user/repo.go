package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/domain/user/repository"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.User {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := Entity2PO(user)
	err := r.db.WithContext(ctx).Create(u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserByName(ctx context.Context, username string) (*entity.User, error) {
	var u *User
	err := r.db.WithContext(ctx).Where("user_name = ?", username).Find(&u).Error
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	var u *User
	err := r.db.WithContext(ctx).Where("id = ?", id).Find(&u).Error
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	return PO2Entity(u), nil
}
