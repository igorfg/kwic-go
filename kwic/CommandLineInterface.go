package kwic

import "os"
import "fmt"
import "log"
import "strings"

type CommandLineInterface struct {
	args []string
}

func (cli *CommandLineInterface) Init(storageManager DataStorageManager,
	indexManager IndexManager, outputManager OutputManager) {

	cli.args = os.Args[1:]

	if len(cli.args) > 0 && cli.args[0] == "-help" {
		printHelpOptions()
	} else if len(cli.args) != 2 {
		wrongArguments()
	} else {
		cli.checkInputMethod(&storageManager)
		cli.checkOutputMethod(&outputManager)
		indexManager = IndexManager{}

		err := storageManager.Init()
		if err != nil {
			log.Fatal(err.Error())
		}

		indexManager.Init()

		sortedWords := indexWords(&storageManager, &indexManager)
		winc := shiftOutputContent(sortedWords, &indexManager)

		showOutput(&outputManager, &winc)
	}
}

func (cli *CommandLineInterface) checkOutputMethod(outputManager *OutputManager) {
	if cli.args[1] == "-terminal" {
		*outputManager = &TerminalOutputManager{}
	} else if cli.args[1] == "-html" {
		*outputManager = &HTMLOutputManager{}
	} else {
		fmt.Println("Invalid output format")
	}
}

func (cli *CommandLineInterface) checkInputMethod(storageManager *DataStorageManager) {
	if cli.args[0] == "-file" {
		*storageManager = &FileBasedStorageManager{}
	} else if cli.args[0] == "-dblp" {
		*storageManager = &DBLPStorageManager{}
	} else {
		log.Fatal("Invalid input format")
	}
}

func showOutput(outputManager *OutputManager, winc *[]string) {
	(*outputManager).Format(*winc)
	err := (*outputManager).Exhibit()

	if err != nil {
		fmt.Println(err)
	}
}

func shiftOutputContent(sortedWords []string, indexManager *IndexManager) []string {
	var winc []string
	wordShift := WordShift{}

	for _, w := range sortedWords {
		for _, tuple := range indexManager.OccurencesOfWord(w) {
			func(line string, pos int) {
				winc = append(winc, (strings.Join(wordShift.Shift(strings.Split(line, " "), pos, 0), " ")))
			}(tuple.First.(string), tuple.Second.(int))
		}
	}
	return winc
}

func indexWords(storageManager *DataStorageManager, indexManager *IndexManager) []string {
	for lineNumber := 0; lineNumber < (*storageManager).Length(); lineNumber++ {
		line := (*storageManager).Line(lineNumber)
		words := strings.Split(line, " ")

		for pos := 0; pos < len(words); pos++ {
			indexManager.Hash(words[pos], line, pos)
		}
	}
	return indexManager.SortedWords()
}

func printHelpOptions() {
	fmt.Println("" +
		"USAGE: kwic-go <input option> <output option>\n\n" +
		"INPUT OPTIONS:\n" +
		"-file Use a file as the input\n" +
		"-dblp Use dblp search criteria\n\n" +
		"OUTPUT OPTIONS:\n" +
		"-terminal Will print output on the terminal\n" +
		"-html Will print out on your browser")
}

func wrongArguments() {
	fmt.Println("Use the flag -help to get a list of commands")
}
