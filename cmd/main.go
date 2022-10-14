package main

import (
	"errors"
	"os"
	"sync"
)

func Process() {
	// TODO: create reporter

	// TODO: create parser

	// TODO: create waitgropup

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// create channels for goroutine comminication

	entry := make(chan []string)
	done := make(chan []bool)

	// TODO: run goroutine to convert data

	// TODO: run goroutine to read csv data

	// wait untill all goroitines finish
	wg.Wait()
}

func main() {
	if len(os.Args[1:]) == 0 {
		err := errors.New("no filepathes were provided!")
		panic(err)
	}

	fileName := os.Args[1]

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		err := errors.New("can't read the directiory")
		panic(err)
	}

	Process()
}
