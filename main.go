package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jgsheppa/learn_go_tests/hello"
	"github.com/jgsheppa/learn_go_tests/injection"
	"github.com/jgsheppa/learn_go_tests/integer"
	"github.com/jgsheppa/learn_go_tests/iteration"
)

func main() {
	fmt.Println(hello.Hello("there", "EN"))
	fmt.Println(integer.Add(3, 3))
	fmt.Println(iteration.Repeat("hey", 5))
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(injection.MyGreeterHandler)))
}