package main

import (
	"errors"
	"os"
	"sync"

	tools "github.com/ssergomol/data-viewer/pkg/process-tools"
)

func Process(filePath string, delim rune) {
	// TODO: create reporter
	reporter := tools.NewReporter(filePath, delim)

	// TODO: create parser
	parser := tools.NewParser(filePath, reporter)
	// TODO: create waitgropup

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// create channels for goroutine comminication
	entries := make(chan []string)
	done := make(chan bool)

	converter := tools.NewConverter(filePath, reporter)

	go converter.ProcessEntry(wg, entries, done)
	go parser.Read(wg, entries, done)

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

	if len(os.Args[2:]) == 0 {
		Process(fileName, ',')
	} else {
		Process(fileName, rune(os.Args[2][0]))
	}
}
