package trim

import (
	"fmt"
	"strings"
)

func T1_trim() {
	sSrc := " !!! Achtung! Achtung! !!! "
	sSep := "! gA"
	// 依次匹配,删除前后所拥有的字符串
	sObj := strings.Trim(sSrc, sSep)
	fmt.Printf("Trim src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)

	sObj = strings.TrimPrefix(sSrc, sSep)
	fmt.Printf("Trim src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)
	sSep = "!!!!"
	sObj = strings.TrimPrefix(sSrc, sSep)
	fmt.Printf("Trim src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)
	sSep = "!!!"
	sObj = strings.TrimPrefix(sSrc, sSep)
	fmt.Printf("Trim src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)
	sSep = " !!! "
	sObj = strings.TrimPrefix(sSrc, sSep)
	fmt.Printf("Trim src:[%s], sep:[%s], obj?:[%q] \n", sSrc, sSep, sObj)
}
