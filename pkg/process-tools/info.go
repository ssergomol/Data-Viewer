package process_tools

type Info struct {
	Headers   []string
	FilePath  string
	Delimeter rune
	FileExt   string
}

func CreateInfo(path string, delim rune, ext string) *Info {
	headers := make([]string, 0)
	return &Info{
		Headers:   headers,
		FilePath:  path,
		Delimeter: delim,
		FileExt:   ext,
	}
}
