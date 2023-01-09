package base

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func All(data []map[string]interface{}) {
	for _, dt := range data {
		for k, v := range dt {
			fmt.Printf("%v: %v\n", k, v)
		}
		fmt.Println("=================================")
	}
}

func InputMenuOption() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("===> select: ")
	option, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	return strings.Trim(option, "\n")
}

func InputIdValue(objectName string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("===> enter %v id: ", objectName)
	input, _ := reader.ReadString('\n')
	idStr := strings.Trim(input, "\n")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("\nĐM! Please enter valid value!!!")
	}
	return id
}

func InputNumberValue() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("===> enter number: ")
	input, _ := reader.ReadString('\n')
	numStr := strings.Trim(input, "\n")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("\nĐM! Please enter valid value!!!")
	}
	return num
}

func SortAndPrint(data map[string]interface{}) {
	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, data[k])
	}
	fmt.Println("=============================")
}
