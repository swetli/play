package main

import "fmt"

func hello(name string) string {
	str := "Hello, " + name + "!"
	return str
}

func main() {
	fmt.Println(hello("Swetli"))
}
