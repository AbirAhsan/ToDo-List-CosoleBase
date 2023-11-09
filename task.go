package main

type Task struct {
	ID   int
	Name string
	Done bool
}

func NewTask(taskId int, taskName string) Task {
	var newTask Task = Task{ID: taskId, Name: taskName}

	return newTask
}
