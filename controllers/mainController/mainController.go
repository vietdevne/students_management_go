package mainController

import (
	"fmt"
	base "student_management/controllers/base"
	classController "student_management/controllers/classController"
	studentController "student_management/controllers/studentController"
	teacherController "student_management/controllers/teacherController"
)

func Start() {
	fmt.Println("1. View list classes")
	fmt.Println("2. View list student")
	fmt.Println("3. View list sort student name ASC")
	fmt.Println("4. View list sort student name DESC")
	fmt.Println("5. View list student of input class")
	fmt.Println("6. View list class input student has join")
	fmt.Println("7. View homeroom teacher of input class")
	fmt.Println("8. Select teachers filter option")
	fmt.Println("9. Exit")
	option := base.InputMenuOption()

	switch option {
	case "1":
		classController.All()
		Start()
	case "2":
		studentController.All()
		Start()
	case "3":
		studentController.SortByNameASC()
		Start()
	case "4":
		studentController.SortByNameDESC()
		Start()
	case "5":
		classId := base.InputIdValue("class")
		classController.WithStudentsById(classId)
		Start()
	case "6":
		studentId := base.InputIdValue("student")
		studentController.WithClassById(studentId)
		Start()
	case "7":
		classId := base.InputIdValue("class")
		classController.WithHomeroomTeacherById(classId)
		Start()
	case "8":
		teacherController.StartFilterOption()
		Start()
	case "9":
		fmt.Println("Thank for use!")
	default:
		fmt.Println("Invalid! Please try again:")
		Start()
	}
}
