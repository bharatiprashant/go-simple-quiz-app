package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")
	flag.Parse()
	// os package interacts with the operating system
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}
	r := csv.NewReader(file)  // Reading the csv type file, package is required
	lines, err := r.ReadAll() // Read all the column at once
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}

	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

/*
This function takes 2D array or we can say slices and return 1D array of type
structure problem
*/
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines { // lines is a 2D array so value of i=0 ,1,2 so on value of line will [5+10, 15] line[0] will be 5+10 and line[1] will 15
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // trims the space from the value
		}
	}
	return ret
}

/*
Using structure we can use any type of file like json or excel
*/

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // Error of some type
}
