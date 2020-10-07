package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Task is DB model
type Task struct {
	id    int
	title string
	desc  string
}


var allTasks []Task

//RemoveFromIndex is used to delete element from slices at specified index
func RemoveFromIndex(s []Task, index int) []Task {
	return append(s[:index], s[index+1:]...)
}

func insertTODO() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Title: ")
	title, _ := reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Desc: ")
	desc, _ := reader.ReadString('\n')

	task := Task{
		id:    len(allTasks),
		title: strings.TrimSpace(title),
		desc:  strings.TrimSpace(desc),
	}
	allTasks = append(allTasks, task)
}

func fetchAllTODO() {
	for i := range allTasks {
		fmt.Printf("%+v\n", allTasks[i])
	}
}

func removeTODO() {
	var index int
	fmt.Println("Enter index of TODO to Delete: ")
	fmt.Scanln(&index)
	allTasks = RemoveFromIndex(allTasks, index)
}

func commands() {
	fmt.Println("1. Add TODO")
	fmt.Println("2. Remove TODO")
	fmt.Println("3. View All")
	fmt.Println("4. Exit")
}

func takeInput() {
	ch := 0

	for ch != 4 {
		fmt.Println("Enter your Choice: ")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			insertTODO()
		case 2:
			removeTODO()
		case 3:
			fetchAllTODO()
		case 4:
			fmt.Println("Task For Using")
		}
	}

}

func main() {
	commands()
	takeInput()
}
