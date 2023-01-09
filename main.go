package main

import (
	"fmt"
	"student_management/controllers/mainController"
)

func main() {
	fmt.Println("=======================================")
	fmt.Println("Welcome to student management program!")
	fmt.Println("=======================================")
	mainController.Start()
}
