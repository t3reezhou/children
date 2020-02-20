package children

import (
	"fmt"

	"github.com/t3reezhou/atest"
)

func init() {
	atest.Regist("88")
	fmt.Println(atest.XX)
}

func Regist(x atest.X) {
	atest.Regist(x)
}
