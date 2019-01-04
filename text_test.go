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
