package replace

import (
	"fmt"
	"strings"
)

func T1_replace() {
	s := "Hello 世界！"
	s = strings.Replace(s, " ", ",", -1)
	fmt.Println(s)
	s = strings.Replace(s, "", "|", -1)
	fmt.Println(s)

}
