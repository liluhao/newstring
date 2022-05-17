# newstring

## 项目介绍

* 针对Go package [strings]包的里函数的缺陷型，即“有很多关于字符串的函数在其他语言里都有，但在Go语言里没有”，newstring包是字符串函数的集合，里面新增了许多函数

* 所有功能都经过充分测试，并经过精心调优以确保性能。

## 函数列表

### 函数列表

| Function         | signature                                                  |      |
| ---------------- | ---------------------------------------------------------- | ---: |
| Center           | func Center(str string,length int,pad string)string        |      |
| Count            | func Count(str,pattern string)int                          |      |
| Delete           | func Delete(str,pattern string)string                      |      |
| ExpandTabs       | func ExpandTabs(str string,tabSize int)string              |      |
| FirstRuneToLower | func FirstRuneToLower(str string)string                    |      |
| FirstRuneToUpper | func FirstRuneToUpper(str string)string                    |      |
| Insert           | func Insert(dst,src string,index int)string                |      |
| LastPartition    | func LastPartition(str,sep string)(head,match,tail string) |      |
| LeftJustify      | func LeftJustify(str string,length int,pad string)string   |      |
| Len              | func Len(str string)int                                    |      |
| Partition        | func Partition(str,sep string)(head,match,tail string)     |      |
| Reverse          | func Reverse(str string)string                             |      |
| RightJustify     | func RightJustify(str string,length int,pad string)string  |      |
| RuneWidth        | func RuneWidth(r rune)int                                  |      |
| Scrub            | func Scrub(str,repl string)string                          |      |
| Shuffle]         | func Shuffle(str string)string                             |      |
| ShuffleSource    | func ShuffleSource(str string,src rand.Source)string       |      |
| Slice            | func Slice(str string,start,end int)string                 |      |
| Squeeze          | func Squeeze(str,pattern string)string                     |      |
| Successor        | func Successor(str string)string                           |      |
| SwapCase         | func SwapCase(str string)string                            |      |
| ToCamelCase      | func ToCamelCase(str string)string                         |      |
| ToKebabCase      | func ToKebabCase(str string) string                        |      |
| ToSnakeCase      | func ToSnakeCase(str string)string                         |      |
| Translate        | func Translate(str,from,to string)string                   |      |
| Width            | func Width(str string)int                                  |      |
| WordCount        | func WordCount(str string)int                              |      |
| WordSplit        | func WordSplit(str string)[]string                         |      |



| type            | signature                                                    |      |
| --------------- | ------------------------------------------------------------ | ---: |
| type Translator | func NewTranslator(from,to string)*Translator                |      |
|                 | func (tr *Translator)HasPattern()bool                        |      |
|                 | func(tr *Translator)Translate(str string)string              |      |
|                 | func(tr *Translator)TranslateRune(r rune)(result rune,translated bool) |      |



### 1.Center

> 如果 str 的字符长度小于length，则 Center 将返回两侧都加有pad的字符串。
> 如果 str 的字符长度大于等于length，则将返回 str 本身。
> 如果 pad 是空字符串，则将返回 str。

```go
func Center(str string,length int,pad string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.Center("hello", 4, " "))    //"hello"
	fmt.Printf("%q\n", newstring.Center("hello", 5, " "))    //"hello"
	fmt.Printf("%q\n", newstring.Center("hello", 10, " "))   //"  hello   "
	fmt.Printf("%q\n", newstring.Center("hello", 10, "123")) //"12hello123"
	fmt.Printf("%q\n", newstring.Center("", 4, "abc"))       //"abab"
	fmt.Printf("%q\n", newstring.Center("ab", 5, ""))        //"ab"
    fmt.Printf("%q\n", newstring.Center("hello", 15, "adc")) //"adcadhelloadcad"
}
```

### 2.Count

> 计算 str 中有多少字符与pattern匹配;
>
> 如果patter或者str是空字符串，则返回0

```go
func Count(str,pattern string)int
```

> 例子

