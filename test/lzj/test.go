package main

import (
	"fmt"
	"runtime"
)

func myTest() *int {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("[%v:%v] in explainFlow\n", file, line)
	var a int = 10
	var b = &a
	return b
}

func main() {
	var c = myTest()
	_ = c
}
