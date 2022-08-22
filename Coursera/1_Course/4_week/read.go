package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func main() {
	var fileName string

	fmt.Println("Introduce name of the file you want to parse")
	fmt.Scan(&fileName)

	fileData, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		personsObj := make([]Person, 2)
		personsString := strings.Split(string(fileData), "\n")

		for _, person := range personsString {
			userData := strings.Split(person, " ")

			personObj := new(Person)
			personObj.fName = userData[0]
			personObj.lName = userData[1]

			personsObj = append(personsObj, *personObj)

			fmt.Println("FirstName: " + personObj.fName + "  LastName: " + personObj.lName)
		}
	}
}
