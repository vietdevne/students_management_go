package studentClass

import (
	"student_management/model/base"
)

var Datas = getData()

type StudentClasses struct {
	StudentClasses []StudentClass `json:"student_classes"`
}

type StudentClass struct {
	StudentId int `json:"student_id"`
	ClassId   int `json:"class_id"`
}

func getData() []map[string]interface{} {
	var data []map[string]interface{}
	var stdClasses StudentClasses
	channel, size := base.ReadData[StudentClasses, StudentClass]("db/student_class.json", "StudentClasses", stdClasses)
	for i := 0; i < size; i++ {
		c := <-channel
		stdClass := map[string]interface{}{
			"student_id": c.StudentId,
			"class_id":   c.ClassId,
		}
		data = append(data, stdClass)
	}

	return data
}

func ByStudentId(studentId int) []map[string]interface{} {
	var result []map[string]interface{}
	for _, data := range Datas {
		if data["student_id"] == studentId {
			result = append(result, map[string]interface{}{"student_id": studentId, "class_id": data["class_id"]})
		}
	}
	return result
}

func ByClassId(classId int) []map[string]interface{} {
	var result []map[string]interface{}

	for _, data := range Datas {
		if data["class_id"] == classId {
			result = append(result, map[string]interface{}{"student_id": data["student_id"], "class_id": classId})
		}
	}
	return result
}
