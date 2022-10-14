package process_tools

import (
	"encoding/csv"
	"io"
	"os"
	"sync"
)

type Parser struct {
	file     string
	reporter *Reporter
}

func NewParser(filePath string, reporter *Reporter) Parser {
	return Parser{
		file:     filePath,
		reporter: reporter,
	}
}

func (p *Parser) parseHeaders(r *csv.Reader) error {
	headers, err := r.Read()
	p.reporter.Headers = headers
	return err
}

func (p *Parser) Read(wg *sync.WaitGroup, entries chan<- []string, done chan<- bool) {
	defer func() {
		close(done)
		close(entries)
		wg.Done()
	}()

	// open file for reading

	file, err := os.OpenFile(p.file, os.O_RDONLY, 755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parse headers
	reader := csv.NewReader(file)

	if err := p.parseHeaders(reader); err != nil {
		panic(err)
	}

	// read line by line
	for {
		entry, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// notify reporter that process is finished
		// c.reporter.RecordProcessed()
		entries <- entry
	}
}
