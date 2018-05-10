package kwic

import "github.com/igorfg/kwic-go/tuple"
import "sort"

type Tuple tuple.Tuple

type IndexManager struct {
	hashTable map[string][]Tuple
}

func (im *IndexManager) init() {
	im.hashTable = make(map[string][]Tuple)
}

func (im *IndexManager) isEmpty() bool {
	return len(im.hashTable) == 0
}

func (im *IndexManager) hash(word string, line string, pos int) {
	tupla := Tuple{First: line, Second: pos}

	if _, exists := im.hashTable[word]; exists {
		im.hashTable[word] = append(im.hashTable[word], tupla)
	} else {
		im.hashTable[word] = []Tuple{tupla}
	}
}

func (im *IndexManager) occurencesOfWord(word string) []Tuple {
	return im.hashTable[word]
}

func (im *IndexManager) sortedWords() []string {
	keys := make([]string, len(im.hashTable))

	for k := range im.hashTable {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
