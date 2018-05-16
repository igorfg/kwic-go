package kwic

import (
	"fmt"
)

type TerminalOutputManager struct {
	lines []string
}

func (h *TerminalOutputManager) Format(winc []string) {
	maxSpacement := BiggerSpace(winc)

	for _, str := range winc {
		size := sizeUntilPipe(str)
		spaceCtn := maxSpacement - size
		h.lines = append(h.lines, StringWithSpace(spaceCtn)+str)
	}
}

func (h *TerminalOutputManager) Exhibit() error {
	for _, str := range h.lines {
		_, err := fmt.Println(str)

		if err != nil {
			panic(err)
		}
	}

	return nil
}

func BiggerSpace(winc []string) int {
	maxLength := 0

	for _, str := range winc {
		size := sizeUntilPipe(str)
		if maxLength < size {
			maxLength = size
		}
	}
	return maxLength
}

func StringWithSpace(s int) string {
	var str string

	for i := 0; i < s; i++ {
		str += " "
	}

	return str
}

func sizeUntilPipe(str string) int {
	var ctnPipe int = 0

	for _, c := range str {
		if c == '|' {
			return ctnPipe
		}
		ctnPipe++
	}
	return ctnPipe
}
