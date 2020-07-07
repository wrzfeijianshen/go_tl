package join

import (
	"fmt"
	"strings"
)

// 将一系列字符串连接为一个字符串，之间用sep来分隔
func T1_join() {
	sSrc := []string{"foo", "bar", "baz"}
	sSep := ","
	sObj := strings.Join(sSrc, sSep)
	fmt.Printf("src:[%s], sSep:[%s], 组合:[%v] \n", sSrc, sSep, sObj)

}
