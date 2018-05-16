package main

import "fmt"
import "os"
import "log"
import "strings"
import "github.com/igorfg/kwic-go/kwic"

func main() {
	var (
		storageManager kwic.DataStorageManager
		indexManager   kwic.IndexManager
		outPutManager  kwic.OutputManager
		winc           []string
	)

	args := os.Args[1:]

	if len(args) > 0 && args[0] == "-help" {
		fmt.Println("" +
			"USAGE: kwic-go <input option> <output option>\n\n" +
			"INPUT OPTIONS:\n" +
			"-file Use a file as the input\n" +
			"-dblp Use dblp search criteria\n\n" +
			"OUTPUT OPTIONS:\n" +
			"-terminal Will print output on terminal\n" +
			"-html Will print out on your browser")
	} else if len(args) != 2 {
		fmt.Println("Use the flag -help to get a list of commands")
	} else {
		if args[0] == "-file" {
			storageManager = &kwic.FileBasedStorageManager{}
			indexManager = kwic.IndexManager{}

			storageManager.Init()
			indexManager.Init()
		} else if args[0] == "-dblp" {
			storageManager = &kwic.DBLPStorageManager{}
			indexManager = kwic.IndexManager{}

			storageManager.Init()
			indexManager.Init()
		} else {
			log.Fatal("Invalid input format")
		}

		for lineNumber := 0; lineNumber < storageManager.Length(); lineNumber++ {
			line := storageManager.Line(lineNumber)
			words := strings.Split(line, " ")

			for pos := 0; pos < len(words); pos++ {
				indexManager.Hash(words[pos], line, pos)
			}
		}

		if args[1] == "-terminal" {
			outPutManager = &kwic.TerminalOutputManager{}
		} else if args[1] == "-html" {
			outPutManager = &kwic.HTMLOutputManager{}
		} else {
			fmt.Println("Invalid output format")
		}

		sortedWords := indexManager.SortedWords()
		wordShift := kwic.WordShift{}

		for _, w := range sortedWords {
			for _, tuple := range indexManager.OccurencesOfWord(w) {
				func(line string, pos int) {
					winc = append(winc, (strings.Join(wordShift.Shift(strings.Split(line, " "), pos, 0), " ")))
				}(tuple.First.(string), tuple.Second.(int))
			}
		}

		outPutManager.Format(winc)

		err := outPutManager.Exhibit()

		if err != nil {
			fmt.Println(err)
		}
	}
}
