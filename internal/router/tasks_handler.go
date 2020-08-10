package router

import (
	"fmt"
	"net/http"

	"github.com/Crang25/httpService/internal/storages"
)

type tasksHandler struct {
	store storages.Store
}

func newTasksHandler(store storages.Store) tasksHandler {
	return tasksHandler{
		store: store,
	}
}

func (th tasksHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	// Если на адрес /tasks пришол нет Get запрос, вернем в заголовке Method Not Allowed
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list, err := th.store.GetTaskList(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get tasks list from store: %v", err)
		return
	}

	writeJSON(w, &list)

}
