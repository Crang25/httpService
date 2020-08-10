package router

import (
	"net/http"

	"github.com/Crang25/httpService/internal/storages"
)

type rootHandler struct {
	tasksHandler tasksHandler
	taskHandler  taskHandler
}

func newRootHandler(store storages.Store) rootHandler {
	return rootHandler{
		tasksHandler: newTasksHandler(store),
		taskHandler:  newTaskHandler(store),
	}
}

func (rh *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "tasks":
		rh.tasksHandler.ServerHTTP(w, r)
	case "task":
		rh.taskHandler.ServerHTTP(w, r)
	default:
		http.NotFound(w, r)
	}

}
