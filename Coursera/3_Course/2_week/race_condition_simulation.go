// Definition: Race condition - is an issue which appears in concurrent programming
//where the outcome depends on the order of the execution of goroutines or threads.

// When it can appear:
// * When we have shared data (variable) and one concurrent task changes it and another reads it;
// * When we want to print 2 different strings in concurrent tasks they may not be printed together
//but in the order we do not know;

package main

import (
	"fmt"
	"time"
)

// Example of race condtion when writing name and surename letter by letter
func main() {
	userName := "Alex"
	userSurename := "Osa"
	go PrintUserNameByLetter(&userName)
	go PrintUserSurenameByLetter(&userSurename)
	time.Sleep(time.Millisecond * 100)
}

func PrintUserNameByLetter(name *string) {
	for _, letter := range *name {
		fmt.Println("Name Letter: ", string(letter))
		time.Sleep(time.Millisecond * 10)
	}
}

func PrintUserSurenameByLetter(surename *string) {
	for _, letter := range *surename {
		fmt.Println("Surename Letter: ", string(letter))
		time.Sleep(time.Millisecond * 10)
	}
}
