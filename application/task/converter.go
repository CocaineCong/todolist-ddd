package task

import (
	"github.com/CocaineCong/todolist-ddd/domain/task/entity"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

func ListResponse(list []*entity.Task, count int64) types.List[*entity.Task] {
	return types.List[*entity.Task]{
		Items: list,
		Count: count,
	}
}
