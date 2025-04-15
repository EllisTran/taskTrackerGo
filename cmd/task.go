package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	todo       TaskStatus = "todo"
	inProgress TaskStatus = "in-progress"
	done       TaskStatus = "done"
)

type Task struct {
	Id          uuid.UUID  `json:"id"` // uuid
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"` //date time
}

func addTask(description string, status TaskStatus) {
	// var tasks []Task
	tasks := loadTasks()
	task := Task{
		Id:          uuid.New(),
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
	tasks = append(tasks, task)

	// Convert to JSON
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	saveToJson(jsonData)
}

func loadFilename(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(data)
	return data
}
func loadTasks() []Task {
	data := loadFilename("tasks.json")
	if data == nil {
		fmt.Println("No Tasks")
		return nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("lala")
		return nil
	}

	return tasks
}

func saveToJson(json []byte) {
	fmt.Println("Saving tasks as json...")

	err := os.WriteFile("tasks.json", json, os.ModePerm)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
