package kwic

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

var fsm DataStorageManager = &FileBasedStorageManager{}
var im IndexManager = IndexManager{}
var numLines int = 0

func TestMain(t *testing.T) {
	fsm.Init()
	im.Init()

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

func TestFileBasedStorageManageInit(t *testing.T) {

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

func TestFileBasedStorageManagerLength(t *testing.T) {

	if fsm.Length() != numLines {
		t.Errorf("Erro, length of file is wrong, got: %d in line: %d", fsm.Length(), numLines)
	}
}

func TestIndexManagerInit(t *testing.T) {

	if !im.IsEmpty() {
		t.Errorf("Erro, type of indexManager is worng, got: %s want: %s", "kwic.IndexManager", reflect.TypeOf(im))
	}
}

func TestIndexManagerHashOccurences(t *testing.T) {
	line := "Teste testando kwic"
	words := strings.Split(line, " ")

	im.Hash(words[1], line, 1)

	tuple := im.OccurencesOfWord("testando")

	if len(tuple) != 1 {
		t.Errorf("Erro, wrong incertion or count words, got: %d want: %d", len(tuple), 1)
	}

	line2 := "Outro testando do kwic"
	words2 := strings.Split(line, " ")

	im.Hash(words2[1], line2, 1)

	tuple = im.OccurencesOfWord("testando")

	if len(tuple) != 2 {
		t.Errorf("Erro, wrong incertion or count words, got: %d want: %d", len(tuple), 2)
	}
}

func TestIndexManagerSortedWords(t *testing.T) {
	im.Init()

	line := "teste kwic igor lindo"
	words := strings.Split(line, " ")

	for pos := 0; pos < len(words); pos++ {
		im.Hash(words[pos], line, pos)
	}

	sortedWords := im.SortedWords()

	lineAlphabetic := "igor kwic lindo teste"
	wordAlphabetic := strings.Split(lineAlphabetic, " ")

	t.Log(sortedWords)

	i := 0
	for _, w := range sortedWords {
		if w != wordAlphabetic[i] {
			t.Errorf("Erro, wrong sort, got: %s want: %s", w, wordAlphabetic[i])
		}
		i++
	}
}
