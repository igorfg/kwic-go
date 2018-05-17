package kwic

import (
	"bufio"
	"os"
	"regexp"
	"sort"

	"github.com/igorfg/kwic-go/tuple"
)

// Tuple : criacao de uma struct do tipo tupla
type Tuple tuple.Tuple

// IndexManager : struct para armazenar o hash de palavras
type IndexManager struct {
	hashTable map[string][]Tuple
	stopWords map[string]bool
}

func (im *IndexManager) Init() {
	file, _ := os.Open("resources/stopwords.txt")
	defer file.Close()

	im.stopWords = make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		im.stopWords[scanner.Text()] = true
	}

	im.hashTable = make(map[string][]Tuple)

}

func (im *IndexManager) IsEmpty() bool {
	return len(im.hashTable) == 0
}

func (im *IndexManager) Hash(word string, line string, pos int) {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	tupla := Tuple{First: line, Second: pos}

	if _, exists := im.hashTable[word]; exists {
		im.hashTable[word] = append(im.hashTable[word], tupla)
	} else if isAlpha(word) && !im.isStopWord(word) {
		im.hashTable[word] = []Tuple{tupla}
	}
}

func (im *IndexManager) OccurencesOfWord(word string) []Tuple {
	return im.hashTable[word]
}

func (im *IndexManager) SortedWords() []string {
	keys := make([]string, 0)

	for k := range im.hashTable {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func (im *IndexManager) isStopWord(word string) bool {
	_, exist := im.stopWords[word]

	if exist {
		return true
	}
	return false
}
