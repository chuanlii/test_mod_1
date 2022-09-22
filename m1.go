package test_mod_1

import "fmt"

func init() {
	fmt.Println("testmod1")
}
func F1() (string, error) {
	return "f1", nil
}
