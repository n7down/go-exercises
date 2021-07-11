package photos

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	layout = "2006-01-02 15:04:05"
)

type FileData struct {
	Extension string
	Name      string
	Timestamp time.Time
	Index     int
}

type FileDataIndexed struct {
	Name      string
	NameIndex int
	Extension string
	Index     int
}

func RenamePhotos(s string) string {
	var (
		fileDataList        []FileData
		fileDataMap         = make(map[string]int)
		fileDataIndexedList []FileDataIndexed
		b                   strings.Builder
	)

	photos := strings.Split(s, "\n")
	for i, photo := range photos {
		data := strings.Split(photo, ",")
		extension := strings.Split(strings.TrimSpace(data[0]), ".")
		str := strings.TrimSpace(data[2])
		timestamp, _ := time.Parse(layout, str)
		fileData := FileData{
			Extension: extension[1],
			Name:      strings.TrimSpace(data[1]),
			Timestamp: timestamp,
			Index:     i,
		}
		fileDataList = append(fileDataList, fileData)
	}

	// sort data by date
	sort.Slice(fileDataList, func(i, j int) bool {
		return fileDataList[i].Timestamp.Before(fileDataList[j].Timestamp)
	})

	// index the data
	for _, f := range fileDataList {
		var i int
		if v, ok := fileDataMap[f.Name]; ok {
			v = v + 1
			fileDataMap[f.Name] = v
			i = v
		} else {
			fileDataMap[f.Name] = 1
			i = 1
		}

		fileDataIndexed := FileDataIndexed{
			Name:      f.Name,
			NameIndex: i,
			Extension: f.Extension,
			Index:     f.Index,
		}

		fileDataIndexedList = append(fileDataIndexedList, fileDataIndexed)
	}

	// sort by index
	sort.Slice(fileDataIndexedList, func(i, j int) bool {
		return fileDataIndexedList[i].Index < (fileDataIndexedList[j].Index)
	})

	for _, f := range fileDataIndexedList {
		b.WriteString(fmt.Sprintf("%s%d.%s\n", f.Name, f.NameIndex, f.Extension))
	}

	return b.String()
}
