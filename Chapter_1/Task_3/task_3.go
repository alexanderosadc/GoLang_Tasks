package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scaner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number of children in First School:")
	scaner.Scan()
	var nrFirstSchoolString = scaner.Text()
	var nrFirstSchoolInt, _ = strconv.Atoi(nrFirstSchoolString)

	fmt.Print("Enter number of children in Second School:")
	scaner.Scan()
	var nrSecondSchoolString = scaner.Text()
	var nrSecondSchoolInt, _ = strconv.Atoi(nrSecondSchoolString)

	var totalNrOfChildren = CalculateChildren(nrFirstSchoolInt, nrSecondSchoolInt)
	fmt.Print("Total number of children: ", totalNrOfChildren)
}

func CalculateChildren(nrFirstSchool, nrSecondSchool int) int {
	return nrFirstSchool + nrSecondSchool + 7
}
