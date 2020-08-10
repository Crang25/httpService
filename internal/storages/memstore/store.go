package memstore

import (
	"context"
	"strconv"
	"sync"

	"github.com/Crang25/httpService/internal/models"
)

type MemStore struct {
	mu     sync.Mutex
	tasks  models.TaskList
	lastID int64
}

func New() *MemStore {
	return &MemStore{}
}

func (ms *MemStore) GetTaskList(ctx context.Context) (models.TaskList, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	tasks := models.TaskList{
		Tasks: append([]models.Task(nil), ms.tasks.Tasks...),
	}
	return tasks, nil
}

func (ms *MemStore) CreateTask(ctx context.Context, tsk models.Task) (models.Task, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.lastID++
	tsk.ID = strconv.FormatInt(ms.lastID, 16)
	ms.tasks.Tasks = append(ms.tasks.Tasks, tsk)
	return tsk, nil
}

func (ms *MemStore) DeleteTask(ctx context.Context, id string) models.Task {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	var task models.Task
	for i := range ms.tasks.Tasks {
		if ms.tasks.Tasks[i].ID == id {
			task = ms.tasks.Tasks[i]
			ms.tasks.Tasks = append(ms.tasks.Tasks[:i], ms.tasks.Tasks[i+1:]...)
			break
		}
	}

	return task
}
