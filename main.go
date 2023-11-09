package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	todoList := TodoList{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("===============TODO LIST===============")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Println("=======================================")
		fmt.Println("Enter Your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Enter task name:")
			scanner.Scan()
			taskName := scanner.Text()
			todoList.AddTask(taskName)
		case "2":
			todoList.ViewTask()
		case "3":
			fmt.Println("Enter task name:")
			scanner.Scan()
			taskIdStr := scanner.Text()

			taskIdInt, err := strconv.Atoi(taskIdStr)

			if err != nil {
				fmt.Println("Invalid Task Id")
				continue
			}
			todoList.MarkTaskAsDone(taskIdInt)

		case "4":
			fmt.Println("Exit Task")
			return
		default:
			fmt.Printf("\nInvalid Choice\n")
		}

	}
}
