package module1

import (
	"fmt"
	fun "github.com/fanghaos/go_mod_t/module1/fun/v2"
)

func Fun1() {
	fmt.Println("module1")
	fun.Module1Fun()
}

func Fun2() {
	fmt.Println("module2")
}

func Fun3() {
	fmt.Println("module3")
}
