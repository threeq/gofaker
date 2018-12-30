package gofaker

import "strings"

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
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, Word())
	}
	return Capitalize(strings.Join(result, " ") + ".")
}

func Capitalize(word string) string {
	return strings.ToUpper(word[0:1]) + word[1:]
}

// 随机生成一段文本
func Paragraph() string {
	return ParagraphN(3,7)
}

func ParagraphN(min, max int) string {
	size := IntN(min, max)
	var result [] string
	for i := 0; i < size; i++ {
		result = append(result, Sentence())
	}
	return strings.Join(result, "")
}