```go
func main() {
	fmt.Println(newstring.Count("hello", "aeiou")) //2
	//a-k：abcdefghijkl
	fmt.Println(newstring.Count("hello", "a-k"))        //2
	fmt.Println(newstring.Count("hellok", "a-k"))       //3
	fmt.Println(newstring.Count("hello", "a-l"))        //4
	fmt.Println(newstring.Count("hello", "^a-k"))       //3
	fmt.Println(newstring.Count("hello", "^a-l"))       //1
	fmt.Println(newstring.Count("hello", "el"))         //3
	fmt.Println(newstring.Count("12hello", "12el"))     //5
	fmt.Println(newstring.Count("12hello", "e12l"))     //5
	fmt.Println(newstring.Count("12hello", "^e12l"))    //2
	fmt.Println(newstring.Count("12hello+-", "e12+-l")) //8
}
```



### 3.Delete

> 删除文本串中与模式串中匹配的字符。

```go
func Delete(str,pattern string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.Delete("hello", "aeiou")) //"hll"
	//a-k：abcdefghijkl
	fmt.Printf("%q\n", newstring.Delete("hello", "a-k"))        //"llo"
	fmt.Printf("%q\n", newstring.Delete("hellok", "a-k"))       //"llo"
	fmt.Printf("%q\n", newstring.Delete("hello", "a-l"))        //"o"
	fmt.Printf("%q\n", newstring.Delete("hello", "^a-k"))       //"he"
	fmt.Printf("%q\n", newstring.Delete("hello", "^a-l"))       //"hell"
	fmt.Printf("%q\n", newstring.Delete("hello", "el"))         //"ho"
	fmt.Printf("%q\n", newstring.Delete("12hello", "12el"))     //"ho"
	fmt.Printf("%q\n", newstring.Delete("12hello", "e12l"))     //"ho"
	fmt.Printf("%q\n", newstring.Delete("12hello", "^e12l"))    //"12ell"
	fmt.Printf("%q\n", newstring.Delete("12hello+-", "e12+-l")) //"o"
}
```



### 4.ExpandTabs

> 可以将 str 中的\t小指定为tabSize。
> 在 str 中出现换行符\n后，column将重置为零。
> 中日韩字符即CJK字符将被视为两个字符。
> 如果 tabSize <= 0，则 ExpandTabs 将panic。

```go
func ExpandTabs(str string,tabSize int)string
```

> 首先要理解\t的使用

```go
func main() {
	//七个空格、六个空格、五个空格、四个空格
	fmt.Println("a\tbc\tdef\tghij\tk") //a       bc      def     ghij    k
	//六个空格
	fmt.Println("aa\tbc") //aa      bc
    //四个空格、六个空格、五个空格、四个空格
	fmt.Println("李李\tbc\tdef\tghij\tk") //李李    bc      def     ghij    k
    //八个空格
	fmt.Println("aaaaaaaa\tbc") //aaaaaaaa        bc
}
```

> 首先要理解\n的使用

```go
func main() {
	//一个空格、七个空格
	fmt.Println("abcdefg\thij\nk\tl")
	/*
		abcdefg hij
		k       l
	*/

	fmt.Println("aa\nbc")
	/*
		aa
		bc
	*/
}
```

> 例子：

```go
func main() {
	//三个空格、二个空格、一个空格、四个空格
	fmt.Printf("%q\n", newstring.ExpandTabs("a\tbc\tdef\tghij\tk", 4)) //"a   bc  def ghij    k"
	//一个空格、三个空格
	fmt.Printf("%q\n", newstring.ExpandTabs("abcdefg\thij\nk\tl", 4)) //"abcdefg hij\nk   l"
	//一个空格、二个空格
	fmt.Printf("%q\n", newstring.ExpandTabs("z中\t文\tw", 4)) //"z中 文  w"
}
```



### 5.FirstRuneToLower

>FirstRuneToLower会将第一个字符转换为小写。

```go
func FirstRuneToLower(str string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.FirstRuneToLower("adsaf"))     //"adsaf"
	fmt.Printf("%q\n", newstring.FirstRuneToLower("ADSSFFDSA")) //"aDSSFFDSA"
	fmt.Printf("%q\n", newstring.FirstRuneToLower("aFSDFSF"))   //"aFSDFSF"
	fmt.Printf("%q\n", newstring.FirstRuneToLower("_aFSDFSF"))  //"_aFSDFSF"
	fmt.Printf("%q\n", newstring.FirstRuneToLower("Aaaaa"))     //"aaaaa"
}
```



