package handlers

import (
	"fmt"

	"github.com/benM4k/go_projects/types"
)

var tasks []types.Tasks

func AddTask(task string) {
	newTask := types.Tasks{TaskName: task, Completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Tasked added")
}

func ListTasks() {
	for i, task := range tasks {
		status := "n"
		if task.Completed {
			status = "d"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

func MarkCompleted(idx int) {
	if idx > 0 && idx <= len(tasks) {
		tasks[idx-1].Completed = true
		fmt.Println("Tasked Marked as completed")
	} else {
		fmt.Println("Invalid task index")
	}
}

func EditTask(idx int, newTask string) {
	if idx > 0 && idx <= len(tasks) {
		tasks[idx-1].TaskName = newTask
		fmt.Println("Tasked edited successfully")
	} else {
		fmt.Println("Invalid task index")
	}
}

func DeleteTask(idx int) {
	if idx > 0 && idx <= len(tasks) {
		tasks = append(tasks[:idx-1], tasks[idx:]...)
		fmt.Println("Tasked deleted")
	} else {
		fmt.Println("Invalid task index")
	}
}
