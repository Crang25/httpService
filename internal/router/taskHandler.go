package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Crang25/httpService/internal/models"

	"github.com/Crang25/httpService/internal/storages"
)

type taskHandler struct {
	store storages.Store
}

func newTaskHandler(store storages.Store) taskHandler {
	return taskHandler{
		store: store,
	}
}

func (th taskHandler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	respTask, err := th.store.CreateTask(r.Context(), task)
	if err != nil {
		fmt.Fprintf(w, "failed to create task in store: %v", err)
		return
	}

	writeJSON(w, respTask)
}

func (th taskHandler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := shiftPath(r.URL.Path)
	if id == "" {
		http.NotFound(w, r)
	}
	task := th.store.DeleteTask(r.Context(), id)

	writeJSON(w, task)
}

func (th taskHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "failed to decode body :%v", err)
			return
		}

		newTask, err := th.store.CreateTask(r.Context(), task)
		if err != nil {
			fmt.Fprintf(w, "failed to create task in store: %v", err)
			return
		}

		writeJSON(w, &newTask)

	case http.MethodDelete:
		th.handleDeleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
