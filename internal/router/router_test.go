package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Crang25/httpService/internal/models"
	"github.com/Crang25/httpService/internal/storages/memstore"

	"github.com/stretchr/testify/require"
)

func TestNotFound(t *testing.T) {
	// Создаем роутер
	r := New(nil)

	// Создаем тестовый сервер на произпольном порту
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	// Делаем запрос на сервер для тестирования ответа
	resp, err := http.Get(srv.URL + "/unknown")
	require.NoError(t, err)

	// Второй параметр Equal - ожидаемая ошибка, третий параметр - фактическая ошибка
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestEmpyList(t *testing.T) {
	r := New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	checkTaskList(t, srv.URL)

	// require.Empty(t, taskList.Tasks)
}

func TestCreateTask(t *testing.T) {
	// Запуск сервера
	r := New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	// Создание и кодирование в json таски
	// Отправляем созданную таску на сервер
	// Получаем созданную ранее таску
	respTask := checkCreateTask(t, srv.URL, "сходить в магазин")

	// Делаем get запрос, чтобы получить список тасок и
	// сохраняем ответ сервера(должен был вернуться список созданных нами тасок(пока что одна))
	// Ожидаем, что сервер вернет нам список из одной таски - созданной нами выше
	checkTaskList(t, srv.URL, respTask)

}

func checkTaskList(t *testing.T, checkURL string, taskList ...models.Task) {
	resp, err := http.Get(checkURL + "/tasks")
	// Делаем get запрос, чтобы получить список тасок
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respTasks models.TaskList
	// Сохраняем ответ сервера(должен был вернуться список созданных нами тасок(пока что одна))
	err = json.NewDecoder(resp.Body).Decode(&respTasks)
	require.NoError(t, err)

	require.Equal(t, taskList, respTasks.Tasks)
}

func checkCreateTask(t *testing.T, checkURL, text string) models.Task {
	task := models.Task{
		Text: text,
	}

	jsTask, err := json.Marshal(&task)
	require.NoError(t, err)

	resp, err := http.Post(checkURL+"/task", "application/json", bytes.NewReader(jsTask))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respTask models.Task
	json.NewDecoder(resp.Body).Decode(&respTask)

	return respTask
}

func TestDeleteTask(t *testing.T) {
	r := New(memstore.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	task := checkCreateTask(t, srv.URL+"/task", "помыть голову")
	delatask := checkDeleteTask(t, srv.URL, task.ID)
	require.Equal(t, task, delatask)
	checkTaskList(t, srv.URL)
}

func checkDeleteTask(t *testing.T, checkURL, id string) models.Task {
	req, err := http.NewRequest(http.MethodDelete, checkURL+"/task/"+id, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var task models.Task

	err = json.NewDecoder(resp.Body).Decode(&task)
	require.NoError(t, err)

	return task

}
