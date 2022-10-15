package process_tools

import (
	"html/template"
	"os"
	"path/filepath"
	"sync"

	"github.com/yookoala/realpath"
)

type Converter struct {
	reporter *Info
	filePath string
}

func NewConverter(path string, reporter *Info) *Converter {
	return &Converter{
		reporter: reporter,
		filePath: path,
	}
}

func (c *Converter) createOutputFile(output Output) error {
	templName := "csv.tmpl"
	if c.reporter.FileExt == ".prn" {
		templName = "prn.tmpl"
	}
	templPath, err := realpath.Realpath("../pkg/templates/" + templName)
	if err != nil {
		panic(err)
	}

	name := filepath.Base(templPath)
	templ, err := template.New(name).ParseFiles(templPath)
	if err != nil {
		panic(err)
	}

	OutputFile, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}
	// Apply template to create html output
	return templ.Execute(OutputFile, output)
}

func (c *Converter) ProcessEntry(wg *sync.WaitGroup, entries <-chan []string, done <-chan bool) {
	defer func() {
		wg.Done()
	}()

	var exit bool
	var data [][]string

	for {
		select {
		case <-done:
			// Reading complete, so we can finish cycle iterating
			exit = true

		case entry := <-entries:
			if len(entry) == 0 {
				// If empty line is found, we should skip it
				continue
			}
			// Append data, which was provided by parser
			data = append(data, entry)
		}

		if exit {
			break
		}
	}

	err := c.createOutputFile(Output{
		FileName:    filepath.Base(c.filePath),
		HeadersNumb: len(c.reporter.Headers),
		Headers:     c.reporter.Headers,
		Data:        data,
	})

	if err != nil {
		panic(err)
	}

}
