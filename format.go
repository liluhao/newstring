package newstring

import (
	"unicode/utf8"
)

//可以将 str 中的\t小指定为tabSize
// 在 str 中出现换行符\n后，column将重置为零。
//中日韩字符即CJK字符将被视为两个字符。
//如果 tabSize <= 0，则 ExpandTabs 将panic。
func ExpandTabs(str string, tabSize int) string {
	if tabSize <= 0 {
		panic("ttabSize必须大于0")
	}

	var r rune
	var i, size, column, expand int
	var output *stringBuilder

	orig := str

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		if r == '\t' {
			expand = tabSize - column%tabSize

			if output == nil {
				output = allocBuffer(orig, str)
			}

			for i = 0; i < expand; i++ {
				output.WriteRune(' ') //这里面是一个空格
			}

			column += expand
		} else {
			if r == '\n' {
				column = 0
			} else {
				column += RuneWidth(r)
			}

			if output != nil {
				output.WriteRune(r)
			}
		}

		str = str[size:]
	}

	if output == nil {
		return orig
	}

	return output.String()
}

//如果str的字符长度小于length，则 RightJustify返回右侧带有 pad 字符串的字符串。
//如果 str 的字符长度大于等于length，则将返回 str 本身。
//如果 pad 是空字符串，则将返回 str。
func LeftJustify(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &stringBuilder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	output.WriteString(str)
	writePadString(output, pad, padLen, remains)
	return output.String()
}

//如果str的字符长度小于length，则 RightJustify返回左侧带有 pad 字符串的字符串。
//如果 str 的字符长度大于等于length，则将返回 str 本身。
//如果 pad 是空字符串，则将返回 str。
func RightJustify(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l
	padLen := Len(pad)

	output := &stringBuilder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains)
	output.WriteString(str)
	return output.String()
}

//如果 str 的字符长度小于length，则 Center 将返回两侧都加有pad的字符串。
//如果 str 的字符长度大于等于length，则将返回 str 本身。
//如果 pad 是空字符串，则将返回 str。
func Center(str string, length int, pad string) string {
	l := Len(str)

	if l >= length || pad == "" {
		return str
	}

	remains := length - l //还需要填充的长度
	padLen := Len(pad)

	output := &stringBuilder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad)) //提前分配,防止扩容重新分配底层;比如对于Center("hello", 15, "adc")"adcadhelloadcad"来说容量是5+(10/3 +1)*3=17
	writePadString(output, pad, padLen, remains/2)      //先填充左侧,注意:如果remain是偶数的话,左右两侧分配的字符数相同;如果remain是奇数的话,右侧需比左侧多分配一个
	output.WriteString(str)                             //再填充中间
	writePadString(output, pad, padLen, (remains+1)/2)  //再填充右侧
	return output.String()
}

func writePadString(output *stringBuilder, pad string, padLen, remains int) {
	var r rune
	var size int

	repeats := remains / padLen //repeat是指pad可以重复的次数,比如pad=adc时,左侧填充adcad则repeats=1

	for i := 0; i < repeats; i++ {
		output.WriteString(pad)
	}

	remains = remains % padLen //比如pad=adc时,左侧填充adcad则remian=2

	if remains != 0 {
		for i := 0; i < remains; i++ {
			r, size = utf8.DecodeRuneInString(pad)
			output.WriteRune(r)
			pad = pad[size:]
		}
	}
}
