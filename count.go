package newstring

import (
	"unicode"
	"unicode/utf8"
)

// 返回字符串长度
func Len(str string) int {
	return utf8.RuneCountInString(str)
}

//返回字符串中的单词数。
//可以包含`'` 和`-`，包含`'` 和`-`的认为是一个单词的连接符。(`'`是英文的单引号)
//但是若以其他字符开头，则忽略它的存在。
func WordCount(str string) int {
	var r rune
	var size, n int

	inWord := false

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		switch {
		case IsAlphabet(r):
			if !inWord { //inWord是false的话就进循环
				inWord = true
				n++
			}
		case inWord && (r == '\'' || r == '-'):
			// 如果遍历过字母后，遇到了'\''或者 '-'不执行任何
		default:
			inWord = false
		}

		str = str[size:]
	}

	return n
}

const minCJKCharacter = '\u3400'

//检查r是不是CJK字符的字母，若是返回true，若不是false
func IsAlphabet(r rune) bool {
	if !unicode.IsLetter(r) { //若不是字母，则返回false
		return false
	}

	switch {
	// 若不是CJK字符，则返回true
	case r < minCJKCharacter:
		return true

	//若是普通的的CJK字符，则返回false
	case r >= '\u4E00' && r <= '\u9FCC':
		return false

	// 若是复杂的CJK，则返回false
	case r >= '\u3400' && r <= '\u4D85':
		return false

	//若是古老的CJK字符，则返回false CJK
	case r >= '\U00020000' && r <= '\U0002B81D':
		return false
	}

	return true
}

//返回字符串宽度
//普通字符被认为是1来计算
//复杂字符再次被认为是2来计算
func Width(str string) int {
	var r rune
	var size, n int

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)
		n += RuneWidth(r) //通过每一个字符的编码值来确定每一个字符的宽度
		str = str[size:]
	}

	return n
}

//返回一个字符宽度
//普通字符被认为是1来计算
//复杂字符再次被认为是2来计算
func RuneWidth(r rune) int {
	switch {
	case r == utf8.RuneError || r < '\x20':
		return 0
		//十六进制的20在ASCII码中代表着空格，
	case '\x20' <= r && r < '\u2000':
		return 1

	case '\u2000' <= r && r < '\uFF61':
		return 2

	case '\uFF61' <= r && r < '\uFFA0':
		return 1

	case '\uFFA0' <= r:
		return 2
	}

	return 0
}
