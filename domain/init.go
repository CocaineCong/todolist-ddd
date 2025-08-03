package domain

import (
	userApp "github.com/CocaineCong/todolist-ddd/application/user"
	userSrv "github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/infra/dbs"
	"github.com/CocaineCong/todolist-ddd/infra/persistence"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)

	// user domain
	userDomain := userSrv.NewUserDomainImpl(repos.User)
	userApp.GetServiceImpl(userDomain)
}