### 6.FirstRuneToUpper

> FirstRuneToUpper 会将第一个字符转换为大写。

```go
func FirstRuneToUpper(str string)string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.FirstRuneToUpper("adsaf"))     //"Adsaf"
	fmt.Printf("%q\n", newstring.FirstRuneToUpper("ADSSFFDSA")) //"ADSSFFDSA"
	fmt.Printf("%q\n", newstring.FirstRuneToUpper("aFSDFSF"))   //"AFSDFSF"
	fmt.Printf("%q\n", newstring.FirstRuneToUpper("_aFSDFSF"))  //"_aFSDFSF"
	fmt.Printf("%q\n", newstring.FirstRuneToUpper("Aaaaa"))     //"Aaaaa"
}
```



### 7.Insert

> 将一个字符串插入到另一个字符串中
> 如果index越界,将panic

```go
func Insert(dst,src string,index int)string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.Insert("hello", "aeio", 2))  //"heaeiollo"
	fmt.Printf("%q\n", newstring.Insert("hello", "aeio", 4))  //"hellaeioo"
	fmt.Printf("%q\n", newstring.Insert("hello", "aeiou", 5)) //"helloaeiou"
	fmt.Printf("%q\n", newstring.Insert("hello", "aeiou", 0)) //"aeiouhello"
	//fmt.Printf("%q\n", newstring.Insert("hello", "aeiou", 6)) //panic
	//fmt.Printf("%q\n", newstring.Insert("hello", "aeiou", -1)) //panic
}
```



### 8.LastPartition

> 将字符传分割成三部分.

```go
func LastPartition(str,sep string)(head,match,tail string)
```

> 例子：

```go
func main() {
	a, b, c := newstring.LastPartition("hello", "ll")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"he","ll","o"
	a, b, c = newstring.LastPartition("hello", "cvb")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"","","hello"
	a, b, c = newstring.LastPartition("hello", "cll")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"","","hello"
}
```



### 9.LeftJustify

> 如果str的字符长度小于length，则 RightJustify返回右侧带有 pad 字符串的字符串。
> 如果 str 的字符长度大于等于length，则将返回 str 本身。
> 如果 pad 是空字符串，则将返回 str。

```go
func LeftJustify(str string,length int,pad string)string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.LeftJustify("hello", 4, " "))    //"hello"
	fmt.Printf("%q\n", newstring.LeftJustify("hello", 5, " "))    //"hello"
	fmt.Printf("%q\n", newstring.LeftJustify("hello", 10, " "))   //"hello     "
	fmt.Printf("%q\n", newstring.LeftJustify("hello", 10, "123")) //"hello12312"
	fmt.Printf("%q\n", newstring.LeftJustify("", 4, "abc"))       //"abca"
	fmt.Printf("%q\n", newstring.LeftJustify("ab", 5, ""))        //"ab"
}
```



### 10.Len

> 返回字符串字符长度

```go
func Len(str string)int
```

> 例子：

```go
func main() {
    fmt.Println(len("eggo世界"))             //10
	fmt.Println(newstring.Len("eggo世界")) //6
}
```

### 11.Partition

> 将字符传分割成三部分.

```go
func Partition(str,sep string)(head,match,tail string)
```

> 例子：

```go
func main() {
	a, b, c := newstring.Partition("hello", "ll")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"he","ll","o"
	a, b, c = newstring.Partition("hello", "cvb")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"hello","",""
	a, b, c = newstring.Partition("hello", "cll")
	fmt.Printf("%q,%q,%q\n", a, b, c) //"hello","",""
}
```

### 12.Reverse

> 反转字符串

```go
func Reverse(str string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.Reverse("adsfasf")) //"fsafsda"
	fmt.Printf("%q\n", newstring.Reverse("123feqf")) //"fqef321"
	fmt.Printf("%q\n", newstring.Reverse("-=qfsf"))  //"fsfq=-"
	fmt.Printf("%q\n", newstring.Reverse("周杰伦"))     //"伦杰周"
	fmt.Printf("%q\n", newstring.Reverse("周杰伦ab1"))  //"1ba伦杰周"
}
```



