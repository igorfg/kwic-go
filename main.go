package main

import "fmt"
import "strings"
import "github.com/igorfg/kwic-go/kwic"

func main() {
	// var fsm kwic.DataStorageManager = &kwic.FileBasedStorageManager{}
	// var im kwic.IndexManager = kwic.IndexManager{}

	// fsm.Init()
	// im.Init()

	// for lineNumber := 0; lineNumber < fsm.Length(); lineNumber++ {
	// 	line := fsm.Line(lineNumber)
	// 	words := strings.Split(line, " ")
	//
	// 	for pos := 0; pos < len(words); pos++ {
	// 		im.Hash(words[pos], line, pos)
	// 	}
	// }
	//
	// sortedWords := im.SortedWords()
	// wordShift := kwic.WordShift{}
	//
	// for _, w := range sortedWords {
	// 	for _, tuple := range im.OccurencesOfWord(w) {
	// 		func(line string, pos int) {
	// 			// winc = append(winc, (strings.Join(wordShift.Shift(strings.Split(line, " "), pos, 0), " ")))
	// 			fmt.Println(strings.Join(wordShift.Shift(strings.Split(line, " "), pos, 0), " "))
	// 		}(tuple.First.(string), tuple.Second.(int))
	// 	}
	// }

	var dblpsm kwic.DataStorageManager = &kwic.DBLPStorageManager{}
	var im kwic.IndexManager = kwic.IndexManager{}
	var h kwic.OutputManager = &kwic.TerminalOutputManager{}

	var winc []string

	dblpsm.Init()
	im.Init()

	for lineNumber := 0; lineNumber < dblpsm.Length(); lineNumber++ {
		line := dblpsm.Line(lineNumber)
		words := strings.Split(line, " ")

		for pos := 0; pos < len(words); pos++ {
			im.Hash(words[pos], line, pos)
		}
	}

	sortedWords := im.SortedWords()
	wordShift := kwic.WordShift{}

	for _, w := range sortedWords {
		for _, tuple := range im.OccurencesOfWord(w) {
			func(line string, pos int) {
				// fmt.Println(strings.Join(wordShift.Shift(strings.Split(line, " "), pos, pos), " "))
				winc = append(winc, (strings.Join(wordShift.Shift(strings.Split(line, " "), pos, pos), " ")))
			}(tuple.First.(string), tuple.Second.(int))
		}
	}

	h.Format(winc)

	err := h.Exhibit("teste")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fim da Execução")
}
