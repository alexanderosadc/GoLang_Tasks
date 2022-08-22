package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	animalName       string
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (animal *Animal) Speak() {
	fmt.Println("Sound: ", animal.spokenSound)
}

func (animal *Animal) Eat() {
	fmt.Println("Food Eaten: ", animal.foodEaten)
}

func (animal *Animal) Move() {
	fmt.Println("Locomotion Method:", animal.locomotionMethod)
}

func main() {
	arrOfAnimals := []Animal{}
	cow := Animal{animalName: "cow", foodEaten: "grass",
		locomotionMethod: "walk", spokenSound: "moo"}
	bird := Animal{animalName: "bird", foodEaten: "worms",
		locomotionMethod: "fly", spokenSound: "peep"}
	snake := Animal{animalName: "snake", foodEaten: "mice",
		locomotionMethod: "slither", spokenSound: "hsss"}

	arrOfAnimals = append(arrOfAnimals, cow, bird, snake)

	for {
		animalName, animalCommand := UserInput()
		if animalName == nil || animalCommand == nil {
			continue
		}

		animalObject := ReturnAnimalIfExist(arrOfAnimals, animalName)
		if animalObject == nil {
			fmt.Println("There is no such animal!")
			continue
		}

		ExecuteCommand(animalObject, animalCommand)
	}

}

// Responsible for executing method of the object Animal based on the animalCommand string
func ExecuteCommand(animalObj *Animal, animalCommand *string) {
	switch *animalCommand {
	case "eat":
		animalObj.Eat()
	case "move":
		animalObj.Move()
	case "speak":
		animalObj.Speak()
	default:
		fmt.Println("There is no such command!")
	}
}

// Responsible for getting user input and returning 2 strings: first is animalName, second animalCommand
func UserInput() (*string, *string) {
	fmt.Println("Introduce command to the prompt")
	fmt.Print(">")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return nil, nil
	}

	commands := strings.Split(command, " ")
	animalName := strings.TrimSpace(commands[0])
	animalCommand := strings.TrimSpace(commands[1])

	return &animalName, &animalCommand
}

// Respinsible for returning animal if animal with animalName exist in the slice of Animal objects
func ReturnAnimalIfExist(animalsSlice []Animal, animalName *string) *Animal {
	for _, animal := range animalsSlice {
		if animal.animalName == *animalName {
			return &animal
		}
	}

	return nil
}
