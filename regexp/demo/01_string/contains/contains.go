package contains

import (
	"fmt"
	"strings"
)

// T1_contains:  判断字符串是否包含子串
func T1_contains() {
	sSrc := "seafood"
	sFind := "foo"
	sObj := strings.Contains(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 查到了?:[%v] \n", sSrc, sFind, sObj)

	sFind = "fg"
	sObj = strings.Contains(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 查到了?:[%v] \n", sSrc, sFind, sObj)

}

func T1_containsAny() {
	sSrc := "seafood"
	sFind := "a"
	sObj := strings.ContainsAny(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 查到了?:[%v] \n", sSrc, sFind, sObj)

	sFind = "a "
	sObj = strings.Contains(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 查到了?:[%v] \n", sSrc, sFind, sObj)

	sSrc = "华夏子民"
	sFind = "夏"
	sObj = strings.Contains(sSrc, sFind)
	fmt.Printf("src:[%s], find:[%s], 查到了?:[%v] \n", sSrc, sFind, sObj)

}
