package gofaker

import (
	"strings"
)

func Word() string {
	return WordN(3, 10)
}

func WordN(min, max int) string {
	size := IntN(min, max)
	var result = ""
	for i := 0; i < size; i++ {
		result += Char("lower")
	}
	return result
}

// 随机生成一个句子，第一个单词的首字母大写
func Sentence() string {
	return SentenceN(12, 18)
}

func SentenceN(min, max int) string {
	min, max = checkRange(min, max)
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, Word())
	}
	return capitalize(strings.Join(result, " ") + ".")
}

// 随机生成一段文本
func Paragraph() string {
	return ParagraphN(3, 7)
}

func ParagraphN(min, max int) string {
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, Sentence())
	}
	return strings.Join(result, "")
}

// 随机生成一句标题，其中每个单词的首字母大写
func Title() string {
	return TitleN(3, 7)
}

func TitleN(min, max int) string {
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, capitalize(Word()))
	}
	return strings.Join(result, " ")
}

// 中文随机字符串
func CWord(pool string) string {
	return CWordN(pool, 3, 10)
}

func CWordN(pool string, min, max int) string {
	min, max = checkRange(min, max)

	var poolRune []rune
	if pool == "" {
		poolRune = hanZiDict
	} else {
		poolRune = []rune(pool)
	}

	size := IntN(min, max)
	var result []rune
	for i := 0; i < size; i++ {
		result = append(result, cword1(poolRune))
	}
	return string(result)
}

func CSentence() string {
	return CSentenceN(12, 18)
}

func CSentenceN(min, max int) string {
	return cwordN(min, max) + "。"
}

func CTitle() string {
	return CTitleN(3, 7)
}

func CTitleN(min, max int) string {
	return cwordN(min, max)
}

func CParagraph() string {
	return CParagraphN(3, 7)
}

func CParagraphN(min, max int) string {
	min, max = checkRange(min, max)
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, CSentence())
	}
	return strings.Join(result, "")
}
