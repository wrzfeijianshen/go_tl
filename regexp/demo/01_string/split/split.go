package split

import (
	"fmt"
	"strings"
)

func T1_split() {
	sSrc := "a,b,c"
	sSep := ","
	sObj := strings.Split(sSrc, sSep)
	fmt.Printf("src:[%s], sep:[%s], obj?:[%v] \n", sSrc, sSep, sObj)

	sSrc = "a man a plan a canal panama"
	sSep = " "
	sObj = strings.Split(sSrc, sSep)
	fmt.Printf("src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)

	sSrc = "jianghu"
	sSep = ""
	sObj = strings.Split(sSrc, sSep)
	fmt.Printf("src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)

	sSrc = "jianghu"
	sSep = "xx"
	sObj = strings.Split(sSrc, sSep)
	// //若分隔符在原字符串中不存在,返回原字符串切片
	fmt.Printf("src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)
}
