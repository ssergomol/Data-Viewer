package process_tools

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type Parser struct {
	file     string
	reporter *Info
}

func NewParser(filePath string, reporter *Info) Parser {
	return Parser{
		file:     filePath,
		reporter: reporter,
	}
}

func (p *Parser) parseCSVHeaders(r *csv.Reader) error {
	headers, err := r.Read()
	p.reporter.Headers = headers
	return err
}

func (p *Parser) parsePRNHeaders(scanner *bufio.Scanner) {
	scanner.Scan()
	p.reporter.Headers = append(p.reporter.Headers, scanner.Text())
}

func (p *Parser) Read(wg *sync.WaitGroup, entries chan<- []string, done chan<- bool) {
	defer func() {
		close(done)
		close(entries)
		wg.Done()
	}()

	file, err := os.OpenFile(p.file, os.O_RDONLY, 755)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	// Parse headers according to defined extension
	if p.reporter.FileExt == ".prn" {
		scanner := bufio.NewScanner(file)
		p.parsePRNHeaders(scanner)

		for scanner.Scan() {
			entries <- []string{scanner.Text()}
		}

	} else {
		reader := csv.NewReader(file)
		reader.LazyQuotes = true
		reader.Comma = p.reporter.Delimeter
		reader.TrimLeadingSpace = true

		if err := p.parseCSVHeaders(reader); err != nil {
			logrus.Fatal(err)
		}

		// read line by line
		for {
			entry, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				logrus.Fatal(err)
			}

			entries <- entry
		}
	}

}
