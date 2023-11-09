package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		fmt.Println("Scaned Text", choice)

		switch choice {
		case "1":
			fmt.Println("Add Task")
		case "2":
			fmt.Println("View Task")
		case "3":
			fmt.Println("Mark Task as Done")
		case "4":
			fmt.Println("Exit Task")
			return
		default:
			fmt.Printf("\nInvalid Choice\n")
		}

	}
}
