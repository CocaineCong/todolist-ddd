package main

import (
	config "github.com/CocaineCong/todolist-ddd/conf"
	"github.com/CocaineCong/todolist-ddd/domain"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/log"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/adapter/initialize"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/dbs"
)

func main() {
	loadingInfra()
	loadingDomain()
	r := initialize.NewRouter()
	_ = r.Run(config.Conf.Server.Port)
}

// loadingInfra 加载基础架构信息
func loadingInfra() {
	config.InitConfig()
	log.InitLog()
	dbs.MySQLInit()
}

// loadingDomain 加载各个领域服务
func loadingDomain() {
	domain.LoadingDomain()
}
