package student

import (
	base "student_management/model/base"
)

var Datas = getData()

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	base.Persions
}

func getData() []map[string]interface{} {
	var data []map[string]interface{}
	var students Students

	channel, size := base.ReadData[Students, Student]("db/students.json", "Students", students)

	for i := 0; i < size; i++ {
		c := <-channel
		student := map[string]interface{}{
			"id":      c.Id,
			"name":    c.Name,
			"age":     c.Age,
			"address": c.Address,
			"gender":  c.Gender,
		}
		data = append(data, student)
	}

	return data
}

func ById(id int) map[string]interface{} {
	return base.ById(id, Datas)
}

func AllName() []string {
	result := make([]string, 0, len(Datas))
	for _, std := range Datas {
		result = append(result, std["name"].(string))
	}
	return result
}
