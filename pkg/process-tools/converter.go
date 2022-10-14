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
	templPath, err := realpath.Realpath("../pkg/templates/table.tmpl")
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
			// reading completed
			exit = true

		case entry := <-entries:
			if len(entry) == 0 {
				continue
			}
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
