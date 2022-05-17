package newstring

import (
	"unicode"
	"unicode/utf8"
)

type runeRangeMap struct {
	FromLo rune // range map的低边界
	FromHi rune // range map的高边界
	ToLo   rune
	ToHi   rune
}

type runeDict struct {
	Dict [unicode.MaxASCII + 1]rune
}

type runeMap map[rune]rune

//转换器可以使用预编译的form模式和to模式转换字符串。
//如果一对from/to模式串需要被使用多于一次，推荐使用转换器的方式，可以重复使用
type Translator struct {
	quickDict  *runeDict       // 快速字典(用于按索引查找字符)；仅适用于拉丁语字符latin runes.
	runeMap    runeMap         // runeMap
	ranges     []*runeRangeMap // 字符的范围
	mappedRune rune            // If mappedRune >= 0, all matched runes are translated to the mappedRune.
	reverted   bool            // 如果to pattern是空的, 所有的匹配字符将被删除
	hasPattern bool
}

// 通过 from/to 模式对创建新的转换器。
func NewTranslator(from, to string) *Translator {
	tr := &Translator{}

	if from == "" {
		return tr
	}

	reverted := from[0] == '^' //判断from的第一个字符，是'^'的话位为false，不是'^'的话返回true
	deletion := len(to) == 0   //如果to是空字符的话，是true

	if reverted {
		from = from[1:]
	}

	var fromStart, fromEnd, fromRangeStep rune
	var toStart, toEnd, toRangeStep rune
	var fromRangeSize, toRangeSize rune
	var singleRunes []rune

	//更新这个to rune range.
	updateRange := func() {
		// 没有更多别的字符需要读进to rune pattern
		if toEnd == utf8.RuneError {
			return
		}

		if toRangeStep == 0 {
			to, toStart, toEnd, toRangeStep = nextRuneRange(to, toEnd)
			return
		}

		// 当前的range不是空的，则从start消耗一个字符
		if toStart != toEnd {
			toStart += toRangeStep
			return
		}

		// 没有更多的字符，则重复最后一个字符
		if to == "" {
			toEnd = utf8.RuneError
			return
		}

		//同时使用开始和结束。从 to 模式中再读两个符文。
		to, toStart, toEnd, toRangeStep = nextRuneRange(to, utf8.RuneError)
	}

	if deletion {
		toStart = utf8.RuneError
		toEnd = utf8.RuneError
	} else {
		//如果from模式恢复，则只会使用to模式中的最后一个字符。
		if reverted {
			var size int

			for len(to) > 0 {
				toStart, size = utf8.DecodeRuneInString(to)
				to = to[size:]
			}

			toEnd = utf8.RuneError
		} else {
			to, toStart, toEnd, toRangeStep = nextRuneRange(to, utf8.RuneError)
		}
	}

	fromEnd = utf8.RuneError

	for len(from) > 0 {
		from, fromStart, fromEnd, fromRangeStep = nextRuneRange(from, fromEnd)

		// fromStart 是单个字符。只需在to模式中使用字符映射它即可
		if fromRangeStep == 0 {
			singleRunes = tr.addRune(fromStart, toStart, singleRunes)
			updateRange()
			continue
		}

		for toEnd != utf8.RuneError && fromStart != fromEnd {
			// 如果映射的字符是单个字符而不是范围，只需移动范围中的第一个字符即可。
			if toRangeStep == 0 {
				singleRunes = tr.addRune(fromStart, toStart, singleRunes)
				updateRange()
				fromStart += fromRangeStep
				continue
			}

			fromRangeSize = (fromEnd - fromStart) * fromRangeStep
			toRangeSize = (toEnd - toStart) * toRangeStep

			// 在to pattern里没有足够多的字符, 需要读更多.
			if fromRangeSize > toRangeSize {
				fromStart, toStart = tr.addRuneRange(fromStart, fromStart+toRangeSize*fromRangeStep, toStart, toEnd, singleRunes)
				fromStart += fromRangeStep
				updateRange()

				// 边缘情况: 如果fromRangeSize == toRangeSize + 1, 最后一个 fromStart value 需要做为一个单个字符被考虑
				if fromStart == fromEnd {
					singleRunes = tr.addRune(fromStart, toStart, singleRunes)
					updateRange()
				}

				continue
			}

			fromStart, toStart = tr.addRuneRange(fromStart, fromEnd, toStart, toStart+fromRangeSize*toRangeStep, singleRunes)
			updateRange()
			break
		}

		if fromStart == fromEnd {
			fromEnd = utf8.RuneError
			continue
		}

		_, toStart = tr.addRuneRange(fromStart, fromEnd, toStart, toStart, singleRunes)
		fromEnd = utf8.RuneError
	}

	if fromEnd != utf8.RuneError {
		tr.addRune(fromEnd, toStart, singleRunes)
	}

	tr.reverted = reverted
	tr.mappedRune = -1
	tr.hasPattern = true

	// 仅在删除或还原模式下Translate RuneError
	if deletion || reverted {
		tr.mappedRune = toStart
	}

	return tr
}

