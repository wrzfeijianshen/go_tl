package repeat

import (
	"fmt"
	"strings"
)

// 返回n个串
func T1_repeat() {
	sSrc := "chicken^"
	sNum := 3
	sObj := strings.Repeat(sSrc, sNum)
	fmt.Printf("src:[%s], num:[%v], out?:[%v] \n", sSrc, sNum, sObj)

}
