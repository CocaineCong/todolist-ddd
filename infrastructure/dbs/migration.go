package dbs

import (
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/task"
	"github.com/CocaineCong/todolist-ddd/infrastructure/persistence/user"
)

// 执行数据迁移
func migration() {
	// 自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&user.User{},
			&task.Task{},
		)
	if err != nil {
		return
	}
}
