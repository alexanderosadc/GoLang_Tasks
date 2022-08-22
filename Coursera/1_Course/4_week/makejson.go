package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var inputData string
	userMap := make(map[string]string)

	fmt.Println("Introduce your name")
	fmt.Scan(&inputData)
	userMap["name"] = inputData

	fmt.Println("Introduce your address")
	fmt.Scan(&inputData)
	userMap["address"] = inputData
	fmt.Println("Introduce your adress")

	userJson, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		fmt.Println(string(userJson))
	}
}
