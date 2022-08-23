package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Speak()
	Eat()
	Move()
	GetAnimalName() *string
}

type Cow struct {
	animalName       string
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

type Bird struct {
	animalName       string
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

type Snake struct {
	animalName       string
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (animal *Cow) GetAnimalName() *string {
	return &animal.animalName
}

func (animal *Cow) Speak() {
	fmt.Println("Sound: ", animal.spokenSound)
}

func (animal *Cow) Eat() {
	fmt.Println("Food Eaten: ", animal.foodEaten)
}

func (animal *Cow) Move() {
	fmt.Println("Locomotion Method:", animal.locomotionMethod)
}

func (animal *Bird) GetAnimalName() *string {
	return &animal.animalName
}

func (animal *Bird) Speak() {
	fmt.Println("Sound: ", animal.spokenSound)
}

func (animal *Bird) Eat() {
	fmt.Println("Food Eaten: ", animal.foodEaten)
}

func (animal *Bird) Move() {
	fmt.Println("Locomotion Method:", animal.locomotionMethod)
}

func (animal *Snake) GetAnimalName() *string {
	return &animal.animalName
}

func (animal *Snake) Speak() {
	fmt.Println("Sound: ", animal.spokenSound)
}

func (animal *Snake) Eat() {
	fmt.Println("Food Eaten: ", animal.foodEaten)
}

func (animal *Snake) Move() {
	fmt.Println("Locomotion Method:", animal.locomotionMethod)
}

func main() {
	arrOfAnimals := []Animal{}

	for {
		queryCommand, animalName, animalCommand := UserInput()
		if animalName == nil ||
			animalCommand == nil ||
			queryCommand == nil {
			continue
		}

		switch *queryCommand {
		case "newanimal":
			newAnimal, err := CreateNewAnimal(animalName, animalCommand)

			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				continue
			}

			arrOfAnimals = append(arrOfAnimals, newAnimal)
		case "query":
			ExecuteAnimalCommand(animalName, animalCommand, arrOfAnimals)
		}
		// animalObject := ReturnAnimalIfExist(arrOfAnimals, animalName)
		// if animalObject == nil {
		// 	fmt.Println("There is no such animal!")
		// 	continue
		// }

		//ExecuteCommand(animalObject, animalCommand)
	}

}
func ExecuteAnimalCommand(animalName *string, animalCommand *string, arrAnimal []Animal) error {
	var err error = nil

	if len(arrAnimal) == 0 {
		err = errors.New("there are no animals registered in the system")
		return err
	}

	for _, animal := range arrAnimal {
		if *animal.GetAnimalName() == *animalName {
			err = ExecuteCommand(animal, animalCommand)
			return err
		}
	}

	return err
}

// Responsible for getting user input and returning 2 strings: first is animalName, second animalCommand
func UserInput() (*string, *string, *string) {
	fmt.Println("Introduce command to the prompt")
	fmt.Print(">")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return nil, nil, nil
	}

	commands := strings.Split(command, " ")
	queryCommand := strings.TrimSpace(commands[0])
	animalName := strings.TrimSpace(commands[1])
	animalCommand := strings.TrimSpace(commands[2])

	return &queryCommand, &animalName, &animalCommand
}

func CreateNewAnimal(animalName *string, animalType *string) (Animal, error) {

	if animalName == nil {
		return nil, errors.New("Animal Name is null")
	}

	switch *animalType {
	case "cow":
		return &Cow{animalName: *animalName, foodEaten: "grass",
				locomotionMethod: "walk", spokenSound: "moo"},
			nil
	case "bird":
		return &Bird{animalName: *animalName, foodEaten: "worms",
				locomotionMethod: "fly", spokenSound: "peep"},
			nil
	case "snake":
		return &Snake{animalName: *animalName, foodEaten: "worms",
				locomotionMethod: "fly", spokenSound: "peep"},
			nil
	default:
		return nil, errors.New("there is no such type of animal")
	}
}

// Responsible for executing method of the object Animal based on the animalCommand string
func ExecuteCommand(animalObj Animal, animalCommand *string) error {
	switch *animalCommand {
	case "eat":
		animalObj.Eat()
		return nil
	case "move":
		animalObj.Move()
		return nil
	case "speak":
		animalObj.Speak()
		return nil
	default:
		return errors.New("there is no such command")
	}
}
