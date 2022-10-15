package process_tools

// Info entity holds some general inforamtion about provided file.
// The parser and converter has the same instance of Info
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
