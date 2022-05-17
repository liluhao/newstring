package newstring

import (
	"math/rand"
	"unicode"
	"unicode/utf8"
)

//是将用空格，下划线和连字符分隔的单词转换为首字母大写
//如果是空字符串，会返回空字符串
func ToCamelCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &stringBuilder{}
	var r0, r1 rune
	var size int

	for len(str) > 0 {
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if !isConnector(r0) { //如果不是'-' 、 '_' 、空白字符，则进入判断语句
			r0 = unicode.ToUpper(r0)
			break //会直接跳出
		}

		buf.WriteRune(r0)
	}

	if len(str) == 0 {
		//对于str里只有一个字符的特殊处理
		if size != 0 {
			buf.WriteRune(r0)
		}

		return buf.String()
	}
	//写两次for循环的原因是进行r1与r2的判断
	for len(str) > 0 {
		r1 = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if isConnector(r0) && isConnector(r1) { //如果本字符以及上一个字符是'-' 、 '_' 、空白字符，则进入判断语句
			buf.WriteRune(r1)
			continue
		}

		if isConnector(r1) { //如果上一个字符是'-' 、 '_' 、空白字符，但是本字符不是，则进入判断语句
			r0 = unicode.ToUpper(r0) //大写字母，这一个判断省略了写入操作
		} else { //只要满足上一个字符不是'-' 、 '_' 、空白字符，就会进入判断语句
			r0 = unicode.ToLower(r0) //
			buf.WriteRune(r1)        //每次写入的都是上一次遍历的字符
		}
	}

	buf.WriteRune(r0) //记得给写进
	return buf.String()
}

// 可以将字符串中的所有大写字符转换小写字符，并根据情况用下划线连接
func ToSnakeCase(str string) string {
	return camelCaseToLowerCase(str, '_')
}

//ToKebabCase:可以将字符串中的所有大写字符转换为小写字符，并根据情况用连接符连接
func ToKebabCase(str string) string {
	return camelCaseToLowerCase(str, '-')
}

func camelCaseToLowerCase(str string, connector rune) string {
	if len(str) == 0 {
		return ""
	}

	buf := &stringBuilder{}
	wt, word, remaining := NextWord(str) //比如str=HTTPServer，则wt=upperCaseWord、word=HTTP、remaining=Server；str=FirstName，则wt=upperCaseWord、word=First、remaining=Name

	for len(remaining) > 0 {
		if wt != connectorWord {
			toLower(buf, wt, word, connector) //
		}

		prev := wt
		last := word
		wt, word, remaining = NextWord(remaining) //比如remaining=Name，则wt=upperCaseWord、word=Name、remaining=""

		switch prev {
		case numberWord:
			for wt == alphabetWord || wt == numberWord {
				toLower(buf, wt, word, connector)
				wt, word, remaining = NextWord(remaining)
			}

			if wt != invalidWord && wt != punctWord {
				buf.WriteRune(connector)
			}

		case connectorWord:
			toLower(buf, prev, last, connector)

		case punctWord:
			//不处理任何

		default:
			if wt != numberWord {
				if wt != connectorWord && wt != punctWord {
					buf.WriteRune(connector) //把连接符写入
				}

				break
			}

			if len(remaining) == 0 {
				break
			}

			last := word
			wt, word, remaining = NextWord(remaining)

			// 遇到数字则进行连接
			// 例如如果connector=是“_"的话，"Bld4Floor" => "bld4_floor"
			if wt != alphabetWord {
				toLower(buf, numberWord, last, connector)

				if wt != connectorWord && wt != punctWord {
					buf.WriteRune(connector)
				}

				break
			}

			// 如果数字后面有一些小写字母，在数字前面添加connector
			// 比如"HTTP2xx" => "http_2xx"
			buf.WriteRune(connector)
			toLower(buf, numberWord, last, connector)

			for wt == alphabetWord || wt == numberWord {
				toLower(buf, wt, word, connector)
				wt, word, remaining = NextWord(remaining)
			}

			if wt != invalidWord && wt != connectorWord && wt != punctWord {
				buf.WriteRune(connector)
			}
		}
	}

	toLower(buf, wt, word, connector) //把word写入buf
	return buf.String()
}

func isConnector(r rune) bool {
	return r == '-' || r == '_' || unicode.IsSpace(r) //IsSpace(r rune) bool：判断一个字符是否是空白字符，以下6个都是空白字符,'\t', '\n', '\v', '\f', '\r', ' '
}

type wordType int

const (
	invalidWord   wordType = iota //标记是无效的rune
	numberWord                    //标记rune是数学字符
	upperCaseWord                 //标记rune是大写字母
	alphabetWord                  //标记rune是CJK字符的字母
	connectorWord                 //标价rune是判断以下8个 : '-'  、 '_' 、'\t', '\n', '\v', '\f', '\r', ' '
	punctWord                     //标记rune是标点符号
	otherWord                     //标记rune是其他字符
)

