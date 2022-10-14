package process_tools

import (
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"sync"
)

type Converter struct {
	reporter Reporter
	filePath string
}

func (c *Converter) createOutputFile(output Output) error {
	templPath := "/pkg/templates/table.tmpl"
	templ := template.Must(template.New("Output").ParseFiles(templPath))
	OutputFile, err := os.OpenFile("output.html", os.O_WRONLY|os.O_CREATE, 0600)
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
				err := errors.New("Can't process entry")
				panic(err)
			}

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
