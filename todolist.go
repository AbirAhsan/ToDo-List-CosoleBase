package main

import "fmt"

type TodoList struct {
	tasks []Task
}

// Add task to todo list
func (td *TodoList) AddTask(taskName string) {

	taskId := len(td.tasks) + 1

	task := NewTask(taskId, taskName)

	td.tasks = append(td.tasks, task)

	fmt.Println("Task added successfully")
}

// View task of this todo list
func (td *TodoList) ViewTask() {
	fmt.Println("==============================")
	for _, task := range td.tasks {
		doneStatus := " "
		if task.Done {
			doneStatus = "x"
		}
		fmt.Printf("[%s] Task #%d: %s\n", doneStatus, task.ID, task.Name)
	}

	fmt.Println("==============================")

}

// Mark task of this todo list as Done
func (td *TodoList) MarkTaskAsDone(taskId int) {
	if taskId < 1 || taskId > len(td.tasks) {
		fmt.Println(taskId, "is invalid Task Id ")
		return
	}
	td.tasks[taskId-1].MarkAsDone()
	fmt.Println("Task Done ", taskId)
}
