package kwic

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/igorfg/kwic-go/kwic"
)

var fsm kwic.DataStorageManager = &kwic.FileBasedStorageManager{}
var im kwic.IndexManager = kwic.IndexManager{}
var numLines int = 0

func TestMain(t *testing.T) {
	fsm.Init()

	file, err := os.Open("../resources/papers.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numLines++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestFileBasedStorageManagerInit(t *testing.T) {

	length := fsm.Length()

	if length != numLines {
		t.Errorf("Erro, the file is empty or the reading are not correct, got: %d, want: %d", length, numLines)
	}
}

func TestFileBasedStorageManagerLine(t *testing.T) {

	for line := 0; line < fsm.Length(); line++ {
		if len(fsm.Line(line)) == 0 {
			t.Errorf("Erro, line length wrong, got: %d in line: %d", fsm.Line(line), line)
			break
		}
	}
}
