package index

import (
	"fmt"
	"strings"
)

func T1_index() {
	sSrc := "chicken"
	sFind := "ken"
	sObj := strings.Index(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 子串的位置?:[%v] \n", sSrc, sFind, sObj)

	sFind = "dmr"
	sObj = strings.Index(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 子串的位置?:[%v] \n", sSrc, sFind, sObj)

}
