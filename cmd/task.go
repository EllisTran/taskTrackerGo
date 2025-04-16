package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type TaskStatus string

const (
	todo       TaskStatus = "todo"
	inProgress TaskStatus = "in-progress"
	done       TaskStatus = "done"
)

type Task struct {
	Id          uint64     `json:"id"` // uuid
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"` //date time
}

func addTask(description string, status TaskStatus) {
	tasks := loadTasks()
	taskId := len(tasks) // Typically this would already handled in db but too lazy to actually check if already exists
	task := Task{
		Id:          uint64(taskId),
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
	fmt.Printf("Task: %v", task)
	tasks = append(tasks, task)

	saveToJson(tasks)
}

func updateTask() {}
func deleteTask(taskId uint64) {
	tasks := loadTasks()

	newTasks := []Task{}
	for idx, val := range tasks {
		if val.Id == taskId {
			fmt.Printf("index:\t%d\ntaskId:\t%d\nvalue:\t%v\n", idx, taskId, val)
			newTasks = append(tasks[:idx], tasks[idx+1:]...)
			break
		}
	}

	if len(newTasks) == 0 {
		fmt.Printf("Task Id: %v does not exist", taskId)
		saveToJson(tasks)
	} else {
		saveToJson(newTasks)
	}
}

func listTasks()           {}
func listDoneTasks()       {}
func listTodoTasks()       {}
func listInProgressTasks() {}

func loadFilename(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
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
		formattedError := fmt.Errorf("Failed to Unmarshal JSON: %w", err)
		fmt.Println(formattedError)
		return nil
	}

	fmt.Printf("%v", string(data))
	return tasks
}

func saveToJson(tasks []Task) {
	// Convert to JSON
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Println("Saving tasks as json...")

	err = os.WriteFile("tasks.json", jsonData, os.ModePerm)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
