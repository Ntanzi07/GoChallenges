package data

import (
	"Ex4/models"
	"errors"
	"sync"
	"time"
)

var (
	tasks []models.Task
	mu    sync.Mutex
)

var lastId = 0

func GetById(id int) (*models.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("not found")
}

func GetAll() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

func Create(title string, description string) (*models.Task, error) {
	mu.Lock()
	defer mu.Unlock()
	lastId++

	newTask := models.Task{
		ID:          lastId,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	return &newTask, nil
}

func Update(updated models.Task) (*models.Task, error) {

	task, err := GetById(updated.ID)
	if err != nil {
		return nil, err
	}

	if updated.Title != "" {
		task.Title = updated.Title
	}

	if updated.Description != "" {
		task.Description = updated.Description
	}

	mu.Lock()
	defer mu.Unlock()

	task.Title = updated.Title
	task.Description = updated.Description
	task.Completed = updated.Completed

	return task, err
}

func Delete(id int) error {
	mu.Lock()
	defer mu.Unlock()

	index, err := findId(id)
	if err != nil {
		return err
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return nil
}

func findId(id int) (int, error) {
	for index, v := range tasks {
		if v.ID == id {
			return index, nil
		}
	}
	return -1, errors.New("id nao encontrado")
}
