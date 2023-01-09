package classController

import (
	"fmt"
	"student_management/controllers/base"
	"student_management/model/class"
	"student_management/model/student"
	studentClass "student_management/model/student_class"
	"student_management/model/teacher"
)

func All() {
	base.All(class.Datas)
}

func WithStudentsById(id int) {
	class := class.ById(id)
	if len(class) == 0 {
		return
	}

	classStudents := studentClass.ByClassId(id)
	var students []map[string]interface{}
	for _, classStudentsData := range classStudents {
		students = append(students, student.ById(classStudentsData["student_id"].(int)))
	}

	fmt.Println("=====================================")
	fmt.Printf("Class %v have students: ", class["name"])
	fmt.Println("=====================================")
	fmt.Println("Students Information:")
	for _, studentsData := range students {
		fmt.Println("Class name: ", class["name"])
		base.SortAndPrint(studentsData)
	}
}

func WithHomeroomTeacherById(id int) {
	class := class.ById(id)
	if len(class) == 0 {
		return
	}

	teacher := teacher.ById(class["homeroom_teacher_id"].(int))

	fmt.Println("===================================")
	fmt.Printf("Class %v has a homeroom teacher %v\n", class["name"], teacher["name"])
	fmt.Println("======> Information: ")
	fmt.Println("===================================")
	base.SortAndPrint(class)
	base.SortAndPrint(teacher)
}
