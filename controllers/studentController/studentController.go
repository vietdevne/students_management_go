package studentController

import (
	"fmt"
	"sort"
	base "student_management/controllers/base"
	"student_management/model/class"
	student "student_management/model/student"
	studentClass "student_management/model/student_class"
)

func All()  {
	base.All(student.Datas)
}

func WithClassById(id int)  {
	student := student.ById(id)
	if len(student) == 0 {
		return
	}

	classStudents := studentClass.ByStudentId(id)
	var classes []map[string]interface{}

	for _, classStudentsData := range classStudents {
		classes = append(classes, class.ById(classStudentsData["class_id"].(int)))
	}

	fmt.Println("===================================")
	fmt.Printf("Student %v had join classes: \n", student["name"])
	fmt.Println("===================================")
	fmt.Println("Class Information:")
	for _, classesData := range classes {
		base.SortAndPrint(classesData)
	}
}

func SortByNameASC()  {
	names := student.AllName()
	sort.Strings(names)

	i:= 1
	for _, name := range names {
		fmt.Printf("\n%v. %v\n", i, name)
		i++
	}
	fmt.Println("===================================")
}

func SortByNameDESC()  {
	names := student.AllName()
	sort.Sort(sort.Reverse(sort.StringSlice(names)))

	i:=1
	for _, name := range names {
		fmt.Printf("\n%v. %v\n", i, name)
		i++
	}
	fmt.Println("===================================")
}
