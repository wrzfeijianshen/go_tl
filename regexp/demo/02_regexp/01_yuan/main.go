package main

import (
	"fmt"
	"regexp"
)

func main() {

	//--1-- 元字符 匹配单个字符
	// 1.匹配除换行符以外的任意字符
	text := "江湖的\n\r\nlas"
	reg := regexp.MustCompile(`.`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // . : ["江" "湖" "的" "\r" "l" "a" "s"]

	// 2.\w 匹配字母或者数字或下划线
	text = "las江2_湖"
	reg = regexp.MustCompile(`\w`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["l" "a" "s" "2" "_"]

	// // \W 匹配汉字
	// reg = regexp.MustCompile(`\W`)
	// fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["江" "湖"]

	// \s 匹配任意的空白符
	text = "l s 1"
	reg = regexp.MustCompile(`\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // : [" " " "]

	// \d 匹配数字
	text = "l s 1 2ds"
	reg = regexp.MustCompile(`\d`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1" "2"]

	// ^ 匹配开始的第一个字符
	text = "ks 1 2ds"
	reg = regexp.MustCompile(`^.`)
	// 	reg = regexp.MustCompile(`^k`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["k"]

	// $ 匹配字符串的结束字符
	text = "ks 1 2ds"
	reg = regexp.MustCompile(`.$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["s"]

	// ==============================
	//--2-- 匹配字符串,有没有子串,查找固定串,查到了返回串
	text = "ks12dos12fd"
	reg = regexp.MustCompile(`s12`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["s12" "s12"]

	text = "ks12dos12fd"
	reg = regexp.MustCompile(`os`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["os"]

	//--3--字符转义 \
	// \. \^ \$ \a \f \t \n \r \v \123 \x7f \x \\  \* \+ \? \{ \} \( \)
	// \[ \] \|
	text = "ks12d.os12fd"
	reg = regexp.MustCompile(`\.o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) //[".o"]

	text = "ks12d^os12fd"
	reg = regexp.MustCompile(`\^o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) //["^o"]

	text = "ks12d$os12fd"
	reg = regexp.MustCompile(`\$o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) //["$o"]

	//--4-- 重复
	// * 重复0次或更多次
	// + 重复1次或更多次
	// ? 重复0次或1次
	// {n} 重复n次
	// {n,} 重复n次或更多次
	// {n,m} 重复n到m次

	// [\d] 匹配数字
	// [^\d] 匹配非数字

	// 反义
	// \W 匹配任意不是字母，数字，下划线的字符
	// \S 匹配不是空白符的字符
	// \D 匹配任意非数字的字符
	// \B 匹配不是单词开头或结束的位置
	// [^x] 匹配除了x以外的任意字符

	// \S+ 匹配不包含空白符的字符串
	text = "ks1 2d$o s12 中 文明 fd"
	reg = regexp.MustCompile(`\S+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) //["ks1" "2d$o" "s12" "中" "文明" "fd"]

	text = "ks12d$os12中文明fd"
	reg = regexp.MustCompile(`[^o]`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//["k" "s" "1" "2" "d" "$" "s" "1" "2" "中" "文" "明" "f" "d"]

	text = "ks12d$os12中文明fd"
	reg = regexp.MustCompile(`\d`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//["k" "s" "1" "2" "d" "$" "s" "1" "2" "中" "文" "明" "f" "d"]

	//--5-- 后向引用 ()

	// 匹配数字开头的任意几位字符 --> 开头字符^
	text = "1332ks12d=os12中文明fd"
	reg = regexp.MustCompile(`^[0-9]*`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1332"]

	// 查找连续的小写字母
	text = "1332ks12d=os12中文明fd"
	reg = regexp.MustCompile(`[a-z]*`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// * 0到多个, 所以匹配了空;先 匹配a-z满足?-->* 第一个字符不是字母,所以
	//["" "" "" "" "ks" "" "d" "os" "" "" "" "" "fd"]

	text = "1332ks12d=os12中文明fd"
	reg = regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["ks" "d" "os" "fd"]

	// 查找连续的非小写字母
	text = "1332ks12d=os12中文明fd"
	// [^xxx] 这里的^ 表示取反
	reg = regexp.MustCompile(`[^a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1332" "12" "=" "12中文明"]

	// 查找连续的单词字母数字
	text = "1332ks12d=os12中文明fd"
	reg = regexp.MustCompile(`[\w]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["1332ks12d" "os12" "fd"]

	// 查找连续的非单词字母,非空白字符
	text = "1332ks1 2d=os 12中文明fd"
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["=" "中文明"]

	// [:alnum:]      字母数字 (相当于 [0-9A-Za-z])
	// [:alpha:]      字母 (相当于 [A-Za-z])
	// [:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
	// [:blank:]      空白占位符 (相当于 [\t ])
	// [:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
	// [:digit:]      数字 (相当于 [0-9])
	// [:graph:]      图形字符 (相当于 [!-~])
	// [:lower:]      小写字母 (相当于 [a-z])
	// [:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
	// [:punct:]      标点符号 (相当于 [!-/:-@[-反引号{-~])
	// [:space:]      空白字符(相当于 [\t\n\v\f\r ])
	// [:upper:]      大写字母(相当于 [A-Z])
	// [:word:]       单词字符(相当于 [0-9A-Za-z_])
	// [:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])

	// 查找连续的大写字母
	text = "1332AAks1 2d=CCos 12中文明fd"
	reg = regexp.MustCompile(`[[:upper:]]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["AA" "CC"]

	// 查找连续的汉字
	text = "1332AA但是ks1 2d=CCos 12中文明fd"
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["但是" "中文明"]

	// 查找 Hello 或 Go
	text = "sdfHello S Gosd 2d=CCos 12中文明fd"

	reg = regexp.MustCompile(`Hello|Go`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	// 查找行首以H开头,以空格结尾的字符串
	text = "HelloS Gosd 2d=CCos 12中文明fd"
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["HelloS Gosd 2d=CCos "]

	reg = regexp.MustCompile(`(?U)^H.*\s`)          // （非贪婪模式）
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["HelloS "]

	// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^.* `) // ["HelloS "]
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	// 查找以空格开头，到行尾结束，中间不包含空格字符串
	reg = regexp.MustCompile(` [^ ]*$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1)) // [" 12中文明fd"]

	// 查找“单词边界”之间的字符串
	reg = regexp.MustCompile(`(?U)\b.+\b`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["HelloS" " " "Gosd" " " "2d" "=" "CCos" " " "12" "中文明" "fd"]

	text = "Helo loS Gosd 2d =o CCos 12o中文明fd"
	// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串  -->?
	reg = regexp.MustCompile(`[^ ]{1,2}o`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello" "Go"]

	text = "1234a qwert789 p1234r"
	// 取到以空格为止的字符串
	reg = regexp.MustCompile(`[^ ]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["1234a" "qwert789" "p1234r"]

	reg = regexp.MustCompile(`[^3]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

}
