package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func readFile(fileNamePtr *string) map[string]string {
	quizDb := make(map[string]string, 0)
	csvfile, err := os.Open(*fileNamePtr)

	if err != nil {
		log.Fatalln("Couldn't open a file", err)
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quizDb[record[0]] = record[1]

	}
	return quizDb
}
func runQuiz(quizDb map[string]string, rightAnswers *int) int {
	// time.Sleep(5 * time.Second)
	var answer int
	intRightAnswer := 0
	for question, correctAnswer := range quizDb {
		fmt.Printf("Question %s ", question)
		fmt.Scanf("%d", &answer)
		fmt.Printf("Answer %s\n", correctAnswer)

		// string to int
		fmt.Sscan(correctAnswer, &intRightAnswer)
		// fmt.Printf("Current right answers %d", *rightAnswers)
		if answer == intRightAnswer {
			*rightAnswers++
		}
	}

	return 0
}

func main() {
	fileNamePtr := flag.String("f", "problems.csv", "quiz filename")
	timerPtr := flag.Int("t", 30, "quiz timer")
	flag.Parse()
	fmt.Println("filename:", *fileNamePtr)

	quizDb := readFile(fileNamePtr)
	var correctAnswerPtr *int
	correctAnswerPtr = new(int)
	c1 := make(chan int, 1)

	go func() {
		text := runQuiz(quizDb, correctAnswerPtr)
		c1 <- text
	}()

	select {
	case  <-c1:
	    fmt.Println("You did it in time!")
	case <-time.After(time.Duration(*timerPtr) * time.Second):
		fmt.Println("\n out of time :(")
	}

	fmt.Printf("Quiz results: Total question %d, right answers: %d ", len(quizDb), *correctAnswerPtr)

}
