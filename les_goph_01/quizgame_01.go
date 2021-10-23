package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileNamePtr := flag.String("f", "problems.csv", "quiz filename")
	flag.Parse()
	fmt.Println("filename:", *fileNamePtr)

	csvfile, err := os.Open(*fileNamePtr)

	if err != nil {
		log.Fatalln("Couldn't open a file", err)
	}
	r := csv.NewReader(csvfile)

	rightAnswers := 0
	totalQuestions :=0
	var answer int
	for {
		record, err := r.Read()
		if err == io.EOF {

			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question %s ", record[0])
		fmt.Scanf("%d", &answer)
		fmt.Printf("Answer %s\n", record[1])
		totalQuestions++

		intRightAnswer := 0
		fmt.Sscan(record[1], &intRightAnswer)
		if answer == intRightAnswer {
			rightAnswers++
		}

	}
	fmt.Printf("Quiz results: Total question %d, right answers: %d ", totalQuestions,rightAnswers)
}

