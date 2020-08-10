package storages

import (
	"context"

	"github.com/Crang25/httpService/internal/models"
)

type Store interface {
	GetTaskList(ctx context.Context) (models.TaskList, error)
	CreateTask(ctx context.Context, tsk models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id string) models.Task
}
