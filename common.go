package newstring

const bufferMaxInitGrowSize = 2048

// 初始化一个stringBuilder
//注意：len(orig)必须大于等于len(cur),否则会panic
//fmt.Printf("%q\n", allocBuffer("hello", "a").String())     //"hell"
//fmt.Printf("%q\n", allocBuffer("hello", "ae").String())    //"hel"
//fmt.Printf("%q\n", allocBuffer("hellok", "aio").String())  //"hel"
//fmt.Printf("%q\n", allocBuffer("hello", "aope").String())  //"h"
//fmt.Printf("%q\n", allocBuffer("hello", "aoper").String()) //""
//fmt.Printf("%q\n", allocBuffer("hello", "apokjl").String()) //panic
func allocBuffer(orig, cur string) *stringBuilder {
	output := &stringBuilder{}
	maxSize := len(orig) * 4

	//避免一次分配太多内存
	if maxSize > bufferMaxInitGrowSize {
		maxSize = bufferMaxInitGrowSize
	}

	output.Grow(maxSize)
	output.WriteString(orig[:len(orig)-len(cur)])
	return output
}
