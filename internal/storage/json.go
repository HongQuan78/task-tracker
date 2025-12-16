package storage

import (
	"encoding/json"
	"os"
	"slices"
	"task-tracker/internal/model"
)

const JSON_PATH = "internal/storage/data.json"

func LoadTasks() ([]model.Task, error) {
	if _, err := os.Stat(JSON_PATH); os.IsNotExist(err) {
		return []model.Task{}, nil
	}
	data, err := os.ReadFile(JSON_PATH)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskIndexById(id int) ([]model.Task, int, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, -1, err
	}

	for i, task := range tasks {
		if task.Id == id {
			return tasks, i, nil
		}
	}

	return nil, -1, os.ErrNotExist
}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(JSON_PATH, data, 0644)
}

func GetNextTaskID() (int, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return 0, err
	}

	maxId := 0
	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	return maxId + 1, nil
}

func AddTask(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	id, err := GetNextTaskID()
	if err != nil {
		return err
	}

	task := model.NewTask(title, id)
	tasks = append(tasks, task)

	return SaveTasks(tasks)
}

func DeleteTask(id int) error {
	tasks, index, err := GetTaskIndexById(id)
	if err != nil {
		return err
	}

	tasks = slices.Delete(tasks, index, index+1)

	if err := SaveTasks(tasks); err != nil {
		return err
	}

	return nil
}

func UpdateTask(id int, title string) error {

	return nil
}
