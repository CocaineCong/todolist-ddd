package domain

import (
	taskApp "github.com/CocaineCong/todolist-ddd/application/task"
	userApp "github.com/CocaineCong/todolist-ddd/application/user"
	taskSrv "github.com/CocaineCong/todolist-ddd/domain/task/service"
	userSrv "github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/auth"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/dbs"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)
	jwtService := auth.NewJWTTokenService()

	// user domain
	userDomain := userSrv.NewUserDomainImpl(repos.User)
	userApp.GetServiceImpl(userDomain, jwtService)

	// task domain
	taskDomain := taskSrv.NewTaskDomainImpl(repos.Task)
	taskApp.GetServiceImpl(taskDomain)
}
