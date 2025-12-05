package main

import "fmt"

func main() {
	println(HelloWorld())
}

func HelloWorld() string {
	return "Hello, world!"
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
