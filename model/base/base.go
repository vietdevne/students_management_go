package base

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

type Persions struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

func ReadJSON(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, errIo := io.ReadAll(jsonFile)
	if errIo != nil {
		fmt.Println(errIo)
	}
	return byteValue
}

func ReadData[T, K any](fileName, fieldName string, instance T) (chan K, int) {
	byteValue := ReadJSON(fileName)
	json.Unmarshal(byteValue, &instance)
	values := reflect.ValueOf(&instance).Elem().FieldByName(fieldName)

	modelChannel := make(chan K)
	go func() {
		for i := 0; i < values.Len(); i++ {
			modelChannel <- values.Index(i).Interface().(K)
		}
		close(modelChannel)
	}()

	return modelChannel, values.Len()
}

func ById(id int, data []map[string]interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	for _, dt := range data {
		if dt["id"] != nil && dt["id"] == id {
			for k, v := range dt {
				result[k] = v
			}
		}
	}

	if len(result) == 0{
		fmt.Println("\nĐM! Data not found!")
	}

	return result
}
