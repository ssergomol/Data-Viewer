package process_tools

type Reporter struct {
	Headers  []string
	FilePath string
}

func NewReporter(path string) *Reporter {
	headers := make([]string, 0)
	return &Reporter{
		Headers:  headers,
		FilePath: path,
	}
}
