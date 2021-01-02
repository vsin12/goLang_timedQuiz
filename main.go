package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {

	// using flag package to parse command line arguments and present user more description for the same
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s", *csvFilename))
	}

	parsedContent := readFile(file)
	problems := createProblemObjects(parsedContent)
	playQuiz(problems)
}

func playQuiz(problems []problem) {

	correct := 0
	incorrect := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answerInput string
		fmt.Scanf("%s \n", &answerInput)
		if answerInput == p.answer {
			fmt.Println("Correct!!")
			correct++
		} else {
			fmt.Println("incorrect!!!")
			incorrect++
		}
	}

	fmt.Printf("You got %d correct \n", correct)
	fmt.Printf("You got %d incorrect\n", incorrect)
}

func createProblemObjects(parsedContent [][]string) []problem {
	problemObject := make([]problem, len(parsedContent))

	for i, line := range parsedContent {
		problemObject[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problemObject
}

func readFile(csvFile *os.File) [][]string {
	r := csv.NewReader(csvFile)
	parsedCSV, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV file")
	}

	return parsedCSV
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
