package main

import "fmt"

type Task struct {
	ID   int
	Name string
	Done bool
}

func (t *Task) MarkAsDone() {
	if t.Done {
		fmt.Println("Task Already Done")
	}

	t.Done = true
	fmt.Println("Task mark as Done")
}

func NewTask(taskId int, taskName string) Task {
	var newTask Task = Task{ID: taskId, Name: taskName}

	return newTask
}
