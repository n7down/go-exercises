package photos

// you can also use imports, for example:
import (
	"strings"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type FileData struct {
	Format string
	Name   string
	Date   string
	Index  int
}

func RenamePhotos(s string) string {
	// write your code in Go 1.4
	var f []FileData
	photos := strings.Split(s, "\n")
	for _, photo := range photos {
		data := strings.Split(photo, ",")
		fileData := FileData{
			Format: strings.TrimSpace(data[0]),
			Name:   strings.TrimSpace(data[1]),
			Date:   strings.TrimSpace(data[2]),
		}
		f = append(f, fileData)
	}

	return ""
}
