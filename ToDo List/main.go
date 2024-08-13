package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/benM4k/go_projects/handlers"
)

func main() {
	var idxInput int
	var taskInput, newTaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark as complete")
	fmt.Println("4. Edit Task")
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter choice (1,2,3,4,5,6):")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Enter task: ")
			scanner.Scan()
			taskInput = scanner.Text()
			handlers.AddTask(taskInput)
		case 2:
			handlers.ListTasks()
		case 3:
			fmt.Println("Enter task index: ")
			scanner.Scan()
			idxInput, _ = strconv.Atoi(scanner.Text())
			handlers.MarkCompleted(idxInput)
		case 4:
			fmt.Println("Enter task index: ")
			scanner.Scan()
			idxInput, _ = strconv.Atoi(scanner.Text())
			fmt.Println("Enter new task: ")
			scanner.Scan()
			newTaskInput = scanner.Text()
			handlers.EditTask(idxInput, newTaskInput)
		case 5:
			fmt.Println("Enter task index: ")
			scanner.Scan()
			idxInput, _ = strconv.Atoi(scanner.Text())
			handlers.DeleteTask(idxInput)
		case 6:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
