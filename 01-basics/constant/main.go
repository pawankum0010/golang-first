package main

import (
	"fmt"
	"log"
	"goLangFirst/utils/myutil"
	"time"
)

const age int = 30
const name string = "Pawan Kumar" //Pawan

func main() {
	const (
		city      string = "Bangalore"
		country   string = "India"
		education string = "B.Tech"
	)

	logger, closeLogger, err := myutil.NewExecutionLogger()
	if err != nil {
		log.Fatalf("failed to setup logger: %v", err)
	}
	defer closeLogger()

	programStart := time.Now()
	logger.Println("program started")

	myutil.RunStep(logger, "print age", func() {
		fmt.Println("This is my age: ", age)
	})
	myutil.RunStep(logger, "print name", func() {
		fmt.Println("This is my name: ", name)
	})
	myutil.RunStep(logger, "print city", func() {
		fmt.Println("This is my city: ", city)
	})
	myutil.RunStep(logger, "print country", func() {
		fmt.Println("This is my country: ", country)
	})
	myutil.RunStep(logger, "print education", func() {
		fmt.Println("This is my education: ", education)
	})
	myutil.RunStep(logger, "print age again", func() {
		fmt.Println("This is my age: ", age)
	})

	logger.Printf("program completed in %s", myutil.FormatDuration(time.Since(programStart)))
}
