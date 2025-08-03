package consts

const (
	DefaultPageSizeMax = 100
	DefaultPageSize    = 10
	DefaultPage        = 1
)

const (
	TaskStatusEmunInit = iota
	TaskStatusEmunFinished
)

const (
	TaskStatusInit     = "未完成"
	TaskStatusFinished = "已完成"
)

var (
	TaskStatusMap = map[int]string{
		TaskStatusEmunInit:     TaskStatusInit,
		TaskStatusEmunFinished: TaskStatusFinished,
	}
)
