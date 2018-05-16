package kwic

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

var fsm DataStorageManager = &FileBasedStorageManager{}
var dblpsmSuccess DataStorageManager = &DBLPStorageManager{}
var im IndexManager = IndexManager{}
var numLines int = 0
var connectionError error

func simulateDblpInput() {
	content := []byte("bonifacio")

	//creating tempfile
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}

	// clean up tempfile after finishing method
	defer os.Remove(tmpfile.Name())

	//writing content on tempfile
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	//saving contents of os.Stdin before assignin tempfile contents
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	connectionError = dblpsmSuccess.Init()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func simulateFsmInput() {
	content := []byte("../resources/papers.txt")

	//creating tempfile
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}

	// clean up tempfile after finishing method
	defer os.Remove(tmpfile.Name())

	//writing content on tempfile
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	//saving contents of os.Stdin before assignin tempfile contents
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	fsm.Init()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(t *testing.T) {
	simulateFsmInput()
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

	simulateDblpInput()
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
			t.Errorf("Erro, line length wrong, got: %v in line: %v", fsm.Line(line), line)
			break
		}
	}
}

func TestFileBasedStorageManagerLength(t *testing.T) {

	if fsm.Length() != numLines {
		t.Errorf("Erro, length of file is wrong, got: %d in line: %d", fsm.Length(), numLines)
	}
}

func TestFileBasedStorageManagerFailInput(t *testing.T) {
	var fsmFail = &FileBasedStorageManager{}
	content := []byte("wrong path")

	//creating tempfile
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}

	// clean up tempfile after finishing method
	defer os.Remove(tmpfile.Name())

	//writing content on tempfile
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	//saving contents of os.Stdin before assignin tempfile contents
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	inputError := fsmFail.Init()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	if inputError.Error() != "Não foi possível abrir o arquivo" {
		t.Errorf("Invalid input did not return an error message, got: %v, expected: %v", inputError.Error(), "Não foi possível abrir o arquivo")
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
	t.Logf("wods: %d", len(words))

	sortedWords := im.SortedWords()

	t.Logf("sortedWords: %s", sortedWords)

	lineAlphabetic := "igor kwic lindo teste"
	wordAlphabetic := strings.Split(lineAlphabetic, " ")

	t.Logf("wordAlphabetic: %s", wordAlphabetic)

	for i, w := range sortedWords {
		if w != wordAlphabetic[i] {
			t.Errorf("Erro, wrong sort, got: %s want: %s", w, wordAlphabetic[i])
		}
	}
}

func TestWordShiftShift(t *testing.T) {
	wordShift := WordShift{}
	line := "igor lindo demais gostoso e sensual"

	final := wordShift.Shift(strings.Split(line, " "), 2, 4)
	correct := []string{"e", "sensual", "igor", "lindo", "|", "demais", "|", "gostoso"}

	for i, _ := range final {
		if final[i] != correct[i] {
			t.Errorf("Erro, Right shifted words wrong, got: %s want: %s", final[i], correct[i])
		}
	}

	final = wordShift.Shift(strings.Split(line, " "), 3, 1)
	correct = []string{"demais", "|", "gostoso", "|", "e", "sensual", "igor", "lindo"}

	for i, _ := range final {
		if final[i] != correct[i] {
			t.Errorf("Erro, Left shifted words wrong, got: %s want: %s", final[i], correct[i])
		}
	}

	final = wordShift.Shift(strings.Split(line, " "), 1, 1)
	correct = []string{"igor", "|", "lindo", "|", "demais", "gostoso", "e", "sensual"}

	for i, _ := range final {
		if final[i] != correct[i] {
			t.Errorf("Erro, Left shifted words wrong, got: %s want: %s", final[i], correct[i])
		}
	}
}

func TestDBLPStorageManagerInit(t *testing.T) {
	dblpsmFailure := new(DBLPStorageManager)
	errorMessage := "Não foi encontrado o registro."
	connectionErrorMessage := "Não foi possível completar a requisição"

	if connectionError != nil && connectionError.Error() != connectionErrorMessage {
		t.Errorf("Error actual = %v, and Expected = %v.", connectionError.Error(), connectionErrorMessage)
	}

	content := []byte("aoisjdaoisdj")

	//creating tempfile
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}

	// clean up tempfile after finishing method
	defer os.Remove(tmpfile.Name())

	//writing content on tempfile
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	//saving contents of os.Stdin before assignin tempfile contents
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	notFoundError := dblpsmFailure.Init()

	if connectionError == nil && notFoundError.Error() != errorMessage {
		t.Errorf("Error actual = %v, and Expected = %v.", notFoundError.Error(), errorMessage)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestDBLPStoreManagerLine(t *testing.T) {
	if connectionError == nil {
		title := "Recipient size estimation for induction heating home appliances based on artificial neural networks."
		line := dblpsmSuccess.Line(0)

		if title != line {
			t.Errorf("Error, the file is empty or the reading is not correct, got: %v, want: %v", line, title)
		}
	}
}

func TestDBLPStorageManagerLength(t *testing.T) {
	if connectionError == nil {
		numLines := 30
		length := dblpsmSuccess.Length()
		t.Logf("Length: %d", dblpsmSuccess.Length())

		if length != numLines {
			t.Errorf("Error, the file is empty or the reading is not correct, got: %d, want: %d", length, numLines)
		}
	}
}
