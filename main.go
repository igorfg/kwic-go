package main

import (
	"fmt"
	"github.com/igorfg/kwic-go/kwic"
	"log"
	"strings"
)

func main() {
	var (
		storageManager kwic.DataStorageManager
		indexManager   kwic.IndexManager
		outputManager  kwic.OutputManager
		winc           []string
	)

	storageManager = kwic.DataStorageManager{}
	outputManager = kwic.OutputManager{}
	indexManager = kwic.IndexManager{}

	err := storageManager.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	indexManager.Init()

	for lineNumber := 0; lineNumber < storageManager.Length(); lineNumber++ {
		line := storageManager.Line(lineNumber)
		words := strings.Split(line, " ")

		for pos := 0; pos < len(words); pos++ {
			indexManager.Hash(words[pos], line, pos)
		}
	}

	sortedWords := indexManager.SortedWords()

	wordShift := kwic.WordShift{}

	for _, w := range sortedWords {
		for _, tuple := range indexManager.OccurencesOfWord(w) {
			func(line string, pos int) {
				winc = append(winc, (strings.Join(wordShift.Shift(strings.Split(line, " "), pos, pos), " ")))
			}(tuple.First.(string), tuple.Second.(int))
		}
	}

	outputManager.Format(winc)
	err = outputManager.Exhibit()

	if err != nil {
		fmt.Println(err)
	}
}
