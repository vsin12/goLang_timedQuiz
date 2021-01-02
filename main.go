package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	// using flag package to parse command line arguments and present user more description for the same
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s", *csvFilename))
	}

	//read file
	parsedContent := readFile(file)
	//create question/anwser structure -- problem
	problems := createProblemObjects(parsedContent)
	//start quiz
	playQuiz(problems, *timeLimit)
}

func playQuiz(problems []problem, timeLimit int) {
	correct := 0
	incorrect := 0
	// start timer
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		// channel to listen to input so that code is not blocked waiting for input while
		// the timer is completed
		answerCh := make(chan string)
		go func() {
			var answerInput string
			fmt.Scanf("%s \n", &answerInput)
			// send signal to the channel to notify for input
			answerCh <- answerInput
		}()

		select {
		// listen for timer signal on completion and print the result
		case <-timer.C:
			fmt.Printf("You got %d correct \n", correct)
			fmt.Printf("You got %d incorrect\n", incorrect)
			return
		// listen for inputSignal and score accordingly
		case answerInput := <-answerCh:
			if answerInput == p.answer {
				fmt.Println("Correct!!")
				correct++
			} else {
				fmt.Println("incorrect!!!")
				incorrect++
			}
		}
	}

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
