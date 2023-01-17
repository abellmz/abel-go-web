package main

import "unicode/utf8"

func main() {
	// 短的，不用换行的，不含双引号的
	println("he said:\" hello go \" ")
	// 长的、复杂的、json串
	println(`he said: "hello, go"
我哈哈哈
`)
	println(len("你好呀"))                     // 字节长度9
	println(utf8.RuneCountInString("你好呀"))  //字符长度 3
	println(utf8.RuneCountInString("你好ab")) //字符长度 4
	println("hello, " + "go")               //字符串拼接
}
