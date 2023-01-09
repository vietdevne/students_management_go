package class

import (
	"student_management/model/base"
)

var Datas = getData()

type Classes struct {
	Classes []Class `json:"classes"`
}

type Class struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	HomeroomTeacherId int    `json:"homeroom_teacher_id"`
}

func getData() []map[string]interface{} {
	var datas []map[string]interface{}

	var classes Classes
	channel, size := base.ReadData[Classes, Class]("db/classes.json", "Classes", classes)
	for i := 0; i < size; i++ {
		c := <-channel
		class := map[string]interface{}{
			"id":                  c.Id,
			"name":                c.Name,
			"homeroom_teacher_id": c.HomeroomTeacherId,
		}
		datas = append(datas, class)
	}

	return datas
}

func ById(id int) map[string]interface{} {
	return base.ById(id, Datas)
}

func ByHomeroomTeacherId(id int) map[string]interface{} {
	var data map[string]interface{}
	for _, dt := range Datas {
		if dt["homeroom_teacher_id"] == id {
			data = dt
			break
		}
	}

	return data
}