###  13.RightJustify

> 如果str的字符长小于length，则 RightJustify返回左侧带有 pad 字符串的字符串。
> 如果 str 的字符长度大于等于length，则将返回 str 本身。
> 如果 pad 是空字符串，则将返回 str。

```go
func RightJustify(str string,length int,pad string)string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.RightJustify("hello", 4, " "))    //"hello"
	fmt.Printf("%q\n", newstring.RightJustify("hello", 5, " "))    //"hello"
	fmt.Printf("%q\n", newstring.RightJustify("hello", 10, " "))   //"     hello"
	fmt.Printf("%q\n", newstring.RightJustify("hello", 10, "123")) //"12312hello"
	fmt.Printf("%q\n", newstring.RightJustify("", 4, "abc"))       //"abca"
	fmt.Printf("%q\n", newstring.RightJustify("ab", 5, ""))        //"ab"
}
```



### 14.RuneWidth

> 返回一个字符宽度。
> 普通字符被认为是1来计算。
> 复杂字符再次被认为是2来计算

```go
func RuneWidth(r rune)int
```

> 例子：

```go
func main() {
	fmt.Println(newstring.RuneWidth(' ')) //1
	fmt.Println(newstring.RuneWidth('a')) //1
	fmt.Println(newstring.RuneWidth('4')) //1
	fmt.Println(newstring.RuneWidth('伦')) // 2
	fmt.Println(newstring.RuneWidth('+')) //1
	fmt.Println(newstring.RuneWidth('-')) //1
}
```

### 15.Scrub

> 使用repl字符串去掉无效的字符。
> 相邻的无效字符仅替换一次。

```go
func Scrub(str,repl string)string
```

> 例子:

```go
func main() {
	fmt.Printf("%q\n", newstring.Scrub("adsfasf", "fa"))   //"adsfasf"
	fmt.Printf("%q\n", newstring.Scrub("adsfasf", "VNBB")) //"adsfasf"
	fmt.Printf("%q\n", newstring.Scrub("adsfasf", "af"))   //"adsfasf"
	fmt.Printf("%q\n", newstring.Scrub("fa", "adsfasf"))   //"fa"
	fmt.Printf("%q\n", newstring.Scrub("", "adsfasf"))     //"fa"
}
```



### 16.Shuffle

> 随机重组字符串中的字符并返回结果。

```go
func Shuffle(str string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.Shuffle("adsaf"))     //"afsad"
	fmt.Printf("%q\n", newstring.Shuffle("ADSSFFDSA")) //"SAFADSSDF"
	fmt.Printf("%q\n", newstring.Shuffle("aFSDFSF"))   //"FFaSDSF"
	fmt.Printf("%q\n", newstring.Shuffle("_aFSDFSF"))  //"SFD_aSFF"
	fmt.Printf("%q\n", newstring.Shuffle("Aaaaa"))     //"aAaaa"
}
```



### 17.ShuffleSource

> 使用被给的随机源rand.Source进行随机化字符

```go
func ShuffleSource(str string,src rand.Source)string
```

> 例子：

```go
func main() {
	a := rand.NewSource(56)                                     //使用给定的种子创建一个伪随机资源。
	fmt.Printf("%q\n", newstring.ShuffleSource("adsaf", a))     //"afsad"
	fmt.Printf("%q\n", newstring.ShuffleSource("ADSSFFDSA", a)) //"FFSDSAASD"
	fmt.Printf("%q\n", newstring.ShuffleSource("aFSDFSF", a))   //"DFSFSaF"
	fmt.Printf("%q\n", newstring.ShuffleSource("_aFSDFSF", a))  //"SDFaSFF_"
	fmt.Printf("%q\n", newstring.ShuffleSource("Aaaaa", a))     //"aaaaA"
}
```



### 18.Slice

> string[i,j]用于切割字节，但是本函数用于分割字符串的字符
> start必须满足 0 <= start <= str的字符数.
> End可以是0也可以是负数也可以是正数
> 如果end >= 0, 比如满足: start <= end <= str的字符数.
> 如果  end < 0, 一直切割到字符串末端