func NextWord(str string) (wt wordType, word, remaining string) {
	if len(str) == 0 {
		return
	}

	var offset int                                      //计算str里str里第一个大写字符到第二个大写字符间的偏移量，比如如果传入的str=FirstName，则offset=5；
	remaining = str                                     //赋值给remaining
	r, size := nextValidRune(remaining, utf8.RuneError) //该函数并不会影响到remaining的值
	offset += size

	if r == utf8.RuneError {
		wt = invalidWord
		word = str[:offset]
		remaining = str[offset:]
		return
	}

	switch {
	case isConnector(r): //判断是否是以下8个 : '-'  、 '_' 、'\t', '\n', '\v', '\f', '\r', ' '
		wt = connectorWord
		remaining = remaining[size:]

		for len(remaining) > 0 {
			r, size = nextValidRune(remaining, r)

			if !isConnector(r) {
				break
			}

			offset += size
			remaining = remaining[size:]
		}

	case unicode.IsPunct(r): //判断是否是标点符号
		wt = punctWord
		remaining = remaining[size:]

		for len(remaining) > 0 {
			r, size = nextValidRune(remaining, r)

			if !unicode.IsPunct(r) {
				break
			}

			offset += size
			remaining = remaining[size:]
		}

	case unicode.IsUpper(r): //判断是否大写字母
		wt = upperCaseWord
		remaining = remaining[size:] //若是大写字母，则进行剪枝

		if len(remaining) == 0 {
			break
		}

		r, size = nextValidRune(remaining, r)

		switch {
		case unicode.IsUpper(r): //判断是否大写字母
			prevSize := size
			offset += size
			remaining = remaining[size:]

			for len(remaining) > 0 {
				r, size = nextValidRune(remaining, r) //此处并没有用DecodeRuneInString函数

				if !unicode.IsUpper(r) { //如果是除了小写的其他字符，则退出for循环
					break
				}

				prevSize = size
				offset += size
				remaining = remaining[size:]
			}
			//当函数传入的str像“HTTPStatus”这样的时候，会进行如下判断
			//以为经过上面代码的处理后，此时remaininig=“tatus“,但实际上S应该被保留的，所以偏移量应该减去S的字节数
			//经过以下处理后，remainnning会成为Status
			if len(remaining) > 0 && IsAlphabet(r) {
				offset -= prevSize
				remaining = str[offset:]
			}

		case IsAlphabet(r): //检查r是不是CJK字符的字母
			offset += size
			remaining = remaining[size:]

			for len(remaining) > 0 {
				r, size = nextValidRune(remaining, r)

				if !IsAlphabet(r) || unicode.IsUpper(r) {
					break
				}

				offset += size
				remaining = remaining[size:]
			}
		}

	case IsAlphabet(r): //检查r是不是CJK字符的字母
		wt = alphabetWord
		remaining = remaining[size:]

		for len(remaining) > 0 {
			r, size = nextValidRune(remaining, r)

			if !IsAlphabet(r) || unicode.IsUpper(r) {
				break
			}

			offset += size
			remaining = remaining[size:]
		}

	case unicode.IsNumber(r): //判断一个字符是否是数学字符
		wt = numberWord
		remaining = remaining[size:]

		for len(remaining) > 0 {
			r, size = nextValidRune(remaining, r)

			if !unicode.IsNumber(r) {
				break
			}

			offset += size
			remaining = remaining[size:]
		}

	default: //其他字符
		wt = otherWord
		remaining = remaining[size:]

		for len(remaining) > 0 {
			r, size = nextValidRune(remaining, r)

			if size == 0 || isConnector(r) || IsAlphabet(r) || unicode.IsNumber(r) || unicode.IsPunct(r) {
				break
			}

			offset += size
			remaining = remaining[size:]
		}
	}

	word = str[:offset]
	return
}

//nextValidRune函数是DecodeRuneInString的加强版，即防止str里出现无效的uft8，返回第一个正常的rune。
//如果str里全部不正常，则返回prev
//如果str是空字符串，则返回prev，返回size=0
//m, n := nextValidRune("abcdf", 'a')
//fmt.Println(string(m)) //a
//fmt.Println(n)         //1
//m, n = nextValidRune("abcdf", 'd')
//fmt.Println(string(m)) //a
//fmt.Println(n)         //1
//m, n = nextValidRune("abcdf", '9')
//fmt.Println(string(m)) //a
//fmt.Println(n)         //1
//m, n = nextValidRune("", 'a')
//fmt.Println(string(m)) //a
//fmt.Println(n)         //0
//m, n = nextValidRune("", '9')
//fmt.Println(string(m)) //9
//fmt.Println(n)         //0
func nextValidRune(str string, prev rune) (r rune, size int) {
	var sz int

	for len(str) > 0 {
		r, sz = utf8.DecodeRuneInString(str)
		size += sz

		if r != utf8.RuneError {
			return
		}

		str = str[sz:]
	}

	r = prev
	return
}

