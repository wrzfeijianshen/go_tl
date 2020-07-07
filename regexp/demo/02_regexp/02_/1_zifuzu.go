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
