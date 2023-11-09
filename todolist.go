package main

import "fmt"

type TodoList struct {
	tasks []Task
}

// Add task to todo list
func (td *TodoList) AddTask(taskName string) {
	fmt.Println("Task added ", taskName)
}

// View task of this todo list
func (td *TodoList) ViewTask() {
	fmt.Println("Task added ", td.tasks)
}

// Mark task of this todo list as Done
func (td *TodoList) MarkTaskAsDone(taskId int) {
	fmt.Println("Task Done ", taskId)
}