func (tr *Translator) addRune(from, to rune, singleRunes []rune) []rune {
	if from <= unicode.MaxASCII {
		if tr.quickDict == nil {
			tr.quickDict = &runeDict{}
		}

		tr.quickDict.Dict[from] = to
	} else {
		if tr.runeMap == nil {
			tr.runeMap = make(runeMap)
		}

		tr.runeMap[from] = to
	}

	singleRunes = append(singleRunes, from)
	return singleRunes
}

func (tr *Translator) addRuneRange(fromLo, fromHi, toLo, toHi rune, singleRunes []rune) (rune, rune) {
	var r rune
	var rrm *runeRangeMap

	if fromLo < fromHi {
		rrm = &runeRangeMap{
			FromLo: fromLo,
			FromHi: fromHi,
			ToLo:   toLo,
			ToHi:   toHi,
		}
	} else {
		rrm = &runeRangeMap{
			FromLo: fromHi,
			FromHi: fromLo,
			ToLo:   toHi,
			ToHi:   toLo,
		}
	}

	// 如果有任何单个字符与字符范围冲突，请清除单个符文记录。
	for _, r = range singleRunes {
		if rrm.FromLo <= r && r <= rrm.FromHi {
			if r <= unicode.MaxASCII {
				tr.quickDict.Dict[r] = 0
			} else {
				delete(tr.runeMap, r)
			}
		}
	}

	tr.ranges = append(tr.ranges, rrm)
	return fromHi, toHi
}

func nextRuneRange(str string, last rune) (remaining string, start, end rune, rangeStep rune) {
	var r rune
	var size int

	remaining = str
	escaping := false
	isRange := false

	for len(remaining) > 0 {
		r, size = utf8.DecodeRuneInString(remaining)
		remaining = remaining[size:]

		//解析特殊的字符
		if !escaping {
			if r == '\\' {
				escaping = true
				continue
			}

			if r == '-' {
				// Ignore slash at beginning of string.
				if last == utf8.RuneError {
					continue
				}

				start = last
				isRange = true
				continue
			}
		}

		escaping = false

		if last != utf8.RuneError {
			// 这是一个start与end都相同的的range
			// 认为它是一个正常的字符
			if isRange && last == r {
				isRange = false
				continue
			}

			start = last
			end = r

			if isRange {
				if start < end {
					rangeStep = 1
				} else {
					rangeStep = -1
				}
			}

			return
		}

		last = r
	}

	start = last
	end = utf8.RuneError
	return
}

func (tr *Translator) Translate(str string) string {
	if !tr.hasPattern || str == "" {
		return str
	}

	var r rune
	var size int
	var needTr bool

	orig := str

	var output *stringBuilder

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)
		r, needTr = tr.TranslateRune(r) //传入的是r，不是str;匹配到的r都被赋值为65533与true
		//第一次需要分配空间
		if needTr && output == nil { //直达遇到needTr是true的时候；虽然strings.Builder不能与nil比较，但是它的指针可以与nil比较
			output = allocBuffer(orig, str) //origin一直都不变;
		}
		//utf8.RuneError=65533
		if r != utf8.RuneError && output != nil {
			output.WriteRune(r)
		}

		str = str[size:]
	}

	//没有字符被转换
	if output == nil {
		return orig
	}

	return output.String()
}

