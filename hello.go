package main

import "fmt"
import "os"

func main() {
	fmt.Println("hello world")
	var l int
	l = len(os.Args)
	fmt.Println("lenght: ", l)
	for i := 0; i < l; i++ {
		fmt.Println("Args[", i, "]: ", os.Args[i])
	}
}
