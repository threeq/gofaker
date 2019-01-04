package gofaker

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNameFirst(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := NameFirst()
		assertion.Contains(firstNames, f)
	})
}

func TestNameLast(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := NameLast()
		assertion.Contains(lastNames, f)
	})
}

func TestName(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := strings.Split(Name(false), " ")
		assertion.Len(f, 2)
		assertion.Contains(firstNames, f[0])
		assertion.Contains(lastNames, f[1])
	})

	batch(1000, func() {
		f := strings.Split(Name(true), " ")
		assertion.Len(f, 3)
		assertion.Contains(firstNames, f[0])
		assertion.Contains(firstNames, f[1])
		assertion.Contains(lastNames, f[2])
	})
}

func TestNameFirstCN(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := CNameFirst()
		assertion.Contains(cnFirstNames, f)
	})
}

func TestNameLastCN(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := CNameLast()
		assertion.Contains(cnLastNames, f)
	})
}

func TestNameCN(t *testing.T) {
	assertion := assert.New(t)
	batch(1000, func() {
		f := CName()
		assertion.True(utf8.RuneCountInString(f) == 2|| utf8.RuneCountInString(f) == 3, f + strconv.Itoa(len(f)))
	})
}
