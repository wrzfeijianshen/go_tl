package main

import (
	"fmt"
	"regexp"
)

// T2_1 : x
func T2_1() {
	// 量词

	// 匹配六位数字构成的字符串,连续的

	fmt.Println("量词...")
	text := "7121"
	reg := regexp.MustCompile(`\d`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["7" "1" "2" "1"] --> 匹配数字的字符 ,字符数组有4个

	text = "7121"
	reg = regexp.MustCompile(`\d\d\d\d\d\d`)

	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// [] --> 没有匹配到

	text = "7121a11aa1234567"
	reg = regexp.MustCompile(`\d\d\d\d\d\d`)

	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["123456"] -->满足6个连续的数字

	// 简写为\d{6} 确定长度为6
	reg = regexp.MustCompile(`\d{6}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["123456"]

	// 一般形式
	// {n} 		: 	之前的元素必须出现n次
	// {m,n}	:	之前的元素最少出现m次,最多出现n次
	// {m,}		: 	之前元素最少出现m次,出现次数没有上限,无数次
	// {0,n}	: 	之前的元素可以没有,但最多出现n次  某些语言写作{,n}

	// a 出现3次
	text = "7121a11aaa123aa4567"
	reg = regexp.MustCompile(`a{3}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["aaa"]

	text = "7121a11aaa123aa456aaaa7"
	reg = regexp.MustCompile(`a{2,3}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["aaa" "aa" "aaa"] --> 从第一个字符开始,判断有没有a,有a有几个,满足2成立,看第三个是否是a
	// 是a则记录此串,然后后面的串判断有没有a,满足a的个数是2-3个,是的留下.
	// aaaa--> 取走3个 ,最后一个a 不满足个数为2-3

	reg = regexp.MustCompile(`a{2,}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["aaa" "aa" "aaaa"]
	// 说明 取连续的a,不连续则继续寻找下一个a,不限次

	reg = regexp.MustCompile(`a{0,3}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["" "" "" "" "a" "" "aaa" "" "" "aa" "" "" "aaa" "a" ""]
}

// T2_2 :
func T2_2() {
	// 常用量词 + ? * 形态虽然不同于{m,n},功能却是相同的,称为量词简记法.

	// 元素也叫子表达式,代表正则表达式的某个部分,比如 [a-z]

	// 常用量词     {m,n}的等价形式      说明
	// *   			{0,}			  可能出现,可能不出现,出现次数没有上限
	// +        	{1,}			  至少出现1次,出现的次数没有上限
	// ?      		{0,1}			  至少出现1次,也可能不出现

	// [^x{0,n}] 取出不是x的字符,也就是取出是x前的字符,换言之是取直到x的字符.
	text := "<html>11aa"
	reg := regexp.MustCompile(`<[^>]{0,}>`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["<html>"]

	reg = regexp.MustCompile(`<[^>]+>`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["<html>"]

	text = "</html>11aa"
	reg = regexp.MustCompile(`<[^>]+>`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["</html>"]

	text = "\"some\""
	reg = regexp.MustCompile(`\"[^\"]*\"$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	text = "aa123456dffd"
	reg = regexp.MustCompile(`\d{6}`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["123456"]

	// . 点号 元字符 , 一般文档都说可以匹配任意字符.
	text = "a"
	reg = regexp.MustCompile(`.$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["a"]
	text = "0"
	reg = regexp.MustCompile(`.$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["0"]
	text = "\r"
	reg = regexp.MustCompile(`.$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["\r"]

	// 不能匹配\n
	text = "\n"
	reg = regexp.MustCompile(`.`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))

	// 单行模式
	reg = regexp.MustCompile(`(?s)^.$`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["\n"]

	// 匹配优先量词 贪婪量词
	// .*

	// 忽略优先量词 懒惰量词 *? 在优先量词 后面添加?

	// 转义  \*\?  \{n}

	// () 这个功能叫做分组 ,把()里的当为一个整体,有+来限制
	// ab+ (ab)+

	// 单词边界 \b  -->  \b\w+\b

}