```go
func Slice(str string,start,end int)string
```

> 例子：

```go
func main() {
    fmt.Printf("%q\n", newstring.Slice("FirstName", 2, 5)) //"rst"
	fmt.Printf("%q\n", newstring.Slice("FirstName", 1, 5)) //"irst"
	fmt.Printf("%q\n", newstring.Slice("FirstName", 1, 9)) //"irstName"
	fmt.Printf("%q\n", newstring.Slice("FirstName", 2, 2))    //""
	fmt.Printf("%q\n", newstring.Slice("FirstName", 1, -100)) //"irstName"
	fmt.Printf("%q\n", newstring.Slice("周杰伦唱歌听听", 1, 2))      //"杰"
	fmt.Printf("%q\n", newstring.Slice("周杰伦唱歌听听", 1, 3))      //"杰伦"
	fmt.Printf("%q\n", newstring.Slice("周杰伦唱歌听听", 1, 7))      //"杰伦唱歌听听"

	//fmt.Printf("%q\n", newstring.Slice("周杰伦唱歌听听", 1, 8))      //panic
	//fmt.Printf("%q\n", newstring.Slice("FirstName", 1, 10)) //panic
	//fmt.Printf("%q\n", newstring.Slice("FirstName", 2, 1)) //panic
	//fmt.Printf("%q\n", newstring.Slice("FirstName", -1, 5)) //panic
	//fmt.Printf("%q\n", newstring.Slice("HTTPServer", 1, 12)) //panic
}
```



### 19.Squeeze

>删除文本串中相邻的字符。
>如果模式串不为空，文本串里只有模式串字串进行Squeeze

```go
func Squeeze(str,pattern string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.Squeeze("hello", "m-z")) //"hello"
	//三个空格、二个空格、一个空格
	fmt.Printf("%q\n", newstring.Squeeze("hello   world", " ")) //"hello world"

	fmt.Printf("%q\n", newstring.Squeeze("hellowwwwwwwworld", "w"))   //"helloworld"
	fmt.Printf("%q\n", newstring.Squeeze("hellowwwwwwwworld", "ww"))  //"helloworld"
	fmt.Printf("%q\n", newstring.Squeeze("hellowwwwwwwworld", "www")) //"helloworld"
}
```

### 20.Successor 

> 如下的符合Alphanumeric:
> a - z
> A - Z
> 0 - 9
>
> 将str里符合最右边的Alphanumeric加1。
>
> 将str里符合最右边的Alphanumeric的是z、Z、9的话，则还会继续再寻找一个Alphanumeric加1，如果这个Alphanumeric仍是z、Z、9的话，则还会再次继续寻找一个Alphanumeric加1。
>
> 将str里全是Alphanumeric，则z、Z、9，则还会在str最左侧进位。

> 如果str是空字符，则结果仍返回字符串。
>
> 如果str里没有Alphanumeric字符，则无论结果是否为有效的字符，将str的最右边的字符都增加1。

```go
func Successor(str string)string
```

> 例子：

