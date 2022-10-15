package main

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	tools "github.com/ssergomol/data-viewer/pkg/process-tools"
)

func startProcessing(filePath string, delim rune, ext string) {
	// TODO: create reporter
	reporter := tools.CreateInfo(filePath, delim, ext)

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

	fileExt := filepath.Ext(fileName)

	delimeter := ','
	if len(os.Args[2:]) != 0 {
		delimeter = rune(os.Args[2][0])
	}
	startProcessing(fileName, delimeter, fileExt)
}
