package process

import (
	"encoding/csv"
	"io"
	"os"
	"sync"

	tools "github.com/ssergomol/data-viewer/pkg/process-tools/reporter"
)

type Parser struct {
	file     string
	reporter tools.Reporter
}

func CreateCSVParser(filePath string, reporter tools.Reporter) Parser {
	return Parser{
		file:     filePath,
		reporter: reporter,
	}
}

func parseHeaders(r *csv.Reader) {

}

func (p *Parser) Read(wg *sync.WaitGroup, entries chan []string, done chan bool) {
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
