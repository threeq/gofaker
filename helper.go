package gofaker

import "strings"

func cword1(pool []rune) rune {
	at := IntN(0, len(pool)-1)
	return pool[at:at+1][0]
}

func cwordN(min, max int) string {
	min, max = checkRange(min, max)
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, CWord(""))
	}
	return strings.Join(result, "")
}

func checkRange(min, max int) (int, int) {
	if min < 1 {
		min = 1
	}

	if max < 1 {
		max = 1
	}

	if min > max {
		min = max
	}
	return min, max
}

func batch(size int, fn func()) {
	for i := 0; i < size; i++ {
		fn()
	}
}

func capitalize(word string) string {
	return strings.ToUpper(word[0:1]) + word[1:]
}