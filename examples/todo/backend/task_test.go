package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"main"
)

type response struct {
	Status  int             `json:"status"`
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func TestGetTasks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.GetTasks)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	var resp response
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
}

func TestCreateTask(t *testing.T) {
	req, _ := http.NewRequest("POST", "/tasks", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.CreateTask)

	oldTasks, _ := main.GetAllTasks()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)

	newTasks, _ := main.GetAllTasks()
	assert.Equal(t, len(oldTasks), len(newTasks))
}

func TestUpdateTask(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/tasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.UpdateTask)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
}

func TestDeleteTask(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.DeleteTask)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode)

	_, err := main.GetTaskById(1)
	assert.IsType(t, main.TaskNotFoundError{}, err)
}

func TestToggleTaskComplete(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/tasks/1/toggle", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.ToggleTaskComplete)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode)
}