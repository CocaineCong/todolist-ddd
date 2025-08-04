package domain

import (
	taskApp "github.com/CocaineCong/todolist-ddd/application/task"
	userApp "github.com/CocaineCong/todolist-ddd/application/user"
	taskSrv "github.com/CocaineCong/todolist-ddd/domain/task/service"
	userSrv "github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/dbs"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)

	// user domain
	userDomain := userSrv.NewUserDomainImpl(repos.User)
	userApp.GetServiceImpl(userDomain)

	// task domain
	taskDomain := taskSrv.NewTaskDomainImpl(repos.Task)
	taskApp.GetServiceImpl(taskDomain)
}
