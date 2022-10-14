package process_tools

type Info struct {
	Headers   []string
	FilePath  string
	Delimeter rune
}

func CreateInfo(path string, delim rune) *Info {
	headers := make([]string, 0)
	return &Info{
		Headers:   headers,
		FilePath:  path,
		Delimeter: delim,
	}
}
