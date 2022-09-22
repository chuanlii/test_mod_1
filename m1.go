package test_mod_1

import (
	"fmt"
	base "github.com/chuanlii/test_mod_base"
)

func init() {
	fmt.Println("testmod1")
}
func F1() (string, error) {
	base.Base()
	return "f1", nil
}
