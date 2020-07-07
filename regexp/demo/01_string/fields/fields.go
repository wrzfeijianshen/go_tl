package fields

import (
	"fmt"
	"strings"
)

func T1_fields() {
	sSrc := "  foo bar  baz   "
	sObj := strings.Fields(sSrc)
	fmt.Printf("fields src:[%s], obj?:[%q] \n", sSrc, sObj)
}
