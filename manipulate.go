package newstring

import (
	"strings"
	"unicode/utf8"
)

// 反转字符串
func Reverse(str string) string {
	var size int

	tail := len(str)
	buf := make([]byte, tail) //预先分配，防止扩容
	s := buf                  //指向同一数组

	for len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		tail -= size
		s = append(s[:tail], []byte(str[:size])...) //底层仍然不会变
		str = str[size:]                            //切割用的也是原来的底层byte数组
	}

	return string(buf)
}

//分割字符串的字符
// start必须满足 0 <= start <= str的字符数.
// End可以是0也可以是负数也可以是正数
// 如果end >= 0, 比如满足: start <= end <= str的字符数.
// If end < 0, 一直切割到字符串末端
func Slice(str string, start, end int) string {
	var size, startPos, endPos int //startPos：起始字节索引；endPos：结束字节索引

	origin := str

	//这里不写start <= str的字符数的原因是，end > len(str) || (end >= 0 && start > end)已经隐含的写了，省略了
	if start < 0 || end > Len(str) || (end >= 0 && start > end) {
		panic("out of range")
	}

	if end >= 0 {
		end -= start //计算切割字符串的长度
	}

	for start > 0 && len(str) > 0 { //计算起始字节索引
		_, size = utf8.DecodeRuneInString(str)
		start--
		startPos += size
		str = str[size:] //并不会一影响到origin
	}
	//如果end < 0, 一直切割到字符串末端
	if end < 0 {
		return origin[startPos:]
	}

	endPos = startPos //进行赋值

	for end > 0 && len(str) > 0 { //计算结束字节索引
		_, size = utf8.DecodeRuneInString(str)
		end--
		endPos += size
		str = str[size:]
	}

	if len(str) == 0 && (start > 0 || end > 0) {
		panic("out of range")
	}

	return origin[startPos:endPos]
}

//将字符传分割成三部分.
func Partition(str, sep string) (head, match, tail string) {
	index := strings.Index(str, sep)

	if index == -1 {
		head = str
		return
	}

	head = str[:index]
	match = str[index : index+len(sep)]
	tail = str[index+len(sep):]
	return
}

//将字符传分割成三部分.
func LastPartition(str, sep string) (head, match, tail string) {
	index := strings.LastIndex(str, sep)

	if index == -1 {
		tail = str
		return
	}

	head = str[:index]
	match = str[index : index+len(sep)]
	tail = str[index+len(sep):]
	return
}

// 将一个字符串插入到另一个字符串中
//如果index越界,将panic
func Insert(dst, src string, index int) string {
	return Slice(dst, 0, index) + src + Slice(dst, index, -1)
}

//使用repl字符串去掉无效的字符
//相邻的无效字符仅替换一次。
func Scrub(str, repl string) string {
	var buf *stringBuilder
	var r rune
	var size, pos int
	var hasError bool

	origin := str

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		if r == utf8.RuneError {
			if !hasError {
				if buf == nil {
					buf = &stringBuilder{}
				}

				buf.WriteString(origin[:pos])
				hasError = true
			}
		} else if hasError {
			hasError = false
			buf.WriteString(repl)

			origin = origin[pos:]
			pos = 0
		}

		pos += size
		str = str[size:]
	}

	if buf != nil {
		buf.WriteString(origin)
		return buf.String()
	}

	//没有无效的字符
	return origin
}

// 分割str成为的单词，并以切片形式返回
//如果str是空字符串，返回nil
//可以包含`'` 和`-`，包含`'` 和`-`的认为是一个单词的连接符。(`'`是英文的单引号)
//但是若以其他字符开头，则忽略它的存在。
func WordSplit(str string) []string {
	var word string
	var words []string
	var r rune
	var size, pos int

	inWord := false

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		switch {
		case IsAlphabet(r):
			if !inWord {
				inWord = true
				word = str
				pos = 0
			}
		case inWord && (r == '\'' || r == '-'):
			// 不执行任何
		default:
			if inWord {
				inWord = false
				words = append(words, word[:pos])
			}
		}

		pos += size
		str = str[size:]
	}

	if inWord {
		words = append(words, word[:pos])
	}

	return words
}
