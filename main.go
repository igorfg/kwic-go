package main

import "fmt"
import "strings"
import "github.com/igorfg/kwic-go/kwic"

func main() {
	var fsm kwic.DataStorageManager = &kwic.FileBasedStorageManager{}
	var im kwic.IndexManager = kwic.IndexManager{}

	fsm.Init()
	im.Init()

	for lineNumber := 0; lineNumber < fsm.Length(); lineNumber++ {
		line := fsm.Line(lineNumber)
		words := strings.Split(line, " ")

		for pos := 0; pos < len(words); pos++ {
			im.Hash(words[pos], line, pos)
		}
	}

	// sortedWords := im.SortedWords()
	// wordShift := kwic.WordShift{}

	// for _, w := range sortedWords {
	// 	for _, tuple := range im.OccurencesOfWord(w) {

	// 	}
	// }

	for i := 0; i < fsm.Length(); i++ {
		fmt.Println(fsm.Line(i))
	}

	fmt.Println("Fim da Execução")
}
