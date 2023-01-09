package teacherController

import (
	"fmt"
	"student_management/controllers/base"
	"student_management/model/class"
	studentClass "student_management/model/student_class"
	"student_management/model/teacher"
)
var teachers = teacher.Datas

func StartFilterOption()  {
	fmt.Println("1. View list teacher")
	fmt.Println("2. View list homeroom teachers")
	fmt.Println("3. View list teacher have more than X students")
	fmt.Println("4. Exit")
	optons := base.InputMenuOption()

	switch optons {
	case "1":
		all()
		StartFilterOption()
	case "2":
		getHomeroomTeachers()
		StartFilterOption()
	case "3":
		num := base.InputNumberValue()
		haveNumOfStudentMore(num)
		StartFilterOption()
	case "4":
		fmt.Println("\nExited to main menu")
	default:
		fmt.Println("\nInvalid! Please try again:")
		StartFilterOption()
	}

}
func all()  {
	base.All(teachers)
}

func getHomeroomTeachers() {
	for _, data := range teachers {
		if data["homeroom_teacher"] == true {
			base.SortAndPrint(data)
		}
	}
}

func haveNumOfStudentMore(num int) {
	size := 0
	for _, teacher := range teachers {
		if teacher["homeroom_teacher"] == true {
			class := class.ByHomeroomTeacherId(teacher["id"].(int))
			stdClasses := studentClass.ByClassId(class["id"].(int))
			if len(stdClasses) > num {
				size++
				fmt.Printf("\n%v has more than %v students\n", teacher["name"], num)
			}
		}
	}

	if size == 0 {
		fmt.Printf("\nNo teacher has more than %v students\n", num)
		fmt.Println("===================================")
	}
}
