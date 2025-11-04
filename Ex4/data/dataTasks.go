package data

import (
	"Ex4/models"
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

var (
	tasks []models.Task
	mu    sync.Mutex
)

var lastId = 0

func LoadTasks() error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open("data/tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []models.Task{}
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	lastId = tasks[len(tasks)-1].ID
	return nil
}

func SaveTasks() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile("data/tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil

}

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

func GetByCompleted(completed bool) []models.Task {
	mu.Lock()
	defer mu.Unlock()

	var filtered []models.Task
	for _, task := range tasks {
		if task.Completed == completed {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func GetAll() []models.Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

func Create(title string, description string) (*models.Task, error) {
	mu.Lock()

	lastId++

	newTask := models.Task{
		ID:          lastId,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	mu.Unlock()
	err := SaveTasks()
	if err != nil {
		return nil, err
	}
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

	task.Title = updated.Title
	task.Description = updated.Description
	task.Completed = updated.Completed

	mu.Unlock()
	err = SaveTasks()
	if err != nil {
		return nil, err
	}
	return task, err
}

func Delete(id int) error {
	mu.Lock()

	index, err := findId(id)
	if err != nil {
		return err
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	mu.Unlock()

	err = SaveTasks()
	if err != nil {
		return err
	}
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
