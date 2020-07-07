package main

import (
	"fmt"
	"strings"
)

func isSlash(r rune) bool {
	return r == '\\' || r == '/' || r == 'H'
}

func T1_Trim() {
	s := "\\///\\HostName\\C\\Windows\\/////\\"
	ts := strings.TrimRightFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "\\\\HostName\\C\\Windows"

	ts = strings.TrimLeftFunc(s, isSlash)
	fmt.Printf("%q\n", ts)
	ts = strings.TrimFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "ostName\\C\\Windows"

}
