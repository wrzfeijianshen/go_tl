### 正则表达式

当你对你正则表达式不太了解,又觉得是及其深奥难以摸索的,可以看一些经典书籍,如正则指引等.

这篇文章根据网上和书籍的资料,汇聚而成,作为基础.

字符组是正则最基本的结构之一,用 "[]"直接列出所有可能出现的字符,如 [ab]  [k,l]

表示在同一个位置可能出现的各种字符

比如原始字符串 text = "helloabcdk,lpapp" 

正则需要匹配的字符 [ab] 从text中依次判断字符,满足不满足正则[]里面的字符,如h 符合[a]字符?不符合去掉,符合b字符? 不符合,依次e l l o a-->此时的a就符合了

简单的讲,从字符串中的每一个字符,判断是不是[]里面的字符,是就取出来

从前依次取出来,先从第一个字符开始,符合留下,不符合丢弃,仅此而已

```
text := "helloabcdk,lpapp"
reg := regexp.MustCompile(`[ab]`)
fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["a" "b" "a"]

// 判断0123456789字符串
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

```

但有人想说 [0-z] 匹配好像也可以,实际是匹配的是ascII的范围,不建议这样用,精确一些用a-z 0-9 A-Z

```
// 精确匹配十六进制 A-F a-f 0-9即可
text = "7f9l2kd4e6f"
reg = regexp.MustCompile(`[0-9a-fA-F]`)
fmt.Printf("%q\n", reg.FindAllString(text, -1))
// ["7" "f" "9" "2" "d" "4" "e" "6" "f"]

```

[] 中的 "-"

```
//[] 中的 "-" ,"-" 也是一个字符,除了范围表示是字符
text = "7f-9l2kd4-e6f"
reg = regexp.MustCompile(`[-]`)
fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["-" "-"]
reg = regexp.MustCompile(`[6-]`)
fmt.Printf("%q\n", reg.FindAllString(text, -1)) // ["-" "-" "6"]
```

排除型字符组: [^]

在当前位置,匹配一个没有列出的字符 : [^x]

```


```