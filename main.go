package main

import (
	"fmt"
	"net/http"
	"strings"
)

var shortGolang = "Watch Go Crash Course"
var fullGolang = "Watch Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	taskList := strings.Join(taskItems, "\n")
	fmt.Fprintf(writer, "My Tasks:\n%s", taskList)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello user, Welcome to Our TodoList App!")
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTasks(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