```go
func main() {
	//将str里符合最右边的Alphanumeric加1
	fmt.Printf("%q\n", newstring.Successor("THX1138"))   //"THX1139"
	fmt.Printf("%q\n", newstring.Successor("abcd"))      //"abce"
	fmt.Printf("%q\n", newstring.Successor("<<koala>>")) //"<<koalb>>"

	//将str里符合最右边的Alphanumeric的是z、Z、9的话，则还会继续再寻找一个Alphanumeric加1，如果这个Alphanumeric仍是z、Z、9的话，则还会再次继续寻找一个Alphanumeric加1
	fmt.Printf("%q\n", newstring.Successor("ab>Z"))    //"ac>A"
	fmt.Printf("%q\n", newstring.Successor("ab>9>Z"))  //"ac>0>A"
	fmt.Printf("%q\n", newstring.Successor("THX1139")) //"THX1140"
	fmt.Printf("%q\n", newstring.Successor("abcz"))    //"abda"
	fmt.Printf("%q\n", newstring.Successor("THX11a9")) //"THX11b0"
	fmt.Printf("%q\n", newstring.Successor("1aZ"))     //"1bA"
	fmt.Printf("%q\n", newstring.Successor("109"))     //"110"
	fmt.Printf("%q\n", newstring.Successor("1999zzz")) //"2000aaa"
    fmt.Printf("%q\n", newstring.Successor("abc周杰伦99"))  //"abd周杰伦00"


	//将str里全是Alphanumeric，则z、Z、9，则还会在str最左侧进位。
	fmt.Printf("%q\n", newstring.Successor("99"))        //"100"
	fmt.Printf("%q\n", newstring.Successor("ZZZ9999"))   //"BAAA0000"
	fmt.Printf("%q\n", newstring.Successor(">>ZZZ9999")) //">>BAAA0000"
    fmt.Printf("%q\n", newstring.Successor("周杰伦ZZZ9999")) //"周杰伦BAAA0000"

	//如果str是空字符，则结果仍返回字符串
	fmt.Printf("%q\n", newstring.Successor("")) //""

	//如果str里没有Alphanumeric字符，则无论结果是否为有效的字符，将str的最右边的字符都增加1。
	fmt.Printf("%q\n", newstring.Successor("***")) //"**+"
	fmt.Printf("%q\n", newstring.Successor("+"))   //","
	fmt.Printf("%q\n", newstring.Successor("-"))   //"."
	fmt.Printf("%q\n", newstring.Successor("*"))   //"+"
	fmt.Printf("%q\n", newstring.Successor("/"))   //"0"
	/*ASCII码:
	* :42
	+ :43
	, :44
	- :45
	. :46
	/ :47
	0: 48
	*/
}
```



### 21.SwapCase

> 小写变大写，大写变小写

```go
func SwapCase(str string)string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.SwapCase("hello"))  //"HELLO"
	fmt.Printf("%q\n", newstring.SwapCase("Hello"))  //"hELLO"
	fmt.Printf("%q\n", newstring.SwapCase("hELLO"))  //"Hello"
	fmt.Printf("%q\n", newstring.SwapCase("_hello")) //"_HELLO"
	fmt.Printf("%q\n", newstring.SwapCase(""))       //""
}
```



### 22.ToCamelCase

> 是将用空格，下划线和连字符分隔的单词转换为只有首字母大写

```go
func ToCamelCase(str string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.ToCamelCase("SOme_words"))  //"SomeWords"
	fmt.Printf("%q\n", newstring.ToCamelCase("Some_words"))  //"SomeWords"
	fmt.Printf("%q\n", newstring.ToCamelCase("SOME+WORDS"))  //"Some+words"
	fmt.Printf("%q\n", newstring.ToCamelCase("http_server")) //"HttpServer"
	fmt.Printf("%q\n", newstring.ToCamelCase("no_https"))    //"NoHttps"
	fmt.Printf("%q\n", newstring.ToCamelCase("_complex__case_")) //"_Complex_Case_"
	fmt.Printf("%q\n", newstring.ToCamelCase("some words"))      //"SomeWords"
	fmt.Printf("%q\n", newstring.ToCamelCase("appfruit"))        //"Appfruit"
	fmt.Printf("%q\n", newstring.ToCamelCase("appfr+uit"))       //"Appfr+uit"
	fmt.Printf("%q\n", newstring.ToCamelCase("appfr++uit"))      //"Appfr++uit"
}

```



### 23.ToKebabCase

> 可以将字符串中的所有大写字符转换为小写字符，并根据情况用连接符连接

```go
func ToKebabCase(str string) string
```

> 例子：

```
func main() {
	fmt.Printf("%q\n", newstring.ToKebabCase("FirstName"))    //"first-name"
	fmt.Printf("%q\n", newstring.ToKebabCase("HTTPServer"))   //"http-server"
	fmt.Printf("%q\n", newstring.ToKebabCase("NoHTTPS"))      //"no-https"
	fmt.Printf("%q\n", newstring.ToKebabCase("GO_PATH"))      //"go-path"
	fmt.Printf("%q\n", newstring.ToKebabCase("GO PATH"))      //"go-path"
	fmt.Printf("%q\n", newstring.ToKebabCase("GO-PATH"))      //"go-path"
	fmt.Printf("%q\n", newstring.ToKebabCase("http2xx"))      //"http-2xx"
	fmt.Printf("%q\n", newstring.ToKebabCase("HTTP20xOK"))    //"http-20x-ok"
	fmt.Printf("%q\n", newstring.ToKebabCase("Duration2m3s")) //"duration-2m3s"
	fmt.Printf("%q\n", newstring.ToKebabCase("Bld4Floor3rd")) //"bld4-floor-3rd"
}
```



