package kwic

import "github.com/igorfg/kwic-go/tuple"
import "sort"
import "os"
import "bufio"
import "regexp"
import "strings"

// Tuple : criacao de uma struct do tipo tupla
type Tuple tuple.Tuple

// IndexManager : struct para armazenar o hash de palavras
type IndexManager struct {
	hashTable map[string][]Tuple
	stopWords []string
}

func (im *IndexManager) Init() {
	file, _ := os.Open("resources/stopwords.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		im.stopWords = append(im.stopWords, scanner.Text())
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
	for _, stopWord := range im.stopWords {
		if strings.EqualFold(word, stopWord) {
			return true
		}
	}
	return false
}
