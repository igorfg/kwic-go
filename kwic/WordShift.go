package kwic

import "strings"

type WordShift struct{}

func (w *WordShift) Shift(words []string, pos int, target int) []string {
	l, r := words[0:pos], words[pos:len(words)]
	if len(strings.Join(l, " ")) < target-5 {
		return shiftRight(l, r, target)
	}
	return shiftLeft(l, r, target)
}

func shiftRight(l []string, r []string, target int) []string {
	if len(r) == 0 {
		return l
	}

	r1, r2 := r[0:len(r)-1], r[len(r)-1:len(r)]
	l1 := append(r2, l...)

	if len(strings.Join(l1, " ")) > target-5 {
		return append(l, r...)
	}
	return shiftRight(l1, r1, target)
}

func shiftLeft(l []string, r []string, target int) []string {
	if len(l) == 0 {
		return r
	}

	l1, l2 := l[0:1], l[1:len(l)]
	r1 := append(r, l1...)

	if len(strings.Join(l2, " ")) < target-5 {
		return append(l2, r1...)
	}
	return shiftLeft(l2, r1, target)
}
