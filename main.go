package main

import "fmt"

func main() {

	fmt.Println("hello world")

	jwt := JWT{}

	jwt.Decode("test")
}