### 24.ToSnakeCase

> 可以将字符串中的所有大写字符转换小写字符，并根据情况用下划线连接

```go
func ToSnakeCase(str string)string
```

> 例子：

```go
func main() {
	fmt.Printf("%q\n", newstring.ToSnakeCase("FirstName"))    //"first_name"
	fmt.Printf("%q\n", newstring.ToSnakeCase("HTTPServer"))   //"http_server"
	fmt.Printf("%q\n", newstring.ToSnakeCase("NoHTTPS"))      //"no_https"
	fmt.Printf("%q\n", newstring.ToSnakeCase("GO_PATH"))      //"go_path"
	fmt.Printf("%q\n", newstring.ToSnakeCase("GO PATH"))      //"go_path"
	fmt.Printf("%q\n", newstring.ToSnakeCase("GO-PATH"))      //"go_path"
	fmt.Printf("%q\n", newstring.ToSnakeCase("http2xx"))      //"http_2xx"
	fmt.Printf("%q\n", newstring.ToSnakeCase("HTTP20xOK"))    //"http_20x_ok"
	fmt.Printf("%q\n", newstring.ToSnakeCase("Duration2m3s")) //"duration_2m3s"
	fmt.Printf("%q\n", newstring.ToSnakeCase("Bld4Floor3rd")) //"bld4_floor_3rd"
}
```



### 26.Translate

> 把str里将定义在to 里的字符替换掉定义在from里的字符。
> Translate将尝试1对映射从from到to
> 如果to是比from更少的,to里的最后一个字符将会被持续映射使用。
> 如果to pattern是空字符串,等效于Delete函数
> 特殊字符：
> 1.'-'   意味着字符的范围。
> 2."a-z" 意味着所有从'a' to 'z'的字符。
> 3."z-a" 意味着所有从'z' to 'a'的字符。
> 4.'^' 作为第一个字符意味着所有被排除在列表里的字符（注意：'^'只在from里起作用，在to里将会被考虑作为一个正常的字符）
> 5."^a-z"意味着除‘a’到'z'之外的所有字符。
> 6.'\'   意味着特殊的字符。
> 7. "abc" 是一个包括'a', 'b' and 'c'的集合。

```go
func Translate(str,from,to string)string
```

> 例子：

```go
func main() {
    fmt.Printf("%q\n", newstring.Translate("hello", "hello", "123")) //"12333"
	fmt.Printf("%q\n", newstring.Translate("hello", "aeiou", "12345")) //"h2ll4"
	//a-k：abcdefghijkl
	fmt.Printf("%q\n", newstring.Translate("hello", "a-z", "A-Z"))        //"HELLO"
	fmt.Printf("%q\n", newstring.Translate("hello", "z-a", "a-z"))        //"svool"
	fmt.Printf("%q\n", newstring.Translate("hello", "aeiou", "*"))        //"h*ll*"
	fmt.Printf("%q\n", newstring.Translate("hello", "^l", "*"))           //"**ll*"
	fmt.Printf("%q\n", newstring.Translate("hello ^ world", `\^lo`, "*")) //"he*** * w*r*d"
}
```



### 27.Width 

>返回字符串宽度。
>普通字节被认为是1来计算。
>复杂字符再次被认为是2来计算

```go
func Width(str string)int
```

> 例子：

```go
func main() {
	fmt.Println(newstring.Width(""))            //0
	fmt.Println(newstring.Width("asf"))         //3
	fmt.Println(newstring.Width("478asf"))      //6
	fmt.Println(newstring.Width("周杰伦"))         //6
	fmt.Println(newstring.Width("周杰伦fdas"))     //10
	fmt.Println(newstring.Width("、"))           //2
	fmt.Println(newstring.Width("=-、/周杰伦fdas")) //10
}
```



