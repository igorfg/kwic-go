package kwic

import "fmt"
import "bufio"
import "os"
import "log"

type FileBasedStorageManager struct {
	lines []string
}

func (f *FileBasedStorageManager) Init() {
	var filePath string

	fmt.Print("Enter the path to the input file: ")
	fmt.Scan(&filePath)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.lines = append(f.lines, scanner.Text())
	}
}

func (f *FileBasedStorageManager) Line(index int) string {
	return f.lines[index]
}

func (f *FileBasedStorageManager) Length() int {
	return len(f.lines)
}
