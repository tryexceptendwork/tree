package main

import "fmt"

func main() {
	fmt.Println(test(4))
}
func test (x int) (out int){
	out = 3
	return x
}