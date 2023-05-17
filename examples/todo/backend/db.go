package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var TaskNotFoundError = errors.New("task not found")

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetAllTasks() ([]Task, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func AddTask(title string) (Task, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return Task{}, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", title, false)
	if err != nil {
		return Task{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Task{}, err
	}

	return Task{ID: int(id), Title: title, Completed: false}, nil
}

func ModifyTask(id int, newTitle string) (Task, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return Task{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE tasks SET title = ? WHERE id = ?", newTitle, id)
	if err != nil {
		return Task{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Task{}, err
	}
	if rowsAffected == 0 {
		return Task{}, TaskNotFoundError
	}

	return Task{ID: id, Title: newTitle}, nil
}

func RemoveTask(id int) error {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return TaskNotFoundError
	}

	return nil
}

func SetTaskCompletion(id int, completed bool) (Task, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return Task{}, err
	}
	defer db.Close()

	result, err := db.Exec("UPDATE tasks SET completed = ? WHERE id = ?", completed, id)
	if err != nil {
		return Task{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Task{}, err
	}
	if rowsAffected == 0 {
		return Task{}, TaskNotFoundError
	}

	return Task{ID: id, Completed: completed}, nil
}