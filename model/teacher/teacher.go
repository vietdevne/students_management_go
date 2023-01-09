package teacher

import (
	base "student_management/model/base"
)

var Datas = getData()

type Teachers struct {
	Teachers []Teacher `json:"teachers"`
}

type Teacher struct {
	base.Persions
	HomeroomTeacher bool `json:"homeroom_teacher"`
}

func getData() []map[string]interface{} {
	var data []map[string]interface{}
	var teachers Teachers
	channel, size := base.ReadData[Teachers, Teacher]("db/teachers.json", "Teachers", teachers)

	for i := 0; i < size; i++ {
		c := <-channel
		teacher := map[string]interface{}{
			"id":               c.Id,
			"name":             c.Name,
			"age":              c.Age,
			"address":          c.Address,
			"gender":           c.Gender,
			"homeroom_teacher": c.HomeroomTeacher,
		}
		data = append(data, teacher)
	}

	return data
}

func ById(id int) map[string]interface{} {
	return base.ById(id, Datas)
}
