package kwic

type WordShift struct{}

func (w *WordShift) Shift(words []string, pos int, target int) []string {
	words = append(words[:pos], append([]string{"|"}, append(words[pos:pos+1], append([]string{"|"}, words[pos+1:]...)...)...)...)

	if pos == target {
		return words
	}
	if pos < target {
		return shiftRight(words, target-pos)
	}
	return shiftLeft(words, pos-target)
}

func shiftRight(words []string, target int) []string {
	for target > 0 {
		words = append(words[len(words)-1:], words[:len(words)-1]...)
		target--
	}
	return words

}

func shiftLeft(words []string, target int) []string {
	for target > 0 {
		words = append(words[1:], words[0])
		target--
	}
	return words

}