//如果传入的buf是空的,wt=upperCaseWord,str=First，connctor=-，则返回后but=First
func toLower(buf *stringBuilder, wt wordType, str string, connector rune) {
	buf.Grow(buf.Len() + len(str))

	if wt != upperCaseWord && wt != connectorWord {
		buf.WriteString(str)
		return
	}

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if isConnector(r) {
			buf.WriteRune(connector)
		} else if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}
}

// 小写变大写，大写变小写
func SwapCase(str string) string {
	var r rune
	var size int

	buf := &stringBuilder{}

	for len(str) > 0 { //对每一个字符都进行操作
		r, size = utf8.DecodeRuneInString(str)

		switch {
		case unicode.IsUpper(r):
			buf.WriteRune(unicode.ToLower(r))

		case unicode.IsLower(r):
			buf.WriteRune(unicode.ToUpper(r))

		default:
			buf.WriteRune(r)
		}

		str = str[size:]
	}

	return buf.String()
}

// FirstRuneToUpper 会将第一个字符转换为大写。
func FirstRuneToUpper(str string) string {
	if str == "" {
		return str
	}

	r, size := utf8.DecodeRuneInString(str)

	//判断runne字符是否是小写字母
	if !unicode.IsLower(r) {
		return str
	}

	buf := &stringBuilder{}
	buf.WriteRune(unicode.ToUpper(r))
	buf.WriteString(str[size:])
	return buf.String()
}

// FirstRuneToLower会将第一个字符转换为小写。
func FirstRuneToLower(str string) string {
	if str == "" {
		return str
	}

	r, size := utf8.DecodeRuneInString(str)
	//判断runne字符是否是大写字母
	if !unicode.IsUpper(r) {
		return str
	}

	buf := &stringBuilder{}
	buf.WriteRune(unicode.ToLower(r))
	buf.WriteString(str[size:])
	return buf.String()
}

// 随机重组字符串中的字符并返回结果
func Shuffle(str string) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	index := 0

	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1) //Intn(n int) int：返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。

		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes) //直接进行转换即可
}

//使用被给的随机源rand.Source进行随机化字符
func ShuffleSource(str string, src rand.Source) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	index := 0
	r := rand.New(src)

	for i := len(runes) - 1; i > 0; i-- {
		index = r.Intn(i + 1)

		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes)
}

//如下的符合Alphanumeric:
//a - z
//A - Z
//0 - 9
//将str里符合最右边的Alphanumeric加1。
//将str里符合最右边的Alphanumeric的是z、Z、9的话，则还会继续再寻找一个Alphanumeric加1，如果这个Alphanumeric仍是z、Z、9的话，则还会再次继续寻找一个Alphanumeric加1。
//将str里全是Alphanumeric，则z、Z、9，则还会在str最左侧进位。
//如果str是空字符，则结果仍返回字符串
//如果str里没有Alphanumeric字符，则无论结果是否为有效的字符，将str的最右边的字符都增加1。
func Successor(str string) string {
	if str == "" {
		return str
	}

	var r rune
	var i int
	carry := ' ' //用于进位
	runes := []rune(str)
	l := len(runes)       //这个l也是str的字符数
	lastAlphanumeric := l //用于记录str中最后一个Alphanumeric的字串的索引

	for i = l - 1; i >= 0; i-- { //倒着对字符开始进行遍历
		r = runes[i]

		if ('a' <= r && r <= 'y') ||
			('A' <= r && r <= 'Y') ||
			('0' <= r && r <= '8') {
			runes[i]++
			carry = ' '
			lastAlphanumeric = i
			break
		}

		switch r { //如果是z、Z、9的话需要特殊讨论，此时还不能break退出，因为还要
		case 'z':
			runes[i] = 'a'
			carry = 'a'
			lastAlphanumeric = i

		case 'Z':
			runes[i] = 'A'
			carry = 'A'
			lastAlphanumeric = i

		case '9':
			runes[i] = '0'
			carry = '0'
			lastAlphanumeric = i
		}
	}

	if i < 0 && carry != ' ' { // 将str里全是Alphanumeric，则z、Z、9，则还会在str最左侧进位； 若进入if语句，需要i=-1，lastAlpanumeric不确定，可能是0也可能是负数，不可能是负数
		buf := &stringBuilder{}
		buf.Grow(len(str) + 1) // 扩容足够的空间用于进位；
		if lastAlphanumeric != 0 {
			for _, r = range runes[:lastAlphanumeric] {
				buf.WriteRune(r) //通过遍历的方式写入
			}
		}

		buf.WriteRune(carry + 1)

		for _, r = range runes[lastAlphanumeric:] {
			buf.WriteRune(r) //通过遍历的方式写入
		}

		return buf.String()
	}

	// 如果str里没有Alphanumeric字符，则无论结果是否为有效的字符，将str的最右边的字符都增加1。
	if lastAlphanumeric == l {
		runes[l-1]++
	}

	return string(runes)
}
