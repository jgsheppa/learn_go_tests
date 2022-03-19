package main

import (
	"fmt"

	"github.com/jgsheppa/learn_go_tests/hello"
	"github.com/jgsheppa/learn_go_tests/integer"
	"github.com/jgsheppa/learn_go_tests/iteration"
)

func main() {
	fmt.Println(hello.Hello("there", "EN"))
	fmt.Println(integer.Add(3, 3))
	fmt.Println(iteration.Repeat("hey", 5))
}