### 28.WordCount

>返回字符串中的单词数。
>可以包含`'` 和`-`，包含`'` 和`-`的认为是一个单词的连接符。(`'`是英文的单引号)
>
>但是若以其他字符开头，则忽略它的存在。

```go
func WordCount(str string)int
```

```go
func main() {
	fmt.Println(newstring.WordCount(""))         //0
	fmt.Println(newstring.WordCount(" abc"))     //1
	fmt.Println(newstring.WordCount(" abc bnb")) //2

	fmt.Printf("%q\n", "-")                      //"-"
	fmt.Println(newstring.WordCount("-"))        //0
	fmt.Println(newstring.WordCount("-abc"))     //1
	fmt.Println(newstring.WordCount("-abc-bnb")) //1

	fmt.Printf("%q\n", "'")                      //"'"
	fmt.Println(newstring.WordCount("'"))        //0
	fmt.Println(newstring.WordCount("'abc"))     //1
	fmt.Println(newstring.WordCount("'abc'bnb")) //1
}
```



### 29.WordSplit

> 割str成为的单词，并以切片形式返回
> 如果str是空字符串，返回nil
> 可以包含`'` 和`-`，包含`'` 和`-`的认为是一个单词的连接符。(`'`是英文的单引号)
> 但是若以其他字符开头，则忽略它的存在。

```go
func WordSplit(str string)[]string
```

> 例子：

```go
func main() {
	fmt.Println(newstring.WordSplit("")) //[]

	fmt.Println(newstring.WordSplit(" "))        //[]
	fmt.Println(newstring.WordSplit(" aaa"))     //[aaa]
	fmt.Println(newstring.WordSplit(" abc bnb")) //[abc bnb]

	fmt.Println(newstring.WordSplit("-"))        //[]
	fmt.Println(newstring.WordSplit("-abc"))     //[abc]
	fmt.Println(newstring.WordSplit("-abc-bnb")) //[abc-bnb]

	fmt.Println(newstring.WordSplit("'"))        //[]
	fmt.Println(newstring.WordSplit("'abc"))     //[abc]
	fmt.Println(newstring.WordSplit("'abc'bnb")) //[abc'bnb]

	fmt.Println(newstring.WordSplit("\\"))         //[]
	fmt.Println(newstring.WordSplit("'\\abc"))     //[abc]
	fmt.Println(newstring.WordSplit("\\abc\\bnb")) //[abc bnb]

	fmt.Println(newstring.WordSplit("bnb25ds33gsg")) //[bnb ds gsg]
}
```



### 30.IsAlphabet

> 检查r是不是CJK字符的字母，若是返回true，若不是false

```
func IsAlphabet(r rune) bool 
```

> 例子：

```go
func main() {
	fmt.Println(newstring.IsAlphabet('a'))  //true
	fmt.Println(newstring.IsAlphabet('a'))  //true
	fmt.Println(newstring.IsAlphabet('a'))  //true
	fmt.Println(newstring.IsAlphabet('z'))  //true
	fmt.Println(newstring.IsAlphabet('1'))  //false
	fmt.Println(newstring.IsAlphabet('-'))  //false
	fmt.Println(newstring.IsAlphabet(','))  //false
	fmt.Println(newstring.IsAlphabet('-'))  //false
	fmt.Println(newstring.IsAlphabet('+'))  //false
	fmt.Println(newstring.IsAlphabet('\\')) //false
	fmt.Println(newstring.IsAlphabet('\'')) //false
	fmt.Println(newstring.IsAlphabet('周'))  //false
	fmt.Println(newstring.IsAlphabet('节'))  //false
}
```



### 31.type Translator

> 以下函数主要用于对内使用，在次不多描述

#### NewTranslator

```go
func NewTranslator(from,to string)*Translator
```

#### HasPattern

```go
func (tr *Translator)HasPattern()bool
```

#### Translate

```go
func(tr *Translator)Translate(str string)string
```

#### TranslateRune

```go
func(tr *Translator)TranslateRune(r rune)(result rune,translated bool)
```



