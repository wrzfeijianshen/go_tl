package main

import (
	"fmt"
	"regexp"
)

// T_1 :
func T_1() {
	// 字符组
	text := "helloabcdk,lpapp"
	reg := regexp.MustCompile(`[ab]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["a" "b" "a"]

	// 判断0-9
	text = "123rt456io7890pkl"

	// 字符重复了也没有关系,
	reg = regexp.MustCompile(`[012134567899]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1" "2" "3" "4" "5" "6" "7" "8" "9" "0"]

	// 精简,从小到大
	reg = regexp.MustCompile(`[0-9]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1" "2" "3" "4" "5" "6" "7" "8" "9" "0"]

	reg = regexp.MustCompile(`[7-9]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["7" "8" "9"]

	// 字母也是如此
	text = "AAaazuilkKh+=-2890dszkKLOP"
	// 我想取出串中的含有aKz的字符
	reg = regexp.MustCompile(`[aKz]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["a" "a" "z" "K" "z" "K"]

	// 取出小写
	reg = regexp.MustCompile(`[a-z]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["a" "a" "z" "u" "i" "l" "k" "h" "d" "s" "k"]

	// 取出小写大写
	reg = regexp.MustCompile(`[a-zA-Z]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["A" "A" "a" "a" "z" "u" "i" "l" "k" "K" "h" "d" "s" "z" "k" "K" "L" "O" "P"]
	// 取出小写数字
	reg = regexp.MustCompile(`[a-z0-9]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["a" "a" "z" "u" "i" "l" "k" "h" "2" "8" "9" "0" "d" "s" "z" "k"]

	// 精确匹配十六进制 A-F a-f 0-9即可
	text = "7f9l2kd4e6f"
	reg = regexp.MustCompile(`[0-9a-fA-F]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "f" "9" "2" "d" "4" "e" "6" "f"]

	//[] 中的 "-" ,"-" 也是一个字符,除了范围表示是字符
	text = "7f-9l2kd4-e6f"
	reg = regexp.MustCompile(`[-]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["-" "-"]
	reg = regexp.MustCompile(`[6-]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["-" "-" "6"]
}

// T_2 :
func T_2() {

	// 排除型字符组
	fmt.Println("排除型字符组...")
	text := "7f-9l2kd4-e6fA"
	reg := regexp.MustCompile(`[^k]`)
	// 匹配的字符取反
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "f" "-" "9" "l" "2" "d" "4" "-" "e" "6" "f" "A"]

	reg = regexp.MustCompile(`[^kfe]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "-" "9" "l" "2" "d" "4" "-" "6" "A"]

	reg = regexp.MustCompile(`[^0-9a-z]`)
	// 匹配的字符不是数字不是小写字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["-" "-" "A"]
}

// T_3 :
func T_3() {
	// 字符组 简记法
	fmt.Println("字符组 简记法...")
	text := "7f-9l2kd4	 -_e6fA中"
	reg := regexp.MustCompile(`[\d]`)

	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "9" "2" "4" "6"]

	// \d 等价于 [0-9] 数字 digit
	reg = regexp.MustCompile(`\d`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "9" "2" "4" "6"]

	// \w 代表单词 word 等价于 [0-9a-zA-Z_]
	reg = regexp.MustCompile(`\w`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "f" "9" "l" "2" "k" "d" "4" "_" "e" "6" "f" "A"]

	reg = regexp.MustCompile(`[0-9a-zA-Z_]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	// \s 代表空白字符 space 等价于[ \t\r\n\v\f],第一个是空格
	reg = regexp.MustCompile(`\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["\t" " "]

	// 字符组简记法可以单独出现,如上,也可以在字符组中出现,如第一个示例
	// [0-9a-zA-z]可以写为[\da-zA-z] 这样是不是就更为容易了

	// 互补的字符 \D \W \S
	// \S能匹配的字符是^\s ,不是匹配\s的字符就是\S,两者相反,合起来便是所有字符
}

// T_4 :
func T_4() {
	// posix字符组
	// 在某些文档中出现 [:digit:] 数字 [:lower:] 小写
	// 这就是posix字符组,在linux某些工具中sed awk grep 使用.
	// 之前的字符组属于perl衍生处理的正则表达式流派 flavor,叫做 PCRE

	// POSIX ,之前的[a-z] 准确名称叫做 posix方括号表达式, 用法基本一致,转义除外,这里暂不介绍
	// pcre 可以跳出[] posix 必须写在[]里面 : [[:xxx:]]
	fmt.Println("字符组 运算...")
	text := "7f-9l2kd4-_e6fA中"
	reg := regexp.MustCompile(`[[:xdigit:]]`)

	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	// posix		说明				ascII字符组		等价的pcre简记法
	// [:alnum:] 字母和数字			[a-zA-Z0-9]
	// [:alpha:] 	字母				[a-zA-Z]
	// [:ASCII::]	ASCII字符			[\x00-\x7F]
	// [:blank:]	空格和制表符		[ \t]
	// [:cntrl:]	控制字符			[\x-\x1F\x7F]
	// [:digit:]	数字				[0-9]			\d
	// [:graph:]	空白字符之外的字符	[x21-\x7E]
	// [:lower:]	小写字母			[a-z]
	// [:print:]	打印字符			[\x20-\x7E]
	// [:punct:]	标点符号
	// [:space:]    空白字符 			[ \t\r\n\v\f]  \s
	// [:upper:]    大写				[A-Z]
	// [:word:]		字母 				[A-Za-z0-9_]  \w
	// [:xdigit:]   十六进制			[A-Fa-f0-9]
}
