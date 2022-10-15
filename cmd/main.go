/*
This programm processes your "csv" or "prn" files and
creates HTML representation. If it is csv file, then
the table in HTML will be created, in other case of "prn" file - the
simple text will be represented in HTML format

Usage:

	main.go [path] [delimeter]

path:
	Path to file with "csv" or "prn" extensions.

delimters:
	Delimeter might be set only if path leads to csv file.
	By default the comma "," delimter is set, but you can define any
	other delimeter except "\n" and "\r".
*/

package main

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	tools "github.com/ssergomol/data-viewer/pkg/process-tools"
)

func startProcessing(filePath string, delim rune, ext string) {
	// Create information reporter, parser and converter for data processing
	reporter := tools.CreateInfo(filePath, delim, ext)
	parser := tools.NewParser(filePath, reporter)
	converter := tools.NewConverter(filePath, reporter)

	// Create wait group to wait untill all goroutines end their execution
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// create channels for goroutine communication
	entries := make(chan []string)
	done := make(chan bool)

	// The parser reads the file line-by-line. After reading the line sends
	// it to converter via channel.

	go converter.ProcessEntry(wg, entries, done)
	go parser.Read(wg, entries, done)

	// wait untill all goroitines finish
	wg.Wait()
}

func main() {
	// Check if the path to file was provided
	if len(os.Args[1:]) == 0 {
		err := errors.New("no filepathes were provided!")
		panic(err)
	}

	// Get Inforamtion about the file
	fileName := os.Args[1]
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	// Check if the file is not a directory
	if fileInfo.IsDir() {
		err := errors.New("can't read the directiory")
		panic(err)
	}

	// Get the extension and provided delimeter
	fileExt := filepath.Ext(fileName)
	if fileExt != ".prn" && fileExt != ".csv" {
		err := errors.New("Wrong file extension")
		panic(err)
	}
	delimeter := ','
	if len(os.Args[2:]) != 0 {
		delimeter = rune(os.Args[2][0])
	}
	startProcessing(fileName, delimeter, fileExt)
}
