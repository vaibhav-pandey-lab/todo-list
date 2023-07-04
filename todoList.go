package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a todo task
type Task struct {
	ID       int
	Title    string
	Completed bool
}

var tasks []Task
var lastID int

// createTask creates a new task with the provided title
func createTask(title string) Task {
	lastID++
	task := Task{
		ID:       lastID,
		Title:    title,
		Completed: false,
	}
	tasks = append(tasks, task)
	return task
}

// getAllTasks returns all tasks
func getAllTasks() []Task {
	return tasks
}

// updateTask updates a task with the provided ID
func updateTask(id int, title string, completed bool) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Completed = completed
			return true
		}
	}
	return false
}

// deleteTask deletes a task with the provided ID
func deleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

// printTasks displays all tasks
func printTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		completed := " "
		if task.Completed {
			completed = "X"
		}
		fmt.Printf("[%s] %d: %s\n", completed, task.ID, task.Title)
	}
}

// readString reads a string from the console input
func readString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// readInt reads an integer from the console input
func readInt(prompt string) int {
	for {
		input := readString(prompt)
		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
}

func main() {
	for {
		fmt.Println("==== Todo List ====")
		fmt.Println("1. Create a task")
		fmt.Println("2. List all tasks")
		fmt.Println("3. Update a task")
		fmt.Println("4. Delete a task")
		fmt.Println("5. Exit")

		choice := readInt("Enter your choice: ")

		switch choice {
		case 1:
			title := readString("Enter task title: ")
			task := createTask(title)
			fmt.Printf("Task created: %d\n", task.ID)

		case 2:
			fmt.Println("All tasks:")
			printTasks()

		case 3:
			id := readInt("Enter task ID to update: ")
			title := readString("Enter new task title: ")
			completed := readInt("Enter completion status (0 - Incomplete, 1 - Complete): ")
			if updateTask(id, title, completed == 1) {
				fmt.Println("Task updated successfully.")
			} else {
				fmt.Println("Task not found.")
			}

		case 4:
			id := readInt("Enter task ID to delete: ")
			if deleteTask(id) {
				fmt.Println("Task deleted successfully.")
			} else {
				fmt.Println("Task not found.")
			}

		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}