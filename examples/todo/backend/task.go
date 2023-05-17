package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vedhavyas/todo/backend"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := backend.GetAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tasks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newTask, err := backend.AddTask(task.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := backend.ModifyTask(task.ID, task.Description); err != nil {
		if _, ok := err.(*backend.TaskNotFoundError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := backend.RemoveTask(id); err != nil {
		if _, ok := err.(*backend.TaskNotFoundError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ToggleTaskComplete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	completedStr := r.URL.Query().Get("completed")
	completed, err := strconv.ParseBool(completedStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := backend.SetTaskCompletion(id, completed); err != nil {
		if _, ok := err.(*backend.TaskNotFoundError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}