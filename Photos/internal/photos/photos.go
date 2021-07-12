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

type FileNameCounter struct {
	fileNameCounterMap map[string]int
}

func NewFileNameCounter() *FileNameCounter {
	f := &FileNameCounter{
		fileNameCounterMap: make(map[string]int),
	}
	return f
}

func (f *FileNameCounter) AddFileName(fileName string) {
	if val, ok := f.fileNameCounterMap[fileName]; ok {
		val = val + 1
		f.fileNameCounterMap[fileName] = val
	} else {
		f.fileNameCounterMap[fileName] = 1
	}
}

func (f *FileNameCounter) GetFormatedIndex(fileName string, index int) string {
	if index >= 10 {
		return fmt.Sprintf("%d", index)
	}

	if val, ok := f.fileNameCounterMap[fileName]; ok {
		if val >= 10 {
			return fmt.Sprintf("0%d", index)
		} else {
			return fmt.Sprintf("%d", index)
		}
	} else {
		return fmt.Sprintf("%d", index)
	}
}

func RenamePhotos(s string) string {
	var (
		fileDataList        []FileData
		fileDataMap         = make(map[string]int)
		fileDataIndexedList []FileDataIndexed
		b                   strings.Builder
		fileNameCounter     *FileNameCounter
	)

	fileNameCounter = NewFileNameCounter()

	photos := strings.Split(s, "\n")
	for i, photo := range photos {
		data := strings.Split(photo, ",")
		extension := strings.Split(strings.TrimSpace(data[0]), ".")
		name := strings.TrimSpace(data[1])
		str := strings.TrimSpace(data[2])
		timestamp, _ := time.Parse(layout, str)
		fileData := FileData{
			Extension: extension[1],
			Name:      name,
			Timestamp: timestamp,
			Index:     i,
		}
		fileDataList = append(fileDataList, fileData)

		fileNameCounter.AddFileName(name)
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
		nameIndex := fileNameCounter.GetFormatedIndex(f.Name, f.NameIndex)
		b.WriteString(fmt.Sprintf("%s%s.%s\n", f.Name, nameIndex, f.Extension))
	}

	return b.String()
}
