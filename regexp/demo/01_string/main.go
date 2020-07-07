package main

import (
	"fmt"
	"strings"

	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/contains"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/fields"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/index"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/join"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/repeat"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/replace"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/split"
	"github.com/wrzfeijianshen/go_tl/regexp/demo/01_string/trim"
)

func main() {
	contains.T1_contains()
	contains.T1_containsAny()

	join.T1_join()
	index.T1_index()
	repeat.T1_repeat()
	split.T1_split()
	trim.T1_trim()
	fields.T1_fields()

	s := "heLLo worLd Ａｂｃ"
	// 转为大写
	us := strings.ToUpper(s)
	// 转为小写
	ls := strings.ToLower(s)
	ts := strings.ToTitle(s)
	fmt.Printf("%q\n", us) // "HELLO WORLD ＡＢＣ"
	fmt.Printf("%q\n", ls) // "hello world ａｂｃ"
	fmt.Printf("%q\n", ts) // "HELLO WORLD ＡＢＣ"

	// 首字母大写
	// func Title(s string) string
	s = "heLLo worLd"
	ts = strings.Title(s)
	fmt.Printf("%q\n", ts) // "HeLLo WorLd"

	T1_Trim()
	replace.T1_replace()
}