func (tr *Translator) TranslateRune(r rune) (result rune, translated bool) {
	switch {
	case tr.quickDict != nil:
		if r <= unicode.MaxASCII {
			result = tr.quickDict.Dict[r]

			if result != 0 {
				translated = true

				if tr.mappedRune >= 0 {
					result = tr.mappedRune
				}

				break
			}
		}

		fallthrough

	case tr.runeMap != nil:
		var ok bool

		if result, ok = tr.runeMap[r]; ok {
			translated = true

			if tr.mappedRune >= 0 {
				result = tr.mappedRune
			}

			break
		}

		fallthrough

	default:
		var rrm *runeRangeMap
		ranges := tr.ranges

		for i := len(ranges) - 1; i >= 0; i-- {
			rrm = ranges[i]

			if rrm.FromLo <= r && r <= rrm.FromHi {
				translated = true

				if tr.mappedRune >= 0 {
					result = tr.mappedRune
					break
				}

				if rrm.ToLo < rrm.ToHi {
					result = rrm.ToLo + r - rrm.FromLo
				} else if rrm.ToLo > rrm.ToHi {
					//如果range是从更高到更低，ToHi可以比ToLo更小
					result = rrm.ToLo - r + rrm.FromLo
				} else {
					result = rrm.ToLo
				}

				break
			}
		}
	}

	if tr.reverted {
		if !translated {
			result = tr.mappedRune
		}

		translated = !translated
	}

	if !translated {
		result = r
	}

	return
}

//如果translator至少有一个pattern则返会true
func (tr *Translator) HasPattern() bool {
	return tr.hasPattern
}

//把str里将定义在to 里的字符替换掉定义在from里的字符。
//Translate将尝试1对映射从from到to
//如果to是比from更少的,to里的最后一个字符将会被持续映射使用。
//如果to pattern是空字符串,等效于Delete函数
//特殊字符：
//1.'-'   意味着字符的范围。
//2."a-z" 意味着所有从'a' to 'z'的字符。
//3."z-a" 意味着所有从'z' to 'a'的字符。
//4.'^' 作为第一个字符意味着所有被排除在列表里的字符（注意：'^'只在from里起作用，在to里将会被考虑作为一个正常的字符）
//5."^a-z"意味着除‘a’到'z'之外的所有字符。
//6.'\'   意味着特殊的字符。
//7. "abc" 是一个包括'a', 'b' and 'c'的集合。
func Translate(str, from, to string) string {
	tr := NewTranslator(from, to)
	return tr.Translate(str)
}

//删除文本串中与模式串中匹配的字符。
func Delete(str, pattern string) string {
	tr := NewTranslator(pattern, "")
	return tr.Translate(str)
}

//计算文本串中有多少字符与模式串匹配。
func Count(str, pattern string) int {
	if pattern == "" || str == "" {
		return 0
	}

	var r rune
	var size int
	var matched bool

	tr := NewTranslator(pattern, "")
	cnt := 0

	for len(str) > 0 { //不断的遍历每一个str里的每一个字符
		r, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		if _, matched = tr.TranslateRune(r); matched {
			cnt++
		}
	}

	return cnt
}

// 删除文本串中相邻的字符。
//如果模式串不为空，文本串里只有模式串字串进行Squeeze
func Squeeze(str, pattern string) string {
	var last, r rune
	var size int
	var skipSqueeze, matched bool
	var tr *Translator
	var output *stringBuilder

	orig := str
	last = -1

	if len(pattern) > 0 {
		tr = NewTranslator(pattern, "")
	}

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		//需要squeeze这个str
		if last == r && !skipSqueeze {
			if tr != nil {
				if _, matched = tr.TranslateRune(r); !matched {
					skipSqueeze = true
				}
			}

			if output == nil {
				output = allocBuffer(orig, str)
			}

			if skipSqueeze {
				output.WriteRune(r)
			}
		} else {
			if output != nil {
				output.WriteRune(r)
			}

			last = r
			skipSqueeze = false
		}

		str = str[size:]
	}

	if output == nil {
		return orig
	}

	return output.String()
}
