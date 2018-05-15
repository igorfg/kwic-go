package kwic

import "github.com/igorfg/kwic-go/tuple"
import "sort"

// Tuple : criacao de uma struct do tipo tupla
type Tuple tuple.Tuple

// IndexManager : struct para armazenar o hash de palavras
type IndexManager struct {
	hashTable map[string][]Tuple
}

func (im *IndexManager) Init() {
	im.hashTable = make(map[string][]Tuple)
}

func (im *IndexManager) IsEmpty() bool {
	return len(im.hashTable) == 0
}

func (im *IndexManager) Hash(word string, line string, pos int) {
	tupla := Tuple{First: line, Second: pos}

	if _, exists := im.hashTable[word]; exists {
		im.hashTable[word] = append(im.hashTable[word], tupla)
	} else {
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
