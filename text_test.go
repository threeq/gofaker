package gofaker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWord(t *testing.T) {
	so := assert.New(t)

	count := len(strings.Split(Word(), ""))
	so.True(3 <= count && count <= 10, fmt.Sprintf("len %d", count))
}

func TestWordN(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		count := len(strings.Split(WordN(4, 4), ""))
		so.Equal(4, count)
	})

	batch(100, func() {
		count := len(strings.Split(WordN(4, 8), ""))
		so.True(4 <= count && count <= 8, fmt.Sprintf("len %d", count))
	})
}

func TestSentence(t *testing.T) {
	so := assert.New(t)

	s := Sentence()
	count := len(strings.Split(s, " "))
	so.True(12 <= count && count <= 18, fmt.Sprintf("len %d", count))
	so.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", s[0:1])
}

func TestSentenceN(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		s := SentenceN(4, 4)
		count := len(strings.Split(s, " "))
		so.Equal(4, count)
		so.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", s[0:1])
	})

	batch(100, func() {
		s := SentenceN(4, 8)
		count := len(strings.Split(s, " "))
		so.True(4 <= count && count <= 8, fmt.Sprintf("len %d", count))
		so.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", s[0:1])
	})
}

func TestParagraph(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		p := Paragraph()
		count := len(strings.Split(p, "."))
		so.True(4 <= count && count <= 8, fmt.Sprintf("len %d", count))
	})
}

func TestTitle(t *testing.T) {
	so := assert.New(t)
	batch(100, func() {
		titleWords := strings.Split(Title(), " ")
		count := len(titleWords)
		so.True(3 <= count && count <= 7, fmt.Sprintf("len %d", count))
		for _, word := range titleWords {
			so.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", word[0:1])
		}

	})
}

func TestCWord(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cWord := CWord("")
		cWordArray := []rune(cWord)
		so.True(3 <= len(cWordArray) && len(cWordArray) <= 10, fmt.Sprintf("%s len %d", cWord, len(cWordArray)))
		for _, c := range cWordArray {
			so.Contains(hanZiDict, c, fmt.Sprintf("%s", cWord))
		}
	})
}

func TestCWordN(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cWord := CWordN("", 1, 3)
		cWordArray := []rune(cWord)
		so.True(1 <= len(cWordArray) && len(cWordArray) <= 3, fmt.Sprintf("%s len %d", cWord, len(cWordArray)))
		for _, c := range cWordArray {
			so.Contains(hanZiDict, c, fmt.Sprintf("%s", cWord))
		}
	})

	cWord := CWordN("", -1, -1)
	cWordArray := []rune(cWord)
	so.Len(cWordArray, 1)

	cWord = CWordN("你好", 5, 4)
	cWordArray = []rune(cWord)
	so.Len(cWordArray, 4)

	cWord = CWordN("你好", 3, 4)
	cWordArray = []rune(cWord)
	so.True(3 <= len(cWordArray) && len(cWordArray) <= 4, fmt.Sprintf("%s len %d", cWord, len(cWordArray)))

	for _, c := range cWordArray {
		so.Contains([]rune("你好"), c, fmt.Sprintf("%s", cWord))
	}
}

func TestCSentence(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cSentence := CSentence()
		cArray := []rune(cSentence)
		so.True(0 < len(cArray))
	})
}

func TestCSentenceN(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cSentence := CSentenceN(1, 3)
		cArray := []rune(cSentence)
		so.True(0 < len(cArray))

	})

}

func TestCTitle(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cSentence := CTitle()
		cArray := []rune(cSentence)
		so.True(0 < len(cArray))
	})
}

func TestCParagraph(t *testing.T) {
	so := assert.New(t)

	batch(100, func() {
		cSentence := CParagraph()
		cArray := []rune(cSentence)
		so.True(0 < len(cArray))
	})
}