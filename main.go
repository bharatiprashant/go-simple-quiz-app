package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")
	timeLimit := flag.Int("Limit", 30, "the time limit for the quiz in seconds")
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
problemLoop: // labels
	for i, p := range problems {
		fmt.Printf("Problem #%d:%s = ", i+1, p.q)
		answerCh := make(chan string) // making a channel https://gobyexample.com/channels

		/*
			this is a go routine and also anonymous
		*/
		go func() { // Go routines are functions or methods that run concurrently with other functions or methods.
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}() // call the go routine function
		select {
		case <-timer.C: // waiting for the mes sage from the channel, this line blocks/stop the program
			fmt.Println()
			// return
			break problemLoop
		case answer := <-answerCh:

			if answer == p.a {
				correct++
			}
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
